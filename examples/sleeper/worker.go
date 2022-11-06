// Package sleeper sleeps in millisecond increments
package sleeper

import (
	"fmt"
	"time"
)

type Input struct{}

type Worker struct {
	Milliseconds int // how many to sleep (1000 = 1 second)
}

func (w Worker) Input(in chan Input) {
	for n := 1; n <= w.Milliseconds; n++ {
		in <- Input{}
	}
}

func (w Worker) Process(i Input) {
	time.Sleep(time.Millisecond)
}

func (w Worker) Done() {
	fmt.Println("done")
}
