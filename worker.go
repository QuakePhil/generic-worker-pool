package worker

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

// Worker state types
type StatisticsWorker struct {
	stuff   []float64
	count   int
	total   float64
	average float64
}

// New worker state (could also use new() or a struct literal)
func NewStatisticsWorker(stuff []float64) (w StatisticsWorker) {
	w.stuff = stuff
	w.count = 0
	w.total = 0
	return
}

// Input() generates Work, e.g. reading from SQL, etc.
func (w StatisticsWorker) Input(in chan Work) {
	for i := range w.stuff {
		in <- Work{number: w.stuff[i]}
	}
}

// Process() consumes Work and produces a Result.
func (w *StatisticsWorker) Process(i Work) Result {
	w.count += 1
	w.total += i.number
	return Result{err: nil, unit: i}
}

// Output() consumes Results and logs, panics, etc. as appropriate.
func (w StatisticsWorker) Output(r Result) {
	fmt.Println("counting", r.unit.number)
}

// Done() runs after the last Process() is done.
func (w *StatisticsWorker) Done() {
	w.average = w.total / float64(w.count)
}
