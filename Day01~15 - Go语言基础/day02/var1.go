/*
@Time : 2019/11/6 0006 下午 3:40
@Author : huanfuan
@File : var1
@Software: GoLand
*/

package main

import "fmt"

func main() {

	//第一种：定义变量，然后进行赋值
	var num1 int
	num1 = 30

	fmt.Printf("num1的数值是：%d\n", num1)

	//写在第一行
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)

	//第二种：类型推断
	var name = "huanfuan"
	fmt.Printf("类型是：%T，数值是:%s\n", name, name)

	//第三种，简短定义，也叫简短声明
	sum := 100
	fmt.Println(sum)

	//多个变量同时定义

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		studentName = "啾咪"
		age         = 18
		sex         = "女"
	)
	fmt.Printf("学生姓名：%s，年龄：%d，性别：%s\n", studentName, age, sex)

}
