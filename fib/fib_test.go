package fib

import "testing"

func TestFib(t *testing.T) {
	input := 0
	got := Fib(input)
	want := 0
	if got != want {
		t.Errorf("Fib(%d): want %d, got %d", input, want, got)
	}

}
