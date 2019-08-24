package client

type httpResponse interface {
	GetStatusCode() int
	GetStatus() string
	GetBody() string
}

type httpResponseImpl struct {
	statusCode int
	status     string
	body       string
}

func newHttpResponse(statusCode int, status string, body string) httpResponse {
	return &httpResponseImpl{statusCode: statusCode, status: status, body: body}
}

func (h *httpResponseImpl) GetStatusCode() int {
	return h.statusCode
}

func (h *httpResponseImpl) GetStatus() string {
	return h.status
}

func (h *httpResponseImpl) GetBody() string {
	return h.body
}
