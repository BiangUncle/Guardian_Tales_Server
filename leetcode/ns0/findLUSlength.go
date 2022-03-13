package ns0

func haveSame(strs []string) int {
	n := len(strs)
	if n == 1 {
		return n
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if strs[i] != strs[j] {
				return len(strs[i])
			}
		}
	}
	return -1
}

func isSubStr(a, b string) int {
	n := len(a)
	m := len(b)
	ptr_a := 0
	ptr_b := 0
	for ptr_b < m && ptr_a < n {
		if a[ptr_a] == b[ptr_b] {
			ptr_a++
			ptr_b++
		} else {
			ptr_a++
		}
	}
	if ptr_b != m {
		return m
	}
	return -1
}

func findLUSlengths(a, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
