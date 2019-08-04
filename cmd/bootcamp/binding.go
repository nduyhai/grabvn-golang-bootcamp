//+build wireinject

package main

import "github.com/google/wire"

type Greeter interface {
	Say() string
}

type MyGreeter string

func (g MyGreeter) Say() string {
	return string(g)
}

func ProviderGreeter() *MyGreeter {
	greeter := new(MyGreeter)
	*greeter = "Hello"
	return greeter
}
func ProviderGreeting(g Greeter) string {
	return g.Say()
}

var Set = wire.NewSet(ProviderGreeter, wire.Bind(new(Greeter), new(*MyGreeter)), ProviderGreeting)


