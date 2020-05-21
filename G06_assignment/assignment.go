package G06_assignment

func SingleAssignment() {
	x := 1
	x = 2
	x *= 2
	x++
	x--
	println(x)
}

func GetMapValueAssignment() {
	m := make(map[string]int)
	m["1"] = 2
	v1, ok1 := m["1"]
	v2, ok2 := m["2"]
	println(v1, ok1) // 2 true
	println(v2, ok2) // 0 false
}

func TypeAssertionAssignment() {
	var x interface{} = 1
	v1, ok1 := x.(int)    // whether x interface holds the concrete int type
	v2, ok2 := x.(string) // whether x interface holds the concrete string type
	println(v1, ok1)      // 1                     true
	println(v2, ok2)      // zero value of string  false
}

func ReceiveChannelAssignment() {
	c1 := make(chan int)

	go func() {
		c1 <- 1
	}()

	v, ok := <-c1
	// v1, ok1 := <-c1  channel already closed, cannot send
	println(v,ok)
}

func BlankIdentifierAssignment() {
	m := make(map[string]int)
	m["1"] = 2
	_, ok1 := m["1"]
	println(ok1) // you cannot pass _ here
}
