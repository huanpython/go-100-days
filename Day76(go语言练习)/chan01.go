/*
@Time : 2019/11/29 0029 上午 10:45
@Author : huanfuan
@File : chan01
@Software: GoLand
*/

//链接 https://www.tuicool.com/articles/YrM7fu3

package main

import (
	"fmt"
)

var (
	myMap1 = make(map[int]int, 10)
)

func cal4(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap1[n] = res
}

func main() {
	for i := 1; i <= 15; i++ {
		go cal4(i)
	}
	for i, v := range myMap1 {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
