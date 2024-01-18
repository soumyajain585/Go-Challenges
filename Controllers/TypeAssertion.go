package Controllers

import "fmt"

func TypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f1 := i.(float64) // panic
	fmt.Println(f1)
}
