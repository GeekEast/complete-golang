package G18_parallel_loop

import "sync"

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

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
