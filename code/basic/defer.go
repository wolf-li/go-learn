package main
import "fmt"

func fu(){
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer3")
}

func main(){
	defer fmt.Println("defer1")
	fu()
	defer fmt.Println("defer4")
}