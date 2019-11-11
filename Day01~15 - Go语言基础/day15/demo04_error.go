/*
@Time : 2019/11/7 0007 下午 5:24
@Author : huanfuan
@File : demo04_error
@Software: GoLand
*/

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")
	//if err != nil && err == filepath.ErrBadPattern{
	//	fmt.Println(err) //syntax error in pattern
	//	return
	//}
	fmt.Println("files:", files)
}
