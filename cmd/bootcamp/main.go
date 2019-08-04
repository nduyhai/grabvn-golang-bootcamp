package main

import (
	"fmt"
	"github.com/google/wire"
	"grabvn-golang-bootcamp/internal/bootcamp"
)

func main() {
	event := InitializeEvent("Say Hi!!!")
	event.Start()
}

func InitializeEvent(msg string) bootcamp.Event {
	msg = wire.Build(bootcamp.NewEvent, bootcamp.NewGreeter, bootcamp.NewMessage)
	fmt.Println(msg)
	return bootcamp.Event{}
}
