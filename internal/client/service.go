package client

import (
	"grabvn-golang-bootcamp/internal/common"
)

type EchoClient interface {
	GetEcho() (common.HttpResponse, error)
}

type EchoClientMiddleware func(EchoClient) EchoClient

type echoClient struct {
	url        string
	httpClient common.HttpClient
}

func NewEchoClient(url string) *echoClient {
	return &echoClient{url: url, httpClient: common.NewHttpClient()}
}

func (e *echoClient) GetEcho() (common.HttpResponse, error) {
	res, err := e.httpClient.Get(e.url)
	if err != nil {
		return res, err
	}
	if 500 <= res.GetStatusCode() {
		return res, newIError(res.GetStatus())
	}
	return res, err
}
