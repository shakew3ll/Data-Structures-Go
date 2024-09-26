package binary_tree

import (
	"fmt"
)

var (
	binaryTreeSlice = []int{4, 3, 5, 2, 6, 1, 7}
	searchValue     = 4
)

func BinaryTree() {
	binaryTree := New(binaryTreeSlice)

	value, parent := binaryTree.Search(searchValue)

	if value != nil {
		fmt.Println("Value:", value.Value)

		if parent != nil {
			fmt.Printf("Parent: %d.", parent.Value)
		} else {
			fmt.Printf("No parent (root node).")
		}
	} else {
		fmt.Println("No such value...")
	}
}
