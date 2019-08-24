package common

import (
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Get(url string) (HttpResponse, error)
}

type httpClientImpl struct {
}

func NewHttpClient() HttpClient {
	return &httpClientImpl{}
}

func (h *httpClientImpl) Get(url string) (HttpResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return NewHttpResponse(501, "cannot execute request", ""), nil
	} else {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyString := string(bodyBytes)

		return NewHttpResponse(resp.StatusCode, resp.Status, bodyString), nil
	}
}
