package ns0

func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	children := make([][]int, n)
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	rem := make([]int, n)

	maxScore := 0
	var dfs func(int) int
	dfs = func(node int) int {
		if rem[node] != 0 {
			return rem[node]
		}
		score, size := 1, n-1
		for _, ch := range children[node] {
			sz := dfs(ch)
			score *= sz
			size -= sz
		}
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		rem[node] = n - size
		return n - size
	}
	dfs(0)
	return
}
