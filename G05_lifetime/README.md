<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Lifetime of Variables](#lifetime-of-variables)
  - [Levels](#levels)
  - [Performance Optimization](#performance-optimization)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Lifetime of Variables
### Levels
- `package-level`: entire execution of the program.
  - **do not declare so many package-level variables**
- `local-level`: from creation(execution) to unreachable status (when pointer is cycled)

### Performance Optimization
- Though golang has `Garbage Collector`, you still need to care about how to use variables
- Bad Example
```go
func BadGCExample() {
	var g *int
	test := func() {
		var x int
		x = 1
		g = &x
	}
	test()
}
```
- Good Example
```go
func GoodGCExample() {
	g := new(int)
	test := func() {
		var x int
		x = 1
		*g = x
	}
	test()
}
```
- Even you use pure function like `Map`, the **heap** can still store a lot of data
- In order to inspect the **head** usage, please using [profiling tool](https://cizixs.com/2017/09/11/profiling-golang-program/)
