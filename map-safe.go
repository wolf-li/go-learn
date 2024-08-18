package main

import (
	"fmt"
	"sync"
	"time"
)

func read(m *sync.Map, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		fmt.Println(m.Load(1))
	}
}

func write(m *sync.Map, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		m.Store(1, time.Now().Format("15:04:05"))
	}
}

func readL(m map[int]string, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	for{
		lock.Lock()
		fmt.Println(m[1])
		lock.Unlock()
	}
}

func writeL(m map[int]string, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	for{
		lock.Lock()
		m[1] = "time"
		lock.Unlock()
	}
}

func main(){
	// 方式一
	// var lock sync.Mutex
	// map1 := make(map[int]string)
	// wg.Add(2)
	// go readL(map1, &wg, &lock)
	// go writeL(map1, &wg, &lock)
	// wg.Wait()

	// 方式二
	var map2 = sync.Map{}
	var wg sync.WaitGroup
	wg.Add(2)
	go read(&map2, &wg)
	go write(&map2, &wg)
	wg.Wait()
}