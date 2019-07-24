package bootcamp

type Callable interface {
	Call() interface{}
}

type Future struct {
	response chan interface{}
	done     bool
}

func (f *Future) Get() interface{} {
	return <-f.response
}

func (f *Future) IsDone() bool {
	return f.done
}

type CallableTask struct {
	Task   Callable
	Handle *Future
}
