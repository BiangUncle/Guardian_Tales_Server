package ns0

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	PlatesNum := make([]int, n)
	plates := 0

	rightCandles := make([]int, n)
	leftCandles := make([]int, n)
	leftidx := -1

	for i := 0; i < n; i++ {
		if s[i] == '|' {
			leftidx = i
		} else {
			plates++
		}
		PlatesNum[i] = plates
		leftCandles[i] = leftidx
	}

	rightidx := -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			rightidx = i
		}
		rightCandles[i] = rightidx
	}

	result := make([]int, len(queries))

	for idx, query := range queries {
		i, j := query[0], query[1]

		// 找到最近的蜡烛
		right_idx := rightCandles[i]
		left_idx := leftCandles[j]

		if right_idx != -1 && left_idx != -1 && right_idx < left_idx {
			result[idx] = PlatesNum[left_idx] - PlatesNum[right_idx]
		} else {
			result[idx] = 0
		}
	}

	return result
}
