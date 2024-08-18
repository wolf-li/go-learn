module 项目创建
1. 创建项目目录  
`mkdir programName`
2. 初始化 module  
`go mod init programName`
3. 创建 main 入口文件
```bash
cat <<EOF> programName/main.go
pacakge main

import (
    "fmt"
)

func main(){
    fmt.Println("hello world")
}
EOF
```
4. 添加依赖
`go get `
5. 运行测试
`go build`

[参考链接](https://blog.csdn.net/taoruicheng1/article/details/104150896)
## 变量定义
```golang
	// 先声明在赋值
	var name string
	name = "adb"
	fmt.Println(name)

	// 声明并赋值
	var name1 string = "abcd"
	fmt.Println(name1)

	// 省略类型
	var name2 = "name2"
	fmt.Println(name2)

	// 简短声明
	name3 := "name3"
	fmt.Println(name3)
```
全局变量、局部变量
```golang
package main

import "fmt"

var userName = "枫枫知道" // 可以不使用,定义全局变量需要使用 var 不可以使用简短变量定义

func main() {
  // var 变量名 类型 = "变量值"
  var name = "枫枫"
  // 在函数体内定义的变量，必须要使用
  fmt.Println(name)
}

```
定义多个变量
```golang
	var (
		name string = "abc"
		userid int = 1
	)
	var v1, v2 = "abc", 123
	v3, v4 := "abdc", 1234
	fmt.Println(name, userid)
	fmt.Println(v1, v2)
	fmt.Println(v3, v4)
```
定义常量(定义时需要赋值，定义后不能修改)
```golang
	const version string = "2.3.1"
```
命名规范
首字母大写的变量，函数，方法，属性可以在包外使用

## 输入输出
常用的输出函数 fmt.Println() 函数
```go
	fmt.Println("hello","world")
	fmt.Println("你好")
```
格式化输出 fmt.Printf() 函数
```go
	fmt.Printf("%s\n","你好")  // 打印字符串
	fmt.Printf("%v %T\n", "你好","你好") // %T 打印类型，%v 按值的文本来输出
	fmt.Printf("%d\n", 123)  // 打印整数
    fmt.Printf("%t", true)   // 打印布尔值
	fmt.Printf("%.2f\n", 234.1234) // 打印浮点数
	fmt.Printf("%#v\n","")   // 输出 Go 语言语法格式的值
```
输入 fmt.Scan() 函数
```go
	fmt.Println("input your name")
	var name string
	fmt.Scan(&name)
    // 整数输入使用 err 判断是否输入正确

	fmt.Println("your name is: ",name)
```
## 基本数据类型
整形
浮点型
复数
布尔值
字符串
### 整形
```go
var n1 uint8 = 2 
var n2 uint16 = 2 
var n3 uint32 = 2 
var n4 uint64 = 2 
var n5 uint = 2 
var n6 int8 = 2 
var n7 int16 = 2 
var n8 int32 = 2 
var n9 int64 = 2 
var n10 int = 2 
```
要点
默认的数字定义类型是int类型
带个u就是无符号，只能存正整数
后面的数字就是2进制的位数
uint8还有一个别名 byte， 一个字节=8个bit位
int类型的大小取决于所使用的平台

### 浮点型
float32 float64
float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32
float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64、
没有显示说明默认为 float64

### 字符型
byte 单字节字符（ascii表）和 rune 多字节字符(UTF-8 编码的 Unicode 码点)
赋值要用单引号
```go
    var c1 = 'a'
	var c2 = 97
	fmt.Println(c1)  // 97
	fmt.Println(c2)  // 97

	fmt.Printf("%c %c\n", c1, c2)  // a a
	var r1 rune = '中'
	fmt.Printf("%c\n", r1)  // 中
```

### 字符串类型
赋值是双引号
```go
    var s string = "你是谁"
	fmt.Println(s)
```

### 转义字符
```go
fmt.Println("枫枫\t知道")              // 制表符
fmt.Println("枫枫\n知道")              // 回车
fmt.Println("\"枫枫\"知道")            // 双引号
fmt.Println("枫枫\r知道")              // 回到行首
fmt.Println("C:\\pprof\\main.exe") // 反斜杠
```

多行字符串 赋值时使用反引号
```go
	var s = `今天天气不错
你觉得怎样`
	fmt.Println(s)
```

### 布尔类型
true 和 false 两个值
布尔类型变量的默认值为false
Go 语言中不允许将整型强制转换为布尔型
布尔型无法参与数值运算，也无法与其他类型进行转换

### 零值问题
如果基本的变量类型只声明不赋值，变量的值为零值
```go
	var i int 
	var f float64
	var s string

	fmt.Printf("%#v\n", i) // 0
	fmt.Printf("%#v\n", f) // 0
	fmt.Printf("%#v\n", s) // ""
```

## 数组、切片、map
### 数组
一组相同类型固定长度的集合
```go
	var arr [3]int = [3]int{1,2,3}
	fmt.Println(arr)

	var arr1 = [3]int{3,4,5}
	fmt.Println(arr1)

    // 简写
	var arr2 = [...]int{5,6,7}
	fmt.Println(arr2)

    // 修改数组中的某个值
    arr2[1] = 10
	fmt.Println(arr2) // [5 10 7]
```

### 切片
数组长度固定，slice 切片声明后长度可变
```go
    var list []string

	fmt.Println(len(list))
    // 添加元素
	list = append(list, "你好")
	list = append(list, "我试试额")
	fmt.Println(list)
	fmt.Println(len(list))
	// 修改某个元素
	list[1] = "bushi"
	fmt.Println(list)
```
make 函数
可以通过make函数创建指定长度的切片
make([]type, length, capacity)
```go
	var list = make([]int, 0)
	fmt.Println(list, len(list), cap(list)) // [] 0 0
	fmt.Println(list == nil)  // false

	list1 := make([]int, 2)
	fmt.Println(list1, len(list1), cap(list1)) // [0 0] 2 2
```
切片排序
```go
var list = []int{4, 5, 3, 2, 7}
fmt.Println("排序前:", list)
sort.Ints(list)
fmt.Println("升序:", list)
sort.Sort(sort.Reverse(sort.IntSlice(list)))
fmt.Println("降序:", list)
```

### map
map 内置的数据结构（类似python 字典 key-value 键值对）
map 中的key 必须是相同类型，value 可以不同
```go
  // 声明
  var m1 map[string]string
  // 初始化1
  m1 = make(map[string]string)
  // 初始化2
  m1 = map[string]string{}
  // 设置值,添加值也是这个方法
  m1["name"] = "枫枫"
  fmt.Println(m1)
  // 取值
  fmt.Println(m1["name"])
  // 删除值
  delete(m1, "name")
  fmt.Println(m1)

  // map 取值
  	var a1 = map[string]int {
		"age": 1,
	}
	age1 := a1["age1"]
	fmt.Println(age1)  // 0
	age2, ok := a1["age1"]
	fmt.Println(age2, ok) // 0 false
```

## if 语句
多选一情况
中断式
```go
func main(){
	// 中断式
	var age int
	fmt.Println("input your age:")
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
		return 
	}

	if age <= 18  {
		fmt.Println("未成年")
		return 
	}

	if age <= 35  {
		fmt.Println("青年")
		return 
	}

	fmt.Println("中年")

}
```
嵌套式
```go
	var age int
	fmt.Println("input your age:")
	fmt.Scan(&age)

	if age <= 18{
		if age > 0{
			fmt.Println("未成年")
		}else{
			fmt.Println("未出生")
		}
	}else{
		if age <= 35{
			fmt.Println("青年")
		}else{
			fmt.Println("中年")
		}
	}
```
多条件
```go
    	var age int
	fmt.Println("input your age:")
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}

	if age >= 0 && age <= 18{
		fmt.Println("青少年")
	}

	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}

	if age > 35 {
		fmt.Println("中年")
	}
```
逻辑短路
&& 第一个为 false 后面不在执行
|| 第一个为 true 后面的不在执行

## switch
```go
	var age int
	fmt.Println("input your age:")
	fmt.Scan(&age)

	switch {
	case age < 0:
		fmt.Println("未出生")
	case age <=18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
```
枚举所有值
```go
	var week int
	fmt.Println("input week number:")
	fmt.Scan(&week)

	switch week{
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("错误")
	}
```

## for 循环
一般写法
```go
for 初始化;条件;操作{
}
```
for 循环五种变体
传统循环
```go
	sum := 0
	for i := 0; i <= 100; i++{
		sum += i
	}
	fmt.Println(sum)
```
死循环
```go
func main(){
	for {
		time.Sleep( 1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
}
```
while 模式
go 没有 while 使用 for 转换
```go
	i, sum := 0, 0
	for i <= 100 {
		sum += i 
		i++
	}
	fmt.Println(sum) // 5050
```
do-while 模式
先执行一次循环在判断
```go
	i, sum := 0, 0
	for {
		sum += i
		i++
		if i == 101{
			break
		}
	}
	fmt.Println(sum) // 5050
```
遍历切片,map
遍历切片
```go
	s := []string{"lala","xix","nihao"}
	// 写法一
	for i := 0; i < len(s); i++{
		fmt.Println(i, s[i])
	}
	// 写法二
	for index, v := range s{
		fmt.Println(index, v)
	}
	// 结果
	// 0 lala
	// 1 xix
	// 2 nihao
```
遍历 map
```go
	dict := map[string]string{
		"name":"lala",
		"age":"12",
		"year": "2024",
	}
	for k,v := range dict{
		fmt.Println(k, v)
	}
	// 结果
	// name lala
	// age 12
	// year 2024
```
break continue
break用于跳出当前循环
continue用于跳过本轮循环
打印九九乘法表
```go
	for i:=1; i <= 9; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
```
使用 continue
```go
	for i := 1; i <= 9; i++ {
    for j := 1; j <= 9; j++ {
      if j > i {
        // 去掉 列比行大的数据
        continue
      }
      fmt.Printf("%d * %d = %d\t", i, j, i*j)
    }
    fmt.Println()
  }
```

## 函数
函数是封装了特定功能的可重用的代码块
函数有输入有输出
函数定义,调用
```go
func hello(){
	fmt.Println("hello")
}

func main(){
	hello()
}
```
函数参数
```go
func fn1(n1 int, n2 int){
	fmt.Println(n1, n2)
}

// 同类型的参数简写
func fn2(n1, n2 float64){
	fmt.Println(n1, n2)
}

// 传递的是一个切片
func fn3(numlist ...int){
	fmt.Println(numlist)
}
```
函数返回值
```go
// 无返回值
func returnNull(){
	
}

// 单返回值
func return1() int{
	return 1
}

// 多返回值
func return2() (int , float64){
	return 1, 23.4
}

// 命名返回值
func returnName(k int, v string) (key int, value string){
	key = k + 1
	value = value + "test"
	return
}

func main(){
	fmt.Println(return1())  // 1 
	fmt.Println(return2())  // 1 23.4
	fmt.Println(returnName(1, "123")) // 2 test
}
```
匿名函数
函数内定义另一个函数使用
```go
	var add = func(a, b int) int {
		return a+b
	}
	fmt.Println(add(1,2))
```
高阶函数（函数的参数可以是函数，返回值也可以是函数）
根据用户输入不同执行不同的操作
```go
	fmt.Printf(`1.登录
2.个人中心
3.注销`)
	fmt.Println("\ninput action:")
	var num int
	fmt.Scan(&num)
	var funcmap = map[int]func(){
		1:func(){
			fmt.Println("登录")
		},
		2:func(){
			fmt.Println("个人中心")
		},
		3:func(){
			fmt.Println("注销")
		},
	}
	funcmap[num]()
```
提取整合
```go
func login(){
	fmt.Println("登录")
}

func userCenter(){
	fmt.Println("用户中心")
}

func logout(){
	fmt.Println("注销")
}

func main(){
	fmt.Printf(`1.登录
2.个人中心
3.注销`)
	fmt.Println("\ninput action:")
	var num int
	fmt.Scan(&num)
	var fummap = map[int]func(){
		1:login,
		2:userCenter,
		3:logout,
	}
	fummap[num]()
}
```
闭包
这些函数可以访问在其主体之外定义的变量。闭包是一个可以捕捉其周围环境状态的函数。这意味着函数可以访问不在其参数列表中或在其主体中定义的变量。闭包函数可以在外部函数返回后访问这些变量。
计数器
```go
func timer() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

func main(){
	counter := timer()
	fmt.Println(counter())
	fmt.Println(counter())
}
```
值传递和值引用
*解析地址  &取地址符号
传递给函数的值，实际上是复制没有对原始变量进行修改
```go
func add(a int){
	a += a
	fmt.Println(a)
}

func main(){
	var a int = 3
	add(a)  		// 6
	fmt.Println(a)  // 3
}
```
修改传递给函数的变量值
```go
func add1(a *int){
	fmt.Printf("func inside: %p\n",a)
	*a = *a + *a
	fmt.Println(*a)
}

func main(){
	var a int = 3
	fmt.Printf("func outside: %p\n",&a)
	add1(&a)
	fmt.Println(a)
}
```
## init 函数 defer 函数
init函数特性
不能被其他函数调用，执行在 main 函数之前，自动调用
init 函数不能作为参数传递
不能有传入参数和返回值
一个 go 文件可以有多个 init 函数，谁在前面执行谁
使用场景：
初始化加载
```go
func init(){
	fmt.Println("init")
}

func init1(){
	fmt.Println("init1")
}

func main(){
	fmt.Println("main")
}
//结果
init 
main
```
defer 函数
defer 用于注册延时调用
这些调用直到 return 前才执行，可以用来做资源清理
多个 defer 语句，按照先进后出方式执行，谁离 return 近谁被先执行
defer 中的变量，在defer声明时就决定了
```go
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
//结果
func
defer3
defer2
defer4
defer1
```

## 结构体
数据封装
结构体定义
```go
type 结构体名称 struct{
	名称 类型 // 成员或属性
}
```
定义结构体，定义结构体函数并使用
```go
type Student struct{
	Age int
	Name string
}

// 结构体绑定一个方法，值类型（对原始变量没有修改）
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
// 结果
name: xixi, age: 12
```
组合
组合是一种机制，可以将较简单的类型组合起来创建更复杂的类型。它是面向对象编程中的一个基本概念，可重用代码并构建更灵活和模块化的程序。
```go
type People struct{
	Name string
}

func (p People) printInfo(){
	fmt.Printf("name: %s\n", p.Name)
}

type Student struct{
	People
	Age int
}

// 方法
func (s Student) printInfo(){
	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}

func main(){
	p := People{
		Name:"lala",
	}
	p.printInfo()
	s1 := Student{
		People: People{Name:"qlql"},
		Age:12,
	}
	s1.printInfo()
}
// 结果
name: lala
name: qlql, age: 12
```
结构体指针
修改结构体内容需要使用结构体指针和指针方法
结构体指针
```go
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
// 结果
{lala 23 F}
{lala 23 F}
{lala 12 F}
```
指针结构体方法
```go
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
// 结果
{lala 23 F}
{lala 23 F}
{lala 12 F}
```
结构体 tag
json 命名格式小驼峰
```go
import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age,omitempty"` // 忽略空值
	Male string `json:"male"`
	Password string `json:"-"` // 不显示
}

