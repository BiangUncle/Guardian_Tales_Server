package ns2

import (
	"fmt"
	"math/rand"
	"server/leetcode/data_struct"
	"time"
)

type MedianFinder struct {
	MinHeap *data_struct.MinHeap
	MaxHeap *data_struct.MaxHeap
}

func Constructor2() MedianFinder {
	m := MedianFinder{}
	m.MaxHeap = new(data_struct.MaxHeap)
	m.MinHeap = new(data_struct.MinHeap)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.MaxHeap.List) == 0 || num <= this.MaxHeap.List[0] {
		this.MaxHeap.Insert(num)
		if len(this.MaxHeap.List)-len(this.MinHeap.List) > 1 {
			top := this.MaxHeap.Pop()
			this.MinHeap.Insert(top)
		}
		return
	}
	this.MinHeap.Insert(num)
	this.Balance()
	if len(this.MinHeap.List) > len(this.MaxHeap.List) {
		top := this.MaxHeap.Pop()
		this.MinHeap.Insert(top)
	}
}

func (this *MedianFinder) Balance() {
	for {
		if len(this.MaxHeap.List)-len(this.MinHeap.List) > 1 {
			num := this.MaxHeap.Pop()
			this.MinHeap.Insert(num)
			continue
		}
		if len(this.MinHeap.List)-len(this.MaxHeap.List) > 0 {
			num := this.MinHeap.Pop()
			this.MaxHeap.Insert(num)
			continue
		}
		break
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (len(this.MaxHeap.List)+len(this.MinHeap.List))%2 == 0 {
		return (float64(this.MaxHeap.List[0]) + float64(this.MinHeap.List[0])) / 2
	}
	return float64(this.MaxHeap.List[0])
}

func MedianFinderExampleCreator(n int) {
	op := []string{"MedianFinder"}
	num := []int{-1}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		r := rand.Intn(10)
		if r < 7 {
			op = append(op, "addNum")
			num = append(num, rand.Intn(1000))
		} else {
			op = append(op, "findMedian")
			num = append(num, -1)
		}
	}

	op_str := "["
	num_str := "["
	for i := 0; i < len(op); i++ {
		op_str = op_str + fmt.Sprintf("\"%s\",", op[i])
		if num[i] == -1 {
			num_str = num_str + "[],"
		} else {
			num_str = num_str + fmt.Sprintf("[%d],", num[i])
		}
	}
	op_str = op_str[:len(op_str)-1] + "]"
	num_str = num_str[:len(num_str)-1] + "]"
	fmt.Println(op_str)
	fmt.Println(num_str)
}
