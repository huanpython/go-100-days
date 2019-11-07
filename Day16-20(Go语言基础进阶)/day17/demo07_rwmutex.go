/*
@Time : 2019/11/7 0007 上午 11:22
@Author : huanfuan
@File : demo07_rwmutex
@Software: GoLand
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main..over...")

}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写：write start。。")
	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "正在写：writing。。。。")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读：read start。。")
	rwMutex.RLock() //读操作上锁
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束：read over。。。")
}
