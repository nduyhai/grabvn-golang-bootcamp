package bootcamp

type QueryService interface {
	User(string) QueryResult
}

type QueryResult struct {
	User string
}

type QueryServiceImpl struct {
}

func NewQueryServiceImpl() *QueryServiceImpl {
	return &QueryServiceImpl{}
}

func (q *QueryServiceImpl) User(user string) QueryResult {
	res := QueryResult{User: user}

	return res
}
