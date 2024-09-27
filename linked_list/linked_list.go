package linked_list

import "fmt"

type LinkedListStruct struct {
	Head *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

func New() *LinkedListStruct {
	return &LinkedListStruct{
		Head: nil,
		Size: 0,
	}
}

func (ll *LinkedListStruct) Prepend(value int) {
	node := &Node{Value: value, Next: ll.Head.Next}

	ll.Head = node
	ll.Size++
}

func (ll *LinkedListStruct) Append(value int) {
	node := &Node{Value: value, Next: nil}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	ll.Head.Next = node
	ll.Size++
}

func (ll *LinkedListStruct) InsertAt(position, value int) error {
	if position < 0 || position > ll.Size {
		return fmt.Errorf("index out of bounds.")
	}

	if position == 0 {
		ll.Prepend(value)

		return nil
	}

	node := &Node{Value: value}
	current := ll.Head

	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	current.Next = node.Next
	current.Next = node
	ll.Size++

	return nil
}

func (ll *LinkedListStruct) RemoveFirst() error {
	if ll.Head == nil {
		return fmt.Errorf("head is empty.")
	}

	ll.Head = ll.Head.Next
	ll.Size--

	return nil
}

func (ll *LinkedListStruct) RemoveLast() error {
	if ll.Head == nil {
		return fmt.Errorf("head is empty.")
	}

	if ll.Head.Next == nil {
		return ll.RemoveFirst()
	}

	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	ll.Size--

	return nil
}

func (ll *LinkedListStruct) Remove(value int) error {
	if ll.Head == nil {
		return fmt.Errorf("head is empty.")
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--

		return nil
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.Size--

			return nil
		}

		current = current.Next
	}

	return fmt.Errorf("value not found.")
}

func (ll *LinkedListStruct) Search(value int) (bool, error) {
	if ll.Head == nil {
		return false, fmt.Errorf("head is empty")
	}

	current := ll.Head

	for current != nil {
		if current.Value == value {
			return true, nil
		}

		current = current.Next
	}

	return false, fmt.Errorf("value not found.")
}

func (ll *LinkedListStruct) Traverse(callback func(value int)) {
	current := ll.Head
	for current != nil {
		callback(current.Value)
		current = current.Next
	}
}

func (ll *LinkedListStruct) Reverse() {
	var prev, next *Node
	current := ll.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}

func (ll *LinkedListStruct) ToSlice() []int {
	slice := make([]int, 0)
	current := ll.Head

	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}

	return slice
}
