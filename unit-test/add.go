package main

import (
	"fmt"
)

func Add[T int|float32|float64](n1, n2 T) T{
	return n1 + n2
}

func main() {
	fmt.Println(Add(1,2))
	fmt.Println(Add(1.1,4.3))
}