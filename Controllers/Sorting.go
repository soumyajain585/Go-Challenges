package Controllers

import (
	"fmt"
	"sort"
)

func SortingNum() {
	// Create an unsorted slice of integers
	numbers := []int{5, 2, 8, 1, 7, 3, 6, 4}

	// Print the unsorted slice
	fmt.Println("Unsorted:", numbers)

	ch := make(chan []int)
	// Use the sort package to sort the slice in accending order
	//sort.Ints(numbers)

	go Accending(ch)
	ch <- numbers
	numbers = <-ch

	fmt.Println("accending Sorted:", numbers)

	go Descending(ch)
	ch <- numbers
	numbers = <-ch

	fmt.Println("descending Sorted:", numbers)
}

func Accending(ch chan []int) {
	numbers := <-ch
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	ch <- numbers
}

func Descending(ch chan []int) {
	numbers := <-ch
	// Reverse the order to get descending order
	n := len(numbers)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	ch <- numbers
}

func Sorting() {
	var num = []int{4, 7, 9}

	for i := 0; i > len(num)-1; i++ {
		if num[i] > num[i+1] {
			fmt.Println(false)
		}
	}

	fmt.Println(true)
}
func Sort() {

	hii := sort.IsSorted(sort.IntSlice{4, 3, 2})
	fmt.Println(hii)

	var s string
	answer := sort.Search(7, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)

	val := sort.SearchInts([]int{1, 2, 3, 4}, 1)
	fmt.Println(val)

}
