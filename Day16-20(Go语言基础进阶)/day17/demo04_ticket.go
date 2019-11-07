/*
@Time : 2019/11/7 0007 上午 11:22
@Author : huanfuan
@File : demo04_ticket
@Software: GoLand
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量，表示票
var ticket = 10 //100张票
func main() {
	/*
		4个goroutine，模拟4个售票口，
	*/
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(5 * time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().Unix())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
	}
}
