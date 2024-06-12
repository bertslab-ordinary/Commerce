package main

import "context"


func main() {

	store := NewStore()
	orderservice := NewService(store)

	orderservice.CreateOrder(context.Background())

}
