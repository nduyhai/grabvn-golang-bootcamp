package client

import (
	"github.com/google/uuid"
	"grabvn-golang-bootcamp/internal/bootcamp/msg"
	"log"
	"net/rpc"
)

const (
	port = ":9000"
)

func Client() {
	log.Print("begin init rpc client....")

	c, err := rpc.Dial("tcp", "localhost"+port)
	if err != nil {
		log.Fatalf("failed to connect server: %v", err)
	}
	var res msg.MessageResponse
	req := msg.Message{Uuid: uuid.New().String(), Data: "Say Hi"}
	err = c.Call("Server.Send", req, &res)
	if err != nil {
		log.Fatalf("failed to receive msg: %v", err)
	} else {
		log.Printf("Received: %v", res)

	}
}
