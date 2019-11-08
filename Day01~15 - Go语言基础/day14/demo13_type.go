/*
@Time : 2019/11/7 0007 下午 5:23
@Author : huanfuan
@File : demo13_type
@Software: GoLand
*/

package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person--->", p.name)
}

//类型别名
type People = Person

func (p People) show2() {
	fmt.Println("People--->", p.name)
}

type Student struct {
	People
	Person
}

func main() {
	var s Student
	s.Person.name = "王二狗"
	s.Person.show()

	fmt.Printf("%T,%T\n", s.Person, s.People)

	s.People.name = "李小花"
	s.People.show()

	s.People.show()
}
