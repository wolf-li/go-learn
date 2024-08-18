package main

import "fmt"

func login(){
	fmt.Println("登录")
}

func userCenter(){
	fmt.Println("用户中心")
}

func logout(){
	fmt.Println("注销")
}

func main(){
	fmt.Printf(`1.登录
2.个人中心
3.注销`)
	fmt.Println("\ninput action:")
	var num int
	fmt.Scan(&num)
	var fummap = map[int]func(){
		1:login,
		2:userCenter,
		3:logout,
	}
	fummap[num]()
}