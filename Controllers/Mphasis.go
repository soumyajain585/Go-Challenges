package Controllers

import "fmt"

func LenCapacity() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[:3]
	s[0] = 10
	fmt.Println("s1", s)
	s = append(s, 6)
	s = append(s, 7)
	s = append(s, 8)
	s[0] = 11
	fmt.Println(len(s), cap(s))
}
