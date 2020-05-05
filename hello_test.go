package main

import "testing"

func TestString(t *testing.T) {
	hello := hello()
	if hello != "Hello, World!" {
		t.Errorf("error")
	}
}
