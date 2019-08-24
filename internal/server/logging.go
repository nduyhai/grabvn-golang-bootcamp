package server

import (
	"github.com/go-kit/kit/log"
	"grabvn-golang-bootcamp/internal/common"
	"time"
)

func loggingMiddleware(logger log.Logger) EchoServiceMiddleware {
	return func(next EchoService) EchoService {
		return &logMw{logger, next}
	}
}

type logMw struct {
	logger log.Logger
	EchoService
}

func (mw *logMw) Echo() (res common.HttpResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"source", "server",
			"method", "Echo",
			"code", res.GetStatusCode(),
			"status", res.GetStatus(),
			"took", time.Since(begin),
		)
	}(time.Now())

	res = mw.EchoService.Echo()
	return
}
