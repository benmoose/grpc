package api

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello (ctx context.Context, msg PingMessage) (*PingMessage, error) {
	log.Printf("Received message: %s", msg.Message)
	return &PingMessage{Message: "Hello from Server!"}, nil
}
