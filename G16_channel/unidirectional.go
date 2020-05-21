package G16_channel


func UnidirectionalChannel() {
	ping := make(chan<- int)
	pong := make(<-chan int)

	ping <-1
	//<- ping // compile error

	<-pong
	//pong<-1 // compile error
}
