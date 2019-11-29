/*
@Time : 2019/11/29 0029 上午 11:02
@Author : huanfuan
@File : chan03
@Software: GoLand
*/
//链接 https://www.tuicool.com/articles/YrM7fu3

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap3 = make(map[int]int, 10)
	//lock是全局互斥锁,synchornized
	lock sync.Mutex
)

func call(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap3[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 15; i++ {
		go call(i)
	}

	time.Sleep(time.Second * 4)

	lock.Lock()
	for i, v := range myMap3 {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
