package linked_list

import "fmt"

func LinkedList() {
	ll := New()

	ll.Prepend(10)
	ll.Prepend(20)
	ll.Prepend(30)

	ll.Append(40)
	ll.Append(50)

	ll.Traverse(func(value int) {
		fmt.Printf("%d\n", value) // Use fmt.Printf instead of fmt.Println
	})

	found, err := ll.Search(20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found: %d\n", found)
	}

	err = ll.InsertAt(25, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("List after inserting 25 at position 2:")
	ll.Traverse(func(value int) {
		fmt.Printf("%d\n", value)
	})

	err = ll.Remove(20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("List after removing 20:")
	ll.Traverse(func(value int) {
		fmt.Printf("%d\n", value)
	})

	err = ll.RemoveFirst()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = ll.RemoveLast()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Final list:")
	ll.Traverse(func(value int) {
		fmt.Printf("%d\n", value)
	})

	slice := ll.ToSlice()
	fmt.Printf("List converted to slice: %v\n", slice)

	ll.Reverse()

	fmt.Println("Reversed list:")
	ll.Traverse(func(value int) {
		fmt.Printf("%d\n", value)
	})
}
