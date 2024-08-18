package main

import (
	"time"
	"fmt"
	"sync"
	"strconv"
)

var MoneyChann = make(chan int) // 初始化并定义一个长度为 0 的信道

func pay(m int, wait *sync.WaitGroup){
	defer wait.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("() 消费金额", strconv.Itoa(m))
	MoneyChann <- m
}

func main(){
	var wg = sync.WaitGroup{}
	now := time.Now()
	wg.Add(3)
	go pay(23, &wg)
	go pay(53, &wg)
	go pay(3, &wg)

	go func(){
		defer close(MoneyChann)
		wg.Wait()
	}()

	// 简写
	for item := range MoneyChann{
		fmt.Println(item)
	}
	
	// for {
	// 	money, ok := <- MoneyChann
	// 	fmt.Println(money, ok)
	// 	if !ok{
	// 		break
	// 	}
	// }
	
	fmt.Println(time.Since(now))
}