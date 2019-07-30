package feedbackclient

import (
	"context"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"time"
)

const (
	port = ":9000"
)

func StartClient() {
	log.Print("begin init rpc client....")

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect server: %v", err)
	}
	defer conn.Close()
	client := feedback.NewFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	res, err := client.GetById(ctx, &feedback.FeedbackRequest{})
	if err != nil {
		log.Fatalf("failed to receive msg: %v", err)
	} else {
		log.Printf("Received: %v", res)

	}
}