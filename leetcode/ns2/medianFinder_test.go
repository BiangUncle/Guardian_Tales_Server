package ns2

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	m := Constructor2()
	m.AddNum(1)
	println(m.FindMedian())
	m.AddNum(2)
	println(m.FindMedian())
	m.AddNum(5)
	println(m.FindMedian())
}

func TestMedianFinderExampleCreator(t *testing.T) {
	MedianFinderExampleCreator(1000)
}
