# [Generic](https://gobyexample.com/generics) worker pool [![Go Reference](https://pkg.go.dev/badge/github.com/quakephil/generic-worker-pool.svg)](https://pkg.go.dev/github.com/quakephil/generic-worker-pool)

Go (1.18+) framework to run a pool of `N` workers
```
go get github.com/quakephil/generic-worker-pool@v0.1.7
```

## Worker
* `State` type
  * used for input and output
* `Input(in chan State)`
```
for ... {
  in <- State{ ... }
}
```
* `Process(i State) State`
```
i.update = ...
return i
```
* `Output(out chan State)`
```
for o := range out {
  fmt.Println(o)
}
```
* `Done()` (optional)

## Pool
* [shared input/output channels](https://gobyexample.com/worker-pools) for `State`
* [`chan bool` synchronization](https://gobyexample.com/channel-synchronization) for
  * concurrent runs of `Process()`
  * completion of `Output()`

Examples: https://github.com/QuakePhil/generic-worker-pool-examples
