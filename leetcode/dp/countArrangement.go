package dp

func countArrangement(n int) int {
	div_map := make(map[int][]int)
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			if i%j == 0 {
				div_map[i] = append(div_map[i], j)
			}
		}
	}

	cur_idx := 1
	for i := 1; i <= n; i++ {
		if i%cur_idx == 0 {

		}

		if cur_idx%i == 0 {

		}
	}

	return 0
}
