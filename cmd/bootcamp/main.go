package main

import (
	"grabvn-golang-bootcamp/internal/bootcamp/client"
	"grabvn-golang-bootcamp/internal/bootcamp/server"
)

func main() {
	go server.Server()
	go client.Client()
}
