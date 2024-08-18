package main

import (
	// "time"
	"fmt"
)

func main(){
	// for
	// sum := 0
	// for i := 0; i <= 100; i++{
	// 	sum += i
	// }
	// fmt.Println(sum)

	// 死循环
	// for {
	// 	time.Sleep( 1 * time.Second)
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// }

	// while
	// i, sum := 0, 0
	// for i <= 100 {
	// 	sum += i 
	// 	i++
	// }
	// fmt.Println(sum)

	// do-while
	// i, sum := 0, 0
	// for {
	// 	sum += i
	// 	i++
	// 	if i == 101{
	// 		break
	// 	}
	// }
	// fmt.Println(sum)

	// for 切片
	// s := []string{"lala","xix","nihao"}
	// for index, v := range s{
	// 	fmt.Println(index, v)
	// }
	// s := []string{"lala","xix","nihao"}
	// for i := 0; i < len(s); i++{
	// 	fmt.Println(i, s[i])
	// }

	// for map
	// dict := map[string]string{
	// 	"name":"lala",
	// 	"age":"12",
	// 	"year": "2024",
	// }
	// for k,v := range dict{
	// 	fmt.Println(k, v)
	// }

	// 9 * 9
	for i:=1; i <= 9; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
