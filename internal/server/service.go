package server

import (
	"grabvn-golang-bootcamp/internal/common"
	"math/rand"
)

type EchoService interface {
	Echo() common.HttpResponse
}

type EchoServiceMiddleware func(EchoService) EchoService

type echoService struct {
}

func NewEchoService() EchoService {
	return &echoService{}
}

func (e *echoService) Echo() common.HttpResponse {
	// 30% chance of failure
	if rand.Intn(100) < 30 {
		return common.NewHttpResponse(500, "error", "a chaos monkey broke your server")
	} else {
		// Happy path
		return common.NewHttpResponse(200, "success", "")
	}
}
