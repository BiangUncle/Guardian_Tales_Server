package lilith

import (
	"fmt"
	"testing"
)

func Test_formatList(t *testing.T) {
	a := &ListNode{Val: 2}
	b := &ListNode{Val: 1}
	c := &ListNode{Val: 2}
	d := &ListNode{Val: 1}
	//e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	//d.Next = e
	re := formatList(a)
	print_listnode(re)
}

func Test_minimum(t *testing.T) {
	a := []int{7, 5, 3, 6, 4, 2, 1}
	re := minimum(a)
	fmt.Println(re)
}

func Test_ans(t *testing.T) {
	arr := []int{3, 1, 1, 2}
	k := 5
	re := ans(arr, k)
	fmt.Println(re)
}
