package main

import (
	"fmt"
)

func main(){
	var arr [3]int = [3]int{1,2,3}
	fmt.Println(arr)

	var arr1 = [3]int{3,4,5}
	fmt.Println(arr1)

	var arr2 = [...]int{5,6,7}
	fmt.Println(arr2)

	arr2[1] = 10
	fmt.Println(arr2)


}