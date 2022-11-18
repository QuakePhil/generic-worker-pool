package pool

import "testing"

// TestPool counts how many bools input() got.
func TestPool(t *testing.T) {
	var result int

	result = New[bool, int](input, nil, output).Wait(10)
	if result != 0 {
		t.Fatalf("Expected 0, got %d", result)
	}

	result = New[bool, int](input, func(_ bool) bool { return true }, output).Wait(10)
	if result != 3 {
		t.Fatalf("Expected 3, got %d", result)
	}
}

func input(in chan<- bool) {
	in <- false
	in <- false
	in <- false
}

func output(in <-chan bool) (count int) {
	for result := range in {
		if result {
			count++
		}
	}
	return
}
