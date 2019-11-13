/*
@Time : 2019/11/13 0013 下午 2:57
@Author : huanfuan
@File : main
@Software: GoLand
*/

package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Run(iris.Addr(":8000"))

}
