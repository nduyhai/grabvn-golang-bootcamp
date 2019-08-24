package client

import (
	"io/ioutil"
	"net/http"
)

type httpClient interface {
	Get(url string) httpResponse
}

type httpClientImpl struct {
}

func newHttpClient() httpClient {
	return &httpClientImpl{}
}

func (h *httpClientImpl) Get(url string) httpResponse {
	resp, err := http.Get(url)
	if err != nil {
		return newHttpResponse(501, "cannot execute request", "")
	} else {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return newHttpResponse(400, "400 read response error", err.Error())
		}
		bodyString := string(bodyBytes)

		return newHttpResponse(resp.StatusCode, resp.Status, bodyString)
	}
}
