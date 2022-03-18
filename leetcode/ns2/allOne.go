package ns2

//type StringListnode struct {
//	Val   string // len = [0, 10]
//	Pre   *StringListnode
//	Next  *StringListnode
//	Count int
//}
//
//type AllOne struct {
//	ValMap  map[string]*StringListnode
//	Head    *StringListnode
//	Tail    *StringListnode
//	LenList []*StringListnode
//}
//
//func Constructor() AllOne {
//	allone := AllOne{}
//	allone.ValMap = make(map[string]*StringListnode)
//	allone.Head = &StringListnode{}
//	allone.Tail = &StringListnode{}
//	allone.Head.Next = allone.Tail
//	allone.Tail.Pre = allone.Head
//	allone.LenList = make([]*StringListnode, 11)
//	return allone
//}
//
//func (this *AllOne) Inc(key string) {
//	list, ok := this.ValMap[key]
//	if ok {
//		list.Count++
//	}
//	list = &StringListnode{Val: key, Count: 1}
//	first := this.Head.Next
//	if len(key) >= len(first.Val) {
//		list.Next = first
//		first.Pre = list
//		this.Head.Next = list
//		list.Pre = this.Head
//		return
//	}
//
//	tail := this.Tail.Pre
//
//	if len(key) <= len(tail.Val) {
//		tail.Next = list
//		list.Pre = tail
//		list.Next = this.Tail
//		this.Tail.Pre = list
//		return
//	}
//
//	this.InsertLenList(list)
//}
//
//func (this *AllOne) InsertLenList(list *StringListnode) {
//	list_len := len(list.Val)
//	if this.LenList[list_len] != nil {
//		aim_list := this.LenList[list_len]
//		list.Next = aim_list.Next
//		list.Next.Pre = list
//		aim_list.Next = list
//		list.Pre = aim_list
//		return
//	}
//	this.LenList[list_len] = list
//	for i := list_len + 1; i <= 10; i++ {
//		if this.LenList[i] != nil {
//			aim_list := this.LenList[i]
//			list.Pre = aim_list.Pre
//			list.Pre.Next = list
//			list.Next = aim_list
//			aim_list.Pre = list
//			return
//		}
//	}
//	return
//}
//
//func (this *AllOne) Dec(key string) {
//	list, _ := this.ValMap[key]
//	list.Count--
//	if list.Count == 0 {
//		list.Pre.Next = list.Next
//		list.Next.Pre = list.Pre
//	}
//}
//
//func (this *AllOne) GetMaxKey() string {
//	if this.Head.Next != this.Tail {
//		return this.Head.Next.Val
//	}
//	return ""
//}
//
//func (this *AllOne) GetMinKey() string {
//	if this.Tail.Pre != this.Head {
//		return this.Tail.Pre.Val
//	}
//	return ""
//}

//type CountList struct {
//	Count   int
//	Next    *CountList
//	Pre     *CountList
//	Strings map[string]int
//}
//
//type AllOne struct {
//	ValMap   map[string]int
//	Head     *CountList
//	Tail     *CountList
//	CountMap map[int]*CountList
//}
//
//func Constructor() AllOne {
//	allone := AllOne{}
//	allone.ValMap = make(map[string]int)
//	allone.CountMap = make(map[int]*CountList)
//	allone.Head = &CountList{}
//	allone.Tail = &CountList{}
//	allone.Head.Next = allone.Tail
//	allone.Tail.Pre = allone.Head
//	return allone
//}
//
//func (this *AllOne) Inc(key string) {
//	_, ok := this.ValMap[key]
//	if !ok {
//		this.ValMap[key] = 1
//		if countList, have := this.CountMap[1]; have {
//			countList.Strings[key] = 1
//		} else {
//			count := CountListCreator(this.ValMap[key], key)
//			this.Head.InsertAfter(count)
//		}
//		return
//	}
//
//	count := this.CountMap[this.ValMap[key]]
//	delete(count.Strings, key)
//	this.ValMap[key]++
//	if count.Next != this.Tail && count.Next.Count == this.ValMap[key] {
//		count.Next.Strings[key] = 1
//	} else {
//		countList := CountListCreator(this.ValMap[key], key)
//		count.InsertAfter(countList)
//		if len(count.Strings) == 0 {
//			DeleteCountList(count)
//		}
//	}
//}
//
//func (this *AllOne) Dec(key string) {
//	v, _ := this.ValMap[key]
//	if v == 1 {
//		delete(this.ValMap, key)
//		countList, _ := this.CountMap[1]
//		delete(, key)
//	}
//}
//
//func CountListCreator(count int, key string) *CountList {
//	countList := &CountList{Count: count}
//	countList.Strings = make(map[string]int)
//	countList.Strings[key] = 1
//	return countList
//}
//
//func DeleteCountListIfEmpty(count int, key string) {
//
//}
//
//func DeleteCountList(count *CountList) {
//	count.Pre.Next = count.Next
//	count.Next.Pre = count.Pre
//}
//
//func (self *CountList) InsertAfter(count *CountList) {
//	count.Next = self.Next
//	count.Next.Pre = count
//	self.Next = count
//	count.Pre = self
//}
//
//func (this *AllOne) GetMaxKey() string {
//	if this.Head.Next != this.Tail {
//		return this.Head.Next.Val
//	}
//	return ""
//}
//
//func (this *AllOne) GetMinKey() string {
//	if this.Tail.Pre != this.Head {
//		return this.Tail.Pre.Val
//	}
//	return ""
//}
