package lilith

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func formatList(head *ListNode) *ListNode {
	// write code here
	h := head
	IsHead := false

	ret := &ListNode{}
	ret.Next = h
	tail := h
	h = h.Next

	for h != nil {
		if IsHead {
			pre := h
			h = h.Next
			pre.Next = ret.Next
			ret.Next = pre
			IsHead = false
		} else {
			tail.Next = h
			tail = tail.Next
			h = h.Next
			tail.Next = nil
			IsHead = true
		}
	}

	return ret.Next
}

func print_listnode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	return
}

func ans(array []int, k int) int64 {
	// write code here
	//pre_map := make(map[int]int64)

	head := &RetListNode{}
	// k - v
	ret := int64(0)
	for _, v := range array {
		re := head.AddNode(v, k)
		ret += re
	}

	return ret
}

type RetListNode struct {
	Val int
	//Pre *RetListNode
	Next *RetListNode
	Num  int64
}

func (self *RetListNode) AddNode(val int, k int) int64 {
	pre := self
	next := self.Next
	ret := int64(0)
	remain := k - val
	for next != nil {
		if next.Val < val {
			//next.Num++
			ret += next.Num
			pre = next
			next = next.Next
		}
		if next.Val == val {
			next.Num++
			ret += next.Num
			return ret
		}
		if next.Val > val {
			break
		}
	}
	node := &RetListNode{Val: val, Num: 1}
	pre.Next = node
	node.Next = next
	ret += 1
	return ret
}

func minimum(a []int) int64 {
	// write code here
	n := len(a)
	pre_sum := make([]int64, n)
	post_sum := make([]int64, n)
	sum := int64(0)
	ret := int64(math.MaxInt64)
	for i, v := range a {
		sum += int64(v)
		pre_sum[i] = sum
	}
	sum = int64(0)
	for i := n - 1; i >= 0; i-- {
		sum += int64(a[i])
		post_sum[i] = sum
	}
	//fmt.Println(pre_sum)
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {

			left := pre_sum[i]
			right := int64(0)
			if j < n-1 {
				right = post_sum[j+1]
			}
			mid := pre_sum[j] - pre_sum[i]
			lr := left + right

			ab := abs(mid, lr)
			if ab < ret {
				ret = ab
			}
		}
	}

	return ret
}

func abs(a, b int64) int64 {
	if a > b {
		return a - b
	}
	if a < b {
		return b - a
	}
	return 0
}
