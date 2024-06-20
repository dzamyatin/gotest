package main

import (
	api "app/user/api/user/proto"
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/interceptor"
	"app/user/internal/di/server"
	_ "app/user/internal/lib"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	//flag.Parse()

	port := flagconfig.GetFlagConfig().Port

	fmt.Printf("Listen: %v \n", port)

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
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
