package Controllers

import (
	"fmt"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func apiCall(ch chan int) {
	//defer wg.Done()

	for i := range ch {
		fmt.Println("API call for", i, "started")
		//time.Sleep(time.Millisecond)
	}

}

func Channel01() {

	numArray := makeRange(0, 500)
	ch := make(chan int)
	//var wg sync.WaitGroup

	start := time.Now()
	//wg.Add(1)
	go apiCall(ch)
	for i := range numArray {
		ch <- i
	}
	close(ch)
	//wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Time taken %s", elapsed)
}
