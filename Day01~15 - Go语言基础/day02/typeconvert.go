/*
@Time : 2019/11/6 0006 下午 3:36
@Author : huanfuan
@File : typeconvert
@Software: GoLand
*/

package main

import "fmt"

func main() {

	var a int8
	a = 10

	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 4.83
	var c int
	c = int(f1)
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)

	sum := f1 + 100
	fmt.Printf("%T,%f\n", sum, sum)

}
