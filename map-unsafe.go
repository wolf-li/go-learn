package main

import (
	"fmt"
	"sync"
)


func read(m map[int]string, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		fmt.Println(m[1])
	}
}

func write(m map[int]string, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		m[1] = "time"
	}
}

func main(){
	map1 := make(map[int]string)
	var wg sync.WaitGroup
	wg.Add(2)
	go read(map1, &wg)
	go write(map1, &wg)
	wg.Wait()
}