<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Interface](#interface)
  - [Declaration](#declaration)
  - [Implement](#implement)
  - [Polymorphism](#polymorphism)
  - [Embedding](#embedding)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Interface
- protcol of behaviours
### Declaration
```go
type Shape interface{
    Area() float64
    Perimeter() float64
}
```

### Implement
- `struct` must implement **all** methods in an `interface`
- a struct can implement multiple interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rect struct {
    width float64
    height float64
}

// implments Area method
func (self Rect) Area() float64 {
    return r.width * r.height
}

// implement Perimeter method
func (self Rect) Perimeter floa64 {
    return 2 * (r.width + r.height)
} 
```

### Polymorphism
- use interface as parameter type
```go
// declaration type vs acutal type
var s Shape = Rect{10,3}
func run(s Shape) {
    fmt.Printlin(s.Area())
} 
```
- down casting
```go
// not *Shape: * is not used for interface
func play(s Shape) {
	// assert as Circle type
	fmt.Println(s.(Circle).Area())
}
```
<!-- TODO: need deeper understanding of use case -->
- type switch
> This is not Generics! (Generics is about element type within a collection)
```go
// empty interface which accept anything
func explain (i interface{}) {
    switch i.(type){
        case string: // 
        case int: // 
        default: //
    }
}
```

### Embedding
```go
type Shape interface {
    Area() float64
}

type Object interface {
    Volume() float64
}

type Material interface {
    Shapre
    Object
}
```

## References
- [GO interface](https://medium.com/rungo/interfaces-in-go-ab1601159b3a)