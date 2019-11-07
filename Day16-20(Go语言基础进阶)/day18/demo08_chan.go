/*
@Time : 2019/11/7 0007 上午 10:01
@Author : huanfuan
@File : demo08_chan
@Software: GoLand
*/

package main

import "fmt"

func main() {
	/*
		channel,通道
	*/
	var a chan int
	fmt.Printf("%T,%v\n", a, a)

	if a == nil {
		fmt.Println("channel是nil的，不能使用，需要先创建通道。。")
		a = make(chan int)
		fmt.Println(a)
	}
	test1(a)
}

func test1(ch chan int) {
	fmt.Printf("%T,%v\n", ch, ch)
}
