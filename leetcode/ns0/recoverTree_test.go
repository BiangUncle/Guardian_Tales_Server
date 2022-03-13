package ns0

import (
	"server/leetcode"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	root := leetcode.create_treeNode("[146,71,-13,55,null,231,399,321,null,null,null,null,null,-33]")

	recoverTree(root)
}
