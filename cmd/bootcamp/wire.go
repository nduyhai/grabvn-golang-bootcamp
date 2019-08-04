//+build wireinject

package main

import (
	"github.com/google/wire"
	"grabvn-golang-bootcamp/internal/bootcamp"
)

func InitializeEvent(phrase string) bootcamp.Event {
	wire.Build(bootcamp.NewEvent, bootcamp.NewGreeter, bootcamp.NewMessage)
	return bootcamp.Event{}
}
