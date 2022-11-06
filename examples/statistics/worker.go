package statistics

import "fmt"

// Input gets sent from Input() to Process()
type Input struct {
	number float64
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

// Input() generates Input, e.g. reading from SQL, etc.
func (w worker) Input(in chan Input) {
	in <- Input{number: 17}
	in <- Input{number: 49}
	in <- Input{number: 25}
}

// Process() consumes Input and produces a Result.
func (w *worker) Process(i Input) {
	w.count += 1
	w.total += i.number

	fmt.Println("counting", i.number)
	if w.count%3 == 0 {
		fmt.Println("average so far:", w.total/float64(w.count))
	}
}
