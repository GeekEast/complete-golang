package G17_performance

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func request() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Printf("sleeping for %d seconds...\n", n)
	time.Sleep(time.Duration(n)*time.Second)
	return 1
}

func query() int {
	chan1 := make(chan int)
	var wg  sync.WaitGroup
	wg.Add(3)
	defer func() {
		close(chan1)
	}()
	go func() {
		defer wg.Done()
		chan1 <- request()
	}()
	go func() {
		defer wg.Done()
		chan1 <- request()
	}()
	go func() {
		defer wg.Done()
		chan1 <- request()
	}()
	return <-chan1
}

func TestGoroutineLeaking() {
	println(runtime.NumGoroutine()) // 1 the main routine
	query()
	time.Sleep(10* time.Second)
	println(runtime.NumGoroutine()) // should be 1 but is 3
}
