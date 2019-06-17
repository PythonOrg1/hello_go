package main

//package hello_go

import (
	"fmt"
	_ "net/http"
)

type mystring string

type user struct {
	name string
	id, age int
	gender string
}

//定义接口
type Reader interface {
	Read(p []byte) (n int, err error)
}


func init() {
	fmt.Println("this is init...")
}

func main()  {
	fmt.Println("Hello, Golang!")

	//client := http.Client{
	//}
	//
	////(resp *Response, err error)
	//fmt.Println( client.Get("http://www.baidu.com") )


	//log("test")

	go log("")
}

func log(msg string)  {
	fmt.Println("i'm dong some logs here now, msg=[" + msg + "]")
}
