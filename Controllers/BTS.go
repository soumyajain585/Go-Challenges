package Controllers

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return curr.Val
		}

		curr = curr.Right
	}

	return -1 // Return -1 if k is out of bounds
}

func main() {
	// Example BST:
	//       3
	//      / \
	//     1   4
	//      \
	//       2
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}

	k := 2
	result := kthSmallest(root, k)

	fmt.Printf("The %dth smallest element is: %d\n", k, result)
}
