package server

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"grabvn-golang-bootcamp/internal/common"
	"net/http"
)

func makeEchoEndpoint(svc EchoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Echo(), nil
	}
}

func decodeEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(_ context.Context, writer http.ResponseWriter, rs interface{}) error {
	res := rs.(common.HttpResponse)
	writer.WriteHeader(res.GetStatusCode())

	if res.GetBody() != "" {
		_, err := writer.Write([]byte(res.GetBody()))
		return err
	} else {
		return nil
	}

}
