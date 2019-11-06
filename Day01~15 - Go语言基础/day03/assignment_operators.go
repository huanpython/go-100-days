/*
@Time : 2019/11/6 0006 下午 4:29
@Author : huanfuan
@File : assignment_operators
@Software: GoLand
*/

package main

import "fmt"

func main() {
	/*
		赋值运算符：
			=,+=,-=,*=,/=,%=,<<=,>>=,&=,|=,^=...
			=，把=右侧的数值，赋值给=左边的变量

			+=, a += b,相当于a = a + b
				a++, a += 1
	*/

	var a int
	a = 3
	fmt.Println(a)

	a += 4 // a = a + 4
	fmt.Println(a)

	a -= 3
	fmt.Println(a)

	a *= 2
	fmt.Println(a)

	a /= 3
	fmt.Println(a)

	a %= 1
	fmt.Println(a)
}
