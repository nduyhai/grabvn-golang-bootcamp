package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go send(message)
	go receive(message)

	select {}
}

func send(message chan<- string) {
	for {
		msg := "do"
		message <- msg
		fmt.Println("Send:", msg)
		time.Sleep(1)
	}
}

func receive(message <-chan string) {
	for {
		msg := <-message
		fmt.Println("Receive:", msg)
		time.Sleep(1)
	}
}
