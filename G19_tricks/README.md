<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Swap in Go](#swap-in-go)
- [Print type](#print-type)
- [Passing by what?](#passing-by-what)
- [Logging](#logging)
- [The essense of Lock](#the-essense-of-lock)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Swap in Go
```go
i,j = j,i // swap value of i,j

func fib(n int) int {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    x, y = y, x+y
  } 
  return x
}
```
### Print type
```go
// type
fmt.Printf("dasdad %T \n", x)
// value
fmt.Printf("dasdad %#v \n", y)
```
### Passing by what?
- You need customize passing by `value` or passing by `reference` in golang
```go
func run(n M) {}  // passing by value
func run(m *M) {} // passing by reference
```

### Logging
- `log.Fatal()` is equivalent to Print() followed by a call to os.Exit(1)


### The essense of Lock
- make `parallel` as `sequential`