<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Error handling in Golang](#error-handling-in-golang)
  - [Error](#error)
    - [Get Error](#get-error)
    - [Generate Error](#generate-error)
  - [Serve Error - Panic](#serve-error---panic)
    - [defer](#defer)
    - [panic](#panic)
    - [recover](#recover)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Error handling in Golang
- Level 1: error
- Level 2: panic (serve error)

### Error
#### Get Error
```go
val,err := myFunction(args...)
if err != nil {
    // handle error
}
```
#### Generate Error
- **General**: use `errors.New()`
```go
func g() (int,error) {
	return 1, errors.New("general error")
}

func TestMakeError() {
	_,e := g()
	fmt.Println(e)
}
```
- **Customizing**: implement error interface
```go
func f() (int,*MyError) {
	return 1, &MyError{}  // must be pointer
}

type MyError struct{}
func (e MyError) Error() string {
	return "Hello Error!"
}

func test() {
    _,err := f()
    fmt.Println(err) // don't use println here
}


// pre-defined global interface in Go
type error interface {
    Error() string
}
```

### Serve Error - Panic
- this is **alternative** to `exception` in other language, there is no `try` or `catch` keywords in golang.
#### defer
- defer is execute **after** `return` 
- `defer` params is eveluated when defer is evaluated (not executed)
```go
func a() {
    i := 0
    defer fmt.Println(i) // i os 0 not 1 
    i++
    return
}
```
- Last in First out (**LIFO**) order
```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)   // 3210
    }
}
```
- useful to modify the `return` value after returning
```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```
- useful to **close** connection, **consume** channel message...
<!-- TODO: add code examples -->
#### panic
- like exception, will be throwed from function up to its goroutine and then main goroutine to crash the program
- won't affect the defer method: **defer is executed before panic go up**
```go
	defer func() {
		println("after panic")
	}()
	panic("throw!")
```
#### recover
- like `catch`, used to catch a panic and **stop** it from crashing the whole program
- can only be used within a defer, because **only defer is not affected by panic**
```go
func f() {
    defer func() {
        if r := recover(); r != nil {  // 2. but here is a panic, so panic will be stop to pass up stack
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    panic("Panic!")  // 1. panic see here is a defer, let defer run normaly
    fmt.Println("Returned normally from g.")
}
```
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-21-00-31-47.png alt="no image found"></p>
