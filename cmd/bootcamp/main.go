package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp/feedbackserver"
)

func main() {
	go feedbackserver.StartServer()

	_, _ = fmt.Scanln()
}
