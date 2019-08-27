package client

import "fmt"

type iError struct {
	status string
}

func newIError(status string) *iError {
	return &iError{status: status}
}

func (i *iError) Error() string {
	return fmt.Sprintf("%s", i.status)
}


