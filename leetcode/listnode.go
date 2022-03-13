package leetcode

import (
	"fmt"
	"server/leetcode/ns0"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func create_list_node(str string) *ListNode {
	str = str[1 : len(str)-1]
	num_str := strings.Split(str, ",")
	head := &ListNode{}
	ret := head
	for _, s := range num_str {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		node := &ListNode{Val: num}
		head.Next = node
		head = head.Next
	}
	return ret.Next
}

func test() {
	n := 1<<12 - 1
	str := "["
	for i := 0; i < n; i++ {
		str = str + fmt.Sprintf("%v", i) + ","
	}
	str = str[:len(str)-1] + "]"
	fmt.Println(str)
}

func sortedListToBST(head *ListNode) *ns0.TreeNode {

	if head == nil {
		return nil
	}

	root := iter(head)

	return root
}

func iter(head *ListNode) *ns0.TreeNode {

	if head == nil {
		return nil
	}

	mid := find_mid_node(head)
	root := &ns0.TreeNode{Val: mid.Val}
	if mid == head {
		return root
	}
	root.Right = iter(mid.Next)
	root.Left = iter(head)

	return root
}

// x - x - x - x - x - x

func find_mid_node(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	h := head
	t := head
	var pre *ListNode

	for t != nil && t.Next != nil {
		pre = h
		h = h.Next
		t = t.Next
		t = t.Next
	}

	if pre != nil {
		pre.Next = nil
	}
	return h
}

func buildTree(inorder []int, postorder []int) *ns0.TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	var build func(int, int) *ns0.TreeNode
	build = func(l int, r int) *ns0.TreeNode {
		if l > r {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &ns0.TreeNode{Val: val}

		aim := l
		for aim < r {
			if inorder[aim] != val {
				aim++
			}
			break
		}

		root.Right = build(aim+1, r)
		root.Left = build(l, aim-1)

		return root
	}

	return build(0, n-1)
}
