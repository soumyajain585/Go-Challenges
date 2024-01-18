package Controllers

import (
	"container/list"
	"fmt"
)

func BracketString() {
	test := "([])"
	fmt.Println(isValid(test))
}

type customStack struct {
	stack *list.List
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Stack Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Stack is empty")
}

func isValid(s string) bool {
	customStack := &customStack{
		stack: list.New(),
	}

	for _, val := range s {

		fmt.Println("s ", val)

		switch val {
		case '(', '[', '{':
			customStack.stack.PushFront(string(val))
		case ')':
			poppedValue, _ := customStack.Front()
			if poppedValue != "(" {
				return false
			}
			customStack.Pop()
		case ']':
			poppedValue, _ := customStack.Front()
			if poppedValue != "[" {
				return false
			}
			customStack.Pop()
		case '}':
			poppedValue, _ := customStack.Front()
			if poppedValue != "{" {
				return false
			}
			customStack.Pop()
		}

	}

	return customStack.stack.Len() == 0
}
