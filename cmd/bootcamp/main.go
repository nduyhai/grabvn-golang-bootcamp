package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp/client"
	"grabvn-golang-bootcamp/internal/bootcamp/server"
)

func main() {
	go server.StartServer()
	go client.StartClient()

	_, _ = fmt.Scanln()
}
