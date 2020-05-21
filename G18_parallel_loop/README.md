### Sync using unbuffered channel
```go
func EmbarrassingConcurrencyChannelSync() {
	lock := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fib(i)
			lock <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-lock
	}
	// then
}
```

### Sync using `sync.awaitGroup`
- `Done()` equals to `Add(-1)`
- `Add() -> Done() -> Wait()` have strict order of execution, Don't parallized them!
```go
func EmbarrassingConcurrencyAwaitGroupSync() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		wg.Add(1) // don't put this inside go func 👇
		go func(i int) {
			defer wg.Done()
			fib(i)
		}(i)
	}

	wg.Wait()
	// then
}
```

### What if you want to sum ?!
- Feel the power of go concurrency pattern
```go
func EmbarrassingConcurrencySyncAndSum() {
	chan1 := make(chan int)
	for i := 1; i < 10; i++ {
		go func(i int) {
			chan1 <- fib(i)
		}(i)
	}
	sum := 0
	for i := 1; i < 10; i++ {
		sum += <-chan1
	}
	println(sum)
}
```