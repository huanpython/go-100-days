/*
@Time : 2019/11/6 0006 下午 2:46
@Author : huanfuan
@File : const1
@Software: GoLand
*/
package main

import "fmt"

func main() {
	fmt.Println(100)
	fmt.Println("hello")

	const path string = "huanfuan@163.com"
	const PI = 3.14
	fmt.Println(path)

	const c1, c2, c3 = 100, 3.14, "haha"
	const (
		a int = 100
		b
		c string = "ruby"
		d
		e
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)
	fmt.Printf("%T,%s\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)

	//5. 枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
