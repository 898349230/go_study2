package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个无缓冲的通道
	results := make(chan *Result)
	// 构建一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup
	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// 启动一个 goroutinne 来执行搜索，这里使用了匿名函数
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// 减少 waitGroup 的计数
			waitGroup.Done()
		}(matcher, feed)
	}
	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()
		// 用关闭通道的方式，通知Display函数，可以退出程序了，这是个内置函数
		close(results)
	}()

	// 启动函数，显示返回的结果，并且在最后一个结果显示完成后返回
	Display(results)
}

// Register调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
