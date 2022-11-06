package statistics2

import "fmt"

// Input gets sent from Input() to Process()
type Input struct {
	Number float64
}

// worker state types
type worker struct {
	in    chan Input
	count int
	total float64
}

// New worker state (could also use new() or a struct literal)
func New(in chan Input) (w worker) {
	w.in = in
	w.count = 0
	w.total = 0
	return
}

// Input() generates Input by chaining to a custom input channel.
func (w worker) Input(in chan Input) {
	for {
		if i, more := <-w.in; !more {
			return
		} else {
			in <- i
		}
	}
}

// Process() consumes Input and produces a Result.
func (w *worker) Process(i Input) {
	w.count += 1
	w.total += i.Number

	fmt.Println("counting", i.Number)
	if w.count%3 == 0 {
		fmt.Println("average so far:", w.total/float64(w.count))
	}
}
