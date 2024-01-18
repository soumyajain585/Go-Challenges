package Controllers

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

import (
	"errors"
	"fmt"
	"sort"
)

func FrequentInteger(num []int, k int) ([]int, error) {
	if len(num) == 0 {
		return []int{}, errors.New("array is empty")
	}
	count := make(map[int]int)

	for _, no := range num {
		count[no]++
	}
	fmt.Println(count)

	var numarray []int
	for val := range count {
		numarray = append(numarray, val)
	}
	fmt.Println(numarray)

	sort.Slice(numarray, func(i, j int) bool {
		return numarray[i] < numarray[j]
	})
	fmt.Println(numarray)
	var result []int
	for i := 0; i < k && i < len(numarray); i++ {
		result = append(result, numarray[i])
	}
	fmt.Println(result)
	return result, nil
}
