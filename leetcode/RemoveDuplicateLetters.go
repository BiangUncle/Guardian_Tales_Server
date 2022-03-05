package leetcode

func removeDuplicateLetters(s string) string {
	n := len(s)
	charmap := make(map[byte]int)
	delmap := make(map[int]byte)
	var stack []int

	for i := 0; i < n; i++ {

		c := s[i]

		for len(stack) > 0 && c >= s[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if idx, ok := charmap[c]; ok {
			if len(stack) > 0 && idx < stack[len(stack)-1] {
				// 删掉当前的
				delmap[i] = c
			} else {
				delmap[idx] = s[idx]
				charmap[c] = i
			}
		} else {
			charmap[c] = i
		}

		stack = append(stack, i)

	}

	result := make([]byte, 0)
	for i := 0; i < n; i++ {
		if _, ok := delmap[i]; ok {
			continue
		}
		result = append(result, s[i])
	}

	return string(result)
}
