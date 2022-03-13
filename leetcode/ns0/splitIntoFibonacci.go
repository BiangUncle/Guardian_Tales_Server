package ns0

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(num string) []int {
	n := len(num)

	ret := make([]int, 0)

	result := make([]int, 0)

	var split_num func(num string, start int) bool
	split_num = func(num string, start int) bool {

		func_ret := false

		if start == n {
			if IsFinish(ret) {
				result = ret
				return true
			}
			return false
		}

		if len(ret) >= 3 {
			if !IsFinish(ret) {
				return false
			}
		}

		if num[start] == '0' {
			ret = append(ret, 0)
			re := split_num(num, start+1)
			if re {
				return re
			}
			ret = ret[:len(ret)-1]
			return re
		}

		for j := start + 1; j <= n; j++ {
			digit, err := strconv.Atoi(num[start:j])
			if err != nil || digit > math.MaxInt32 {
				return false
			}
			ret = append(ret, digit)
			re := split_num(num, j)
			if re {
				func_ret = re
				break
			}
			ret = ret[:len(ret)-1]
		}

		return func_ret
	}

	split_num(num, 0)

	return result
}

func IsFinish(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	for i := 2; i < len(nums); i++ {
		if nums[i-2]+nums[i-1] != nums[i] {
			return false
		}
	}

	return true
}
