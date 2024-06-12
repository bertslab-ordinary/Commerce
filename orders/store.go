// Similar to the repository
package main

import "context"

import "fmt"
/* TODO: DataStructure for MongoDB.*/
type store {
	// MongoDB instance
}

func NewService() *store {
	return &store{}
}

func (s *store) CreateOrder(context.Context) {
	// TODO: Create Order
	fmt.Println("Persistance layer activated.")
	
	return nil
}