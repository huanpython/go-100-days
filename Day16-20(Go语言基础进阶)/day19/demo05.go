/*
@Time : 2019/11/7 0007 上午 11:27
@Author : huanfuan
@File : demo05
@Software: GoLand
*/

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {

	s1 := Student{"王二狗", 19, "千峰教育"}
	//通过反射，更改对象的数值：前提也是数据可以被更改
	fmt.Printf("%T\n", s1) //main.Student
	p1 := &s1
	fmt.Printf("%T\n", p1)
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name)

	//改变数值
	value := reflect.ValueOf(&s1)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("啾咪")
		f3 := newValue.FieldByName("School")
		f3.SetString("铁憨憨")
		fmt.Println(s1)
	}
}
