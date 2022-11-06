# Generic worker pool

Go (1.18+) framework to run a pool of `N` workers

Worker specifies its own:
* `Input` type
* `Input(chan Input)`
* `Process(Input)`
* `Done()` (optional)

Pool handles:
* shared channel for `Input`
* `sync.WaitGroup` logic to run multiple `Process(Input)` concurrently

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
