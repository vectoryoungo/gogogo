package basicfeature

//consumer
func Consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv: ", x)
	}

	done <- true
}

//producer
func Producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}
