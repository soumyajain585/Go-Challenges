package Controllers

import (
	"fmt"
	"strconv"
)

func ReplaceMiddleWithCount(s string) {
	if len(s) <= 2 {
		fmt.Println(s)
		return
	}

	firstChar := string(s[0])
	lastChar := string(s[len(s)-1])
	middleCount := strconv.Itoa(len(s) - 2)

	fmt.Println(firstChar + middleCount + lastChar)
}
