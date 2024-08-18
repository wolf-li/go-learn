package main

import (
	"fmt"
	"sync"
)

// 范围太小没有效果，要使用大数
func add(n *int, wait *sync.WaitGroup){
	defer wait.Done()
	for i:=0; i<100000; i++{
		*n += i
	}
}

func sub(n *int, wait *sync.WaitGroup){
	defer wait.Done()
	for i:=0; i<100000; i++{
		*n -= i
	}
}

func main(){
	var wg sync.WaitGroup
	var n int = 0
	wg.Add(2)
	go add(&n, &wg)
	go sub(&n, &wg)
	wg.Wait()
	fmt.Println(n)
}