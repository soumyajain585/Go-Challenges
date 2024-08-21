package Controllers

import (
	"fmt"
	"sort"
	"strings"
)

func TopKFrequent(nums []int, k int) {

	// Step 1: Count the frequency of each element
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Create a slice to sort the elements by frequency
	type pair struct {
		num   int
		count int
	}
	freqPairs := make([]pair, 0, len(freqMap))
	for num, count := range freqMap {
		freqPairs = append(freqPairs, pair{num, count})
	}

	// Step 3: Sort the pairs by frequency in descending order
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i].count > freqPairs[j].count
	})

	// Step 4: Extract the top k frequent elements
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, freqPairs[i].num)
	}

	fmt.Println(result[k-1])
}

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:

// Input: s = "rat", t = "car"
// Output: false1
//------------------------------------------------

func IsAnagram(s, t string) {

	if len(s) != len(t) {
		fmt.Println(false)
		return
	}
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	for _, char := range t {
		if count, ok := charCount[char]; !ok || count == 0 {
			fmt.Println(false)
			return
		}
		charCount[char]--
	}

	for _, count := range charCount {
		if count != 0 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.
//     Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true

// Example 2:

// Input: s = "()[]{}"
// Output: true

// Example 3:

// Input: s = "(]"
// Output: false

func CheckBracketValid(s string) {

	stack := []rune{}
	matchBracket := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if strings.ContainsRune("({[", char) {
			//push
			stack = append(stack, char)
		} else if strings.ContainsRune(")}]", char) {
			if len(stack) == 0 || stack[len(stack)-1] != matchBracket[char] {
				fmt.Println(false)
				return
			}
			//pop
			stack = stack[:len(stack)-1]

		}
	}
	//empty
	fmt.Println(len(stack) == 0)

}
