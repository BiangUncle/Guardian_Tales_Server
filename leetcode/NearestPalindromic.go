package leetcode

import (
	"math"
	"strconv"
)

const MaxFloat64 = 1.797693134862315708145274237317043567981e+308

func nearestPalindromic(nums string) string {
	n := len(nums)

	mid := -1

	if n%2 == 0 {
		mid = n/2 - 1
	} else {
		mid = n / 2
	}

	result := make([]byte, n)

	h := 0
	t := n - 1

	for h < mid {
		result[h] = nums[h]
		result[t] = nums[h]
		h++
		t--
	}

	min_ab := MaxFloat64
	min_val := nums[h]

	// 第一次判定

	result[h] = nums[h]
	if h != t {
		result[t] = nums[h]
	}

	i, _ := strconv.Atoi(string(result[h:]))
	j, _ := strconv.Atoi(nums[h:])

	ab := math.Abs(float64(i - j))
	if ab != 0 && ab < min_ab {
		min_ab = ab
		min_val = nums[h]
	}

	// 第二次判定

	if nums[h] != '9' {
		result[h] = nums[h] + 1
		if h != t {
			result[t] = nums[h] + 1
		}

		i, _ = strconv.Atoi(string(result[h:]))
		j, _ = strconv.Atoi(nums[h:])

		ab = math.Abs(float64(i - j))
		if ab != 0 && ab < min_ab {
			min_ab = ab
			min_val = nums[h] + 1
		}
	}

	// 第三次判定

	if nums[h] != '0' {
		result[h] = nums[h] - 1
		if h != t {
			result[t] = nums[h] - 1
		}

		i, _ = strconv.Atoi(string(result[h:]))
		j, _ = strconv.Atoi(nums[h:])

		ab = math.Abs(float64(i - j))
		if ab != 0 && ab < min_ab {
			min_ab = ab
			min_val = nums[h] - 1
		}
	}

	result[h] = min_val
	if h != t {
		result[t] = min_val
	}

	return string(result)
}

/*
"12932"
"99800"
"12120"
"1234"
"23"
"31"
"123"
*/
