package data_struct

import "testing"

func TestMinHeap(t *testing.T) {
	h := MinHeap{}
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
