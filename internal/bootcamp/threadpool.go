package bootcamp

import (
	"log"
)

type ThreadPool struct {
	queueSize   int64
	workers     int
	pool        chan chan FutureTask
	job         chan FutureTask
	closeHandle chan bool
}
type Task interface {
	Execute() map[string]int
}

type Future struct {
	Data chan map[string]int
}

type FutureTask struct {
	Handler  Task
	Response *Future
}

func NewFixedThreadPool(poolSize int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{workers: poolSize, queueSize: queueSize}

	threadPool.job = make(chan FutureTask, queueSize)
	threadPool.pool = make(chan chan FutureTask, queueSize)
	threadPool.closeHandle = make(chan bool)

	threadPool.createPool()

	return threadPool
}

func (p *ThreadPool) Submit(task Task) *Future {
	response := &Future{make(chan map[string]int, 1)}
	future := FutureTask{task, response}
	p.job <- future
	return future.Response
}

func (p *ThreadPool) createPool() {
	log.Println("Create pool.....", p.workers)
	for i := 0; i < p.workers; i++ {
		go func() {
			for {
				p.pool <- p.job
				select {
				case job := <-p.job:
					log.Println("Executing")
					result := job.Handler.Execute()
					log.Println("Executed")
					job.Response.Data <- result

					close(job.Response.Data)
				case <-p.closeHandle:
					return
				}
			}
		}()
	}
}
