# Generic worker pool

Go (1.18+) framework to run a pool of `N` workers

Worker specifies its own:
* `State` type
  * used for input and output
* `Input(in chan State)`
```
for ... {
  in <- State{ ... }
}
```
* `Process(i *State) State`
```
i.update = ...
return i
```
* `Output(o State)`
```
fmt.Println(o)
```
* `Done()` (optional)

Pool handles:
* shared input/output channels for `State`
* `sync.WaitGroup` logic to run `Process(State)` concurrently
* a `chan bool` to synchronize with completion of output

## Example
Check out some demos:
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
