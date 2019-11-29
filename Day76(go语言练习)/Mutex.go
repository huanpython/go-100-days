/*
@Time : 2019/11/29 0029 上午 11:28
@Author : huanfuan
@File : Mutex01
@Software: GoLand
*/

package main

//连接 https://www.tuicool.com/articles/n6VJ7rJ

/*
读锁不能阻塞读锁

读锁需要阻塞写锁，直到所有读锁都释放

写锁需要阻塞读锁，直到所有写锁都释放

写锁需要阻塞写锁
*/

import (
	"fmt"
	"sync"
)

var (
	count1 int
	lock   sync.Mutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				lock.Lock()
				count1++
				lock.Unlock()
			}
			fmt.Println(count1)
		}()
	}
	fmt.Scanf("\n") //等待子线程全部结束
}
