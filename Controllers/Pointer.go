package Controllers

import "fmt"

func Pointer() {
	demoStruct1 := demo_func1()
	fmt.Println("DemoStruct1 ", demoStruct1)

	demoStruct2 := demo_func2()
	fmt.Println("DemoStruct2 ", demoStruct2)

	demo_func3(demoStruct2)
}

type DemoStruct struct {
	Val int
}

//call by value
func demo_func1() DemoStruct {
	return DemoStruct{Val: 1}
}

//call by reference
func demo_func2() *DemoStruct {
	return &DemoStruct{Val: 2}
}

//call by reference but just assign value
func demo_func3(s *DemoStruct) {
	s.Val = 1
	fmt.Println("DemoStruct3 ", s.Val)
}
