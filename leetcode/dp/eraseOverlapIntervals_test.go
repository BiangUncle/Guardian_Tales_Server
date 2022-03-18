package dp

import (
	"server/leetcode"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	intervals := leetcode.CreatRandomRange(10000)
	leetcode.Print2DimAarray(intervals)
	//intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	eraseOverlapIntervals(intervals)
}
