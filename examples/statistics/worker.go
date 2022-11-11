package statistics

import "fmt"

// Input gets sent from Input() to Process()
type State float64

// worker state types
type worker struct{}

// New worker state (could also use new() or a struct literal)
func New() (w worker) {
	return
}

// Input() generates Input, e.g. reading from SQL, etc.
func (w worker) Input(in chan State) {
	in <- 17
	in <- 49
	in <- 25
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
