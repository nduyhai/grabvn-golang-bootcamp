package server

import (
	"context"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/msg"
	"log"
	"net"
)

const (
	port = ":9000"
)

type Server struct{}

func (s *Server) Send(ctx context.Context, in *msg.Message) (*msg.MessageResponse, error) {
	log.Printf("Received: %v", in)
	return &msg.MessageResponse{Uuid: in.Uuid, Status: true}, nil
}

func StartServer() {
	log.Print("begin init rpc server....")

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	msg.RegisterMessageServiceServer(server, &Server{})
	if server == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := server.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
