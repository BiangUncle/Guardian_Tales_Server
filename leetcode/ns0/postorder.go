package ns0

func postorder(root *Node) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*Node, 0)

	stack = append(stack, root)

	m := make(map[*Node]int)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if _, ok := m[node]; ok {
			stack = stack[:len(stack)-1]
			ret = append(ret, node.Val)
			continue
		}
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		m[node] = 1

	}
	return ret
}
