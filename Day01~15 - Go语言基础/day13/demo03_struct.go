/*
@Time : 2019/11/7 0007 下午 5:22
@Author : huanfuan
@File : demo03_struct
@Software: GoLand
*/

package main

import "fmt"

/*
	匿名结构体和匿名字段：

	匿名结构体：没有名字的结构体，
		在创建匿名结构体时，同时创建对象
		变量名 := struct{
			定义字段Field
		}{
			字段进行赋值
		}

	匿名字段：一个结构体的字段没有字段名


	匿名函数：

*/
func main() {
	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("hello world...")
	}()

	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  18,
	}
	fmt.Println(s2.name, s2.age)

	w2 := Worker{"李小华", 32}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

type Worker struct {
	//name string
	//age int
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字，那么匿名字段的类型就不能重复，否则会冲突
	//string
}

type Student struct {
	name string
	age  int
}
