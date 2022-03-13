package ns0

// [2,3,1,4,0]

// 3 -> [4] 为起点
// 3 -> [3] 为终点

// [1] -> [4] = (1 - 4 + 5) % 5 = 2
// [1] -> [3] = (1 - 3 + 5) % 5 = 3

func bestRotation(nums []int) int {
	n := len(nums)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		j := nums[i]
		start := (i + 1) % n
		end := (i - j + n) % n
		count[start]++
		if end < n-1 {
			count[end+1]--
		}
	}

	//fmt.Println(count)
	res_num := -1
	res_rot := 0

	res := 0

	for i := 0; i < n; i++ {
		res += count[i]
		if res > res_num {
			res_num = res
			res_rot = i
		}
	}

	return res_rot
}
