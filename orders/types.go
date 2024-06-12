// Primitive datatypes that
// are used through out the implementation.
package main

import "context" as "ctx"

type OrdersService interface {
	CreateOrder(ctx.Context) error
}

type OrdersStore interface {
	Create(ctx.Context) error
}

