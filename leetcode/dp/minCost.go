package dp

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = min(dp[1][i-1]+costs[i-1][0], dp[2][i-1]+costs[i-1][0])
		dp[1][i] = min(dp[0][i-1]+costs[i-1][1], dp[2][i-1]+costs[i-1][1])
		dp[2][i] = min(dp[0][i-1]+costs[i-1][2], dp[1][i-1]+costs[i-1][2])
	}
	return min(dp[0][n], min(dp[1][n], dp[2][n]))
}
