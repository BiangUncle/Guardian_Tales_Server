package ns0

func pathSum(root *TreeNode, targetSum int) int {

	ret := 0

	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {

		ret_list := make([]int, 0)

		if root == nil {
			return ret_list
		}

		left_ret := dfs(root.Left)
		for _, v := range left_ret {
			s := v + root.Val
			if s == targetSum {
				ret++
			}
			ret_list = append(ret_list, s)
		}

		right_ret := dfs(root.Right)
		for _, v := range right_ret {
			s := v + root.Val
			if s == targetSum {
				ret++
			}
			ret_list = append(ret_list, s)
		}

		if targetSum == root.Val {
			ret++
		}

		ret_list = append(ret_list, root.Val)

		return ret_list
	}
	dfs(root)

	return ret
}
