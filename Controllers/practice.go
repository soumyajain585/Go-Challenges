package Controllers

import "fmt"

//1,0,3,0,5,0,2,0
//Expected OP: 1,3,5,2,0,0,0,0

func ShiftZerosAtLast() {
	arr := []int{1, 0, 3, 0, 5, 0, 2, 0}

	shift(arr)

}

func shift(arr []int) {

	arrZero := []int{}
	arrNonZero := []int{}

	for _, v := range arr {
		if v == 0 {
			arrZero = append(arrZero, v)
		} else {
			arrNonZero = append(arrNonZero, v)
		}
	}

	fmt.Println("arrZero", arrZero)
	fmt.Println("arrNonZero", arrNonZero)

	// mergeSlice := []int{}

	// mergeSlice = mergeSlice(arrZero, arrNonZero)

	// fmt.Println("newArr", mergeSlice)
}
