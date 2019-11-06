/*
@Time : 2019/11/6 0006 下午 3:32
@Author : huanfuan
@File : stringtype
@Software: GoLand
*/

package main

import "fmt"

func main() {

	//1.定义字符串
	var s1 string
	s1 = "huanfuan"
	fmt.Printf("%T,%s\n", s1, s1)

	s2 := "hello world"
	fmt.Printf("%T,%s\n", s2, s2)

	//2.区别：'A',"A"
	v1 := 'A'
	v2 := 'A'
	fmt.Printf("%T,%d\n", v1, v1)
	fmt.Printf("%T,%s\n", v2, v2)

	v3 := 'A'
	fmt.Printf("%T,%d,%c,%q\n", v3, v3, v3, v3)

	//3.转义字符
	fmt.Println("\"HelloWorld\"")
	fmt.Println("Hello\nWor\tld")

	fmt.Println(`He"lloWor"ld`)
	fmt.Println("Hello`Wor`ld")

}
