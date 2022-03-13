package ns0

import (
	"fmt"
	"testing"
)

func Test_validUtf8(t *testing.T) {
	data := []int{1, 196, 130}
	re := validUtf8(data)
	fmt.Println(re)
}
