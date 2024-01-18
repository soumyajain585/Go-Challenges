package Controllers

import (
	"fmt"
	"strings"
)

func StringUpper() {

	s := "sOuPya JaIN"
	ss := strings.Split(s, "")
	len := len(ss)
	var new []string
	for i, _ := range ss {
		if i == 0 {
			new = append(new, strings.ToUpper(ss[0]))
		} else {
			new = append(new, strings.ToLower(ss[len-(len-i)]))
		}

	}
	new1 := strings.Join([]string(new), "")
	fmt.Println(new1)

}

func CapitalizeName(name string) {

	// Convert the entire name to lowercase
	name = strings.ToLower(name)

	// Capitalize the first letter
	fmt.Println(strings.Title(name))
}

func ReverseWords(input string) {
	// Split the input string into words
	words := strings.Fields(input)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words back into a string
	reversed := strings.Join(words, " ")

	fmt.Println(reversed)
}

func ReverseString(input string) {
	words := strings.Fields(input)
	var lenOfword []int
	for _, word := range words {
		lenOfword = append(lenOfword, len(word))
	}
	fmt.Println(lenOfword)
	joindString := strings.Join(words, "")
	joinarray := []rune(joindString)
	// Reverse the order of runes
	for i, j := 0, len(joinarray)-1; i < j; i, j = i+1, j-1 {
		joinarray[i], joinarray[j] = joinarray[j], joinarray[i]
	}
	// Convert the reversed slice of runes back to a string
	reversed := string(joinarray)
	fmt.Println(reversed)
	result := make([]string, len(lenOfword))
	index := 0
	for i, length := range lenOfword {
		if index+length > len(reversed) {
			fmt.Println("not enough characters in the input for the specified lengths")
		}
		result[i] = reversed[index : index+length]
		index += length
	}
	reverstring := strings.Join(result, " ")
	fmt.Println(reverstring)
}
