/*
@Time : 2019/11/7 0007 上午 11:25
@Author : huanfuan
@File : demo17_after
@Software: GoLand
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		 2. func After(d Duration) <-chan Time
				返回一个通道：chan，存储的是d时间间隔之后的当前时间

			相当于：return NewTimer(d).C
	*/
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch) //<-chan time.Time
	fmt.Println(time.Now())

	time2 := <-ch
	fmt.Println(time2)
}
