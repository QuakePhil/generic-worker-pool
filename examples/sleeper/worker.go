// Package sleeper sleeps in millisecond increments
package sleeper

import (
	"fmt"
	"time"
)

type Work struct{}

type Result struct{}

type Worker struct {
	Milliseconds int // how many to sleep (1000 = 1 second)
}

func (w Worker) Input(in chan Work) {
	for n := 1; n <= w.Milliseconds; n++ {
		in <- Work{}
	}
}

func (w Worker) Process(i Work) (r Result) {
	time.Sleep(time.Millisecond)
	return
}

// TODO: make the generic Output(O) optional in pool.go...
// how to test any(p.w).(interface{Output(O)}) ?
func (w *Worker) Output(r Result) {
}

func (w *Worker) Done() {
	fmt.Println("done")
}
