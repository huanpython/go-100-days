/*
@Time : 2019/11/27 0027 下午 4:45
@Author : huanfuan
@File : demo15_panic
@Software: GoLand
*/

//  https://www.tuicool.com/articles/2ARvMbZ

package main

func main() {
	defer println("defer in main")
	go func() {
		defer println("defer in goroutine")
		panic("panic in goroutine")
	}()
}
