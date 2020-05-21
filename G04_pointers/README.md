<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Pointer](#pointer)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Pointer
- an `address` of a **variable**, not of a value
- not all values have an address, for example, primitive value
- but all variables must have an address
```go
x := 1
y := &x // not &1
```