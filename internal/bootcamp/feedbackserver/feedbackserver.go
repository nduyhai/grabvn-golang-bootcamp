package feedbackserver

import (
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"net"
)

const (
	port = ":9000"
)

func StartServer() {
	log.Print("begin init rpc server....")

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	feedback.RegisterFeedbackServiceServer(rpcServer, &server{})
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
