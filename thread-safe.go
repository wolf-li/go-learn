package main

import (
	"fmt"
	"sync"
)


// 范围太小没有效果，要使用大数
func add(n *int, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	lock.Lock()
	for i:=0; i<100000; i++{
		*n += i
	}
	lock.Unlock()
}

func sub(n *int, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	lock.Lock()
	for i:=0; i<100000; i++{
		*n -= i
	}
	lock.Unlock()
}

func main(){
	var wg sync.WaitGroup
	var l sync.Mutex
	var n int = 0
	wg.Add(2)
	go add(&n, &wg, &l)
	go sub(&n, &wg, &l)
	wg.Wait()
	fmt.Println(n)
}