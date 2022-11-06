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
There are a couple of demo workers, run them with `go test`:
```
go test -v example_test.go
```
```
=== RUN   ExampleSleeper
--- PASS: ExampleSleeper (1.29s)
=== RUN   ExampleTenSleepers
--- PASS: ExampleTenSleepers (0.13s)
=== RUN   Example
--- PASS: Example (0.00s)
=== RUN   ExampleCustomInputChannel
--- PASS: ExampleCustomInputChannel (0.00s)
PASS
ok  	command-line-arguments	1.518s
```
