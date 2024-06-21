package main

import (
	api "app/user/api/user/proto"
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/server"
	"app/user/internal/di/server/interceptor"
	"app/user/internal/di/static"
	_ "app/user/internal/lib"
	"app/user/internal/lib/profiler"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" //Register debug request listeners
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fc := flagconfig.GetFlagConfig()
	if fc.Help {
		fmt.Printf("Listen and serve grpc requests \n\n Available flags:\n")
		fc.PrintHelp()
		os.Exit(0)
	}

	prof := static.GetProfiler()
	prof.Start()

	port := flagconfig.GetFlagConfig().Port

	fmt.Printf("Listen: %v \n", port)
	fmt.Printf("Process: %v \n", os.Getpid())

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	defer listener.Close()

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.SessionServerInterceptor,
			interceptor.InterceptorProfiler,
			interceptor.InterceptorTrace,
		),
	)
	registerServers(grpcServer)

	startMetricServer(ctx)

	graceFullShutdown(grpcServer, listener, ctx, cancel, prof)

	log.Print("Grpc is running...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("GRPC failed to serve: %v", err)
	}

	prof.Stop()

	os.Exit(0)
}

func registerServers(grpcServer grpc.ServiceRegistrar) {
	api.RegisterUserServer(grpcServer, server.GetUserServer())
}

func startMetricServer(ctx context.Context) {
	mx := http.DefaultServeMux

	//prometheus runtime metrics,
	mx.Handle("/metrics", promhttp.Handler())

	metricPort := flagconfig.GetFlagConfig().MetricPort
	fmt.Printf("Metrics: %d\n", metricPort)

	handler := newMiddlewareAuth(mx, func(writer http.ResponseWriter, request *http.Request, h http.Handler) {
		if strings.HasPrefix(request.RequestURI, "/debug/") || strings.HasPrefix(request.RequestURI, "/metrics/") {
			if !checkAuth(request) {
				writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted Area"`)
				writer.WriteHeader(401)
				writer.Write([]byte("Resource portected"))
				return
			}
		}

		h.ServeHTTP(writer, request)
	})

	//run server
	go func() {
		metricServ := &http.Server{
			Addr:    fmt.Sprintf(":%d", metricPort),
			Handler: handler,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		}

		errProm := metricServ.ListenAndServe()
		if errProm != nil {
			log.Fatalf("Failed to listen prometeus: %v", errProm)
			return
		}
	}()
}

func checkAuth(r *http.Request) bool {
	u, pwd, ok := r.BasicAuth()

	if !ok {
		return false
	}

	if u == "test" && pwd == "test" {
		return true
	}

	return false
}

type middlewareAuth struct {
	h  http.Handler
	fn func(writer http.ResponseWriter, request *http.Request, h http.Handler)
}

func newMiddlewareAuth(h http.Handler, fn func(writer http.ResponseWriter, request *http.Request, h http.Handler)) middlewareAuth {
	return middlewareAuth{h: h, fn: fn}
}

func (m middlewareAuth) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.fn(writer, request, m.h)
}

func graceFullShutdown(
	grpcServer *grpc.Server,
	lister net.Listener,
	ctx context.Context,
	cancelFunc context.CancelFunc,
	prof profiler.IProfiler,
) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		cancelFunc()
		select {
		case <-time.Tick(5 * time.Second):
		case <-ctx.Done():
		}
		grpcServer.GracefulStop()
		lister.Close()
		prof.Stop()

		fmt.Printf("Gracefull shutdown \n")
		os.Exit(0)
	}()
}
