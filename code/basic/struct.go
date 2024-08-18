package main
import "fmt"

type Student struct{
	Age int
	Name string
}

// 结构体绑定一个方法
func (s Student) PrintInfo(){
	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}

func main(){
	s := Student{
		Age: 12,
		Name: "lala",
	}
	s.Name = "xixi"
	s.PrintInfo()
}