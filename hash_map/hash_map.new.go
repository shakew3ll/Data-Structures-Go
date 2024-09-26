package hash_map

import "fmt"

var (
	size      = 4
	searchKey = "four"
)

func HashMap() {
	hashMap := New[string, int](size)

	hashMap.Insert("one", 1)
	hashMap.Insert("two", 2)
	hashMap.Insert("three", 3)
	hashMap.Insert("four", 4)

	//hashMap.Delete("four")

	if value, found := hashMap.Search(searchKey); found {
		fmt.Printf("Found: key=%s, value=%d\n", searchKey, value)
	} else {
		fmt.Println("Key not found")
	}
}
