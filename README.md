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

Check out `example_test.go` and run it with `go test` and friends:
```
go test -coverprofile=/tmp/c.out && go tool cover -func=/tmp/c.out
```
