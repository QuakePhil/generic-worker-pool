// Package pool implements a generic worker pool.
package pool

import "sync"

//import "log"

type Worker[I, O any] interface {
	Input(chan I)
	Process(I) O
	Output(O)
	// Optional:
	// Done()
}

type Pool[I, O any] struct {
	in  chan I
	out chan O
	w   Worker[I, O]
}

// New creates a generic worker pool.
func New[I, O any](w Worker[I, O]) (p Pool[I, O]) {
	p.w = w

	// input channel and method
	p.in = make(chan I)
	go func() {
		p.w.Input(p.in)
		close(p.in)
	}()

	// output channel and method
	p.out = make(chan O)
	go func() {
		for r := range p.out {
			p.w.Output(r)
		}
	}()
	return
}

// Wait spawns a number of worker processes
// and consumes the shared input channel.
func (p Pool[I, O]) Wait(num int) {
	var wg sync.WaitGroup

	for id := 1; id <= num; id++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range p.in {
				p.out <- p.w.Process(i)
			}
		}()
	}
	wg.Wait()

	// optional Done method
	if w, ok := p.w.(interface{ Done() }); ok {
		w.Done()
	}
}
