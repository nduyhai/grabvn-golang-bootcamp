package client

import (
	"grabvn-golang-bootcamp/internal/response"
	"io/ioutil"
	"net/http"
)

type httpClient interface {
	Get(url string) response.HttpResponse
}

type httpClientImpl struct {
}

func newHttpClient() httpClient {
	return &httpClientImpl{}
}

func (h *httpClientImpl) Get(url string) response.HttpResponse {
	resp, err := http.Get(url)
	if err != nil {
		return response.NewHttpResponse(501, "cannot execute request", "")
	} else {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return response.NewHttpResponse(400, "400 read response error", err.Error())
		}
		bodyString := string(bodyBytes)

		return response.NewHttpResponse(resp.StatusCode, resp.Status, bodyString)
	}
}
