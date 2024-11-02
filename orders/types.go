package main

import ("context"

       pb "commons/api"
)

type OrdersService interface {
    CreateOrder(context.Context) error
    validateOrder(context.Context,*pb.CreateOrderRequest) error
}

type OrdersStore interface {
    Create(context.Context) error
}

