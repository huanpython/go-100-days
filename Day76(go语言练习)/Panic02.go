/*
@Time : 2019/11/29 0029 上午 10:36
@Author : huanfuan
@File : Panic02
@Software: GoLand
*/

//链接 https://www.tuicool.com/articles/6bEFJbB

package main

import (
	"fmt"
	"time"
)

func handlepanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

func myfun1() {
	defer handlepanic()
	fmt.Println("Welcome to Function 1")
	go myfun2()
	time.Sleep(10 * time.Second)
}

func myfun2() {
	fmt.Println("Welcome to Function 2")
	panic("Panicked!!")
}

func main() {
	myfun1()
	fmt.Println("Return successfully from the main function")
}
