package leetcode

import (
	"math/rand"
	"server/leetcode/ns0"
	"strconv"
	"time"
)

func create_random_list(nums_len int) []int {
	rand.Seed(time.Now().Unix())

	re := make([]int, nums_len)
	for i := 0; i < nums_len; i++ {
		randNum := rand.Intn(nums_len)
		re[i] = randNum
	}

	return re
}

func nums_list_2_str(nums []int) string {

	re := "["

	for _, num := range nums {
		re = re + strconv.Itoa(num) + ","
	}

	re = re[:len(re)-1] + "]"

	return re
}

func CreateTestExample(str_len int, query_len int) (string, [][]int) {

	rand.Seed(time.Now().Unix())
	s := ""

	for i := 0; i < str_len; i++ {
		randNum := rand.Intn(100)
		if randNum < 60 {
			s = s + "*"
		} else {
			s = s + "|"
		}
	}

	queries := make([][]int, query_len)

	for i := 0; i < query_len; i++ {
		left := -1
		right := -2
		for right < left {
			left = rand.Intn(str_len)
			right = rand.Intn(str_len)
		}
		queries[i] = []int{left, right}
	}

	return s, queries
}

//func postorderTraversal(root *TreeNode) []int {
//	res := make([]int, 0)
//	if root == nil {
//		return res
//	}
//	stack := make([]*TreeNode, 0)
//	var pre *TreeNode
//
//	for pre != nil || len(stack) > 0 {
//
//		node := stack[len(stack)-1]
//		root = node
//		stack = stack[:len(stack)-1]
//
//		if pre == nil || {
//			stack = append(stack, root)
//			root = root.Left
//		}
//
//		if node.Right != nil {
//			stack = append(stack, node)
//			root = node.Right
//			continue
//		}
//		res = append(res, node.Val)
//		root = node
//	}
//
//	return res
//}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func morrisPostorderTraversal(root *ns0.TreeNode) (res []int) {
	addPath := func(node *ns0.TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
