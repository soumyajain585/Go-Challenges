package Controllers

import (
	"fmt"
	"sync"
	"time"
)

func PinPong() {
	var wg sync.WaitGroup
	ch := make(chan int) // Unbuffered channel

	// Start player A and player B goroutines
	wg.Add(2)
	go playerA(ch, &wg)
	go playerB(ch, &wg)

	// Start the ball passing after a brief pause to ensure goroutines are started
	time.Sleep(time.Second)

	// Send initial ball to player A
	ch <- 1

	// Wait for both goroutines to finish
	wg.Wait()
}

func playerA(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Receive ball from the channel
		_, ok := <-ch
		if !ok {
			return
		}

		// Simulate ball passing
		fmt.Println("Ball is at player A")
		time.Sleep(2 * time.Second) // Simulate ball being held for 2 seconds

		// Pass the ball to player B
		ch <- 1
	}
}

func playerB(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Receive ball from the channel
		_, ok := <-ch
		if !ok {
			return
		}

		// Simulate ball passing
		fmt.Println("Ball is at player B")
		time.Sleep(2 * time.Second) // Simulate ball being held for 2 seconds

		// Pass the ball to player A
		ch <- 1
	}
}

func FindVowelsFromString(str string) {

	s := []rune{}

	for _, v := range str {

		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {

			s = append(s, v)
		}

	}

	fmt.Println(s)
}
