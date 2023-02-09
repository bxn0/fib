package fib

import "testing"

func TestFib(t *testing.T) {
	input = 1
	got := Fib(input)
	want := 1
	if got != want {
		t.Errorf("Fib(%d): want %d, got %d", input, want, got)
	}
}
