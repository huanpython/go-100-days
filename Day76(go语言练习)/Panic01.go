/*
@Time : 2019/11/29 0029 上午 10:29
@Author : huanfuan
@File : Panic01
@Software: GoLand
*/

//链接 https://www.tuicool.com/articles/6bEFJbB

package main

import (
	"fmt"
	"time"
)

func catchErr(num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[recover]", err, num)
		}
	}()
	fmt.Println("goroutine", num)
	panic("panic occurred ...")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main recover: ", err)
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("main goroutine", i)
		go catchErr(i)
		time.Sleep(time.Second * 1)
	}
start:
	goto start
}
