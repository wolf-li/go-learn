package main

import (
	"fmt"
)

type myMap[K string|int, V any] map[K]V

func main(){
	var dict = make(myMap[int,string])
	dict[1] = "lala"
	dict[2] = "lala"
	fmt.Println(dict)	
}