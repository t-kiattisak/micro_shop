package grpc

import (
	"context"
	"order-service/internal/repository"
	"order-service/proto"
)

type OrderGRPCServer struct {
	proto.UnimplementedOrderServiceServer
	repo repository.OrderRepository
}

func NewOrderGRPCServer(repo repository.OrderRepository) *OrderGRPCServer {
	return &OrderGRPCServer{repo: repo}
}

func (s *OrderGRPCServer) CheckOrderExists(ctx context.Context, req *proto.CheckOrderRequest) (*proto.CheckOrderResponse, error) {
	order, err := s.repo.GetOrderByID(uint(req.OrderId))
	if err != nil {
		return &proto.CheckOrderResponse{Exists: false}, nil
	}
	return &proto.CheckOrderResponse{Exists: order != nil}, nil
}
