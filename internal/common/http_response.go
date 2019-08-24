package common

type HttpResponse interface {
	GetStatusCode() int
	GetStatus() string
	GetBody() string
}

type httpResponse struct {
	statusCode int
	status     string
	body       string
}

func NewHttpResponse(statusCode int, status string, body string) HttpResponse {
	return &httpResponse{statusCode: statusCode, status: status, body: body}
}

func (h *httpResponse) GetStatusCode() int {
	return h.statusCode
}

func (h *httpResponse) GetStatus() string {
	return h.status
}

func (h *httpResponse) GetBody() string {
	return h.body
}
