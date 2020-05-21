## Theory
### Occurence
- **Happens Before**
  - `x` finishes **before** `y` starts
  - `D1`,`D2`,`D3` finishes before `y` need them to start
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-21-01-24-45.png alt="no image found"></p>

- **Happens Concurrently**: 
  - Cannot **guarantee** the **order** of occurence of `x` and `y`

## Channel
- **Synchronize**: unbufferred channel can used to synchronize goroutines, which essentially is **making concurrency sequential**
- **Decouple**: bufferred channel can be used to decouple `sender` and `receiver` goroutine, avoiding **goroutine leaking**
- **Data transfer**: transfer data between goroutines

### Operation
#### Create
> happens outside either `sender` or `receiver`
```go
ch1 := make(chan int)    // buffer as 0, default value
ch1 := make(chan int, 3) // buffer as 3
```
- unidirectional channel (check type duing compile time)
  - `uni` can be **converted** to `bi` automatically but **not** vice versa
```go
ping := make(chan<- int) // you can only send data to ping
pong := make(<-chan int) // you can only retrieve data from pong
```

#### Send
- **only one** sender is working at one time
```go
ch <- x
```
- <div style="color:blue">send to closed channel will cause panic (exception)</div>

```go
func main() {
    closed := make(chan int)
    close(closed)
    closed <- 1 // panic!
}
```

#### Close
- close channel only when you need to broadcast a signal
```go
close(ch)
```
- In most cases, you don't need to close channels bease GC will cycle them automatically
- <div style="color:blue">try to close a closed channel will cause panic</div>
- <div style="color:blue">try to close nil channel will cause panic</div>


#### Receive
- **only one** receiver is working at one time
```go
x := <- ch       // declare and assignment
x = <- ch        // just assignment
<- ch            // result is discared
```
- <div style="color:blue">retrieve value from closed channel will get zero type of channel element</div>

```go
ch := make(chan int)
close(ch)
value,ok := <-ch
fmt.Println(value,ok) // 0, false
```

#### Drain
- clean channel buffer after closing: `for...break`
```go
func LoopOverChannel() {
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
    close(queue)
    
	for {
		if elem,ok := <- queue; ok {
			println(elem)
		} else {
			break
		}
	}
}
```
- clean channel buffer after closing: `for...range`
```go
func RangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
```

### Types
#### Unbufferred Channel
- **create**
```go
// no buffer
ch := make(chan int)
```
- **blocking rule**:
  - no `sender` **OR** no `receiver`
- **sync single goroutine**
```go
func main() {
    ch := make(chan struct{}) // save memory

    go func() {
        ch <- struct{}{} // send finishing signal
    }()

    <-ch // block until receive finishing signal
}
```
- **sync multiple goroutines**
```go
func main() {
    ch := make(chan struct{}) 
    go func() {
        ch <- struct{}{}
    }()

    go func() {
        ch <- struct{}{}
    }()
    <-ch
    <-ch
}
```

#### Bufferred Channels
- **create**
```go
ch1 := make(chan int, 3) // 3 int cache
```
- **inspect buffer**
```go
cap(ch1) // get the cache size
len(ch1) // get the number of elements cached
```
- **blocking rules**:
  - the `sender` will be **blocked** when the channel is `full`
  - the `receiver` will be **blocked** when the channel is `empty`
- **Situation 1**: continues data transferring
  - useful to **streamline** data transferring when `sending speed >= receiving speed`
  - **not** useful when `sending speed < receiving speed`
- **Situation 2**: limit number of sending or receiving
  - `S <= R` is always GC **safe**
  - `S > R` might cause **Goroutine Leaking**
    - Golden Rule: `buffer size >= (S-R)` 
```go
func query() int {
    // declouple sender goroutine and receiver goroutine
	bc := make(chan int, 1)
	go func() {
		bc <- 1
	}()
	go func() {
		bc <- 2
	}()
	return <- bc
}
```


#### Pipelines
- Channels can be used to connect goroutines together so that the **output of one is the input to another**.
- This is about **where do you place** the channel code within a goroutine functuon
```go
func PipelineSimple() {
	naturals := make(chan int)
	squares := make(chan int)
	// generate 0,1,... and send to channel - naturals
	go func() {
		for i:=0;;i++ {
			naturals <- i
		}
	}()

	// receive 0,1,... from naturals channel, square and send to square channel
	go func() {
		for {
			x :=<- naturals
			squares <- x * x
		}
	}()

	for {
		println(<- squares)
	}
}
```
- more
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-21-01-53-40.png alt="no image found"></p>