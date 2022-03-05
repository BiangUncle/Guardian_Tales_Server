package leetcode

import (
	"fmt"
	"testing"
)

func Test_nearestPalindromic(t *testing.T) {
	nums := "123"
	result := nearestPalindromic(nums)
	fmt.Println(nums, "->", result)
}
