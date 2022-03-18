package ns2

import "server/leetcode"

type GraphNode struct {
	Val   int
	Nodes []*GraphNode
}

type SumNode struct {
	Val      int
	Left     *SumNode
	Right    *SumNode
	LeftSum  []int
	RightSum []int
}

func maxPathSum(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	ret := -(1 << 63)

	var dfs_creat func(root *leetcode.TreeNode, rootSum *SumNode, dir bool)
	dfs_creat = func(root *leetcode.TreeNode, rootSum *SumNode, dir bool) {
		if root == nil {
			return
		}

		new_rootSum := &SumNode{Val: root.Val}
		dfs_creat(root.Left, new_rootSum, false)
		dfs_creat(root.Right, new_rootSum, true)

		result := make([]int, 0)
		if root.Val > ret {
			ret = root.Val
		}
		result = append(result, root.Val)
		for _, v := range new_rootSum.LeftSum {
			cur := v + root.Val
			if cur > ret {
				ret = cur
			}
			result = append(result, cur)
		}
		for _, v := range new_rootSum.RightSum {
			cur := v + root.Val
			if cur > ret {
				ret = cur
			}
			result = append(result, cur)
		}
		if dir {
			rootSum.RightSum = result
		} else {
			rootSum.LeftSum = result
		}

	}
	sumNode := &SumNode{Val: root.Val}
	dfs_creat(root, sumNode, false)
	dfs_creat(root, sumNode, true)

	var dfs_sum func(sumNode *SumNode)

	dfs_sum = func(sumNode *SumNode) {

	}

	dfs_sum(sumNode)

	return ret

}

//graphNode := &GraphNode{Val: root.Val}
//AddGraphNode(root.Left, graphNode)
//AddGraphNode(root.Right, graphNode)
func AddGraphNode(root *leetcode.TreeNode, graphNode *GraphNode) {
	if root == nil {
		return
	}
	new_graphNode := &GraphNode{Val: root.Val}
	graphNode.Nodes = append(graphNode.Nodes, new_graphNode)
	new_graphNode.Nodes = append(new_graphNode.Nodes, graphNode)

	AddGraphNode(root.Left, graphNode)
	AddGraphNode(root.Right, graphNode)
}

type LayerTreeNode struct {
	Layer int
	Node  *leetcode.TreeNode
}

func rightSideView(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	q := make([]*LayerTreeNode, 0)
	q = append(q, &LayerTreeNode{Node: root, Layer: 0})

	for len(q) > 0 {
		top := q[len(q)-1]
		q = q[:len(q)-1]

		curLayer := top.Layer
		curNode := top.Node
		if len(result) < curLayer+1 {
			result = append(result, curNode.Val)
		} else {
			result[curLayer] = curNode.Val
		}
		if curNode.Right != nil {
			q = append(q, &LayerTreeNode{Node: curNode.Right, Layer: curLayer + 1})
		}
		if curNode.Left != nil {
			q = append(q, &LayerTreeNode{Node: curNode.Left, Layer: curLayer + 1})
		}

	}

	return result
}

//func countNodes(root *TreeNode) int {
//
//	if root == nil {
//		return 0
//	}
//
//	layer := 1
//	cur := root
//	left := 1
//	if cur.Right != nil {
//		layer++
//		left = 1<<layer - 1
//		cur = cur.Right
//	}
//	left++
//	layer++
//	right := 1<<(layer) - 2
//
//	for left < right {
//		mid := (left + right) / 2
//
//	}
//	return 0
//}

//func findMid(root *TreeNode, aim int) bool {
//	idx := 1
//	if root.Left != nil {
//
//	}
//}
