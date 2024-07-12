// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: orders.proto

package orders

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Orders_Accept_FullMethodName  = "/orders.Orders/Accept"
	Orders_Deliver_FullMethodName = "/orders.Orders/Deliver"
	Orders_List_FullMethodName    = "/orders.Orders/List"
	Orders_Return_FullMethodName  = "/orders.Orders/Return"
)

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersClient interface {
	// Accept принимает заказ от курьера
	Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*AcceptResponse, error)
	// Deliver выдаёт заказы клиенту
	Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error)
	// List получает список заказов
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Return возвращает заказ курьеру
	Return(ctx context.Context, in *ReturnRequest, opts ...grpc.CallOption) (*ReturnResponse, error)
}

type ordersClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersClient(cc grpc.ClientConnInterface) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*AcceptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptResponse)
	err := c.cc.Invoke(ctx, Orders_Accept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliverResponse)
	err := c.cc.Invoke(ctx, Orders_Deliver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Orders_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) Return(ctx context.Context, in *ReturnRequest, opts ...grpc.CallOption) (*ReturnResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReturnResponse)
	err := c.cc.Invoke(ctx, Orders_Return_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
// All implementations must embed UnimplementedOrdersServer
// for forward compatibility
type OrdersServer interface {
	// Accept принимает заказ от курьера
	Accept(context.Context, *AcceptRequest) (*AcceptResponse, error)
	// Deliver выдаёт заказы клиенту
	Deliver(context.Context, *DeliverRequest) (*DeliverResponse, error)
	// List получает список заказов
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Return возвращает заказ курьеру
	Return(context.Context, *ReturnRequest) (*ReturnResponse, error)
	mustEmbedUnimplementedOrdersServer()
}

// UnimplementedOrdersServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServer struct {
}

func (UnimplementedOrdersServer) Accept(context.Context, *AcceptRequest) (*AcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (UnimplementedOrdersServer) Deliver(context.Context, *DeliverRequest) (*DeliverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}
func (UnimplementedOrdersServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrdersServer) Return(context.Context, *ReturnRequest) (*ReturnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}
func (UnimplementedOrdersServer) mustEmbedUnimplementedOrdersServer() {}

// UnsafeOrdersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServer will
// result in compilation errors.
type UnsafeOrdersServer interface {
	mustEmbedUnimplementedOrdersServer()
}

func RegisterOrdersServer(s grpc.ServiceRegistrar, srv OrdersServer) {
	s.RegisterService(&Orders_ServiceDesc, srv)
}

func _Orders_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orders_Accept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Accept(ctx, req.(*AcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orders_Deliver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Deliver(ctx, req.(*DeliverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orders_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orders_Return_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Return(ctx, req.(*ReturnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orders_ServiceDesc is the grpc.ServiceDesc for Orders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Accept",
			Handler:    _Orders_Accept_Handler,
		},
		{
			MethodName: "Deliver",
			Handler:    _Orders_Deliver_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Orders_List_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _Orders_Return_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders.proto",
}
