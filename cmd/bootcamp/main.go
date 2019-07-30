package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp/feedbackclient"
	"grabvn-golang-bootcamp/internal/bootcamp/feedbackserver"
)

func main() {
	go feedbackserver.StartServer()
	go feedbackclient.StartClient()

	_, _ = fmt.Scanln()
}
