## Pointer
- an `address` of a **variable**, not of a value
- not all values have an address, for example, primitive value
- but all variables must have an address
```go
x := 1
y := &x // not &1
```