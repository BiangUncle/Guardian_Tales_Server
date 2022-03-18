package dp

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	l := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		step := nums[i]
		if step == 0 {
			l[i] = 1 << 32
			continue
		}
		cur_step := 1 << 31
		for j := min(i+step, n-1); j > i; j-- {
			if l[j]+1 < cur_step {
				cur_step = l[j] + 1
			}
		}
		l[i] = cur_step
	}
	fmt.Println(l)
	return l[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
