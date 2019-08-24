package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/client"
	"grabvn-golang-bootcamp/internal/server"
)

func main() {
	go server.StartEchoServer()

	go client.StartEchoClient()

	_, _ = fmt.Scanln()
}
