package G05_lifetime



func BadGCExample() {
	var g *int
	test := func() {
		var x int
		x = 1
		g = &x
	}
	test()
}


func GoodGCExample() {
	g := new(int)
	test := func() {
		var x int
		x = 1
		*g = x
	}
	test()
}