// Package primes counts how many primes there are in batches of "step" size
package primes

import "fmt"

// State gets sent from Input() to Process() to Output()
type State struct {
	Start  int
	End    int
	Primes int
}

// worker has shared state e.g. handles, connections, settings, etc
type worker struct {
	max  int
	step int
	In   chan State
}

// New worker state (could also use new() or a struct literal)
func New(max, step int) (w worker) {
	w.max = max
	w.step = step
	return
}

func NewWithChannel() (w worker) {
	w.In = make(chan State)
	return
}

// Input() generates State, e.g. reading from SQL, etc.
func (w worker) Input(in chan State) {
	if w.In != nil {
		// better way to chain channels? maybe in <- <-w.In
		for {
			if i, more := <-w.In; !more {
				return
			} else {
				in <- i
			}
		}
	} else {
		for i := 1; i <= w.max; i += w.step {
			in <- State{i, i + w.step, 0}
		}
	}
}

// Process() works on State, counting how many primes there are in a given range
func (w worker) Process(i State) State {
	for j := i.Start; j < i.End; j++ {
		if isPrime(j) {
			i.Primes += 1
		}
	}
	return i
}

// Output() works on State results from Process(), computing a total sum of primes
func (w worker) Output(out chan State) {
	count := 0
	for o := range out {
		count += o.Primes
	}
	fmt.Printf("Found %d primes\n", count)
}
