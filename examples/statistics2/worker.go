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

// worker state types
type worker struct {
	in    chan Work
	count int
	total float64
}

// New worker state (could also use new() or a struct literal)
func New(in chan Work) (w worker) {
	w.in = in
	w.count = 0
	w.total = 0
	return
}

// Input() generates Work by chaining to a custom input channel.
func (w worker) Input(in chan Work) {
	for {
		if i, more := <-w.in; !more {
			return
		} else {
			in <- i
		}
	}
}

// Process() consumes Work and produces a Result.
func (w *worker) Process(i Work) Result {
	w.count += 1
	w.total += i.Number
	return Result{err: nil, unit: i}
}

// Output() consumes Results and logs, panics, etc. as appropriate.
func (w worker) Output(r Result) {
	fmt.Println("counting", r.unit.Number)
	if w.count%3 == 0 {
		fmt.Println("average so far:", w.total/float64(w.count))
	}
}
