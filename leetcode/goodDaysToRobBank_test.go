package leetcode

import (
	"fmt"
	"testing"
)

func Test_goodDaysToRobBank(t *testing.T) {

	security := []int{1, 2, 5, 4, 1, 0, 2, 4, 5, 3, 1, 2, 4, 3, 2, 4, 8}
	time := 2
	re := goodDaysToRobBank(security, time)
	fmt.Println(re)
	fmt.Println("-------------------------------------------------")
	//
	//security = []int{1, 1, 1, 1, 1}
	//time = 0
	//re = goodDaysToRobBank(security, time)
	//fmt.Println(re)
	//fmt.Println("-------------------------------------------------")
	//
	//security = []int{1, 2, 3, 4, 5, 6}
	//time = 2
	//re = goodDaysToRobBank(security, time)
	//fmt.Println(re)
	//fmt.Println("-------------------------------------------------")
	//
	//security = []int{1}
	//time = 5
	//re = goodDaysToRobBank(security, time)
	//fmt.Println(re)
}
