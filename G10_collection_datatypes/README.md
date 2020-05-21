<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Array](#array)
- [Slice](#slice)
- [array vs slice](#array-vs-slice)
- [Map](#map)
- [Struct](#struct)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### Array
- fixed length of sequence 
```go
var a [3]int
var b [3]int = [3]int{1,2,3}
```
- **quick declaration**
```go
p := [...]int{1,2,3} // detect the length
q = [4]int{1,2,3,4}
```
- compare using `==`, when element is comparable


### Slice
- variable length of sequence
```go
var s []int
s = []int{1,2,3}
```
- **quick declaration**
```go
s := []int{1,2,3}
```
- range
```go
s[0,1] // [0,1) 左闭右开
```
- append
```go
s = append(s,4) // immutably
```
- cannot directly compare

### array vs slice
- `[...]` vs `[]`

### Map
- key: must be same type
- value: must be the same type
- declare
```go
var m map[string]int
m = map[string] int {
    "alice":31,
}
```
-  **quick declare**
```go
p := map[string]int {
    "alice": 30,
}
q := make(map[string]int)
```
- get
```go
p["alice"]
age,ok := p["alice"] // return zero value amd false if key doens't exits
```
- set
```go
p["alice"] = 20
```
- delete
```go
delete(p,"alice")
```
### Struct
- collection of serveral `attributes`
- can have `behaviours`, **but** can only be declared as `attributes`
```go
type CalcFunc() func(int) int // valid

type Car struct {
    width int
    height int
    run CalcFunc
}
```
- **behaviour as attribute won't implement any interface**
- compare using `==` when elements are comparable
- declare
```go
type Employee struct {
    ID int
    Name string
}
```
- create
```go
var e1 Employee
e1.ID = 1
```
- **quick create**
```go
e2 := Employee{1,"James"}
e3 := Employee{ID: 1, Name: "James"}
```
- embed
```go
type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

w1 := Wheel{Circle{Point{1,2},3},4}
```
- add meta-data
```go
type Employee struct {
    firstName string `json:"firstName"`
    lastName string `json:"lastName"`
    salary int `json:"salary"`
    fullTime bool `json:"fulltime,omitempty"`
}
```

