package worker

import (
	"code/worker/pool"
	"fmt"
)

func Example() {
	worker := NewStatisticsWorker([]float64{17, 49, 25})
	pool.New[Work, Result](&worker).Wait(1)
	fmt.Println(worker.average)
	// Output:
	// counting 17
	// counting 49
	// counting 25
	// 30.333333333333332
}
