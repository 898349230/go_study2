package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool管理一组可以安全地在多个goroutine间共享的资源
// 被管理的资源必须实现 io.Closer 接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	close     bool
}

// ErrPoolClosed 表示请求了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed")

// 创建一个用来管理资源的池，这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 从翅中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

		// 因为没有空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire: ", "New Resource")
		return p.factory()
	}
}

// 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和 close 操作的安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果池已经被关闭，销毁这个资源
	if p.close {
		r.Close()
		return
	}
	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release: ", "In Queue")
		// 如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	// 保证本操作和 release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	if p.close {
		return
	}
	// 关闭
	p.close = true
	// 在清空通道的资源之前，将通道关闭，如果不这样，会发生思索
	close(p.resources)
	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
