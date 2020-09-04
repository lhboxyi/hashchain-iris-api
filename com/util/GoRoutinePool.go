package util

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

/**
go 的 goroutine 提供了一种较线程而言更廉价的方式处理并发场景, go 使用二级线程的模式,
将 goroutine 以 M:N 的形式复用到系统线程上, 节省了 cpu 调度的开销, 也避免了用户级线程（协程）进行系统调用时阻塞整个系统线程的问题。
但 goroutine 太多仍会导致调度性能下降、GC 频繁、内存暴涨, 引发一系列问题。在面临这样的场景时, 限制 goroutine 的数量、重用 goroutine 显然很有价值。

使用任务池和原生 goroutine 性能相近（略好于原生）

使用任务池比直接 goroutine 内存分配节省 7000 倍左右, 内存分配次数减少 2700 倍左右

https://github.com/wazsmwazsm/mortar/blob/master/examples/simple.go
*/
var (
	// ErrInvalidPoolCap return if pool size <= 0
	ErrInvalidPoolCap = errors.New("invalid pool cap")
	// ErrPoolAlreadyClosed put task but pool already closed
	ErrPoolAlreadyClosed = errors.New("pool already closed")
)

const (
	// RUNNING pool is running
	RUNNING = 1
	// STOPED pool is stoped
	STOPED = 0
)

// Task task to-do
type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

// Pool task pool
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	state          int64
	taskC          chan *Task
	closeC         chan bool
	PanicHandler   func(interface{})
}

// NewPool init pool
func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
		closeC:   make(chan bool),
	}, nil
}

// GetCap get capacity
func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}

// GetRunningWorkers get running workers
func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

// Put put a task to pool
func (p *Pool) Put(task *Task) error {

	if p.state == STOPED {
		return ErrPoolAlreadyClosed
	}

	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	p.taskC <- task

	return nil
}

func (p *Pool) run() {
	p.incRunning()

	go func() {
		defer func() {
			p.decRunning()
			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("Worker panic: %s\n", r)
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskC:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			case <-p.closeC:
				return
			}
		}
	}()
}

// Close close pool graceful
func (p *Pool) Close() {
	p.state = STOPED // stop put task

	for len(p.taskC) > 0 { // wait all task be consumed
	}

	p.closeC <- true
	close(p.taskC)
}
func test()  {
	pool, err := NewPool(10)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		pool.Put(&Task{
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
			Params: []interface{}{i},
		})
	}
	time.Sleep(1e9)
}
