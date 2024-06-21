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
	"os"
	"os/signal"
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
	mx := http.NewServeMux()
	mx.Handle("/metrics", promhttp.Handler())
	go func() {
		metricPort := "8998"
		fmt.Println("Metrics: " + metricPort)

		metricServ := &http.Server{
			Addr:    ":" + metricPort,
			Handler: mx,
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
