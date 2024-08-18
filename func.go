package main

import "fmt"

func hello(){
	fmt.Println("hello")
}

func fn1(n1 int, n2 int){
	fmt.Println(n1, n2)
}

// 同类型的参数简写
func fn2(n1, n2 float64){
	fmt.Println(n1, n2)
}

// 传递的是一个切片
func fn3(numlist ...int){
	fmt.Println(numlist)
}

// 无返回值
func returnNull(){
	
}

// 单返回值
func return1() int{
	return 1
}

// 多返回值
func return2() (int , float64){
	return 1, 23.4
}

// 命名返回值
func returnName(k int, v string) (key int, value string){
	key = k + 1
	value = value + "test"
	return
}

func main(){
	// fn1(1,23)
	// fn2(1,32)
	// fn3(1,2,3,4,5)
	// returnNull()
	// fmt.Println(return1())
	// fmt.Println(return2())
	// fmt.Println(returnName(1, "123"))

	// var add = func(a, b int) int {
	// 	return a+b
	// }
	// fmt.Println(add(1,2))

	
	fmt.Printf(`1.登录
2.个人中心
3.注销`)
	fmt.Println("\ninput action:")
	var num int
	fmt.Scan(&num)
	var funcmap = map[int]func(){
		1:func(){
			fmt.Println("登录")
		},
		2:func(){
			fmt.Println("个人中心")
		},
		3:func(){
			fmt.Println("注销")
		},
	}
	funcmap[num]()
}