<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table Of Content

- [Scope](#scope)
  - [Categories](#categories)
  - [for..loop](#forloop)
  - [if..else](#ifelse)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Scope
- also called `lexical scope` or `static scope`
- regions that one piece of code can have access to
- determined **statically** by the `location` of that piece of code during the `compiling stage`
- basic rules:
  - inner can access to outer 
  - outer cannot access inner
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-20-21-48-22.png alt="no image found"></p>

### Categories
- with `{}`
  - func() {}
  - for ... {}
  - {} itself
- without `{}`
  - () in `condition` and `loop`
  - `file`
  - `package`
  - `global`

### for..loop
```go
for _,x := range m       // scope1
{
    x := x + 'A' - 'a'   // scope 2
}
```
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-20-21-55-41.png alt="no image found"></p>

### if..else
```go
if x := f(); x == 0  // scope 1 
{                    // scope 2
    fmt.Println(x)  
} 
else if y := g(x); x == y // scope 3
{                         // scope 4
    fmt.Println(x, y) 
} 
else                      // scope 5
{                         // scope 6
    fmt.Println(x, y) 
} 
fmt.Println(x, y) // compile error: x and y are not visible here
```
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-20-21-57-03.png alt="no image found"></p>