package ns0

import (
	"fmt"
	"testing"
)

func Test_find132pattern(t *testing.T) {
	nums := []int{3, 5, 0, 3, 4}
	re := find132pattern(nums)
	fmt.Println(re)
}
