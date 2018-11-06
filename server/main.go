package main

import (
	"fmt"
	"log"
	"net"

	"github.com/benjaminhadfield/grpc/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
