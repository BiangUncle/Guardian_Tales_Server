package dp

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {

	n := len(intervals)
	if n == 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		}
		if intervals[i][1] > intervals[j][1] {
			return false
		}
		if intervals[i][1] == intervals[j][1] && intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	//Print2DimAarray(intervals)
	//left := intervals[0][0]
	right := intervals[0][1]

	ret := 0
	for i := 1; i < n; i++ {
		rang := intervals[i]
		if rang[0] >= right {
			right = rang[1]
			continue
		}
		ret++
	}

	return ret
}
