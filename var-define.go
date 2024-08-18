package main

import (
	"fmt"
	"example.com/go-study/version"
)

// name4:= "123" // 语法问题全局变量需要使用 var 定义
var name4 = "name4"

func main(){
	// // 先声明在赋值
	// var name string
	// name = "adb"
	// fmt.Println(name)

	// // 声明并赋值
	// var name1 string = "abcd"
	// fmt.Println(name1)

	// // 省略类型
	// var name2 = "name2"
	// fmt.Println(name2)

	// // 简短声明
	// name3 := "name3"
	// fmt.Println(name3)

	// fmt.Println(name4)

	// 多个变量
	// var (
	// 	name string = "abc"
	// 	userid int = 1
	// )
	// var v1, v2 = "abc", 123
	// v3, v4 := "abdc", 1234
	// fmt.Println(name, userid)
	// fmt.Println(v1, v2)
	// fmt.Println(v3, v4)

	// 定义常量
	// const version string = "2.3.1"
	fmt.Println(version.Version)
}