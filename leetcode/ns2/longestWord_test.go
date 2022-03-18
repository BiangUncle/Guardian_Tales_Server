package ns2

import (
	"fmt"
	"testing"
)

func Test_longestWord(t *testing.T) {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	re := longestWord(words)
	fmt.Println(re)
}
