package client

import (
	"github.com/avast/retry-go"
	"github.com/go-kit/kit/log"
	"grabvn-golang-bootcamp/internal/common"
	"time"
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
		retry.RetryIf(func(err error) bool {
			err, ok := err.(iError)
			if ok {
				return true
			} else {
				return false
			}
		}),
		retry.OnRetry(func(n uint, err error) {
			_ = e.log.Log("source", "client", "method", "retry", "error", err.Error(), "num", n)
		}),
		retry.Attempts(5),
		retry.Delay(100*time.Millisecond),
		retry.LastErrorOnly(true),
	)

	return res, err
}
