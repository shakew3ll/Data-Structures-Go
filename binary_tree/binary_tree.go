package binary_tree

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func New(slice []int) *Node {
	var root *Node

	for _, val := range slice {
		root = root.insert(val)
	}

	return root
}

func (n *Node) insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}

	if value < n.Value {
		if n.left == nil {
			n.left = &Node{Value: value}
		} else {
			n.left = n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{Value: value}
		} else {
			n.right = n.right.insert(value)
		}
	}

	return n
}

func (n *Node) Search(value int) (*Node, *Node) {
	if n == nil || value == n.Value {
		return n, nil
	}

	parent := n
	for n != nil && value != n.Value {
		parent = n

		if value < n.Value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return n, parent
}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.left = n.left.Delete(value)
	} else if value > n.Value {
		n.right = n.right.Delete(value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minRight := n.right.minNode()
		n.Value = minRight.Value
		n.right = n.right.Delete(minRight.Value)
	}

	return n
}

func (n *Node) minNode() *Node {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return n.left
	}

	return n
}

func (n *Node) Traverse() {
	if n == nil {
		return
	}

	n.left.Traverse()
	fmt.Printf("%d ", n.Value)
	n.right.Traverse()
}
