/*
@Time : 2019/11/27 0027 下午 4:55
@Author : huanfuan
@File : Goroutine01
@Software: GoLand
*/

//

package main

import "fmt"

/*
Goroutine 的调用时立刻返回的，并执行下一行代码。而且 Goroutine 的返回值将会被忽略
主Goroutine正在运行，其他任何Goroutine才能运行。如果主Goroutine终止，则程序将终止，并且其他Goroutine将不会运行。
*/

func hello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go hello() // 开启一个新的 goroutine
	fmt.Println("main function")
}
