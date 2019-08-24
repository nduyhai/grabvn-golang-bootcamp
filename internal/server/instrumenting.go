package server

import (
	"github.com/go-kit/kit/metrics"
	"grabvn-golang-bootcamp/internal/common"
	"time"
)

func instrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
) EchoServiceMiddleware {
	return func(next EchoService) EchoService {
		return &instMw{requestCount, requestLatency, next}
	}
}

type instMw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	EchoService
}

func (mw *instMw) Echo() (res common.HttpResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Echo"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())

	}(time.Now())

	res = mw.EchoService.Echo()
	return
}
