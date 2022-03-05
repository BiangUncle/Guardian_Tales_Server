package leetcode

func removeKdigits(num string, k int) string {

	n := len(num)
	if n == k {
		return "0"
	}

	count := 0
	var s []byte

	for i := 0; i < n; i++ {

		for len(s) > 0 && count < k && num[i] < s[len(s)-1] {
			s = s[:len(s)-1]
			count++
		}

		s = append(s, num[i])
	}

	for count < k {
		s = s[:len(s)-1]
		count++
	}

	cut := 0
	for cut < len(s) {
		if s[cut] == '0' {
			cut++
		} else {
			break
		}
	}

	if cut == len(s) {
		return "0"
	}
	return string(s[cut:])

}
