package data_struct

import "testing"

func TestMaxHeap_Insert(t *testing.T) {
	h := MaxHeap{}
	h.Insert(3)
	h.Insert(1)
	h.Insert(9)
	h.Insert(-2)
	h.Insert(9)
	h.Insert(3)
	h.Insert(1)
	println(h.Pop())
	println(h.Pop())
	println(h.Pop())
	println(h.Pop())
	println(h.Pop())
	println(h.Pop())
	println(h.Pop())
}
