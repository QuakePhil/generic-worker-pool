# [Generic](https://gobyexample.com/generics) worker pool [![Go Reference](https://pkg.go.dev/badge/github.com/quakephil/generic-worker-pool.svg)](https://pkg.go.dev/github.com/quakephil/generic-worker-pool)

Go (1.18+) framework to run a pool of `N` workers
```
go get github.com/quakephil/generic-worker-pool@v0.2.0
```

## Pool
* [shared input/output channels](https://gobyexample.com/worker-pools)
```
workers := pool.New[I, O](Input, Worker{ ... })
```
* [channel synchronization](https://gobyexample.com/channel-synchronization) for
  * concurrent runs of `Process()`
  * completion of `Output()`
```
result := workers.Wait(concurrency)
```

## Input
* `func Input(in chan I)`
```
for ... {
  in <- I{ ... }
}
```

## Worker
* `func (w Worker) Process(i I) I`
```
i.update = ...
return i
```
* `func (w Worker) Output(processed chan I) (out O)`
```
for o := range processed {
  out.update += ...
}
return
```

Examples: https://github.com/QuakePhil/generic-worker-pool-examples