func main(){
	user := Student{Name:"rase", Age:12, Male:"m"}
	byteDate, _ := json.Marshal(user)
	fmt.Println(string(byteDate))
}
// 没有tag 输出
{"Name":"rase","Age":12,"Male":"m"}
// 结果
{"name":"rase","age":12,"male":"m"}
```
## 自定义数据类型
自定义类型指的是使用 type 关键字定义的新类型，它可以是基本类型的别名，也可以是结构体、函数等组合而成的新类型。自定义类型可以帮助我们更好地抽象和封装数据，让代码更加易读、易懂、易维护
自定义类型
例子：客户端响应码
```go
type Code int

const (
	SuccessCode Code = 0
	InvalidCode Code = 10003
	FailCode Code = 11111
)

func (c Code)GetMsg()(string){
	switch c{
	case SuccessCode:
		return "成功"
	case InvalidCode:
		return "无效参数"
	case FailCode:
		return "失败操作"
	default:
		return "未知"
	}
}

func (c Code) CodeStatus()(Code, string){
	return c,c.GetMsg()
}

func main(){
	var c1 Code = 1
	fmt.Println(c1.CodeStatus())
}
// 结果
1 未知
```
类型别名
和自定义类型有区别
* 不能绑定方法
* 打印类型还是原始类型
* 和原始类型比较，不需要转换
```go
package main

