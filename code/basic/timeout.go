package main

import (
	"time"
	"fmt"
)

var done = make(chan struct{})

func event(){
	defer close(done)
	fmt.Println("start")
	time.Sleep(time.Millisecond * 3)
	fmt.Println("end")
}

func main(){
	go event()

	select{
	case <- done:
		fmt.Println("执行完毕")
	case <- time.After(time.Millisecond * 4):
		fmt.Println("超时")
	}
}