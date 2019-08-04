//+build wireinject

package bootcamp

import (
	"fmt"
)

type Message string

func NewMessage(msg string) Message {
	return Message("msg")
}

type Greeter struct {
	Message Message
}

func NewGreeter(message Message) Greeter {
	return Greeter{Message: message}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(greeter Greeter) Event {
	return Event{Greeter: greeter}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
