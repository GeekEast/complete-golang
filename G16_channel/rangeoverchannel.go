package G16_channel

import "fmt"



func LoopOverChannel() {
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for {
		if elem,ok := <- queue; ok {
			println(elem)
		} else {
			break
		}
	}
}

func RangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
