package ns2

var m map[int]int

func init() {

	m = make(map[int]int)
	idx := 0
	num := 1
	for i := 0; i < 14; i++ {
		m[num] = idx
		idx++
		num <<= 1
	}

}

func countMaxOrSubsets(nums []int) int {
	max_val := 0
	for _, v := range nums {
		max_val |= v
	}
	ret := 0
	//n := len(nums)
	//bits := make([]int, 14)
	// 111 - 101 -> 010
	bit_map := make(map[int]int)

	for _, v := range nums {
		bit_map[v]++
		for key, value := range bit_map {
			if v|key == max_val {
				ret += value
			}
			bit_map[v|key]++
		}
	}
	return ret
}

func lowbit(x int) int {
	a := x & (-x)
	return a
}

func get1Idx(x int) []int {

	re := make([]int, 0)

	for x > 0 {
		a := lowbit(x)
		re = append(re, m[a])
		x -= a
	}

	return re
}
