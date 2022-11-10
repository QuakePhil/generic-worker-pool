// Package sleeper sleeps in millisecond increments
package sleeper

import (
	"fmt"
	"time"
)

type State struct{}

type Worker struct {
	Milliseconds int // how many to sleep (1000 = 1 second)
}

func (w Worker) Input(in chan State) {
	for n := 1; n <= w.Milliseconds; n++ {
		in <- State{}
	}
}

func (w Worker) Process(i State) State {
	time.Sleep(time.Millisecond)
	return i
}

func (w Worker) Output(o State) {
}

func (w Worker) Done() {
	fmt.Println("done")
}
