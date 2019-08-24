package client

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"time"
)

const (
	url        = "http://localhost:8080/test"
	cmd        = "echo_cmd"
	hystrixAdd = "localhost:8088"
)

func StartEchoClient() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	hystrix.ConfigureCommand(cmd, hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(hystrixAdd, hystrixStreamHandler)

	var client EchoClient
	client = NewEchoClient(url)
	client = retryMiddleware(logger)(client)
	client = circuitBreakerMiddleware(cmd, logger)(client)

	for {
		res, err := client.GetEcho()
		if err == nil {
			_ = logger.Log("source", "client", "method", "echo", "code", res.GetStatusCode(), ", status: ", res.GetStatus())
		} else {
			_ = logger.Log("source", "client", "method", "echo", "error", err.Error())
		}

		time.Sleep(100 * time.Millisecond)
	}

}
