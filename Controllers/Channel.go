package Controllers

import "fmt"

func workerRoutine(ch chan struct{}) {
	// Receive message from main program.
	<-ch
	fmt.Println("Signal Received")

	// Send a message to the main program.
	close(ch)
}

func Channel() {
	//Create channel
	ch := make(chan struct{})

	//define workerRoutine
	go workerRoutine(ch)

	// Send signal to worker goroutine
	ch <- struct{}{}

	// Receive a message from the workerRoutine.
	<-ch
	fmt.Println("Signal Received")
}

func workerRoutine1(ch chan IntStruct) {
	// Receive message from main program.
	val := <-ch
	fmt.Println("Signal Received - 1 ", val)

	// Send a message to the main program.
	//ch <- IntStruct{Value: 1}
	close(ch)
}

type IntStruct struct {
	Value int
}

func Channel1() {
	//Create channel
	ch := make(chan IntStruct)

	//define workerRoutine
	go workerRoutine1(ch)

	// Send signal to worker goroutine
	ch <- IntStruct{Value: 1}

	// Receive a message from the workerRoutine.
	val := <-ch
	fmt.Println("Signal Received - 2 ", val)
}

func workerRoutine2(ch chan interface{}) {
	// Receive message from main program.
	val := <-ch
	fmt.Println("Signal Received int", val)

	val2 := <-ch
	fmt.Println("Signal Received string", val2)

	// Send a message to the main program.
	ch <- struct{ Name string }{"John Doe"}
	close(ch)
}

func Channel2() {
	//Create channel
	ch := make(chan interface{})

	//define workerRoutine
	go workerRoutine2(ch)

	// Send int signal to worker goroutine
	ch <- 1

	// Send int signal to worker goroutine
	ch <- "hello"

	// Receive a message from the workerRoutine.
	val := <-ch
	fmt.Println("Signal Received - 2 ", val)
}

func workerRoutine3(ch chan []string) {
	// Receive message from main program.
	val := <-ch
	for _, v := range val {
		fmt.Println("Signal Received ", v)
	}

	// Send a message to the main program.
	ch <- val
	close(ch)
}

func Channel3() {
	//Create channel
	ch := make(chan []string)

	//define workerRoutine
	go workerRoutine3(ch)

	// Send int signal to worker goroutine
	ch <- []string{"hell0", "hii", "bye"}

	// Receive a message from the workerRoutine.
	val := <-ch
	fmt.Println("Signal Received ", val)
}

func BufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
