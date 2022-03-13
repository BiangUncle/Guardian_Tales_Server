package ns0

import (
	"fmt"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	a := &Node{Val: 1}
	a.Children = []*Node{&Node{Val: 2}, &Node{Val: 3}, &Node{Val: 4}}
	re := levelOrder(a)
	for _, num := range re {
		fmt.Println(num)
	}
}
