# [Generic](https://gobyexample.com/generics) worker pool [![Go Reference](https://pkg.go.dev/badge/github.com/quakephil/generic-worker-pool.svg)](https://pkg.go.dev/github.com/quakephil/generic-worker-pool)

Go (1.18+) framework to run a pool of `N` workers
```
go get github.com/quakephil/generic-worker-pool@v0.1.5
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
* [`sync.WaitGroup` logic](https://gobyexample.com/waitgroups) to run `Process(State)` concurrently
* [a `chan bool` to synchronize](https://gobyexample.com/channel-synchronization) completion of `Output()`

Examples: https://github.com/QuakePhil/generic-worker-pool-examples
