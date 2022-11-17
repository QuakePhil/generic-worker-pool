// Package pool implements a generic worker pool with shared channels.
package pool

// Worker implements input-specific processing and aggregation.
type Worker[I, O any] interface {
	Process(I) I
	Output(chan I) O
}

// Pool contains channels for the generic worker.
type Pool[I, O any] struct {
	in        chan I
	processed chan I
	out       chan O
	w         Worker[I, O]
}

// New creates the channels and kicks off the input producer and the output consumer.
func New[I, O any](w Worker[I, O], input func(chan I)) (p Pool[I, O]) {
	p.w = w

	// input channel and method
	p.in = make(chan I)
	go func() {
		input(p.in)
		close(p.in)
	}()

	// processing and output channels
	p.processed = make(chan I)
	p.out = make(chan O, 1)
	go func() {
		p.out <- p.w.Output(p.processed)
	}()

	return
}

// Wait kicks off the input consumers and returns the result of output,
// using an intermediary processed channel.
func (p Pool[I, O]) Wait(concurrency int) O {
	limiter := make(chan bool, concurrency)
	for i := range p.in {
		limiter <- true
		go func(i I) {
			p.processed <- p.w.Process(i)
			<-limiter
		}(i)
	}

	for wait := 1; wait <= concurrency; wait++ {
		limiter <- true
	}
	close(p.processed) // safe to close, as only Process() writes here
	return <-p.out     // return the result of Output()
}
