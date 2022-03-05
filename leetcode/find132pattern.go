package leetcode

func find132pattern(nums []int) bool {
	n := len(nums)
	minRight := make([]int, n)
	var s []int

	for i := n - 1; i >= 0; i-- {

		for len(s) > 0 && nums[i] < nums[s[len(s)-1]] {
			s = s[:len(s)-1]
		}

		if len(s) == 0 {
			minRight[i] = -1
		} else {
			minRight[i] = s[len(s)-1]
		}
		s = append(s, i)

	}

	min_val := nums[0]

	for i := 1; i < n; i++ {

		// 左边没有较小值
		if nums[i] < min_val {
			min_val = nums[i]
			continue
		}

		// 右边没有较小值
		if minRight[i] == -1 {
			continue
		}

		if min_val < nums[minRight[i]] && nums[minRight[i]] < nums[i] {
			return true
		}

	}

	return false
}