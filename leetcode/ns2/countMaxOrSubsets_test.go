package ns2

import (
	"fmt"
	"testing"
)

func Test_countMaxOrSubsets(t *testing.T) {
	nums := []int{3, 2, 1, 5}
	re := countMaxOrSubsets(nums)
	fmt.Println(re)
}

func Test_lowbit(t *testing.T) {
	lowbit(1)
	lowbit(2)
	lowbit(3)
	lowbit(4)
}

func Test_get1Idx(t *testing.T) {
	a := 5 & (^4)
	fmt.Println(a)
	//re := get1Idx(0)
	//fmt.Println(re)
	//re = get1Idx(1)
	//fmt.Println(re)
	//re = get1Idx(2)
	//fmt.Println(re)
	//re = get1Idx(3)
	//fmt.Println(re)
	//re = get1Idx(4)
	//fmt.Println(re)
	//re = get1Idx(5)
	//fmt.Println(re)
	//re = get1Idx(10000)
	//fmt.Println(re)
}
