package statistics2

import "fmt"

// Input gets sent from Input() to Process()
type State float64

// worker state types
type worker struct {
	in    chan State
	count int
	total float64
}

func New(in chan State) (w worker) {
	w.in = in
	w.count = 0
	w.total = 0
	return
}

// Input() generates Input by chaining to a custom input channel.
func (w worker) Input(in chan State) {
	for {
		if i, more := <-w.in; !more {
			return
		} else {
			in <- i
		}
	}
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
