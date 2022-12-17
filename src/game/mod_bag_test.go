package game

import (
	"os"
	"testing"
)

func init() {
	err := os.Chdir("/Users/biang/Desktop/github.com/BiangUncle/Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
}

func TestModBag_Save(t *testing.T) {
	m := NewBag()
	err := m.Save()
	if err != nil {
		t.Error(err)
	}
}
