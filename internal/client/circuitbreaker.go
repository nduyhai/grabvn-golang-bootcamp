package client

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"grabvn-golang-bootcamp/internal/common"
)

func circuitBreakerMiddleware(cmd string, log log.Logger) EchoClientMiddleware {
	return func(next EchoClient) EchoClient {
		return &cbMw{cmd, log, next}
	}
}

type cbMw struct {
	cnd string
	log log.Logger
	EchoClient
}

func (e *cbMw) GetEcho() (res common.HttpResponse, err error) {

	if err := hystrix.Do(e.cnd, func() (err error) {
		res, err = e.EchoClient.GetEcho()
		return err
	}, nil); err != nil {
		_ = e.log.Log("source", "client", "method", "echo", "fallback", err.Error())
		return nil, err
	}
	return res, nil
}
