package main

import(
	"fmt"
	// "time"
)

func main(){
	messageCh := make(chan int, 10)
	for i:=0; i< 10;i++{
		messageCh <- i
	}
	close(messageCh)
	for item := range messageCh{
		// time.Sleep(time.Second)
		fmt.Println(item)
	}
}