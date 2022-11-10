// Package pool implements a generic worker pool.
package pool

import (
	"sync"
	//"time"
)

type Worker[S any] interface {
	Input(chan S)
	Process(S) S
	Output(S)
	// Optional:
	// Done()
}

type Pool[S any] struct {
	in   chan S
	out  chan S
	done chan bool
	w    Worker[S]
}

// New creates a generic worker pool.
func New[S any](w Worker[S]) (p Pool[S]) {
	// If you get "cannot use ... method has pointer receiver"
	// then try "pool.New(&worker)" instead of "pool.New(worker)"
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
		for o := range p.out {
			p.w.Output(o)
		}
		p.done <- true
	}()

	return
}

// Wait spawns a number of worker processes
// and consumes the shared input channel.
func (p Pool[I]) Wait(concurrency int) {

	wg := &sync.WaitGroup{}
	wg.Add(concurrency)
	for id := 1; id <= concurrency; id++ {
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
	close(p.out)
	<-p.done
}
