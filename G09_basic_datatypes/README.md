<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Basic Data Type](#basic-data-type)
  - [Integer](#integer)
    - [conversion](#conversion)
    - [Operation](#operation)
  - [Float](#float)
  - [Complex](#complex)
  - [Boolean](#boolean)
  - [String](#string)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Basic Data Type
### Integer
- **signed**: range`[0,2^(n-1)]`, `99% usage`
  - int8
  - int 16
  - int 32 (rune)
  - int 64
  - `int`(auto determined as 32 or 64)
- unsigned: range`[-2^(n-1),2^(n-1)-1]`, `biwise or special usage`
  - uint8(byte)
  - uint16
  - uint32
  - uint64
  - uint(auto determined as 32 or 64)
- uintptr
  - small but enough to hold a pointer

#### conversion
- is **compulsory** in golang even from `int32` to `int64`
```go
f := 3.143
i := int(f) // 3
```

#### Operation
```go
5 / 3    // 1
5 % 3    // 2
5 / 3.0  // 1.6
```


### Float
- **float64**: 99% usage
- `float32`: just don't use this
```go
fmt.Printf("%.3f",3213.32) // 3位小数
```
- special value
```go
var z float64
fmt.Println(
    z,        //  0
    -z,       // -0
    1/z,      //  Inf
    -1/z,     // -Inf
    z/z       // NaN
)
```


### Complex
- complex64(based on float32)
- complext128(based on float64)

### Boolean
- `true`
- `false`

### String
- `immutable`
```go
x := "hi"
len(x)  // get the numeber of chars
x += " world" // new string generated behind the scene
```

