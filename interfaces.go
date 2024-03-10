package main

import "fmt"

// interface
type Storer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDbStore struct {
	// some value
}

func (mongoDbStore MongoDbStore) GetAll() ([]int, error) {
	return []int{1, 3, 4, 5, 6}, nil
}
func (mongoDbStore MongoDbStore) Put(int) error {
	return nil
}

type ApiServer struct {
	store Storer
}

func interfaces() {
	apiServer := ApiServer{store: MongoDbStore{}}
	numbers, error := apiServer.store.GetAll()
	if error == nil {
		fmt.Println(numbers)
	}
}
