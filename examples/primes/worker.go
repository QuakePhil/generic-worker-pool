// Package primes counts how many primes there are in batches of "step" size
package primes

import "fmt"

// State gets sent from Input() to Process()
type State struct {
	Start  int
	End    int
	Primes int
}

// worker state types
type worker struct {
	start                int
	end                  int
	step                 int
	optionalInputChannel chan State
}

// New worker state (could also use new() or a struct literal)
func New(start, end, step int) (w worker) {
	w.start = start
	w.end = end
	w.step = step
	return
}

func NewWithChannel(in chan State) (w worker) {
	w.optionalInputChannel = in
	return
}

// Input() generates State, e.g. reading from SQL, etc.
func (w worker) Input(in chan State) {
	if w.optionalInputChannel != nil {
		// better way to chain channels? maybe in <- <-w.optionalInputChannel
		for {
			if i, more := <-w.optionalInputChannel; !more {
				return
			} else {
				in <- i
			}
		}
	} else {
		for i := w.start; i <= w.end; i += w.step {
			in <- State{i, i + w.step, 0}
		}
	}
}

// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
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
