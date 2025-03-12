// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: inventory.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InventoryService_CheckStock_FullMethodName  = "/inventory.InventoryService/CheckStock"
	InventoryService_ReduceStock_FullMethodName = "/inventory.InventoryService/ReduceStock"
	InventoryService_GetPrice_FullMethodName    = "/inventory.InventoryService/GetPrice"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error)
	ReduceStock(ctx context.Context, in *ReduceStockRequest, opts ...grpc.CallOption) (*ReduceStockResponse, error)
	GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckStockResponse)
	err := c.cc.Invoke(ctx, InventoryService_CheckStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) ReduceStock(ctx context.Context, in *ReduceStockRequest, opts ...grpc.CallOption) (*ReduceStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReduceStockResponse)
	err := c.cc.Invoke(ctx, InventoryService_ReduceStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPriceResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility.
type InventoryServiceServer interface {
	CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error)
	ReduceStock(context.Context, *ReduceStockRequest) (*ReduceStockResponse, error)
	GetPrice(context.Context, *GetPriceRequest) (*GetPriceResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStock not implemented")
}
func (UnimplementedInventoryServiceServer) ReduceStock(context.Context, *ReduceStockRequest) (*ReduceStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReduceStock not implemented")
}
func (UnimplementedInventoryServiceServer) GetPrice(context.Context, *GetPriceRequest) (*GetPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedInventoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_CheckStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CheckStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CheckStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CheckStock(ctx, req.(*CheckStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_ReduceStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReduceStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).ReduceStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_ReduceStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).ReduceStock(ctx, req.(*ReduceStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetPrice(ctx, req.(*GetPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckStock",
			Handler:    _InventoryService_CheckStock_Handler,
		},
		{
			MethodName: "ReduceStock",
			Handler:    _InventoryService_ReduceStock_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _InventoryService_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
