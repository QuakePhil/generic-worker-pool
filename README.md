# Generic worker pool

Go (1.18+) framework to run a pool of `N` workers

Worker specifies its own:
* `Work` and `Result` types
* `Input(chan Work)`
* `Process(Work) Result`
* `Output(Result)`
* `Done()`

Pool handles:
* channels for `Work` and `Result`
* `sync.WaitGroup` logic to run `N` goroutines

## Example
There's a demonstration `statistics` worker that calculates an average, run it with `go test`
```
go test example_test.go
```
