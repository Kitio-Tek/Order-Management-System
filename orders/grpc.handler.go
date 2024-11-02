package main

import (
    "context"
    "log"
    pb "commons/api"
    "google.golang.org/grpc"
)

type grpcHandler struct {
    pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
    handler := &grpcHandler{}
    pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
    log.Println("New order received!")
    o := &pb.Order{
        ID: "42",
    }
    return o, nil
}
