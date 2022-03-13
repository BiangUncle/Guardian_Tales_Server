package ns0

func sumNumbers(root *TreeNode) int {
	sum := 0
	ret := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ret += sum
		}
		dfs(root.Left)
		dfs(root.Right)
		sum = (sum - root.Val) / 10
	}
	dfs(root)
	return ret
}
