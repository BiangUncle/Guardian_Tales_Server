package ns0

import (
	"fmt"
	"testing"
	"time"
)

type Build struct {
}

func (b *Build) PrintName() {
	fmt.Println("Build")
}

type LandMarkBuild interface {
	GetFunction()
}

type Inn struct {
	Build
}

func (i *Inn) GetFunction() {
	fmt.Println("Function")
}

func TestInterface(t *testing.T) {

}

func TestTime(t *testing.T) {
	now := time.Now().Unix()
	sd, _ := time.ParseDuration("-24h")

	past := time.Now().Add(sd)

	// 24 * 60 * 60
	remain := now - past.Unix() - 20
	var hour = remain / (60 * 60)
	var minus = remain / 60 % 60
	var second = remain % 60
	fmt.Println(fmt.Sprintf("%d:%d:%d", hour, minus, second))
}
