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



---
https://www.liwenzhou.com/posts/Go/flag/

