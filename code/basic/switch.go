package main

import "fmt"

func main(){
	// var age int
	// fmt.Println("input your age:")
	// fmt.Scan(&age)

	// switch {
	// case age < 0:
	// 	fmt.Println("未出生")
	// case age <=18:
	// 	fmt.Println("未成年")
	// case age <= 35:
	// 	fmt.Println("青年")
	// default:
	// 	fmt.Println("中年")
	// }

	var week int
	fmt.Println("input week number:")
	fmt.Scan(&week)

	switch week{
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("错误")
	}
}