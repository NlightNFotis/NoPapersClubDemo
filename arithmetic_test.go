package main

import "testing"

func TestMultiply(t *testing.T) {

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 4, 8},
		// Bug: This test is failing if b == 12
		{2, 12, 24},
	}

	for _, test := range tests {
		got := Multiply(test.a, test.b)
		if got != test.expected {
			t.Errorf("got %d want %d", got, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	got := Add(3, 2)
	expected := 5

	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}
}
