package client

import "testing"

func TestClient_Register(t *testing.T) {
	err := NewClient().Register()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Get(t *testing.T) {
	err := NewClient().GetServe()
	if err != nil {
		t.Error(err)
	}
}
