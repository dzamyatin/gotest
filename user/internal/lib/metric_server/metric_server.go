package metric_server

import (
	"app/user/internal/config/flagconfig"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" //Register debug request listeners
	"strings"
)

func StartMetricServer(ctx context.Context) {
	mx := http.DefaultServeMux

	//prometheus runtime metrics,
	mx.Handle("/metrics", promhttp.Handler())

	metricPort := flagconfig.GetFlagConfig().MetricPort

	help()

	handler := newMiddlewareAuth(mx, func(writer http.ResponseWriter, request *http.Request, h http.Handler) {
		if strings.HasPrefix(request.RequestURI, "/debug/") || strings.HasPrefix(request.RequestURI, "/metrics/") {
			if !checkAuth(request) {
				writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted Area"`)
				writer.WriteHeader(401)
				_, err := writer.Write([]byte("Resource portected"))
				if err != nil {
					log.Println(err)
				}
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
			log.Fatalf("Failed to listen metrics: %v", errProm)
			return
		}
	}()
}

func help() {
	fmt.Printf("**********\n")
	fmt.Printf("Metrics: port: %d; user: %s; pwd: %s;\n", flagconfig.GetFlagConfig().MetricPort, flagconfig.GetFlagConfig().MetricAuthUser, flagconfig.GetFlagConfig().MetricAuthPwd)
	url := fmt.Sprintf("%s:%d/debug/pprof/trace?seconds=10", flagconfig.GetFlagConfig().ServiceExternalUrl, flagconfig.GetFlagConfig().MetricPort)
	fmt.Printf("trace: %s\n", url)
	fmt.Printf(
		"curl -o trace.out -u %s:%s %s \n",
		flagconfig.GetFlagConfig().MetricAuthUser,
		flagconfig.GetFlagConfig().MetricAuthPwd,
		url,
	)
	fmt.Printf("**********\n")
}

func checkAuth(r *http.Request) bool {
	u, pwd, ok := r.BasicAuth()

	if !ok {
		return false
	}

	if u == flagconfig.GetFlagConfig().MetricAuthUser && pwd == flagconfig.GetFlagConfig().MetricAuthPwd {
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
