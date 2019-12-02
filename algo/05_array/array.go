/*
@Time : 2019/12/2 0002 上午 10:53
@Author : huanfuan
@File : array
@Software: GoLand
*/
package _5_array

import (
	"errors"
	"fmt"
)

/*
	1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
*/

type Array struct {
	data   []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (this *Array) len() uint {
	return this.length
}

//判断索引是否越界

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

//插入数值到索引index上
func (this *Array) Insert(index uint, v int) error {
	if this.len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.len(), v)
}

//删除索引index上的值
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

//打印数列
func (this *Array) print() {
	var format string
	for i := uint(0); i < this.len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Printf(format)
}
