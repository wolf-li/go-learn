package main

import (
	"time"
	"fmt"
	"strconv"
)

func shopping(name string){
	fmt.Println("() 开始购物", name)
	time.Sleep(time.Second)
	fmt.Println("() 结束购物", name)
}

func main(){
	// 未使用 go 
	now := time.Now()
	shopping("1")
	shopping("2")
	shopping("3")
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	now = time.Now()
	// 并发 go 写法，匿名函数方式
	for i:=0; i < 3; i++{
		go func(name string){
			fmt.Println("() 开始购物", name)
			time.Sleep(time.Second)
			fmt.Println("() 结束购物", name)
		}(strconv.Itoa(i))
	}
	time.Sleep(time.Millisecond * 1200)
	fmt.Println(time.Since(now))
	
}