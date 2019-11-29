/*
@Time : 2019/11/29 0029 上午 11:38
@Author : huanfuan
@File : RWMutex
@Software: GoLand
*/

package main

//连接 https://www.tuicool.com/articles/n6VJ7rJ

import (
	"fmt"
	"sync"
)

var (
	count2 int
	rwLock sync.RWMutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				rwLock.Lock()
				count2++
				rwLock.Unlock()
			}
			fmt.Println(count2)
		}()
	}

	fmt.Scanf("\n") //等待子线程全部结束
}
