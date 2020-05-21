package G12_error_handling


func TestPanic() {
	defer func() {
		println("after panic")
	}()
	panic("throw!")
	println("2nd line")
}