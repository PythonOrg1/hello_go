package main

import (
	"fmt"
	"log"
	"os"

	"hello_go/searcher/search"

	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
)

func Match(matcher search.Matcher, feed, item, results) {

}

//before main
func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

//main starter
func main() {
	fmt.Println("Seracher is running... ")
	//matchers.Start("xxxx")

	//使用特定的项做搜索
	search.Run("boss")
}
