package ns1

import "fmt"

type Item struct {
	Node  *Node
	layer int
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	stack := make([]*Item, 0)
	stack = append(stack, &Item{Node: root, layer: 0})
	result := root

	for len(stack) > 0 {
		item := stack[0]
		stack = stack[1:]
		cur_layer := item.layer
		cur_node := item.Node
		if cur_node.Left != nil {
			stack = append(stack, &Item{Node: cur_node.Left, layer: cur_layer + 1})
		}
		if cur_node.Right != nil {
			stack = append(stack, &Item{Node: cur_node.Right, layer: cur_layer + 1})
		}

		if len(stack) > 0 && cur_layer == stack[0].layer {
			cur_node.Next = stack[0].Node
		}
	}

	return result
}

func test() {
	n := 1<<12 - 1
	count := 0
	num := -1000
	str := "["
	for i := 0; i < n; i++ {
		str = str + fmt.Sprintf("%v", num) + ","
		count++
		if count == 3 {
			count = 0
			num++
		}
	}
	str = str[:len(str)-1] + "]"
	fmt.Println(str)
}
