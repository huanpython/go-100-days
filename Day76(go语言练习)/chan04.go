/*
@Time : 2019/11/29 0029 上午 11:15
@Author : huanfuan
@File : chan04
@Software: GoLand
*/

package main

/*

为什么需要管道？

（1）主线程在等待所有协程全部完成的时间很难确定；

（2）如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能协程还处于工作状态，这时也会随着主协程的结束而销毁；

（3）通过全局变量加锁同步来实现通讯，也并不利于多个协程对全局变量的读写操作；
*/

/*
管道的介绍：

（1）管道的本质就是一种数据结构--队列；

（2）数据先进先出；

（3）线程安全，多协程访问时，不需要加锁；

（4）管道只能存储相同的数据类型；
*/

import (
	"fmt"
)

var (
	myMap = make(map[int]int, 10)
)

func cal(n int) map[int]int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
	return myMap
}

func write(myChan chan map[int]int) {
	for i := 0; i <= 15; i++ {
		myChan <- cal(i)
		fmt.Println("writer data：", cal(i))
	}
	close(myChan)
}

func read(myChan chan map[int]int, exitChan chan bool) {
	for {
		v, ok := <-myChan
		if !ok {
			break
		}
		fmt.Println("read data：", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	var myChan chan map[int]int
	myChan = make(chan map[int]int, 20)
	var exitChan chan bool
	exitChan = make(chan bool, 1)
	go write(myChan)
	go read(myChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
