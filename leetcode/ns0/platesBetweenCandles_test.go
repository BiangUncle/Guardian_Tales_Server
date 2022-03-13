package ns0

import (
	"fmt"
	"server/leetcode"
	"testing"
)

func Test_platesBetweenCandles(t *testing.T) {
	s := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	re := platesBetweenCandles(s, queries)
	fmt.Println(re)

	s = "***|**|*****|**||**|*"
	queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	re = platesBetweenCandles(s, queries)
	fmt.Println(re)
}

func TestCreateTestExample(t *testing.T) {
	s, queries := leetcode.CreateTestExample(10000, 30)
	fmt.Println(s)

	q_str := "["
	for _, q := range queries {
		q_str = q_str + "[" + fmt.Sprint(q[0]) + "," + fmt.Sprint(q[1]) + "],"
	}
	q_str = q_str[:len(q_str)-1] + "]"
	fmt.Println(q_str)
}
