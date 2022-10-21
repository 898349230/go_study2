package work

import "sync"

//  Worker 必须满足接口类型 才能使用工作池
type Worker interface {
	Task()
}

//  Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// 创建一个工作池
func New(maxGotoutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGotoutine)
	for i := 0; i < maxGotoutine; i++ {
		go func() {
			// 这里 for range 循环会一直阻塞，直到从 work 通道收到一个 Worker 接口值，
			// 收到值后就会执行这个值得 Task 方法
			for w := range p.work {
				w.Task()
			}
		}()
	}
	return &p
}

//  提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	// 关闭 work 通道，导致所有池里的goroutine 停止工作
	close(p.work)
	// 让等待所有goroutine终止
	p.wg.Wait()
}
