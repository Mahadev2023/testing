package main

import "testing"

func TestSquare(t *testing.T) {
	x := square(5)

	if x != 25 {
		t.Error("Expected ", 25, "Got", x)
	}
}
