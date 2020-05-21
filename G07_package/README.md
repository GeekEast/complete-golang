## Package and Files
- files within one folder should only have **one** package declared.
```
- Project Folder
  - Package1 Folder
    - file1.go  
    - file2.go
  - Package2 Folder
  - main.go
```
### Initialization
- `init()`: special function will run at the very begining of `main()` **before any other things**.
- will be run **only once** even the package is imported many times.
- **very useful to build database connections..**
- you can declare many `init()`,execution simply follows the `declaration` order
```go
func init() {
    println("start initializing 1")
}

func init() {
    println("start initializing 2")

}
```