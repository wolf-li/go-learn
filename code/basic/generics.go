package main

import (
	"fmt"
)

func add[T int|float32|float64](a,b T)T{
	return a + b
}

// 类型约束
type AddType interface{
	 int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float32|float64
}

func add1[T AddType](a, b T) T{
	return a + b
}

func main(){
	fmt.Println(add1(1,2))
	fmt.Println(add1(1.3,2.4))
}