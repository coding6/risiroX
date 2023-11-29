package service

import (
	"errors"
	"fmt"
	"risirox/risirox/src/conf"
	"risirox/risirox/src/data"
	"sync/atomic"
)

const (
	RUNNING = 1
	STOPPED = 2
)

// WorkPool 写线程工作池
/*

 */
type WorkPool struct {
	poolSize     uint64
	maxPoolSize  uint32
	runningTasks uint64
	state        int32
	taskQueue    chan *data.Task
	close        chan bool
	panicHandler func(interface{})
}

var WorkPoolObj *WorkPool

func NewWorkPool() *WorkPool {
	if conf.GlobalConfigObj.WorkerPoolSize < 0 {
		return nil
	}
	return &WorkPool{
		poolSize:    conf.GlobalConfigObj.WorkerPoolSize,
		maxPoolSize: conf.GlobalConfigObj.MaxWorkerPoolSize,
		state:       RUNNING,
		taskQueue:   make(chan *data.Task),
		close:       make(chan bool),
	}
}

func (pool *WorkPool) Submit(task *data.Task) error {
	if pool.state == STOPPED {
		return errors.New("pool is stopped")
	}
	if pool.getWorkers() < pool.poolSize {
		pool.run()
	}
	pool.taskQueue <- task
	return nil
}

func (pool *WorkPool) getWorkers() uint64 {
	return atomic.LoadUint64(&pool.runningTasks)
}

func (pool *WorkPool) incrRunningTask() {
	atomic.AddUint64(&pool.runningTasks, 1)
}

func (pool *WorkPool) delRunningTask() {
	atomic.AddUint64(&pool.runningTasks, ^uint64(0))
}

func (pool *WorkPool) run() {
	pool.incrRunningTask()
	go func() {
		defer func() {
			pool.delRunningTask()
		}()
		for {
			select {
			case task, ok := <-pool.taskQueue:
				if !ok {
					fmt.Printf("data is not ready")
					return
				}
				task.Handler(task.Param)
			case <-pool.close:
				return
			}
		}
	}()

}

func (pool *WorkPool) Close() {
	pool.state = STOPPED
	for len(pool.taskQueue) > 0 {

	}
	pool.close <- true
	close(pool.taskQueue)
}

func init() {
	WorkPoolObj = NewWorkPool()
}
