package Controllers

import (
	"errors"
	"fmt"
	"time"
)

//Using anonymous functions as closures in Golang

func Anonymous() {
	number := 0
	increment := func() int {
		number++
		return number
	}
	fmt.Println(increment())
	fmt.Println(increment())

	err := func(err error) {
		fmt.Println(err)
	}
	err(errors.New("this is my error"))
}

//Passing anonymous functions as arguments
func Anonymous1() {
	numbers := []int{1, 2, 3, 4, 5}
	forEach(numbers, func(number int) {
		fmt.Println(number)
	})
}

func forEach(numbers []int, callback func(int)) {
	for _, number := range numbers {
		callback(number)
	}
}

//Returning anonymous functions from functions in Golang

func makeGreeter() func(string) string {
	return func(name string) string {
		return "Hello " + name
	}
}

func Anonymous2() {
	greeter := makeGreeter()
	fmt.Println(greeter("Kevin"))
}

//Using anonymous functions as goroutines in Golang

func Anonymous3() {
	go func() {
		fmt.Println("Hello from goroutine")
	}()
	time.Sleep(1 * time.Second)
}
