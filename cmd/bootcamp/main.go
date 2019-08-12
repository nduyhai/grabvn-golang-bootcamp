package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp"
)

func main() {
	go bootcamp.StartServer()
	_, _ = fmt.Scanln()
}
