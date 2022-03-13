package ns0

import "fmt"

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	up := make([]int, n)
	down := make([]int, n)
	cur := 0

	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			cur++
			down[i] = cur
		} else {
			cur = 0
		}
	}

	cur = 0
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			cur++
			up[i] = cur
		} else {
			cur = 0
		}
	}

	fmt.Println(up)
	fmt.Println(down)

	result := make([]int, 0)

	for i := 0; i < n; i++ {
		if down[i] >= time && up[i] >= time {
			result = append(result, i)
		}
	}

	return result
}
