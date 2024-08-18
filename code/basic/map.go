package main

import (
	"fmt"
)

func main(){
	var userName = map[int]string{
		1: "zhangsan",
		2: "lisi",
		3: "wa",
	}
	fmt.Println(userName)
	fmt.Println(userName[1])
	userName[5] = "lala"
	fmt.Println(userName)

	var a1 = map[string]int {
		"age": 1,
	}
	age1 := a1["age1"]
	fmt.Println(age1)
	age2, ok := a1["age1"]
	fmt.Println(age2, ok)
}