import "fmt"

type AliasCode = int
type MyCode int

const (
  SuccessCode      MyCode    = 0
  SuccessAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
func (m AliasCode) MyAliasCodeMethod() {

}

func main() {
  // 类型别名，打印它的类型还是原始类型
  fmt.Printf("%T %T \n", SuccessCode, SuccessAliasCode) // main.MyCode int
  // 可以直接和原始类型比较
  var i int
  fmt.Println(SuccessAliasCode == i)
  fmt.Println(int(SuccessCode) == i) // 必须转换之后才能和原始类型比较
}

```
## 接口
接口是一组仅包含方法名、参数、返回值未实现具体方法的集合
接口本身不能绑定方法
接口是值类型，保存的是：值+原始类型
实现接口：一个类型实现了接口的所有方法
```go

type Animal interface{
	sing()
}

type Chicken struct{
	Name string
}

func (ck Chicken)sing(){
	fmt.Printf("%s 在唱歌\n",ck.Name)
}

type Cat struct{
	Name string
}

func (c Cat)sing(){
	fmt.Println("唱歌")
}

func (c Cat)jump(){
	fmt.Println("跳")
}

func (c Cat)rap(){
	fmt.Println("rap")
}

func sing(obj Animal){
	obj.sing()
}

func main(){
	chicken1 := Chicken{Name:"caixukun"}
	sing(chicken1)
	cat1 := Cat{Name:"mimi"}
	sing(cat1)
}
// 结果
caixukun 在唱歌
唱歌
```
类型断言
使用接口时，想要知道具体类型是什么可以使用接口断言
```go
func sing(obj Animal){
	switch obj.(type){
	case Chicken:
		fmt.Println("鸡")
	case Cat:
		fmt.Println("猫")
	}
	obj.sing()
}
```
或者断言某个类型
```go
func sing(obj Animal) {
  c, ok := obj.(Chicken) // 两个参数    断言之后的类型   是否是对应类型
  fmt.Println(c, ok)
  d := obj.(Cat) // 一个参数 就是断言之后的类型,注意，类型不对是要报错的 main.Animal is main.Chicken, not main.Cat
  fmt.Println(d)
}
```
空接口
空接口可以接受任何类型
any 等价于 空接口 interface{}


## 并发
golang 推荐使用 CSP并发模型
不同于传统的多线程通过共享内存来通信，CSP是以通信的方式来共享内存
Go 通过一些机制减少了开发代码的难度（并不能提高性能）
* Goroutine
* Chan 、Select
* WaitGroup、Mutex、Context
并发和并行：
* 并发多个任务是相互抢占资源
* 并行多个任务互不抢占资源
异步编程是让程序并发的一种手段
goroutine is a function 由 golang scheduler 调度
goroutine 并不是并行

goroutine 入门示例
主线程中开启一个 goroutine ,该协程每隔 1s 输出 hello world 10次
在主线程中每隔 1s 输出 hello golang  10次
要求 主线程和 goroutine 同时进行
```go
import (
	"time"
	"fmt"
	"strconv"
)

func test(){
	for i:= 0; i<10;i++{
		fmt.Println("test () hello world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	go test()

	for i:=0;i<10;i++{
		fmt.Println("main () hello golang", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
```
goroutine 和 waitGroup 结合使用
```go
package main
import(
	"fmt"
	"time"
	"sync"
)

func shopping(name string){
	fmt.Println("() 开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() 结束购物", name)
}

func shoppingWait(name string){
	defer wait.Done()
	fmt.Println("() shoppingWait开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() shoppingWait结束购物", name)
}
var wait = sync.WaitGroup{}

func shoppingWaitP(name string, wait *sync.WaitGroup){
	defer wait.Done()
	fmt.Println("() shoppingWaitP开始购物", name)
	time.Sleep(time.Millisecond)
	fmt.Println("() shoppingWaitP结束购物", name)
}

func main(){
	// 未使用 go 
	now := time.Now()
	shopping("1")
	shopping("2")
	shopping("3")
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// 使用 go, 主线程结束，go 协程未开始
	now = time.Now()
	go shopping("1")
	go shopping("2")
	go shopping("3")
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// waitGroup 等待 goroutine 执行完毕后在执行主线程
	now = time.Now()
	wait.Add(3)
	go shoppingWait("1")
	go shoppingWait("2")
	go shoppingWait("3")
	wait.Wait()
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
	// waitGroup 作为参数进行函数传递
	now = time.Now()
	var wait1 = sync.WaitGroup{}
	wait1.Add(3)
	go shoppingWaitP("1", &wait1)
	go shoppingWaitP("2", &wait1)
	go shoppingWaitP("3", &wait1)
	wait1.Wait()
	fmt.Println(time.Since(now))
	fmt.Println("------------------")
}
```
goroutine 使用匿名函数
```go
		go func(name string){
			fmt.Println("() 开始购物", name)
			time.Sleep(time.Second)
			fmt.Println("() 结束购物", name)
		}(strconv.Itoa(i))
```

### go 与 chan
go function介绍
* 一种受golang并发机制中的调度器控制的方法
go 的用处
	* 性能高
	* 利用多核
chan 的用处：goroutine 之间通信
chan 类型
	* 根据传入数据类型划分
	* unbuffer 与 buffer 划分
	* 只读 <-chan 、只写 chan<-、双向 chan
channel 示例代码
```go
import (
	"time"
	"fmt"
	"sync"
	"strconv"
)

var MoneyChann = make(chan int) // 初始化并定义一个长度为 0 的信道

func pay(m int, wait *sync.WaitGroup){
	defer wait.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("() 消费金额", strconv.Itoa(m))
	MoneyChann <- m
}

func main(){
	var wg = sync.WaitGroup{}
	now := time.Now()
	wg.Add(3)
	go pay(23, &wg)
	go pay(53, &wg)
	go pay(3, &wg)

	go func(){
		defer close(MoneyChann)
		wg.Wait()
	}()

	// 简写
	for item := range MoneyChann{
		fmt.Println(item)
	}
	
	// for {
	// 	money, ok := <- MoneyChann
	// 	fmt.Println(money, ok)
	// 	if !ok{
	// 		break
	// 	}
	// }
	
	fmt.Println(time.Since(now))
}
```
select 流程控制
接受多个 channel 的数据
```go
var moneyChann = make(chan int)
var nameChann = make(chan string)
var doneChan = make(chan struct{})

func send(m int, n string, wait *sync.WaitGroup){
	defer wait.Done()
	fmt.Println("{} 开始购物", n)
	time.Sleep(time.Millisecond)
	fmt.Println("{} 结束购物", n)

	moneyChann <- m
	nameChann <- n
}

func main(){
	var wg = &sync.WaitGroup{}
	now := time.Now()
	wg.Add(3)
	go send(23, "lala", wg)
	go send(3, "xixi", wg)
	go send(5, "hehe", wg)

	go func(){
		defer close(moneyChann)
		defer close(nameChann)
		defer close(doneChan)
		wg.Wait()
	}()

	var moneyList []int
	var nameList []string
	
	var event = func(){
		for {
			select{
			case money := <- moneyChann:
				moneyList = append(moneyList, money)
			case name := <- nameChann:
				nameList = append(nameList, name)
			case <- doneChan:
				return
			}
		}
	}

	event()
	fmt.Println(nameList, moneyList)
	fmt.Println(time.Since(now))
}
```
另一个写法
```go
package main

import (
	"fmt"
	"time"
)

type MessageServer struct {
	messageCh chan string
	quitCh    chan struct{}
}

func NewMessageServer() *MessageServer {
	return &MessageServer{
		quitCh:    make(chan struct{}),
		messageCh: make(chan string, 100),
	}
}

func sendMessage(s *MessageServer, message string, number int) {
	for i := 0; i < number; i++ {
		s.messageCh <- fmt.Sprintf("mes: %s: %d", message, i+1)
	}
}

func (s *MessageServer) handleMessage(message string) {
	fmt.Println("get message : ", message)
}

// loop 死循环 工作循环
func (s *MessageServer) work() {
messageloop:
	for {
		select {
		case message := <-s.messageCh:
			s.handleMessage(message)
		case <-s.quitCh:
			fmt.Println("quiting server")
			break messageloop
			// default:
			// fmt.Println("nothing")
		}
	}
	fmt.Println("server is down")
}

func (s *MessageServer) quit() {
	s.quitCh <- struct{}{}
}

func main() {
	server := NewMessageServer()
	go func() {
		time.Sleep(time.Second * 3)
		// server.quitCh <- struct{}{} // 0 byte
		server.quit()
	}()
	sendMessage(server, "house message", 10)
	server.work()
}
```
协程超时处理
```go
var done = make(chan struct{})

func event(){
	defer close(done)
	fmt.Println("start")
	time.Sleep(time.Millisecond * 3)
	fmt.Println("end")
}

func main(){
	go event()

	select{
	case <- done:
		fmt.Println("执行完毕")
	case <- time.After(time.Millisecond * 4):
		fmt.Println("超时")
	}
}
```

Mutex / Atomic 解决数据竞争问题
协程不安全产生数据竞争,最好使用 go test 进行测试 go run -race xx.go
```go
// 范围太小没有效果，要使用大数
func add(n *int, wait *sync.WaitGroup){
	defer wait.Done()
	for i:=0; i<100000; i++{
		*n += i
	}
}

func sub(n *int, wait *sync.WaitGroup){
	defer wait.Done()
	for i:=0; i<100000; i++{
		*n -= i
	}
}

func main(){
	var wg sync.WaitGroup
	var n int = 0
	wg.Add(2)
	go add(&n, &wg)
	go sub(&n, &wg)
	wg.Wait()
	fmt.Println(n)
}
// 结果
4678446137
```

Mutex 使用同步锁
```go
// 范围太小没有效果，要使用大数
func add(n *int, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	lock.Lock()
	for i:=0; i<100000; i++{
		*n += i
	}
	lock.Unlock()
}

func sub(n *int, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	lock.Lock()
	for i:=0; i<100000; i++{
		*n -= i
	}
	lock.Unlock()
}

func main(){
	var wg sync.WaitGroup
	var l sync.Mutex
	var n int = 0
	wg.Add(2)
	go add(&n, &wg, &l)
	go sub(&n, &wg, &l)
	wg.Wait()
	fmt.Println(n)
}
```
线程安全下的 map
如果在一个协程函数下，读写 map 会出现错误
fatal error: concurrent map read and map write
```go
func read(m map[int]string, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		fmt.Println(m[1])
	}
}

func write(m map[int]string, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		m[1] = "time"
	}
}

func main(){
	map1 := make(map[int]string)
	var wg sync.WaitGroup
	wg.Add(2)
	go read(map1, &wg)
	go write(map1, &wg)
	wg.Wait()
}
```
解决方式：
方式一 加锁
方式二 使用 sync.map
```go
// 方式一
func readL(m map[int]string, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	for{
		lock.Lock()
		fmt.Println(m[1])
		lock.Unlock()
	}
}

func writeL(m map[int]string, wait *sync.WaitGroup, lock *sync.Mutex){
	defer wait.Done()
	for{
		lock.Lock()
		m[1] = "time"
		lock.Unlock()
	}
}
func main(){
	var lock sync.Mutex
	map1 := make(map[int]string)
	wg.Add(2)
	go readL(map1, &wg, &lock)
	go writeL(map1, &wg, &lock)
	wg.Wait()
}
```
```go
// 方式二
func read(m *sync.Map, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		fmt.Println(m.Load(1))
	}
}

func write(m *sync.Map, wait *sync.WaitGroup){
	defer wait.Done()
	for{
		m.Store(1, time.Now().Format("15:04:05"))
	}
}
func main(){
	// 方式二
	var map2 = sync.Map{}
	var wg sync.WaitGroup
	wg.Add(2)
	go read(&map2, &wg)
	go write(&map2, &wg)
	wg.Wait()
}
```

## 异常处理
golang 中没有捕获机制，导致每调一个函数都要接一下这个函数的error
常见的异常处理方式
方式一：向上抛出异常
错误交给上一级进行处理
```go
func div(a, b int)(res int ,err error){
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a/b
	return
}

func server()(res int, err error){
	res, err = div(1, 3)
	if err != nil {
		return
	}

	// 正常逻辑
	res++
	
	return 
}

func main(){
	res, err := server()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
```

方式二：中断程序（panic）
遇到错误直接终止程序
一般用于初始化
```go
import (
	"fmt"
	"os"
)

func init(){
	_, err := os.ReadFile("xx")
	if err != nil {
		panic(err.Error())
	}
}

func main(){
	fmt.Println("main")
}
```

方式三：遇到异常继续执行
使用 defer 函数对 panic进行捕获
```go
import (
	"fmt"
	"runtime/debug"
)

func read(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()

	// 问题的代码
	var list = []int{1,2}
	fmt.Println(list[2])
}

func main(){
	read()
}
```

## 泛型
golang 1.18 开始支持泛型
不使用具体的数据类型，使用通用数据类型进行程序设计的方法，减少代码量
求和函数使用泛型
```go
// 类型约束
type AddType interface{
	 int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float32|float64
}

func add1[T AddType](a, b T) T{
	return a + b
}

// 多个泛型
func myPrint[T int, K string|int](n1 T,n2 K){

}

func main(){
	fmt.Println(add1(1,2))
	fmt.Println(add1(1.3,2.4))
}
```
泛型结构体
使用泛型做解析
```go
type User struct {
    Name string `json:"name"`
  }

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response[T any] struct {
  Code int    `json:"code"`
  Msg  string `json:"msg"`
  Data T      `json:"data"`
}

func main() {

// 执行下面的代码时 Response 结构体 Data 字段设置为 any，删除泛型
//   user := Response{
// 	Code:0,
// 	Msg:"ok",
// 	Data:new(User),
//   }
//   user.Data = User{Name:"123"}
//   byteData, _ := json.Marshal(user)
//   fmt.Println(string(byteData))
//   userInfo := Response{
//    Code: 0,
//    Msg:  "成功",
//    Data: UserInfo{
//      Name: "枫枫",
//      Age:  24,
//    },
//   }
//   byteData, _ = json.Marshal(userInfo)
//   fmt.Println(string(byteData))

  var userResponse Response[User]
  json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫"}}`), &userResponse)
  fmt.Println(userResponse.Data.Name)
  var userInfoResponse Response[UserInfo]
  json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫","age":24}}`), &userInfoResponse)
  fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}
```

泛型切片
```go
type mySlice[T any] []T

func main(){
	var m1 mySlice[string]
	var m2 mySlice[int]
	m1 = append(m1,"123")
	m2 = append(m2, 1)
	fmt.Println(m1,m2)
}
```
泛型 map
```go
type myMap[K string|int, V any] map[K]V

func main(){
	var dict = make(myMap[int,string])
	dict[1] = "lala"
	dict[2] = "lala"
	fmt.Println(dict)	
}
```

## 文件操作
文件读取
一次性读取
```go
import (
	"fmt"
	"os"
)

func main(){	
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0][2:])
		return
	}
	byteData, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteData))
}
// 核心
// func ReadFile(name string) ([]byte, error)  返回值为 []byte 切片
	byteData, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteData))
