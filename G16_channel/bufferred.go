package G16_channel



func query() int {
	bc := make(chan int, 3)

	go func() {
		bc <- 1
	}()

	go func() {
		bc <- 2
	}()

	return <- bc
}
func BufferedChannel(){
	query()
}
