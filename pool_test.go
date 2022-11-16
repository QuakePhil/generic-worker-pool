package pool

import "testing"

// TestPool counts how many bool Input() got.
func TestPool(t *testing.T) {
	worker := test{make(chan int, 10)}
	New[bool](worker).Wait(1)
	finalResult := <-worker.result
	if finalResult != 3 {
		t.Fatalf("Expected 3, got %d", finalResult)
	}
}

type test struct {
	result chan int
}

func (t test) Input(in chan bool) {
	in <- true
	in <- true
	in <- true
}

func (t test) Process(i bool) bool {
	return i
}

func (t test) Output(out chan bool) {
	count := 0
	for _ = range out {
		count++
	}
	t.result <- count
}

func (t test) Done() {
}