```
分片读取
```go
file, _ := os.Open()
defer file.Close()
for {
	buf := make([]byte, 1)
	_, err = file.Read(buf)
	if err == io.EOF {
		break
	}
	fmt.Printf("%s",buf)
}
```
按行读取
```go
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}
```

按指定分隔符读取
```go
	// 指定分隔符，按照单词读取
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
```

文件写入

os.OpenFile()
文件打开方式
// 如果文件不存在就创建
os.O_CREATE|os.O_WRONLY
// 追加写
os.O_APPEND|os.O_WRONLY
// 可读可写
os.O_RDWR

完整模式
const (
  O_RDONLY int = syscall.O_RDONLY // 只读
  O_WRONLY int = syscall.O_WRONLY // 只写
  O_RDWR   int = syscall.O_RDWR   // 读写
  
  O_APPEND int = syscall.O_APPEND // 追加
  O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
  O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
  O_SYNC   int = syscall.O_SYNC   // 同步打开
  O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件
)

文件的权限
主要用于linux系统，在windows下这个参数会被无视，代表文件的模式和权限位，

三个占位符

第一个是文件所有者所拥有的权限

第二个是文件所在组对其拥有的权限

第三个占位符是指其他人对文件拥有的权限

根据UNIX系统的权限模型，文件或目录的权限模式由三个数字表示，分别代表 所有者（Owner） 、组（Group） 和 其他用户(Other) 的权限。每个数字由三个比特位组成，分别代表读、写和执行权限。因此，对于一个mode参数值，它的每一个数字都是一个八进制数字，代表三个比特位的权限组合

一次性写
```go
	err := os.WriteFile("hello", []byte("Hello, Gophers!\n"), 0000)
	if err != nil {
		log.Fatal(err)
	}
