package ns2

import (
	"fmt"
	"server/leetcode"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	root := leetcode.create_treeNode("[1,2,3,null,5,null,4]")
	re := rightSideView(root)
	fmt.Println(re)
}

//func Test_countNodes(t *testing.T) {
//	root := create_treeNode("[1,2,3,4,5,6]")
//	countNodes(root)
//}
