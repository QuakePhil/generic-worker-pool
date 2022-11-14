package pool

import (
	"code/generic-worker-pool/examples/primes"
	"code/generic-worker-pool/examples/sleep"
)

// 1 sleeper finish in about a second
func ExampleSleeper() {
	NewPool[sleeper.State](sleeper.Worker{10}).Wait(1)
	// Output: done
}

// 10 sleepers finish in about 1/10th of a second
func ExampleTenSleepers() {
	NewPool[sleeper.State](sleeper.Worker{10}).Wait(10)
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
