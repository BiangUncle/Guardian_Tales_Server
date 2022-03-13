package ns0

import (
	"fmt"
	"server/leetcode"
	"testing"
)

func Test_pathSum(t *testing.T) {
	str := "[10,5,-3,3,2,null,11,3,-2,null,1]"
	root := leetcode.create_treeNode(str)
	targetSum := 8
	re := pathSum(root, targetSum)
	fmt.Println(re)

	str = "[5,4,8,11,null,13,4,7,2,null,null,5,1]"
	root = leetcode.create_treeNode(str)
	targetSum = 22
	re = pathSum(root, targetSum)
	fmt.Println(re)
}
