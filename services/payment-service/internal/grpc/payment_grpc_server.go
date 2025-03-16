package grpc

import (
	"context"
	"payment-service/internal/usecase"
	"payment-service/proto"
)

type PaymentGRPCServer struct {
	proto.UnimplementedPaymentServiceServer
	usecase *usecase.PaymentUseCase
}

func NewPaymentGRPCServer(usecase *usecase.PaymentUseCase) *PaymentGRPCServer {
	return &PaymentGRPCServer{usecase: usecase}
}

func (s *PaymentGRPCServer) UpdatePaymentStatus(ctx context.Context, req *proto.UpdatePaymentStatusRequest) (*proto.UpdatePaymentStatusResponse, error) {
	err := s.usecase.UpdatePaymentStatus(uint(req.OrderId), req.Status)
	if err != nil {
		return nil, err
	}
	return &proto.UpdatePaymentStatusResponse{Message: "update payment successfully"}, nil
}
