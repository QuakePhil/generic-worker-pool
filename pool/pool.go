package pool

import "sync"

type Worker[I, O any] interface {
	Input(chan I)
	Process(I) O
	Output(O)
	Done()
}

type Pool[I, O any] struct {
	in  chan I
	out chan O
	w   Worker[I, O]
}

// New creates a generic worker pool.
func New[I, O any](w Worker[I, O]) (p Pool[I, O]) {
	p.in = make(chan I)
	p.out = make(chan O)
	p.w = w
	go func() {
		p.w.Input(p.in)
		close(p.in)
	}()
	go func() {
		for r := range p.out {
			p.w.Output(r)
		}
	}()
	return
}

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
	p.w.Done()
}
