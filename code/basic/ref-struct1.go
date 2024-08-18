package main

import (
  "fmt"
  "reflect"
//   "strings"
)

type Student struct {
  Name1 string `big:"-"`
  Name2 string
}

func main() {
  s := Student{
    Name1: "fengfeng",
    Name2: "zhangsan",
  }
  t := reflect.TypeOf(s)
  v := reflect.ValueOf(&s).Elem()

  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    bigField := field.Tag.Get("big")
    // 判断类型是不是字符串
    if field.Type.Kind() != reflect.String {
      continue
    }
    if bigField == "" {
      continue
    }
    // 修改值
    valueFiled := v.Field(i)
    valueFiled.SetString("xx")
  }
  fmt.Println(s)
}
