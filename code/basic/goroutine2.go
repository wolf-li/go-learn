package main
import(
	"fmt"
	"time"
	"sync"
)

func shopping(name string){
	fmt.Println("() 开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() 结束购物", name)
}

func shoppingWait(name string){
	defer wait.Done()
	fmt.Println("() shoppingWait开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() shoppingWait结束购物", name)
}
var wait = sync.WaitGroup{}

func shoppingWaitP(name string, wait *sync.WaitGroup){
	defer wait.Done()
	fmt.Println("() shoppingWaitP开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() shoppingWaitP结束购物", name)
}

func main(){
	// 未使用 go 
	now := time.Now()
	shopping("1")
	shopping("2")
	shopping("3")
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// 使用 go, 主线程结束，go 协程未开始
	now = time.Now()
	go shopping("1")
	go shopping("2")
	go shopping("3")
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// waitGroup 等待 goroutine 执行完毕后在执行主线程
	now = time.Now()
	wait.Add(3)
	go shoppingWait("1")
	go shoppingWait("2")
	go shoppingWait("3")
	wait.Wait()
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// waitGroup 作为参数进行函数传递
	now = time.Now()
	var wait1 = sync.WaitGroup{}
	wait1.Add(3)
	go shoppingWaitP("1", &wait1)
	go shoppingWaitP("2", &wait1)
	go shoppingWaitP("3", &wait1)
	wait1.Wait()
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
}