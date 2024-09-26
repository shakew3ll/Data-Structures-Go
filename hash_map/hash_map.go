package hash_map

import "fmt"

type HashTable[K comparable, V any] struct {
	buckets map[K]*Node[K, V]
	size    int
}

type Node[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](size int) *HashTable[K, V] {
	return &HashTable[K, V]{
		buckets: make(map[K]*Node[K, V], size),
		size:    size,
	}
}

func (ht *HashTable[K, V]) hash(key K) int {
	hash := 0

	keyStr := fmt.Sprintf("%v", key)
	for _, char := range keyStr {
		hash = (31 * hash) + int(char)
	}

	return hash % ht.size
}

func (ht *HashTable[K, V]) hash2(key K) int {
	hash := 0
	keyStr := fmt.Sprintf("%v", key)
	for _, char := range keyStr {
		hash = (17 * hash) + int(char)
	}

	return (hash % (ht.size - 1)) + 1
}

func (ht *HashTable[K, V]) Insert(key K, value V) {
	index := ht.hash(key)
	step := ht.hash2(key)
	i := 1

	for ht.buckets[key] != nil {
		//index = (index + i) % len(ht.buckets)
		//index = (index + i*i) % ht.size
		index = (index + i*step) % ht.size
	}

	ht.buckets[key] = &Node[K, V]{key: key, value: value}
}

func (ht *HashTable[K, V]) Search(key K) (V, bool) {
	index := ht.hash(key)
	step := ht.hash2(key)
	i := 1
	var zeroValue V

	for ht.buckets[key] != nil {
		if ht.buckets[key].key == key {
			return ht.buckets[key].value, true
		}
		i++
		//index = (index + i) % len(ht.buckets)
		//index = (index + i*i) % ht.size
		index = (index + i*i + step) % ht.size
	}

	return zeroValue, false
}

func (ht *HashTable[K, V]) Delete(key K) bool {
	index := ht.hash(key)
	step := ht.hash2(key)
	i := 1

	for ht.buckets[key] != nil {
		if ht.buckets[key].key == key {
			ht.buckets[key] = nil
			return true
		}
		i++
		//index = (index + i) % len(ht.buckets)
		//index = (index + i*i) % ht.size
		index = (index + i*i + step) % ht.size
	}

	return false
}
