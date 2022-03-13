package ns0

import (
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 2}
	c := &TreeNode{Val: 3}
	d := &TreeNode{Val: 4}
	b.Left = a
	b.Right = d
	d.Left = c

	ser := Constructor()
	re := ser.serialize(b)
	//fmt.Println(re)
	ret := ser.deserialize(re)
	midorder(ret)

}
