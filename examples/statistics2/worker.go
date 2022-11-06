package statistics2

import "fmt"

// Input gets sent from Input() to Process()
type Input float64

type w struct {
	in    chan Input
	count int
	total float64
}

func New(in chan Input) (w w) {
	w.in = in
	w.count = 0
	w.total = 0
	return
}

// Input() generates Input by chaining to a custom input channel.
func (w w) Input(in chan Input) {
	for {
		if i, more := <-w.in; !more {
			return
		} else {
			in <- i
		}
	}
}

// Process() consumes Input and produces a Result.
func (w *w) Process(i Input) {
	w.count += 1
	w.total += float64(i)

	fmt.Println("counting", i)
	if w.count%3 == 0 {
		fmt.Println("average so far:", w.total/float64(w.count))
	}
}
