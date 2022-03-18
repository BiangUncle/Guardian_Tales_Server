package dp

import "testing"

func Test_minCost(t *testing.T) {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	minCost(costs)
}
