package main

import (
	"io"
	"log"
	"math/rand"
	"study2/ch7/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// 要使用的 goroutine 的数量
	maxGoroutine = 25
	// 池中的资源的数量
	pooledResources = 2
)

// 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便dbConnection 可以被池管理。close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的id
var idCounter int32

// 工厂函数，当需要一个新链接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	// atomic.AddInt32()  可以安全的增加包级变量 idCounter 的值
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	// 创建管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	// 使用池内的连接来完成查询
	for query := 0; query < maxGoroutine; query++ {
		// 每个 goroutine 需要自己复制一份要查询的副本，不然所有的查询会共享同一个查询变量
		go func(q int) {
			// 调用 performQueries 函数
			performQueries(q, p)
			// 退出
			wg.Done()
		}(query)
	}
	// 等待 goroutine 结束
	wg.Wait()
	// 关闭池
	log.Println("Shutdown Program.")
	p.Close()
}

// 用来测试连接的连接池
func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	// 用等待模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
