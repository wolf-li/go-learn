package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChann = make(chan int)
var nameChann = make(chan string)
var doneChan = make(chan struct{})

func send(m int, n string, wait *sync.WaitGroup){
	defer wait.Done()
	fmt.Println("{} 开始购物", n)
	time.Sleep(time.Millisecond)
	fmt.Println("{} 结束购物", n)

	moneyChann <- m
	nameChann <- n
}

func main(){
	var wg = &sync.WaitGroup{}
	now := time.Now()
	wg.Add(3)
	go send(23, "lala", wg)
	go send(3, "xixi", wg)
	go send(5, "hehe", wg)

	go func(){
		defer close(moneyChann)
		defer close(nameChann)
		defer close(doneChan)
		wg.Wait()
	}()

	var moneyList []int
	var nameList []string
	
	var event = func(){
		for {
			select{
			case money := <- moneyChann:
				moneyList = append(moneyList, money)
			case name := <- nameChann:
				nameList = append(nameList, name)
			case <- doneChan:
				return
			}
		}
	}

	event()
	fmt.Println(nameList, moneyList)
	fmt.Println(time.Since(now))
}
