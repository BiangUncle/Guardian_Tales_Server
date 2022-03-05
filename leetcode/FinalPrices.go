package leetcode

import "fmt"

func finalPrices(prices []int) []int {
	n := len(prices)
	minRight := make([]int, n)
	var s []int

	for i := n - 1; i >= 0; i-- {
		price := prices[i]

		for len(s) > 0 && price < s[len(s)-1] {
			s = s[:len(s)-1]
		}

		if len(s) == 0 {
			minRight[i] = 0
		} else {
			minRight[i] = s[len(s)-1]
		}

		s = append(s, price)
	}

	fmt.Println(minRight)

	result := make([]int, n)
	for i, price := range prices {
		result[i] = price - minRight[i]
	}

	return result
}
