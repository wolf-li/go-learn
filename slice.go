package main

import (
	"fmt"
	"sort"
)

func main(){
	var list []string

	fmt.Println(len(list))
	list = append(list, "你好")
	list = append(list, "我试试额")
	fmt.Println(list)
	fmt.Println(len(list))
	// 修改某个元素
	list[1] = "bushi"
	fmt.Println(list)


	var ints = []int{234,3361,23,5426,47346,1234}
	sort.Ints(ints)
	fmt.Println(ints) // 升序
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // 降序
	fmt.Println(ints) // 降序
}