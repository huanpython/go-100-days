/*
@Time : 2019/11/27 0027 下午 4:50
@Author : huanfuan
@File : panic_recover
@Software: GoLand
*/

package main

import "fmt"

func main() {
	defer println("in main")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
			panic("panic test")
		}()
	}()
}
