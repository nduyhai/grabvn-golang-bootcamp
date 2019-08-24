package server

import (
	"github.com/go-kit/kit/log"
	"grabvn-golang-bootcamp/internal/response"
	"time"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next EchoService) EchoService {
		return &logMw{logger, next}
	}
}

type logMw struct {
	logger log.Logger
	EchoService
}

func (mw *logMw) Echo() (res response.HttpResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Echo",
			"code", res.GetStatusCode(),
			"status", res.GetStatus(),
			"took", time.Since(begin),
		)
	}(time.Now())

	res = mw.EchoService.Echo()
	return
}
