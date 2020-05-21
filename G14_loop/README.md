## Loop
- for loop
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
- while loop
```go
n := 0
for n < 5 {
    fmt.Println(n)
    n++
}
```
- infinite loop
```go
m := 0
for {
    fmt.Println("Infinite loop")
    if m <= 8 {
        continue
    }
    if m == 8 {
        break
    }
    m++
}
```
- range loop
```go
strings := []string{"hello", "world"}
// range loop
for i, s := range strings {
    fmt.Println(i, s)
}
```
