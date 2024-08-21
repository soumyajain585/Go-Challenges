package Controllers

import (
	"fmt"
	"sync"
)

// Problem: Concurrency - Number Crunching
// Write a Go program that calculates the sum of squares of a range of numbers concurrently using goroutines and channels. The program should take an integer n as input and calculate the sum of squares of all numbers from 1 to n. Each goroutine should calculate the sum of squares for a subrange of numbers, and the main goroutine should aggregate the results.

func NumberCrunching() {
	var wg sync.WaitGroup
	n := 5
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	squareChan := make(chan int, n)
	sumChan := make(chan int)

	wg.Add(1)
	go SquareNum(nums, squareChan, &wg)

	wg.Add(1)
	go AggregateSquare(squareChan, sumChan, &wg)

	go func() {
		wg.Wait()
		close(squareChan)
	}()

	sum := <-sumChan
	fmt.Println("Sum of Squares:", sum)
}

func SquareNum(numbers []int, squareChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range numbers {
		square := num * num
		squareChan <- square
	}
}

func AggregateSquare(squareChan chan int, sumChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for square := range squareChan {
		sum += square
	}
	sumChan <- sum
}
