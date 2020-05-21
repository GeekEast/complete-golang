package G04_pointers

func CreatePointers() {
	// using &
	x := 1
	y := &x
	println(y) // 0xc....

	// using new
	p := new(int)
	println(p)
	println(*p) // initial value of int is 0
}


func GetValueOfPointer() {
	p := new(int)
	println(*p)
}

func SetValueByPointer() {
	p :=new(int)
	*p = 1
	println(*p)
}

func DeclarePointer(){
	var x *int
	y := 2
	//*x = 1  // not initialized
	x = &y
	println(*x)
}

func ValueOfEmptyPointer() {
	var x *int
	println(x)  // nil
}
