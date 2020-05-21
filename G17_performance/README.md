<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Performance](#performance)
  - [Goroutine Leaking](#goroutine-leaking)
    - [Try and Error](#try-and-error)
    - [Solution 1: drain channel](#solution-1-drain-channel)
    - [Solution 2: bufferred channel](#solution-2-bufferred-channel)
    - [Solution 3: select and broadcast cancellation](#solution-3-select-and-broadcast-cancellation)
  - [Rules](#rules)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Performance
### Goroutine Leaking
- **Reason**: `blocked` goroutine won't be **cycled** by the GC automatically
- **Solution**: **resolve** the `blocking` status of goroutines.
- This following example causes **2** goroutine leaking:
<!-- TODO: how channel is cycled in golang? -->
```go
func query() int {
    chan1 := make(chan int)

    // slow goroutine not be garbage collected
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
    }()
    
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
    }()
    
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
    }()

    // return the quickest response
	return <-chan1
}
```
#### Try and Error
- **method**: try to **close** the `channel`
- **result**: panic!
- **reason**: 
  - send message to `closed` channel causes **panic**
  - concentrate on resloving blocking **goroutine** rather than **channel**
```go
func query() int {
	chan1 := make(chan int)
	defer close(chan1)
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	return <-chan1
}
```

#### Solution 1: drain channel
```go
func query() int {
    chan1 := make(chan int)

    // core: clean up the channel after return
	defer func() {
		<- chan1
		<- chan1
    }()
    
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	return <-chan1
}
```

#### Solution 2: bufferred channel
- buffered channel essentially **decouples** the `sender` and `receiver` 
```go
func query() int {
	chan1 := make(chan int, 3)
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	return <-chan1
}
```
- **not always necessarily** have to **drain** a channel or **close** the channel 
- because `GC will do that automatically` [More](https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open)
- but it's a good practice to **drain** or **close** channel explicitly
```go
func query() int {
    chan1 := make(chan int, 3)
    
    // method 1: just close
    defer func() {
        close(chan1)
    }
    // method 2: drain and close ≈ method 1
    defer func() {
        <- chan1
        <- chan1
        close(chan1)
    }

    // method 3: close and drain with extra operations
    defer func() {
        close(chan1)
        for elem := range chan1 {
            // more opeartions on elements
        }
    }

	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		chan1 <- 1
	}()
	return <-chan1
}
```
#### Solution 3: select and broadcast cancellation
- select will **ignore** `awaiting` channel thus won't block the goroutine!
- **retrieve message from close channel will get zero type and error, not panic!**
- `put select inside blocked goroutine to save its life`
```go
func request() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Printf("sleeping for %d seconds...\n", n)
	time.Sleep(time.Duration(n)*time.Second)
	return 1
}

func query() int {
    chan1 := make(chan int)
    // core: closing channel here serves as a broadcasting signal!
    cancel := make(chan bool)
    // close 'cancel' channel after return 1st response
	defer func() {
		close(cancel)
	}()
	go func() {
		select {
		case chan1 <- request():
		case <- cancel:
		}
	}()
	go func() {
		select {
		case chan1 <- request():
		}
	}()
	go func() {
		select {
		case chan1 <- request():
		case <-cancel:
		}
	}()
	return <-chan1
}
```

### Rules
- **main goroutine ends, all goroutines destroyed**
- `S > R` is not necessarily GC safe
  - goroutine leaking
  - panic
- `S <= R` is always safe
  - error 
