package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp/feedbackproxy"
	"grabvn-golang-bootcamp/internal/bootcamp/feedbackserver"
)

func main() {
	go feedbackserver.StartServer()
	go feedbackproxy.StartProxy()

	_, _ = fmt.Scanln()
}
