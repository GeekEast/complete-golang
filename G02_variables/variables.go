package G02_variables

func VariableNames() {
	println("start with [a-zA-Z] || _")
	println("continue with [a-ZA-Z] || [0-9] || _")
	println("case sensitive")
	println("prefer camel case naming")
}


func DeclareVariables() {
	println("Define local,package,global-level variables")
	var v0 string
	var v1 string = "James"  // explicit type
	var v2 = "james"         // type detecting
	var v3, v4 = 1, 2        // multiple with type detecting
	println(v0,v1)
	print(v2, v3, v4)
}

func DeclareShort(){
	println("Quickly define local variables")
	x := 1
	println(x)
}

func DeclareConstants() {
	const c1 = 1     // type detecting
	const c2 int = 1 // explicit type
	const (
		c3 = 2
		c4 = 2
	)
	println(c1,c2,c3,c4)
}

func DeclareEnum(){
	// manual
	const (
		c0 = iota
		c1 = iota
	)
	println("manual spreading")
	println(c0,c1)

	// auto increment
	const (
		c3 = iota
		c4
	)	// manual
	const (
		c0 = iota
		c1 = iota
	)
	println("manual spreading")
	println(c0,c1)

	// auto increment
	const (
		c3 = iota
		c4
	)
	println("auto spreading")
	println(c3,c4)

	// start from 1
	const (
		c5 = iota + 1
		c6
	)
	println("start from 1")
	println(c5,c6)

	// skip value
	const (
		c7 = iota + 1
		_
		c8
	)
	println("skip values")
	println(c7,c8)
	println("auto spreading")
	println(c3,c4)

	// start from 1
	const (
		c5 = iota + 1
		c6
	)
	println("start from 1")
	println(c5,c6)

	// skip value
	const (
		c7 = iota + 1
		_
		c8
	)
	println("skip values")
	println(c7,c8)
}

// idiomatic way (best practice)
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
