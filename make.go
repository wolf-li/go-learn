package main
import (
	"fmt"
)

func main(){
	var list = make([]int, 0)
	fmt.Println(list, len(list), cap(list))
	fmt.Println(list == nil)

	list1 := make([]int, 2)
	fmt.Println(list1, len(list1), cap(list1))
}