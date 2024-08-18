package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Class struct{
	Name string `orm:"name"`
	Class int `orm:"class"`
}

// 参数 "name = ? and age = ?", "lal", 23
func Find(obj any, query ...any)(sql string, err error){
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		err = errors.New("非结构体")
		return
	}

	// 拼接条件
	var where string
	if len(query) > 0 {
		q := query[0]
		if strings.Count(q.(string),"?")+1 != len(query){
			err = errors.New("参数个数不对")
			return
		}

		for _, a := range query[1:]{
			at := reflect.TypeOf(a)
			switch at.Kind() {
			case reflect.Int:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("%d",a.(int)),1)
			case reflect.String:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("%s",a.(string)),1)
			}
		}
		where += "where" + q.(string)
	}

	var column []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		f := field.Tag.Get("orm")
		column = append(column, f)
	}

	name := strings.ToLower(t.Name()) + "s"

	sql = fmt.Sprintf("select %s from %s %s", strings.Join(column, ","), name, where)
	return
}

func main() {
	sql, err := Find(Class{}, "name ? and class ?", "lala", 23)
	fmt.Println(sql, err)
}