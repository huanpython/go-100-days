/*
@Time : 2019/11/29 0029 上午 10:59
@Author : huanfuan
@File : chan02
@Software: GoLand
*/

//链接 https://www.tuicool.com/articles/YrM7fu3

package main

import (
	"fmt"
	"sync"
)

var (
	myMapp = make(map[int]int, 10)
	//lock是全局互斥锁,synchornized
	lock sync.Mutex
)

func calp(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMapp[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 15; i++ {
		go calp(i)
	}

	for i, v := range myMapp {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
