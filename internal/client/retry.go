package client

import (
	"github.com/avast/retry-go"
	"github.com/go-kit/kit/log"
	"grabvn-golang-bootcamp/internal/common"
)

func retryMiddleware(log log.Logger) EchoClientMiddleware {
	return func(next EchoClient) EchoClient {
		return &rMw{log, next}
	}
}

type rMw struct {
	log log.Logger
	EchoClient
}

func (e *rMw) GetEcho() (common.HttpResponse, error) {
	var res common.HttpResponse
	var err error
	err = retry.Do(
		func() error {
			res, err = e.EchoClient.GetEcho()
			return err
		},
		retry.OnRetry(func(n uint, err error) {
			_ = e.log.Log("source", "server", "method", "retry", "error", err.Error())
		}),
	)

	return res, err
}
