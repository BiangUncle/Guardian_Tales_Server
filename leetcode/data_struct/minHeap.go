package data_struct

type MinHeap struct {
	List []int
}

func (self *MinHeap) Insert(val int) {
	self.List = append(self.List, val)
	cur_idx := len(self.List) - 1
	for {
		parent_idx := (cur_idx - 1) / 2
		if parent_idx >= 0 {
			if self.List[parent_idx] > self.List[cur_idx] {
				self.List[parent_idx], self.List[cur_idx] = self.List[cur_idx], self.List[parent_idx]
				cur_idx = parent_idx
				continue
			}
		}
		break
	}
}

func (self *MinHeap) Pop() int {
	if len(self.List) <= 0 {
		panic("Empty Heap Pop!")
	}
	ret := self.List[0]
	n := len(self.List)
	self.List[0] = self.List[n-1]
	self.List = self.List[:len(self.List)-1]
	n -= 1
	cur_idx := 0
	for {
		child1, child2, aim_idx := cur_idx*2+1, cur_idx*2+2, cur_idx
		if child1 < n && self.List[child1] < self.List[aim_idx] {
			aim_idx = child1
		}
		if child2 < n && self.List[child2] < self.List[aim_idx] {
			aim_idx = child2
		}
		if aim_idx != cur_idx {
			self.List[aim_idx], self.List[cur_idx] = self.List[cur_idx], self.List[aim_idx]
			cur_idx = aim_idx
		}
		break
	}
	return ret
}
