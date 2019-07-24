package bootcamp

type ThreadPool struct {
	queueSize   int64
	workers     int
	pool        chan chan interface{}
	job         chan interface{}
	closeHandle chan bool
}

func NewFixedThreadPool(poolSize int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{workers: poolSize, queueSize: queueSize}

	threadPool.job = make(chan interface{}, queueSize)
	threadPool.pool = make(chan chan interface{}, queueSize)
	threadPool.closeHandle = make(chan bool)

	threadPool.createPool()

	return threadPool
}

func (p *ThreadPool) Submit(task Callable) *Future {
	handle := &Future{response: make(chan interface{})}
	futureTask := CallableTask{Task: task, Handle: handle}

	p.job <- futureTask

	return futureTask.Handle
}

func (p *ThreadPool) createPool() {
	for i := 0; i < p.workers; i++ {
		go func() {
			for {
				p.pool <- p.job
				select {
				case job := <-p.job:
					task, _ := job.(CallableTask)

					res := task.Task.Call()

					task.Handle.done = true
					task.Handle.response <- res

				case closed := <-p.closeHandle:
					if closed {
						return
					}
				}
			}
		}()
	}
}
