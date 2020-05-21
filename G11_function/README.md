## Function
- `first-class` member: can be defined as normal type
```go
type CalcFunc func(int int) int
```
  
### Declaration
- no return
```go
func play(x int) {
    fmt.Println(x)
}
```
- has return
```go
func play(x int) int {
    return x
}
```
- value return
```go
func getTextValue() string {
	return "Hello World"
}
```
- pointer return
```go
func getTextPointer() *string {
	x := "Hello World"
	return &x
}
```

- multiple return 
```go
func play(x int, y int) (int int) {
    return x,y
}
```
- attach to `receiver`
  - receiver can be `type`, `struct` but **cannot** be `interface`
```go
type Employee struct {
    Name string
    age int32
}

func (self Employee) play() int {
    return self.age
}
```
- variable parameters
```go
func sum(vals...int) int {
    total := 0
    for _,val := range vals {
        total += val
    }
    return total
}
```
- `override` is **not** supported in go
    - same name, diff params -> collision
    - same name, diff params, diff return type -> collision
    - **same name**, same params, **diff receiver** -> `no collision`
    - same name, diff params. same receiver -> collision
- declare function inside function
```go
func ConvertTypes() (func(c Celsius) Fahrenheit) {
	CToF := func(c Celsius) Fahrenheit {
		return Fahrenheit(c*9/5 + 32)
	}
    return CToF
}
```
