package Controllers

import (
	"container/list"
	"fmt"
	"sync"
)

//input
func main1() {
	array := []int{1, 2, 3, 4, 5}
	for _, element := range array {
		go func() {
			fmt.Println(element)
		}()
	}
}

//outut
func main2() {
	array := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, element := range array {
		wg.Add(1)
		go func(element int) {
			defer wg.Done()
			fmt.Println(element)
		}(element)
	}
	wg.Wait()
}

//input
func main3() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

//output
func main4() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

func CheckValid(s string) bool {
	customStack := &customStack{
		stack: list.New(),
	}

	for _, val := range s {
		switch val {
		case '(', '[', '{':
			customStack.stack.PushFront(string(val))
		case ')', ']', '}':
			if customStack.stack.Len() == 0 {
				return false
			}
			poppedValue, _ := customStack.Front()
			if !isValidPair(poppedValue, string(val)) {
				return false
			}
			customStack.Pop()
		}
	}

	return customStack.stack.Len() == 0
}

func isValidPair(open, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "[":
		return close == "]"
	case "{":
		return close == "}"
	default:
		return false
	}
}

func Bracket() {
	fmt.Println(CheckValid("([]))"))
}
