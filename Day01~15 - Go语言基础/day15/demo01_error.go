/*
@Time : 2019/11/6 0006 下午 4:03
@Author : huanfuan
@File : demo01_error
@Software: GoLand
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")
}
