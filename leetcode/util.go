package leetcode

import (
	"fmt"
	"math/rand"
	"server/leetcode/ns0"
	"strconv"
	"time"
)

const MIN = -50000

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

func CreatRandomRange(n int) [][]int {
	ret := make([][]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		ret[i] = make([]int, 2)
		v := ret[i]
		v[0] = MIN + rand.Intn(100000)
		v[1] = v[0] + 1 + rand.Intn(50000-v[0])
	}
	return ret
}

func Print2DimAarray(nums [][]int) {
	str := "["
	for i := 0; i < len(nums); i++ {
		str = str + fmt.Sprintf("[%d,%d],", nums[i][0], nums[i][1])
	}
	str = str[:len(str)-1] + "]"
	fmt.Println(str)
}

func CreateRandomPoint(num, n int) [][]int {
	ret := make([][]int, 0)
	rand.Seed(time.Now().Unix())

	for i := 0; i < num; i++ {
		point := make([]int, 2)
		point[0] = rand.Intn(n)
		point[1] = rand.Intn(n)
		ret = append(ret, point)
	}
	return ret
}
