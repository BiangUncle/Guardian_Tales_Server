package leetcode

import (
	"fmt"
	"server/leetcode/ns0"
	"strconv"
	"strings"
)

func create_treeNode(str string) *ns0.TreeNode {
	str = str[1 : len(str)-1]
	nums := strings.Split(str, ",")
	if len(nums) == 0 {
		return nil
	}

	val, _ := strconv.Atoi(nums[0])
	root := &ns0.TreeNode{Val: val}

	node_q := make([]*ns0.TreeNode, 0)
	node_q = append(node_q, root)

	i := 1
	for i < len(nums) && len(node_q) > 0 {
		cur := node_q[0]
		node_q = node_q[1:]

		if i < len(nums) && nums[i] != "null" {
			num, _ := strconv.Atoi(nums[i])
			node := &ns0.TreeNode{Val: num}
			node_q = append(node_q, node)
			cur.Left = node
		}
		i++

		if i < len(nums) && nums[i] != "null" {
			num, _ := strconv.Atoi(nums[i])
			node := &ns0.TreeNode{Val: num}
			node_q = append(node_q, node)
			cur.Right = node
		}
		i++

	}

	return root
}

func preorder_print(root *ns0.TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	str := "["
	var preorder func(root *ns0.TreeNode)
	preorder = func(root *ns0.TreeNode) {
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
