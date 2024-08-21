package Controllers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// PART 1
// Write a program that takes two input strings and find if they are anagram
// or not. Anagram - a word or phrase that is made by arranging the letters
// of another word or phrase in a different order.
// Example: ARMY - MARY
// Example: Input 1: A decimal point; Input 2: Im a dot in place
// BARE - BEAR; CALM - CLAM;

func AreAnagram(s1, s2 string) {

	if Sorting1(s1) == Sorting1(s2) {
		fmt.Println("this strings are anamgram")
	} else {
		fmt.Println("this strings are not anamgram")
	}

}

func Sorting1(input string) string {
	str1 := strings.ToUpper(input)
	char := []rune(str1)
	sort.Slice(char, func(i, j int) bool {
		return char[i] < char[j]
	})
	return string(char)
}

//second approach
// func FindIsAnagram() {
// 	str1 := "ARMY"
// 	str2 := "MARY"

// 	result := find(str1, str2)
// 	fmt.Println("result", result)
// }

// func find(str1 string, str2 string) bool {

// 	rune1 := []rune(str1)
// 	rune2 := []rune(str2)

// 	for i := 0; i < len(rune1); i++ {
// 		for j := i; j < len(rune1)-1; j++ {
// 			if rune1[j] > rune1[j+1] {
// 				rune1[j], rune1[j+1] = rune1[j+1], rune1[j]
// 			}
// 		}
// 	}

// 	fmt.Println("str1", string(rune1))

// 	for i := 0; i < len(rune2); i++ {
// 		for j := i; j < len(rune2)-1; j++ {
// 			if rune2[j] > rune2[j+1] {
// 				rune2[j], rune2[j+1] = rune2[j+1], rune2[j]
// 			}
// 		}
// 	}

// 	fmt.Println("str2", string(rune2))

// 	if string(rune1) == string(rune2) {
// 		return true
// 	} else {
// 		return false
// 	}

// }

//PART 2
//Find the second largest number possible with the given input of digits.
//For example, if input is 47923, output should be 97423;
//If Input is 86046427, output should be 87664402. 87664420

func SecordLargest(digits string) {
	perms := Permute(digits)

	uniNum := make(map[string]struct{})
	for _, perm := range perms {
		uniNum[perm] = struct{}{}
	}
	nums := make([]string, 0, len(uniNum))
	for num := range uniNum {
		nums = append(nums, num)
	}

	sort.Slice(nums, func(i, j int) bool {
		numI, _ := strconv.Atoi(nums[i])
		numJ, _ := strconv.Atoi(nums[j])

		return numI > numJ

	})

	if len(nums) < 2 {
		fmt.Println("")
		return
	}

	fmt.Println(nums[1])

}

func Permute(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	var perms []string
	for i, c := range s {
		for _, perm := range Permute(s[:i] + s[i+1:]) {
			perms = append(perms, string(c)+perm)
		}
	}
	return perms
}
