package main

import (
	"context"
	"fmt"
	"github.com/benjaminhadfield/grpc/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	message := fmt.Sprintf("Hello from client at %v", time.Now().UTC())
	response, err := c.SayHello(context.Background(), &api.PingMessage{Message: message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}
	log.Printf("Response from server: %v", response)
}
