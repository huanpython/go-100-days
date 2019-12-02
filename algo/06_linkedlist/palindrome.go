/*
@Time : 2019/12/2 0002 下午 3:38
@Author : huanfuan
@File : palindrome
@Software: GoLand
*/
package _6_linkedlist

/*
思路1：开一个栈存放链表前半段
时间复杂度：O(N)
空间复杂度：O(N)
*/

func isPalindrome1(l *LinkedList) bool {
	llen := l.length
	if llen == 0 {
		return false
	}
	if llen == 1 {
		return true
	}

	s := make([]string, 0, llen/2)
	cur := l.head
	for i := uint(1); i <= llen; i++ {
		cur = cur.next
		if llen%2 != 0 && i == (llen/2+1) { //如果链表有奇数个节点，中间的直接忽略
			continue
		}
		if i <= llen/2 { //前一半入栈
			s = append(s, cur.GetValue().(string))
		} else { //后一半与前一半进行对比
			if s[llen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/

func isPalindrome2(l *LinkedList) bool {
	llen := l.length
	if llen == 0 {
		return false
	}
	if llen == 1 {
		return true
	}

	var isPalindrome = true
	step := llen / 2
	var pre *ListNode = nil
	cur := l.head.next
	next := l.head.next.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	mid := cur
	var left, right *ListNode = pre, nil
	if llen%2 != 0 {
		right = mid.GetNext()
	} else {
		right = mid
	}

	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre

	return isPalindrome
}
