package bootcamp

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

type QueryService interface {
	User(string) QueryResult
}

type QueryResult struct {
	Success bool
	ID      *int64
	User    string
	URL     *string
}

type QueryServiceImpl struct {
	client *github.Client
}

func NewQueryServiceImpl() *QueryServiceImpl {
	client := github.NewClient(nil)
	return &QueryServiceImpl{client: client}
}

func (q *QueryServiceImpl) User(u string) QueryResult {
	res := QueryResult{User: u}
	user, response, err := q.client.Users.Get(context.Background(), u)
	if err == nil {
		log.Print(user, response)
		res.ID = user.ID
		res.URL = user.URL
		res.Success = true
	} else {
		log.Print(err)
		res.Success = false
	}
	return res
}
