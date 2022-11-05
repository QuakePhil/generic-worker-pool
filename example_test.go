package main

import (
	"code/worker/pool"
	"code/worker/statistics"
	"code/worker/statistics2"
	"fmt"
)

func Example() {
	worker := statistics.New()
	pool.New[statistics.Work, statistics.Result](&worker).Wait(1)
	fmt.Println(worker.Average)
	// Output:
	// counting 17
	// counting 49
	// counting 25
	// 30.333333333333332
}

func ExampleCustomInputChannel() {
	in := make(chan statistics2.Work)
	go func() {
		in <- statistics2.Work{17}
		in <- statistics2.Work{49}
		in <- statistics2.Work{25}
		close(in)
	}()
	worker := statistics2.New(in)
	pool.New[statistics2.Work, statistics2.Result](&worker).Wait(1)
	fmt.Println(worker.Average)
	// Output:
	// counting 17
	// counting 49
	// counting 25
	// 30.333333333333332
}
