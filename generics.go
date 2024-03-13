package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}
func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func generics() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("moh", 2)
	fmt.Println(*m1)
}
