package Controllers

import (
	"fmt"
	"os"
	"time"
)

const (
	one = 1
	two = iota
	three
)

var res = []int{three, two, one}

type some struct{}

func Unlimit() {

	//1
	arr := []int{1, 2, 3, 4, 5, 6}

	for i := range arr {
		go fmt.Println(i)

	}

	//2
	data := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			data <- fmt.Sprintf("%d", i)
		}(i)

	}
	for d := range data {
		fmt.Println(d)

	}

	//3
	var tmp interface{}
	fmt.Println(tmp, tmp == nil)
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	panic("error")

	//4
	defer func() {
		<-time.After(time.Second)
		fmt.Println(0)
	}()
	os.Exit(0)
	fmt.Println(res[:])

	//5
	s1 := &some{}
	s2 := &some{}
	fmt.Println(fmt.Sprintf("%p", s1) == fmt.Sprintf("%p", s2))

	//6
	var f chan struct{}
	s := make(chan struct{})
	select {
	case <-f:
	case <-s:
	}
}
