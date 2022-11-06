package statistics

import "fmt"

// Work gets sent from Input() to Process()
type Work struct {
	number float64
}

// Result gets sent from Process() to Output()
type Result struct {
	err  error
	unit Work
}

// worker state types
type worker struct {
	count int
	total float64
}

// New worker state (could also use new() or a struct literal)
func New() (w worker) {
	w.count = 0
	w.total = 0
	return
}

var database []float64 = []float64{17, 49, 25}

// Input() generates Work, e.g. reading from SQL, etc.
func (w worker) Input(in chan Work) {
	for i := range database {
		in <- Work{number: database[i]}
	}
}

// Process() consumes Work and produces a Result.
func (w *worker) Process(i Work) Result {
	w.count += 1
	w.total += i.number
	return Result{err: nil, unit: i}
}

// Output() consumes Results and logs, panics, etc. as appropriate.
func (w worker) Output(r Result) {
	fmt.Println("counting", r.unit.number)
}

// Done() runs after the last Process() is done.
func (w *worker) Done() {
	fmt.Println(w.total / float64(w.count))
}
