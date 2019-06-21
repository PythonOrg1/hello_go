package main

//package hello_go

import (
	"fmt"
	_ "net/http"
	"unsafe"
)

type mystring string

type User struct {
	name string
	id, age int
	gender string
}

//定义接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

var user User

func init() {
	fmt.Println("this is init...")

	user.age = 10
	user.id = 1
	user.name = "zhangsan"
	user.gender = "man"

}

func main()  {
	fmt.Println("Hello, Golang!")

	//client := http.Client{
	//}
	//
	////(resp *Response, err error)
	//fmt.Println( client.Get("http://www.baidu.com") )


	//log("test")

	log("")
}

const(
	U = 1
	F = 2
	M = 3
)

func log(msg string)  {
	fmt.Println("i'm dong some logs here now, msg=[" + msg + "]")
	fmt.Println(user)

	var a = "ddd"
	fmt.Println("sss", a)

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(len(a))
	//fmt.Println(cap(  ))

	//fmt.Println(21*0.12)
	//
	//fmt.Println(23*0.12)



}
