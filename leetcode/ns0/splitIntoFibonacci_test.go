package ns0

import (
	"fmt"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	num := "74912134825162255812723932620170946950766784234934"
	re := splitIntoFibonacci(num)
	fmt.Println(re)
}
