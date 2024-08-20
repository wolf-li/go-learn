# 常用标准模块

## time 模块
time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。示例代码如下：
```go
    now := time.Now() //获取当前时间
    fmt.Printf("current time:%v\n", now)

    year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
```
时间戳
时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
---
https://www.topgoer.com/%E9%A1%B9%E7%9B%AE/github%E5%BA%93%E5%9C%B0%E5%9D%80.html

## flag 模块
flag 包实现了命令行参数的解析，开发 cli（command line interface）利器

## strconv 模块
strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itoa()、parse系列、format系列、append系列。

## context 模块
context 在 1.7 中引入
规范
conntext 接口
```go
type Context interface{
    Deadline()(deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any)any
}
```
context 用处
取消 channl
超时 deadline
上下文值 value

数据传递
```go
func GetUser(ctx context.Context){
	fmt.Println(ctx.Value("name"))
}

func main() {
	ctx := context.Background() // root context 
	ctx = context.WithValue(ctx, "name", "lala")
	GetUser(ctx)
}
```

通过 type 定义不同类型获取传递值
```go
package main

import (
	"context"
	"fmt"
)

const RequestID = "requestID"

func main() {
	ctx := context.Background()
	SendContext(ctx)
}

type CtxSendKey string

func SendContext(ctx context.Context) {
	key := CtxSendKey(RequestID)
	ctx = context.WithValue(ctx, key, "123")
	ReciverContext(ctx)
}

type CtxReciverKey string

func ReciverContext(ctx context.Context) {
	key := CtxReciverKey(RequestID)
	ctx = context.WithValue(ctx, key, "3434")
	LoggerContext(ctx)
}

func LoggerContext(ctx context.Context) {
	fmt.Println("Send", CtxSendKey(RequestID), ctx.Value(CtxSendKey(RequestID)))
	fmt.Println("Receiver", CtxReciverKey(RequestID), ctx.Value(CtxReciverKey(RequestID)))
}
```


---
https://www.liwenzhou.com/posts/Go/flag/

