package main

import (
	"fmt"
	"newproject/types"
)

type Putter interface {
	Put(id int, val any) error
}
type Storage interface {
	Putter
	Get(id int) (any, error)
}
type FooStorage struct{}

func (f *FooStorage) Get(id int) (any, error) {
	return nil, nil
}
func (f *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (f *BarStorage) Get(id int) (any, error) {
	return nil, nil
}
func (f *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, putter Putter) error {
	return putter.Put(id, val)
}
func main() {
	user := types.User{
		Username: "aymen",
		Age:      33,
	}
	fmt.Println("user")
	fmt.Println(user)
	s := &Server{
		store: &FooStorage{},
	}
	//s.store.Get(1)
	//s.store.Put(1, 3)
	err := updateValue(1, 3, s.store)
	if err != nil {
		return
	}
	//fmt.Println(s.store.Put(1, 3))
}
