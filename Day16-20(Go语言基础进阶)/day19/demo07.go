/*
@Time : 2019/11/7 0007 上午 11:27
@Author : huanfuan
@File : demo07
@Software: GoLand
*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//函数的反射
/*
	思路：函数也是看做接口变量类型
	step1：函数--->反射对象，Value
	step2：kind-->func
	step3：call()
*/
func main() {
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s, type :%s\n", value.Kind(), value.Type()) //kind:func, type :func()
	value2 := reflect.ValueOf(fun2)

	value3 := reflect.ValueOf(fun3)
	fmt.Printf("kind:%s,type:%s\n", value2.Kind(), value2.Type()) //kind:func,type:func(int, string)
	fmt.Printf("kind:%s,type:%s\n", value3.Kind(), value3.Type()) //kind:func,type:func(int, string) string

	//通过反射调用函数
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(1000), reflect.ValueOf("啾咪")})

	resultValue := value3.Call([]reflect.Value{reflect.ValueOf(2000), reflect.ValueOf("铁憨憨")})
	fmt.Printf("%T\n", resultValue)
	fmt.Println(len(resultValue))                                                 //1
	fmt.Printf("kind:%s,type:%s\n", resultValue[0].Kind(), resultValue[0].Type()) //kind:string,type:string

	s := resultValue[0].Interface().(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

}

func fun1() {
	fmt.Println("我是函数fun1(),无参的...")
}

func fun2(i int, s string) {
	fmt.Println("我是函数fun2(),有参的。。", i, s)
}

func fun3(i int, s string) string {
	fmt.Println("我是函数fun3()，有参的，也有返回值。。", i, s)
	return s + strconv.Itoa(i)
}
