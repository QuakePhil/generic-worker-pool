package statistics

import "fmt"

// Input gets sent from Input() to Process()
type State float64

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

// Input() generates Input, e.g. reading from SQL, etc.
func (w worker) Input(in chan State) {
	in <- 17
	in <- 49
	in <- 25
}

// Process() consumes Input and produces a Result.
func (w *worker) Process(i State) State {
	w.count += 1
	w.total += float64(i)
	return State(w.total / float64(w.count))
}

// Note: w.state can be changed from start of Output() to finish, by concurrent Process()
// use o. instead of w.
func (w worker) Output(o State) {
	fmt.Println("average so far:", o)
}
