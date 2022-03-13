package ns0

import (
	"testing"
)

/*
[4,2,4,2,3]
[1,2,3,4,5]
[9,0,1,6]
[1,2]
[1]
*/

func Test_finalPrices(t *testing.T) {
	prices := []int{8, 4, 6, 2, 3}

	finalPrices(prices)

	//prices = []int{1, 2, 3, 4, 5}
	//finalPrices(prices)
	//
	//prices = []int{10, 1, 1, 6}
	//finalPrices(prices)
	//
	//prices = []int{1, 2}
	//finalPrices(prices)
	//
	//prices = []int{1}
	//finalPrices(prices)
}
