package dp

import (
	"server/leetcode"
	"testing"
)

func Test_orderOfLargestPlusSign(t *testing.T) {
	n := 20
	mines := leetcode.CreateRandomPoint(10, n)
	leetcode.Print2DimAarray(mines)
	orderOfLargestPlusSign(n, mines)
}
