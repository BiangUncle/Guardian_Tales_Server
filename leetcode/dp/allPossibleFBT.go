package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {

	if n%2 == 0 {
		return nil
	}
	treeNode := make(map[int][]*TreeNode)
	treeNode[1] = append(treeNode[1], &TreeNode{})

	for i := 3; i <= n; i += 2 {

		for k, nodes := range treeNode {

			if k >= i {
				continue
			}

			remain := i - 1 - k

			for _, node := range nodes {

				right_nodes, _ := treeNode[remain]

				for _, right_node := range right_nodes {
					new_node := &TreeNode{}
					new_node.Left = node
					new_node.Right = right_node
					treeNode[i] = append(treeNode[i], new_node)
				}
			}

		}
	}

	return treeNode[n]
}

/*
[0,0,0,null,null,0,0,0,0,0,0],[0,0,0,null,null,0,0,0,0,0,0],

[0,0,0,null,null,0,0,0,0,null,null,null,null,0,0],[0,0,0,null,null,0,0,0,0,null,null,null,null,0,0],

[0,0,0,null,null,0,0,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0,null,null,0,0],

[0,0,0,null,null,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,null,null,0,0,null,null,0,0],

[0,0,0,null,null,0,0,null,null,0,0,0,0],[0,0,0,null,null,0,0,null,null,0,0,0,0],

[0,0,0,0,0,0,0,null,null,null,null,null,null,0,0],[0,0,0,0,0,0,0,null,null,null,null,null,null,0,0],

[0,0,0,0,0,0,0,null,null,null,null,0,0],
[0,0,0,0,0,0,0,null,null,0,0],
[0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,null,null,0,0,0,0],
[0,0,0,0,0,null,null,0,0,null,null,null,null,0,0],
[0,0,0,0,0,null,null,0,0,null,null,0,0],
[0,0,0,0,0,null,null,null,null,0,0,null,null,0,0],
[0,0,0,0,0,null,null,null,null,0,0,0,0]

[0,0,0,0,0,0,0,null,null,null,null,0,0],
[0,0,0,0,0,0,0,null,null,0,0],
[0,0,0,0,0,0,0,0,0],
[0,0,0,0,0,null,null,null,null,0,0,null,null,0,0],
[0,0,0,0,0,null,null,null,null,0,0,0,0],
[0,0,0,0,0,null,null,0,0,0,0],
[0,0,0,0,0,null,null,0,0,null,null,null,null,0,0],
[0,0,0,0,0,null,null,0,0,null,null,0,0]
*/
