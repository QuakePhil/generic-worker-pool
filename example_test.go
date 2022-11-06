package main

import (
	"code/generic-worker-pool/examples/sleeper"
	"code/generic-worker-pool/examples/statistics"
	"code/generic-worker-pool/examples/statistics2"
	"code/generic-worker-pool/pool"
)

func ExampleSleeper() {
	worker := sleeper.Worker{1000}
	// 1 sleeper finish in about a second
	pool.New[sleeper.Work, sleeper.Result](&worker).Wait(1)
	// Output: done
}

func ExampleTenSleepers() {
	worker := sleeper.Worker{1000}
	// 10 sleepers finish in about 1/10th of a second
	pool.New[sleeper.Work, sleeper.Result](&worker).Wait(10)
	// Output: done
}

func Example() {
	worker := statistics.New()
	pool.New[statistics.Work, statistics.Result](&worker).Wait(1)
	// Output:
	// counting 17
	// counting 49
	// counting 25
	// average so far: 30.333333333333332
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
	// Output:
	// counting 17
	// counting 49
	// counting 25
	// average so far: 30.333333333333332
}
