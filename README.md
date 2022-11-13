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

Pool handles:
* shared input/output channels for `State`
* `sync.WaitGroup` logic to run `Process(State)` concurrently
* a `chan bool` to synchronize with completion of output

## Example
Check out some demos:
```
go test -v -race example_test.go
```
```
=== RUN   ExampleSleeper
--- PASS: ExampleSleeper (1.02s)
=== RUN   ExampleTenSleepers
--- PASS: ExampleTenSleepers (0.10s)
=== RUN   ExamplePrimes
--- PASS: ExamplePrimes (1.39s)
=== RUN   ExampleManyPrimes
--- PASS: ExampleManyPrimes (0.32s)
=== RUN   ExampleCustomStateChannel
--- PASS: ExampleCustomStateChannel (0.34s)
PASS
ok  	command-line-arguments	3.406s
```
