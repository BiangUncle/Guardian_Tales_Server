package ns0

import (
	"server/leetcode"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	//head := create_list_node("[-10,4,-3,0,5,9]")
	//root := sortedListToBST(head)
	//preorder_print(root)
	leetcode.test()
}

func Test_buildTree(t *testing.T) {
	inorder := []int{9, 10, 3, 15, 20, 7}
	postorder := []int{10, 9, 15, 7, 20, 3}
	root := leetcode.buildTree(inorder, postorder)
	leetcode.preorder_print(root)
}
