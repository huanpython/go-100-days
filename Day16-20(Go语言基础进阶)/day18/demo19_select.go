/*
@Time : 2019/11/7 0007 上午 11:26
@Author : huanfuan
@File : demo19_select
@Software: GoLand
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	select {
	case <-ch1:
		fmt.Println("case1可以执行。。")
	case <-ch2:
		fmt.Println("case2可以执行。。")
	case <-time.After(3 * time.Second):
		fmt.Println("case3执行。。timeout。。。")

	}
}
