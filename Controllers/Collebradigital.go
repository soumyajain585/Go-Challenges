package Controllers

import "fmt"

//1,0,3,0,5,0,2,0
//Expected OP: 1,3,5,2,0,0,0,0

func ShiftZero(arr []int) {
	fmt.Println("current arr ", arr)
	index := 0
	for _, val := range arr {
		if val != 0 {
			arr[index] = val
			index++
			fmt.Println(index-1, " running arr ", arr)
		}
	}

	for index < len(arr) {
		arr[index] = 0
		index++
	}
	fmt.Println("new arr ", arr)

}
