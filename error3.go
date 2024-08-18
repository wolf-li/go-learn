package main

import (
	"fmt"
	"runtime/debug"
)

func read(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()

	// 问题的代码
	var list = []int{1,2}
	fmt.Println(list[2])
}

func main(){
	read()
}