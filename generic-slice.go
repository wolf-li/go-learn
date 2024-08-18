package main

import (
	"fmt"
)

type mySlice[T any] []T

func main(){
	var m1 mySlice[string]
	var m2 mySlice[int]
	m1 = append(m1,"123")
	m2 = append(m2, 1)
	fmt.Println(m1,m2)
}