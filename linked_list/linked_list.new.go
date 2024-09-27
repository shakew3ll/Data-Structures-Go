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
		fmt.Println(value)
	})

	found, err := ll.Search(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found:", found)
	}

	err = ll.InsertAt(25, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("List after inserting 25 at position 2:")
	ll.Traverse(func(value int) {
		fmt.Println(value)
	})

	err = ll.Remove(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List after removing 20:")
	ll.Traverse(func(value int) {
		fmt.Println(value)
	})

	err = ll.RemoveFirst()
	if err != nil {
		fmt.Println(err)
	}

	err = ll.RemoveLast()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Final list:")
	ll.Traverse(func(value int) {
		fmt.Println(value)
	})

	slice := ll.ToSlice()
	fmt.Println("List converted to slice:", slice)

	ll.Reverse()

	fmt.Println("Reversed list:")
	ll.Traverse(func(value int) {
		fmt.Println(value)
	})
}
