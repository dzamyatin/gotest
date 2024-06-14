package main

import (
	api "app/user/api/user/proto"
	"app/user/internal/config"
	_ "app/user/internal/lib"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 9999, "Port to listen")
)

func main() {
	flag.Parse()
	fmt.Printf("Listen: %v \n", *port)

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	registerServers(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("GRPC failed to serve: %v", err)
	}

	os.Exit(0)
}

func registerServers(grpcServer grpc.ServiceRegistrar) {
	api.RegisterUserServer(grpcServer, config.GetUserServer())
}
