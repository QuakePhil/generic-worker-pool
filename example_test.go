package pool

import (
	"github.com/quakephil/generic-worker-pool/examples/primes"
	"github.com/quakephil/generic-worker-pool/examples/sleep"
)

// 10 units of sleep = 1 second.
// 1 worker takes 1 second
func ExampleSleeper() {
	NewPool[sleep.State](sleep.Worker{10}).Wait(1)
	// Output: done
}

// 10 workers take 1/10th of a second
func ExampleTenSleepers() {
	NewPool[sleep.State](sleep.Worker{10}).Wait(10)
	// Output: done
}

func ExamplePrimes() {
	NewPool[primes.State](primes.New(20000000, 100000)).Wait(1)
	// Output: Found 1270607 primes
}

func ExampleManyPrimes() {
	NewPool[primes.State](primes.New(20000000, 100000)).Wait(1000)
	// Output: Found 1270607 primes
}

func ExampleCustomInputChannel() {
	worker := primes.NewWithChannel()
	go func() {
		for i := 1; i <= 20000000; i += 100000 {
			worker.In <- primes.State{i, i + 100000, 0}
		}
		close(worker.In)
	}()
	NewPool[primes.State](worker).Wait(1000)
	// Output: Found 1270607 primes
}
