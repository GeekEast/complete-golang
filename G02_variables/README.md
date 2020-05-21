## Declarations
- `local` variables are initialized when its line is encountered **during function exectution**
- `package-level` variables are initialized **before `main()` executes**
### `var`
- used to define `local`, `package-level` or `global` variables
```go
var name string = "Hello world"
```
- Usage 1: stress on type
```go
var x float = 1
```
- Usage 2: Initialize with Empty value
```go
var x string
```
- Usage 3: Multiple Variables
```go
// declare multiples
var i,j,k int 
var b,f,s = true,2.3,"fout" // no type!

// receive multiple returns
var f,err = os.Open(name)
```
#### `:=` Short Variable Declarations
- **only** used to declare `local` variables
- must have a **explicit** initialized value (for referencing type)

```go
func play() string {
  text := "hello world"
  return text
}
```
- `:=` is a **declaration** not **assignment**, 
```go
var cwd string
func init() {
  cwd, err := os.Getwd() // error because name collision
}
```
- but if one variable is already declared, `:=` equals to `=`
```go
x := 1
x := 2 // error!
x = 2 // fix 1: general method

// fix 2: only use this when you have more than 1
x := 1 
x, y := 1, 2  // right! at least one new variable along with x
```
### const
- basic
```go
	// manual
	const (
		c0 = iota
		c1 = iota
	)

	// auto increment
	const (
		c3 = iota
		c4
	)

	// start from 1
	const (
		c5 = iota + 1
		c6
	)

	// skip value
	const (
		c7 = iota + 1
		_
		c8
	)
	println("skip values")
	println(c7,c8)
```
- best practice (idiomatic way)
```go
type WeekDay int // like a class

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// WeekDay implements String() interface
func (w WeekDay) String() string {
	// ... auto set length of an array [fixed length of sequence]
	return [...]string{"Monday","Tuesday", "Wednesday","Thursday","Friday","Saturday","Sunday"}[w]
}
```

> [The Different Between `Var` and `:=`](https://www.geeksforgeeks.org/difference-between-var-keyword-and-short-declaration-operator-in-golang/)


### `type`
- used to represent `real-world` object based on **underlying types**
- **type is a wrapper over basic data type**
#### declare
```go
type Celsius float64
type Fahrenheit float64
```
#### convert
```go
func CToF(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit)  Celsius {
	return Celsius((f - 32) * 5 / 9)
}
```
#### compare
```go
var c Celsius
var f Fahrenheit
res := f - CToF(c)
println(res == Fahrenheit(30))
```
#### attach function
```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
```

<!-- TODO: -->
### `func`
- declare with value return type
```go
func getTextValue() string {
	return "Hello World"
}
```
- declare with pointer return type
```go
func getTextPointer() *string {
	x := "Hello World"
	return &x
}
```
