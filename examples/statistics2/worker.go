package statistics2

import "fmt"

// Input gets sent from Input() to Process()
type State float64

// worker state types
type worker struct {
	in chan State
}

func New(in chan State) (w worker) {
	w.in = in
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
func (w worker) Process(i State) State {
	return i
}

func (w worker) Output(out chan State) {
	count := 0
	total := 0.0
	for o := range out {
		count += 1
		total += float64(o)
		fmt.Println("average so far:", total/float64(count))
	}
}
