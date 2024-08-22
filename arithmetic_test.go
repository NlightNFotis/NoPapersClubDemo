package main

import "testing"

func TestMultiply(t *testing.T) {
	// Bug: This test is failing if b == 12
	got := Multiply(2, 4)
	want := 8

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
