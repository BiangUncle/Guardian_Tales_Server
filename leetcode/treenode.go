package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func create_treeNode(str string) *TreeNode {
	str = str[1 : len(str)-1]
	nums := strings.Split(str, ",")
	if len(nums) == 0 {
		return nil
	}

	val, _ := strconv.Atoi(nums[0])
	root := &TreeNode{Val: val}

	node_q := make([]*TreeNode, 0)
	node_q = append(node_q, root)

	i := 1
	for i < len(nums) && len(node_q) > 0 {
		cur := node_q[0]
		node_q = node_q[1:]

		if i < len(nums) && nums[i] != "null" {
			num, _ := strconv.Atoi(nums[i])
			node := &TreeNode{Val: num}
			node_q = append(node_q, node)
			cur.Left = node
		}
		i++

		if i < len(nums) && nums[i] != "null" {
			num, _ := strconv.Atoi(nums[i])
			node := &TreeNode{Val: num}
			node_q = append(node_q, node)
			cur.Right = node
		}
		i++

	}

	return root
}

func preorder_print(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	str := "["
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		str = str + fmt.Sprintf("%v", root.Val) + ","
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	str = str[:len(str)-1] + "]"

	fmt.Println(str)
}
