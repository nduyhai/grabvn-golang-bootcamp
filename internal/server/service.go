package server

import (
	"grabvn-golang-bootcamp/internal/response"
	"math/rand"
)

type EchoService interface {
	Echo() response.HttpResponse
}

type ServiceMiddleware func(EchoService) EchoService

type echoService struct {
}

func NewEchoService() EchoService {
	return &echoService{}
}

func (e *echoService) Echo() response.HttpResponse {
	// 30% chance of failure
	if rand.Intn(100) < 30 {
		return response.NewHttpResponse(500, "error", "a chaos monkey broke your server")
	} else {
		// Happy path
		return response.NewHttpResponse(200, "success", "")
	}
}
