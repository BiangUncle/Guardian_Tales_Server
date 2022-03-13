package ns0

import (
	"fmt"
	"testing"
)

func Test_removeDuplicateLetters(t *testing.T) {
	//re := removeDuplicateLetters("bcabc")
	//fmt.Println(re)

	re := removeDuplicateLetters("cbacdcbc")
	fmt.Println(re)
}
