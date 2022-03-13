package ns0

import (
	"fmt"
	"server/leetcode"
	"testing"
)

func Test_bestRotation(t *testing.T) {
	//nums := []int{2, 3, 1, 4, 0}
	nums := leetcode.create_random_list(9999)
	nums_str := leetcode.nums_list_2_str(nums)
	fmt.Println(nums_str)
	re := bestRotation(nums)
	fmt.Println(re)
}
