package server

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/dogstatsd"
	trans "github.com/go-kit/kit/transport/http"
	"net/http"
	"os"
	"time"
)

const (
	statsHostPort = "127.0.0.1:8126"
)

func StartEchoServer() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	stats := dogstatsd.New("echo_service", logger)

	var es EchoService
	es = NewEchoService()
	es = loggingMiddleware(logger)(es)
	es = instrumentingMiddleware(stats.NewCounter("echo_counter", 1), stats.NewHistogram("echo_latency", 1))(es)

	echoHandler := trans.NewServer(
		makeEchoEndpoint(es),
		decodeEchoRequest,
		encodeResponse,
	)

	http.Handle("/", echoHandler)
	_ = http.ListenAndServe(":8080", nil)

	stats.SendLoop(context.Background(), time.Tick(2*time.Second), "tcp", statsHostPort)
}
