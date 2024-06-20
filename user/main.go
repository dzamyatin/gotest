package main

import (
	api "app/user/api/user/proto"
	"app/user/internal/di/interceptor"
	"app/user/internal/di/server"
	_ "app/user/internal/lib"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	port       = flag.Int("port", 9999, "Port to listen")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write mem profile to file")
)

func main() {
	flag.Parse()
	fmt.Printf("Listen: %v \n", *port)

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.SessionServerInterceptor),
	)
	registerServers(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("GRPC failed to serve: %v", err)
	}

	os.Exit(0)
}

func registerServers(grpcServer grpc.ServiceRegistrar) {
	api.RegisterUserServer(grpcServer, server.GetUserServer())
}