```
使用 os.OpenFile函数
```go
	file, err := os.OpenFile(os.Args[1], os.O_CREATE| os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	file.Write([]byte("你好"))
```
文件复制
```go
// 核心函数
// io.Copy(dst Writer, src Reader) (written int64, err error)

func errHandler(err error){
	if err != nil {
		log.Fatal(err)
		return 
	}
}

func main(){
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s sourcefile destfile\n", os.Args[0][2:])
		return
	}
	
	read, err := os.Open(os.Args[1])
	errHandler(err)
	defer read.Close()

	write, err := os.Create(os.Args[2])
	errHandler(err)
	defer write.Close()

	n, err := io.Copy(write, read) 
	errHandler(err)
	
	fmt.Println(n)
}
```

目录操作
读取目录下内容
```go
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0][2:])
		return
	}
	dir, err := os.ReadDir(os.Args[1])
	ErrHandler(err)

	for _, entry := range dir{
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
```

## 单元测试
源文件 case.go
```go
package main

func hello()string{
	return "hello"
}

func main() {
	hello()
}
```
测试文件 case_test.go
```go
package main

import "testing"

func Test_hello(t *testing.T) {
	if ans := hello(); ans != "hello"{
		t.Errorf("return not hello")
	}
}
```
go test // 可以运行某个包下的所有测试用例
-v 参数会显示每个用例的测试结果
-run参数可以指定测试某个函数
单元测试框架提供的日志方法

方 法	|备 注|	测试结果
--|--|--
Log	打印日志，同时结束测试	PASS
Logf	格式化打印日志，同时结束测试	PASS
Error	打印错误日志，同时结束测试	FAIL
Errorf	格式化打印错误日志，同时结束测试	FAIL
Fatal	打印致命日志，同时结束测试	FAIL
Fatalf	格式化打印致命日志，同时结束测试	FAIL

### 子测试
给一个函数调用不同的测试用例，可以使用子测试
子测试里测试用例失败不会中断程序
```go
package main

import (
  "testing"
)

func TestAdd(t1 *testing.T) {
  t1.Run("add1", func(t *testing.T) {
    if ans := Add(1, 2); ans != 3 {
      // 如果不符合预期，那就是测试不通过
      t.Fatalf("1 + 2 expected be 3, but %d got", ans)
    }
  })
  t1.Run("add2", func(t *testing.T) {
    if ans := Add(-10, -20); ans != -30 {
      t.Fatalf("-10 + -20 expected be -30, but %d got", ans)
    }
  })

}
```
优雅子测试写法
```go
func TestAdd(t *testing.T) {
  cases := []struct {
    Name           string
    A, B, Expected int
  }{
    {"a1", 2, 3, 5},
    {"a2", 2, -3, -1},
    {"a3", 2, 0, 2},
  }

  for _, c := range cases {
    t.Run(c.Name, func(t *testing.T) {
      if ans := Add(c.A, c.B); ans != c.Expected {
        t.Fatalf("%d * %d expected %d, but %d got",
          c.A, c.B, c.Expected, ans)
      }
    })
  }
}
```
泛型子测试写法
```go
type args[T int|float32|float64] struct {
	n1 T
	n2 T
}

type testCase[T int|float32|float64] struct{
	name string
	args args[T]
	want T
}

func testAdd[T int|float32|float64](t *testing.T, testCases []testCase[T]) {  
	for _, tt := range testCases {  
		t.Run(tt.name, func(t *testing.T) {  
			if got := Add(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {  
				t.Errorf("Add() = %v, want %v", got, tt.want)  
			}  
		})  
	}  
}  
func TestAdd(t *testing.T) {

	testint := []testCase[int] {
		{
			name: "AddInt",  
			args: args[int]{n1: 2, n2: 3},  
			want: 5, 
		},
		{
			name: "AddInt",  
			args: args[int]{n1: 1000, n2: 3},  
			want: 1003, 
		},
	}
	testAdd(t, testint)
	

	testfloat32 := []testCase[float32] {
		{
			name: "AddInt",  
			args: args[float32]{n1: 2.0, n2: -0.2},  
			want: 1.8, 
		},
		{
			name: "AddInt",  
			args: args[float32]{n1: 10.81, n2: 3},  
			want: 13.81, 
		},
	}

	testAdd(t, testfloat32)

}
```
testMain函数实现测试流程的生命周期
```go
package main

import (
  "fmt"
  "os"
  "testing"
)

// 测试前执行
func setup() {
  fmt.Println("Before all tests")
}

// 测试后执行
func teardown() {
  fmt.Println("After all tests")
}

func Test1(t *testing.T) {
  fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
  fmt.Println("I'm test2")
}

// 必须叫这个名字  测试主入口
func TestMain(m *testing.M) {
  // 测试前执行
  setup()
  code := m.Run()
  // 测试后执行
  teardown()
  os.Exit(code)
}
```

## 反射
反射（Reflection）在编程中通常被定义为在运行时检查程序的能力。这种能力使得一个程序能够操纵像变量、数据结构、方法和类型这样的对象的各种属性和行为。这一机制在Go中主要通过reflect标准库实现。

应用场景
动态编程: 通过反射，你可以动态地创建对象，调用方法，甚至构建全新的类型。
框架与库开发: 很多流行的Go框架，如Gin、Beego等，都在内部使用反射来实现灵活和高度可定制的功能。
元编程: 你可以写出可以自我分析和自我修改的代码，这在配置管理、依赖注入等场景中尤为有用。

判断变量类型
```go
func getType(obj any){
	t := reflect.TypeOf(obj)
	fmt.Println(t, t.Kind())
	switch t.Kind() {
	case reflect.Float32:
		fmt.Println("float32")
	case reflect.Int:
		fmt.Println("int")
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.String:
		fmt.Println("string")
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	name := "lala"
	getType(name)
	getType(struct{Name string}{Name:"xix"})
}
```
通过反射获取值
```go
func refValue(obj any){
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind(){
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.String:
		fmt.Println(value.String())
	}
}

func main(){
	refValue("lala")
	refValue(3)
}
//结果
lala string
lala
3 int
3
```
通过反射修改值
如果需要通过反射修改值，必须要传指针，在反射中使用Elem取指针对应的值
```go
func refSetValue(obj any){
	objValueOf := reflect.ValueOf(obj)
	elem := objValueOf.Elem()
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("haha")
	}
}

name := "lala"
refSetValue(&name)
fmt.Println(name)
// 结果
haha
```
结构体反射
读取结构体 tag 没有就使用结构体属性
```go
type User struct{
	Age int
	Name string `json:name`
}

func main() {
	u1 := User{
		Name: "x",
		Age: 12,
	}
	t := reflect.TypeOf(u1)
	v := reflect.ValueOf(u1)

	// NumField() 结构体属性数量
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		jsonFiled := filed.Tag.Get("json")
		if jsonFiled == "" {
			jsonFiled = filed.Name
		}
		fmt.Printf("name: %s type: %s json: %s value: %v\n", filed.Name, filed.Type, jsonFiled, v.Field(i))
	}
}
//
name: Age type: int json: Age value: 12
name: Name type: string json: Name value: x
```

修改结构体中的某些值
```go

```

