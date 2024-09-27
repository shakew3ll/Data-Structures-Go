package main

import (
	"fmt"

	"github.com/shakew3ll/data_structures_go.git/binary_tree"
	"github.com/shakew3ll/data_structures_go.git/hash_map"
	"github.com/shakew3ll/data_structures_go.git/linked_list"
)

func main() {
	fmt.Println("################")
	fmt.Println("# Binary Tree: #")
	fmt.Println("################")
	binary_tree.BinaryTree()

	fmt.Println("################")
	fmt.Println("# Hash Map: #")
	fmt.Println("################")
	hash_map.HashMap()

	fmt.Println("################")
	fmt.Println("# Linked List: #")
	fmt.Println("################")
	linked_list.LinkedList()
}
