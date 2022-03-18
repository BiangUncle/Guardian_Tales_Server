package dp

func stoneGame(piles []int) bool {

	n := len(piles)
	pre := make([][]int, n)
	post := make([][]int, n)

	for i := 0; i < n; i++ {
		pre[i] = make([]int, n)
		post[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pre[i][j] = -1
			post[i][j] = -1
		}
	}
	for idx, v := range piles {
		pre[idx][idx] = v
		post[idx][idx] = 0
	}

	var dfs_pre func(i, j int) int
	var dfs_post func(i, j int) int
	dfs_pre = func(i, j int) int {
		if pre[i][j] != -1 {
			return pre[i][j]
		}
		pre[i][j] = max(piles[i]+dfs_post(i+1, j), piles[j]+dfs_post(i, j-1))
		return pre[i][j]
	}
	dfs_post = func(i, j int) int {
		if post[i][j] != -1 {
			return post[i][j]
		}
		post[i][j] = max(dfs_pre(i+1, j), dfs_pre(i, j-1))
		return post[i][j]
	}
	alice := dfs_pre(0, n-1)
	bob := dfs_post(0, n-1)
	return alice > bob
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
