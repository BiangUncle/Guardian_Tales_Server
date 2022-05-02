package logger

import "testing"

func TestInit(t *testing.T) {
	Init()

	Info("Hello")

	t.Log(&buf)

	Info("Hello * 2")

	t.Log(&buf)
}
