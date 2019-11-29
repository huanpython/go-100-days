/*
@Time : 2019/11/29 0029 下午 2:11
@Author : huanfuan
@File : main
@Software: GoLand
*/

package main

import (
	_ "github.com/go-100-days/Day77/boot"
	_ "github.com/go-100-days/Day77/router"
	"github.com/gogf/gf/frame/g"
)

func main()  {
	g.Server().Run()
}