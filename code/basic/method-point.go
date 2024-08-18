package main
import "fmt"

type Student struct{
	Name string
	Age int
	Male string
}

func (s Student)changeAge(age int){
	s.Age = age
}

func (s *Student)changeAge1(age int){
	s.Age = age
}

func main(){
	s1 := Student{
		Name:"lala",
		Age: 23,
		Male:"F",
	}
	fmt.Println(s1)
	s1.changeAge(12)
	fmt.Println(s1)
	s1.changeAge1(12)
	fmt.Println(s1)
}