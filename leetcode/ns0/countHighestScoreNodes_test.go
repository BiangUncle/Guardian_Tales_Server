package ns0

import "testing"

func Test_countHighestScoreNodes(t *testing.T) {
	parents := []int{-1, 2, 0, 2, 0}
	countHighestScoreNodes(parents)
}
