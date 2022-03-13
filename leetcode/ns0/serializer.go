package ns0

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	re := ""

	preorder(1, &re, root)

	return re[1:]
}

func preorder(idx int, str *string, root *TreeNode) {
	if root == nil {
		return
	}

	*str = *str + "|" + fmt.Sprintf("%d:%d", idx, root.Val)
	preorder(idx*2, str, root.Left)
	preorder(idx*2+1, str, root.Right)
}

//func (this *Codec) serialize(root *TreeNode) string {
//	pre := make([]int, 0)
//	mid := make([]int, 0)
//	midorder(&mid, root)
//	preorder(&pre, root)
//	pre_str := fmt.Sprintf("%v", pre)
//	mid_str := fmt.Sprintf("%v", mid)
//
//	return pre_str[1:len(pre_str)-1] + "|" + mid_str[1:len(mid_str)-1]
//}

//func midorder(nums *[]int, root *TreeNode) {
//	if root == nil {
//		return
//	}
//	midorder(nums, root.Left)
//	*nums = append(*nums, root.Val)
//	midorder(nums, root.Right)
//}
//
//func preorder(nums *[]int, root *TreeNode) {
//	if root == nil {
//		return
//	}
//	*nums = append(*nums, root.Val)
//	preorder(nums, root.Left)
//	preorder(nums, root.Right)
//}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	map_str := strings.Split(data, "|")
	nums_map := make(map[int]int)

	for _, str := range map_str {
		kv := strings.Split(str, ":")
		k, _ := strconv.Atoi(kv[0])
		v, _ := strconv.Atoi(kv[1])
		nums_map[k] = v
	}

	//for k, v := range nums_map {
	//	fmt.Println(k, v)
	//}

	return ordercreate(1, nums_map)
}

func ordercreate(idx int, nums_map map[int]int) *TreeNode {
	k, ok := nums_map[idx]
	if !ok {
		return nil
	}

	root := &TreeNode{Val: k}
	root.Left = ordercreate(idx*2, nums_map)
	root.Right = ordercreate(idx*2+1, nums_map)

	return root
}

func midorder(root *TreeNode) {
	if root == nil {
		return
	}
	midorder(root.Left)
	fmt.Println(root.Val)
	midorder(root.Right)
}

//func (this *Codec) deserialize(data string) *TreeNode {
//	strs := strings.Split(data, "|")
//	pre := strs[0]
//	mid := strs[1]
//	nums_pre := str_2_nums(pre)
//	nums_mid := str_2_nums(mid)
//	fmt.Println(nums_pre)
//	fmt.Println(nums_mid)
//	return nil
//}

//func str_2_nums(str string) []int {
//	nums_str := strings.Split(str, " ")
//	n := len(nums_str)
//	re := make([]int, n)
//	for i := 0; i < n; i++ {
//		re[i], _ = strconv.Atoi(nums_str[i])
//	}
//	return re
//}

//func binary_find(aim int, nums []int) int {
//	l := 0
//	r := len(nums) - 1
//	for l <= r {
//		mid := (l + r) / 2
//		if nums[mid] == aim {
//			return mid
//		}
//		if nums[mid] > aim {
//			r = mid - 1
//			continue
//		}
//		if nums[mid] < aim {
//			l = mid + 1
//		}
//	}
//	return -1
//}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
