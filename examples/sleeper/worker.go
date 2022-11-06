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

// TODO: make the generic Output(O) optional in pool.go...
// how to test any(p.w).(interface{Output(O)}) ..?
// maybe factor out Input() into examples,
// and use w.Input and w.Output for state instead of func()
// and test for nil before sending to output
// and maybe replace Work/Result with InputType/OutputType
//func (w *Worker) Output(r Result) {
//}

func (w *Worker) Done() {
	fmt.Println("done")
}
