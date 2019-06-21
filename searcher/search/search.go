package search

import (
	"fmt"
	"log"
	"sync" //提供同步goroutine的功能

	//"github.com/goinaction/code/chapter2/sample/search"

)

//注册用于匹配搜索的匹配器映射
var matchers = make(map[string]Matcher)

func init() {

}

//执行搜索逻辑
func Run(item string) {
	fmt.Println("What u wanna to search is [ " + item + "]")

	//数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果  【通道channel，引用类型，是一组用于在 goroutine之间传递的值，】
	results := make(chan *Result)

	//借助sync，构造一个 处理所有数据的 wait group （感觉类似于一个队列）【sync的WatiGroup能跟踪所有启动的goroutine工作是否结束】
	var waitGroup sync.WaitGroup
	//设置休要等待处理的每个数据源的 goroutine 数量
	waitGroup.Add(len(feeds))

	// 遍历数据， 逐个开启协程
	for _, feed := range feeds {
		//for i, feed := range feeds {
		//fmt.Println(i)
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个goroutine执行搜索业务， 循环里开启了很多个，可以并发的执行
		// 以下是一个 "匿名函数"
		// 参数变量：'*Feed' -- Feed 类型的指针
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, item, results)
			waitGroup.Done()		//完成后调用，
		}(matcher, feed)
	}

	// go func  定义方法并直接执行
	// 启动一个goroutine专门监控所有的任务 是否完成
	go func() {
		// 等待所有任务执行完成  阻塞式，阻塞当前的goroutine，直到WaitGroup内部计数到0，
		waitGroup.Wait()

		//关闭 通道， 通知Display函数，可以退出程序了
		close(results)
	}()

	// 显示返回结果
	Display(results)
}
