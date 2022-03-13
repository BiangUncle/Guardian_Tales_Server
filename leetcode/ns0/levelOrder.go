package ns0

type Node struct {
	Val      int
	Children []*Node
}

type Queue struct {
	Node  *Node
	Layer int
}

func levelOrder(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	node_list := make([]*Queue, 0)
	sub_list := make([]int, 0)
	cur := 0

	node_list = append(node_list, &Queue{root, cur + 1})
	for len(node_list) > 0 {
		q := node_list[0]
		node_list = node_list[1:]
		layer := q.Layer
		if layer != cur {
			if len(sub_list) != 0 {
				ret = append(ret, sub_list)
			}
			sub_list = make([]int, 0)
			cur += 1
		}
		sub_list = append(sub_list, q.Node.Val)
		for _, node := range q.Node.Children {
			node_list = append(node_list, &Queue{node, cur + 1})
		}
	}
	if len(sub_list) != 0 {
		ret = append(ret, sub_list)
	}
	return ret
}
