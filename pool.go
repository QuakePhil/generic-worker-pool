// Package pool implements a generic worker pool with shared channels.
package pool

// Pool contains channels for the generic worker() to run at limited concurrency.
type Pool[I, O any] struct {
	in      chan I
	results chan I
	out     chan O
	worker  func(I) I
}

// New creates the channels and kicks off the input producer and the output consumer.
func New[I, O any](
	input func(chan<- I),
	worker func(I) I,
	output func(<-chan I) O,
) (p Pool[I, O]) {

	p.in = make(chan I)
	go func() {
		input(p.in)
		close(p.in)
	}()

	p.out = make(chan O, 1)
	p.worker = worker
	p.results = make(chan I)
	go func() {
		p.out <- output(p.results)
	}()
	return
}

// Wait runs input consumers concurrently and returns the result of output.
func (p Pool[I, O]) Wait(concurrency int) O {
	sem := make(chan bool, concurrency)
	for i := range p.in {
		sem <- true
		go func(i I) {
			if p.worker != nil {
				i = p.worker(i)
			}
			p.results <- i
			<-sem
		}(i)
	}
	for ; concurrency > 0; concurrency-- {
		sem <- true
	}
	close(p.results)
	return <-p.out // return the result of output()
}
