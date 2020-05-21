<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Start Go on MacOS](#start-go-on-macos)
- [Hello Golang](#hello-golang)
- [Go CMD](#go-cmd)
- [Go Module](#go-module)
- [Future Plan](#future-plan)
- [Learning Resources](#learning-resources)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->



## Start Go on MacOS
- **Install** `brew install go`
- **IDE**
  - `VScode`: lightweight choice 
    - `go`: you will install other dependencies
    - `go snippets`
  - **GoLand**: standard choice
- **Run** `go run hello.go`


## Hello Golang
```go
// main.go
package main // required

import "fmt" // required

const text = "Hello Golang" // package-level var

func getText() string {
  text := "hello Golang" // local to getText()
  return text
}

func main() {  // required
  fmt.Println(text)
  fmt.Println(getText())
}

```


## Go CMD
```go
go env         // environment variables
go build       // build to binary
go get         // install packages
go version     // version of go
```
- **GOROOT**: includes `cmd`, `standard lib` and `go compiler`
- **GOPATH**: after 1.11, you don't have to worry about this 👇

## Go Module
- Before `1.11`, you will always to change `GOPATH` to your current project folder.
- Now, you can use `go mod` to **map** dependencies in `GOPATH` and don't need to change `GOPATH` anymore
```sh
export GO111MODULE=off   # turn off module system
go mod init essential-go # init module
go list -m all           # list dependencies
go mod tidy              # remove unused 
go get -u                # update all
```

## Future Plan
- New insights will be added to `G19_tricks` folder
- Chapter 11: Testing
- Chapter 12: Reflection
- Chapter 9: Concurrency with Share Variables
- Chapter 10: Package and Tool

## Learning Resources
- [Go Examples](https://gobyexample.com/interfaces)
- [Go Playground](https://goplay.space/)
- [Go Introduction](https://medium.com/rungo/go-introductory-tutorials-896aeda0fb8a)
- [Introcution to OOP in GO](https://code.egym.de/introduction-to-oop-in-golang-e4841a9c4e3e)
- [Go Bible Code Samples](https://github.com/adonovan/gopl.io)