package main

import "context"

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{
		store: store,
	}
}

// Here is where all the bussiness logic
// is happening.

func (s *service) CreateOrder(context.Context) {
	return nil
}
