// Package pool implements a generic worker pool with shared input and output channels.
package pool

// Worker is an interface to another package that implements a specific process.
type Worker[S any] interface {
	Input(chan S)
	Process(S) S
	Output(chan S)
	// Optional:
	// Done()
}

// Pool contains channels for the generic type, as well as a done channel
// to signal when output is complete.
type Pool[S any] struct {
	in   chan S
	out  chan S
	done chan bool
	w    Worker[S]
}

// New creates the channels and kicks off the Input() and Output() methods.
func New[S any](w Worker[S]) (p Pool[S]) {
	p.w = w

	// input channel and method
	p.in = make(chan S)
	go func() {
		p.w.Input(p.in)
		close(p.in)
	}()

	// output channel and method
	p.out = make(chan S)
	p.done = make(chan bool)
	go func() {
		p.w.Output(p.out)
		p.done <- true // signal output is finished
	}()

	return
}

// Wait spawns a number of worker processes and consumes the shared input channel.
func (p Pool[S]) Wait(concurrency int) {
	processed := make(chan bool, concurrency)
	for i := range p.in {
		processed <- true
		go func(i S) {
			p.out <- p.w.Process(i)
			<-processed
		}(i)
	}

	for wait := 1; wait <= concurrency; wait++ {
		processed <- true
	}
	close(p.out) // safe to close, as only Process() writes here

	// optional Done method
	if w, ok := p.w.(interface{ Done() }); ok {
		w.Done()
	}
	<-p.done // wait until all of output has been consumed
}
