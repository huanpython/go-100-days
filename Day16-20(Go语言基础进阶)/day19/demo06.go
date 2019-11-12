/*
@Time : 2019/11/7 0007 上午 11:27
@Author : huanfuan
@File : demo06
@Software: GoLand
*/

package main

import (
	"fmt"
	"reflect"
)

/*
	通过反射来进行方法的调用
	思路：
	step1：接口变量-->对象反射对象：Value
	step2：获取对应的方法对象：MethodByName()
	step3：将方法对象进行调用：Call()
*/

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p.Name, p.Age, p.Sex)
}

func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func main() {
	p1 := Person{"啾咪", 20, "女"}
	value := reflect.ValueOf(p1)

	fmt.Printf("kind : %s, type:%s\n", value.Kind(), value.Type()) //kind : struct, type:main.Person

	methodValue1 := value.MethodByName("PrintInfo")
	fmt.Printf("kind:%s,type:%s\n", methodValue1.Kind(), methodValue1.Type()) //kind:func,type:func()

	//没有参数，进行调用
	methodValue1.Call(nil) //没有参数，直接写nil

	args1 := make([]reflect.Value, 0) //或者创建一个空的切片也可以
	methodValue1.Call(args1)

	methodValue2 := value.MethodByName("Say")
	fmt.Printf("kind:%s, type:%s\n", methodValue2.Kind(), methodValue2.Type()) //kind:func, type:func(string)
	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := value.MethodByName("Test")
	fmt.Printf("kind:%s,type:%s\n", methodValue3.Kind(), methodValue3.Type()) //kind:func,type:func(int, int, string)
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("hello world")}
	methodValue3.Call(args3)

}
