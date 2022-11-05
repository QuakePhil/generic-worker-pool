package statistics2

import "fmt"

// Work gets sent from Input() to Process()
type Work struct {
	Number float64
}

// Result gets sent from Process() to Output()
type Result struct {
	err  error
	unit Work
}

// Worker state types
type Worker struct {
	in      chan Work
	count   int
	total   float64
	Average float64
}

// New worker state (could also use new() or a struct literal)
func New(in chan Work) (w Worker) {
	w.in = in
	w.count = 0
	w.total = 0
	return
}

// Input() generates Work by chaining to a custom input channel.
func (w Worker) Input(in chan Work) {
	for {
		if i, more := <-w.in; !more {
			return
		} else {
			in <- i
		}
	}
}

// Process() consumes Work and produces a Result.
func (w *Worker) Process(i Work) Result {
	w.count += 1
	w.total += i.Number
	return Result{err: nil, unit: i}
}

// Output() consumes Results and logs, panics, etc. as appropriate.
func (w Worker) Output(r Result) {
	fmt.Println("counting", r.unit.Number)
}

// Done() runs after the last Process() is done.
func (w *Worker) Done() {
	w.Average = w.total / float64(w.count)
}
