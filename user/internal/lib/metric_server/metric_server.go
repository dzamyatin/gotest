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
	fmt.Printf("Metrics: port: %d; user: %s; pwd: %s;\n", metricPort, flagconfig.GetFlagConfig().MetricAuthUser, flagconfig.GetFlagConfig().MetricAuthPwd)

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
			log.Fatalf("Failed to listen metrics: %v", errProm)
			return
		}
	}()
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
