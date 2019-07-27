package server

import (
	"context"
	"grabvn-golang-bootcamp/internal/bootcamp/msg"
	"log"
	"net"
	"net/rpc"
)

const (
	port = ":9000"
)

type server struct{}

func (s *server) Send(ctx context.Context, in *msg.Message) (*msg.MessageResponse, error) {
	log.Printf("Received: %v", in)
	return &msg.MessageResponse{Uuid: in.Uuid, Status: true}, nil
}

func Server() {
	log.Print("begin init rpc server....")
	err := rpc.Register(new(server))
	if err != nil {
		log.Fatalf("failed to register server: %v", err)

	}
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
