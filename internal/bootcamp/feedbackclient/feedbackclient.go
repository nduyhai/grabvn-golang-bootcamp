package feedbackclient

import (
	"context"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/configuration"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"time"
)

func StartClient() {
	var config configuration.Conf
	config.LoadConf()

	log.Print("begin init rpc client....")

	conn, err := grpc.Dial("localhost"+config.RPC.Port, grpc.WithInsecure())

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
