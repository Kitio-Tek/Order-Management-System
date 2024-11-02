package main

import (
	pb "commons/api"
	"context"

	"log"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

    service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
        service: service,
    }
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New order received! Order %v", req)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
