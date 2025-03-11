package grpc

import (
	"context"
	"order-service/internal/usecase"
	"order-service/proto"
)

type OrderGRPCServer struct {
	proto.UnimplementedOrderServiceServer
	usecase *usecase.OrderUseCase
}

func NewOrderGRPCServer(usecase *usecase.OrderUseCase) *OrderGRPCServer {
	return &OrderGRPCServer{usecase: usecase}
}

func (s *OrderGRPCServer) CheckOrderExists(ctx context.Context, req *proto.CheckOrderRequest) (*proto.CheckOrderResponse, error) {
	exists, err := s.usecase.CheckOrderExists(uint(req.OrderId))
	if err != nil {
		return &proto.CheckOrderResponse{Exists: false}, nil
	}
	return &proto.CheckOrderResponse{Exists: exists}, nil
}

func (s *OrderGRPCServer) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*proto.UpdateOrderStatusResponse, error) {
	err := s.usecase.UpdateOrderStatus(uint(req.OrderId), req.Status)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateOrderStatusResponse{Message: "Order status updated successfully"}, nil
}
