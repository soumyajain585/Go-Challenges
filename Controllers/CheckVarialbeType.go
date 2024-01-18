package Controllers

import "fmt"

func CheckVarialbeType() {
	var param interface{} = 0
	switch v := param.(type) {
	default:
		fmt.Printf("Unexpected type %T", v)
	case uint64:
		fmt.Println("Integer type")
	case string:
		fmt.Println("String type")
	}
}
