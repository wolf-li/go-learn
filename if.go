package main

import (
	"fmt"
)

// <=0     未出生
// 1-18    未成年
// 18-35   青年
// >=35    中年

// func main(){
// 	// 中断式
// 	var age int
// 	fmt.Println("input your age:")
// 	fmt.Scan(&age)

// 	if age <= 0 {
// 		fmt.Println("未出生")
// 		return 
// 	}

// 	if age <= 18  {
// 		fmt.Println("未成年")
// 		return 
// 	}

// 	if age <= 35  {
// 		fmt.Println("青年")
// 		return 
// 	}

// 	fmt.Println("中年")

// }

// func main(){
	// var age int
	// fmt.Println("input your age:")
	// fmt.Scan(&age)

	// if age <= 18{
	// 	if age > 0{
	// 		fmt.Println("未成年")
	// 	}else{
	// 		fmt.Println("未出生")
	// 	}
	// }else{
	// 	if age <= 35{
	// 		fmt.Println("青年")
	// 	}else{
	// 		fmt.Println("中年")
	// 	}
	// }

// }

// 多条件
func main(){
	var age int
	fmt.Println("input your age:")
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}

	if age >= 0 && age <= 18{
		fmt.Println("青少年")
	}

	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}

	if age > 35 {
		fmt.Println("中年")
	}
}