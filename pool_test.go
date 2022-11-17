package pool

import "testing"

// TestPool counts how many bools input() got.
func TestPool(t *testing.T) {
	result := New[bool, int](test{}, input).Wait(10)
	if result != 3 {
		t.Fatalf("Expected 3, got %d", result)
	}
}

// Input
func input(in chan bool) {
	in <- true
	in <- true
	in <- true
}

// Worker
type test struct{}

func (t test) Process(i bool) bool {
	return i
}

func (t test) Output(processed chan bool) (count int) {
	for _ = range processed {
		count++
	}
	return
}
