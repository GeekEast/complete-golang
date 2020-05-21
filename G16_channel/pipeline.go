package G16_channel


func PipelineSimple() {
	naturals := make(chan int)
	squares := make(chan int)
	// generate 0,1,... and send to channel - naturals
	go func() {
		for i:=0;;i++ {
			naturals <- i
		}
	}()

	// receive 0,1,... from naturals channel, square and send to square channel
	go func() {
		for {
			x :=<- naturals
			squares <- x * x
		}
	}()

	for {
		println(<- squares)
	}
}