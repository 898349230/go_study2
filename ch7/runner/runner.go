package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal
	// complete 通道报告处理任务已完成
	complete chan error
	// timeout 报告处理已超时
	timeout <-chan time.Time
	// tasks 持有一组以索引顺序依次执行的任务
	tasks []func(int)
}

// ErrTimeout 会在执行任务超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		// 初始化缓冲区容量为1的通道，可以保证通道至少能接收一个来自语言运行时的os.Signal值。
		// 确保语言运行时发送这个事件的时候不会被阻塞
		interrupt: make(chan os.Signal, 1),
		//当执行任务的goroutine完成时，会向这个通道发送一个error类型的值或者nil值，之后等待main函数接受这个值
		complete: make(chan error),
		// 到期后向这个通道发送一个 time.Time 的值
		timeout: time.After(d),
	}
}

// Add 将一个任务附加到Runner上。这个任务是一个接收 int 类型的 ID 作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接受所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	// select 语句 阻塞等待两个事件中的任意一个
	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
		// 当任务处理完成时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已经注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		fmt.Printf("task %d run... \n", id)
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已注册的任务
		// fmt.Printf("task %d 准备执行... \n", id)
		task(id)
	}
	return nil
}

// 让 goroutine 检查中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
