/*
@Time : 2019/11/7 0007 下午 5:23
@Author : huanfuan
@File : demo12_type
@Software: GoLand
*/

package main

import "time"

func main() {

}

type MyDuration = time.Duration

func (m MyDuration) SimpleSet() { //cannot define new methods on non-local type time.Duration

}
