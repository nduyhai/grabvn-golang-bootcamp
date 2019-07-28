package client

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/msg"
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
	c := msg.NewMessageServiceClient(conn)

	req := msg.Message{Uuid: uuid.New().String(), Data: "Say Hi"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Send(ctx, &req)
	if err != nil {
		log.Fatalf("failed to receive msg: %v", err)
	} else {
		log.Printf("Received: %v", res)

	}
}
