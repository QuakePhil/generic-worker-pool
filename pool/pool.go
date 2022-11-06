// Package pool implements a generic worker pool.
package pool

import "sync"

//import "log"

type Worker[I any] interface {
	Input(chan I)
	Process(I)
	// Optional:
	// Done()
}

type Pool[I any] struct {
	in chan I
	w  Worker[I]
}

// New creates a generic worker pool.
func New[I any](w Worker[I]) (p Pool[I]) {
	p.w = w

	// input channel and method
	p.in = make(chan I)
	go func() {
		p.w.Input(p.in)
		close(p.in)
	}()
	return
}

// Wait spawns a number of worker processes
// and consumes the shared input channel.
func (p Pool[I]) Wait(num int) {
	var wg sync.WaitGroup

	for id := 1; id <= num; id++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range p.in {
				p.w.Process(i)
			}
		}()
	}
	wg.Wait()

	// optional Done method
	if w, ok := p.w.(interface{ Done() }); ok {
		w.Done()
	}
}
