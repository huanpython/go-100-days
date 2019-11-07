/*
@Time : 2019/11/7 0007 上午 10:48
@Author : huanfuan
@File : demo03_race
@Software: GoLand
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		临界资源：
	*/
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine中。。", a)
	}()

	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine...", a)
}
