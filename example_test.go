package main

import (
	"code/generic-worker-pool/examples/average"
	"code/generic-worker-pool/examples/sleep"
	"code/generic-worker-pool/pool"
)

func ExampleSleeper() {
	// 1 sleeper finish in about a second
	pool.New[sleeper.State](sleeper.Worker{10}).Wait(1)
	// Output: done
}

func ExampleTenSleepers() {
	// 10 sleepers finish in about 1/10th of a second
	pool.New[sleeper.State](sleeper.Worker{10}).Wait(10)
	// Output: done
}

func Example() {
	worker := average.New()
	pool.New[average.State](&worker).Wait(1)
	// Output:
	// average so far: 17
	// average so far: 33
	// average so far: 30.333333333333332
}

func ExampleCustomStateChannel() {
	in := make(chan average.State)
	worker := average.NewWithChannel(in)
	go func() {
		in <- 17
		in <- 49
		in <- 25
		close(in)
	}()

	pool.New[average.State](&worker).Wait(1)
	// Output:
	// average so far: 17
	// average so far: 33
	// average so far: 30.333333333333332
}
