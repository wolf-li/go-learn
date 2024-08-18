package main
import "fmt"

type Student struct{
	Name string
	Age int
	Male string
}

func setAge(s Student) {
	s.Age = 12
}

func setAge1(s *Student) {
	s.Age = 12
}

func main(){
	s1 := Student{
		Name:"lala",
		Age: 23,
		Male:"F",
	}
	fmt.Println(s1)
	setAge(s1)
	fmt.Println(s1)
	setAge1(&s1)
	fmt.Println(s1)
}