package ns0

import "fmt"

func recoverTree(root *TreeNode) {

	l := make([]int, 0)

	var p func(root *TreeNode)
	p = func(root *TreeNode) {
		if root == nil {
			return
		}

		p(root.Left)
		l = append(l, root.Val)
		p(root.Right)
	}
	p(root)

	fmt.Println(l)

	err_key := make(map[int]int)
	stack := make([]int, 0)
	for i := 0; i < len(l); i++ {

		for len(stack) > 0 && stack[len(stack)-1] > l[i] {
			err_key[stack[len(stack)-1]] = 1
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, l[i])
	}

	fmt.Println("===============")
	for k, _ := range err_key {
		fmt.Println(k)
	}

	var node1 *TreeNode
	var node2 *TreeNode

	var fac func(root *TreeNode)
	fac = func(root *TreeNode) {
		if root == nil {
			return
		}

		fac(root.Left)
		fac(root.Right)
		_, ok := err_key[root.Val]
		if ok {
			if node1 == nil {
				node1 = root
				return
			}
			if node2 == nil {
				node2 = root
			}
		}
	}
	fac(root)

	node1.Val, node2.Val = node2.Val, node1.Val
}
