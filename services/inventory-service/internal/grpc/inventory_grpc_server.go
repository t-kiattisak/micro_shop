package grpc

import (
	"context"
	"inventory-service/internal/usecase"
	"inventory-service/proto"
)

type InventoryGRPCServer struct {
	proto.UnimplementedInventoryServiceServer
	usecase *usecase.InventoryUseCase
}

func NewInventoryGRPCServer(uc *usecase.InventoryUseCase) *InventoryGRPCServer {
	return &InventoryGRPCServer{usecase: uc}
}

func (s *InventoryGRPCServer) CheckStock(ctx context.Context, req *proto.CheckStockRequest) (*proto.CheckStockResponse, error) {
	err := s.usecase.CheckStock(req.Product, int(req.Quantity))
	if err != nil {
		return &proto.CheckStockResponse{
			Available: false,
			Message:   err.Error(),
		}, nil
	}
	return &proto.CheckStockResponse{
		Available: true,
		Message:   "Stock available",
	}, nil
}

func (s *InventoryGRPCServer) ReduceStock(ctx context.Context, req *proto.ReduceStockRequest) (*proto.ReduceStockResponse, error) {
	err := s.usecase.ReduceStock(req.Product, int(req.Quantity))
	if err != nil {
		return &proto.ReduceStockResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	return &proto.ReduceStockResponse{
		Success: true,
		Message: "Stock reduced successfully",
	}, nil
}

func (s *InventoryGRPCServer) GetPrice(ctx context.Context, req *proto.GetPriceRequest) (*proto.GetPriceResponse, error) {
	price, err := s.usecase.GetPrice(req.Product)
	if err != nil {
		return nil, err
	}
	return &proto.GetPriceResponse{Price: price}, nil
}
