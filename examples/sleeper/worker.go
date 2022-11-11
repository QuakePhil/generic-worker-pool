// Package sleeper sleeps 100 milliseconds at a time
package sleeper

import (
	"fmt"
	"time"
)

type State struct{}

type Worker struct {
	Units int // how much input to generate
}

func (w Worker) Input(in chan State) {
	for n := 1; n <= w.Units; n++ {
		in <- State{}
	}
}

func (w Worker) Process(i State) State {
	time.Sleep(100 * time.Millisecond) // placeholder for busywork
	return i
}

func (w Worker) Output(out chan State) {
	for _ = range out {
	}
}

func (w Worker) Done() {
	fmt.Println("done")
}
