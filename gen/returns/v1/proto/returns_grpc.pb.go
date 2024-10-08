// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: returns.proto

package returns

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
	Returns_AcceptReturn_FullMethodName = "/returns.Returns/AcceptReturn"
	Returns_ListReturns_FullMethodName  = "/returns.Returns/ListReturns"
)

// ReturnsClient is the client API for Returns service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReturnsClient interface {
	// AcceptReturn принимает возврат от клиента
	AcceptReturn(ctx context.Context, in *AcceptReturnRequest, opts ...grpc.CallOption) (*AcceptReturnResponse, error)
	// ListReturns получает список возвратов
	ListReturns(ctx context.Context, in *ListReturnsRequest, opts ...grpc.CallOption) (*ListReturnsResponse, error)
}

type returnsClient struct {
	cc grpc.ClientConnInterface
}

func NewReturnsClient(cc grpc.ClientConnInterface) ReturnsClient {
	return &returnsClient{cc}
}

func (c *returnsClient) AcceptReturn(ctx context.Context, in *AcceptReturnRequest, opts ...grpc.CallOption) (*AcceptReturnResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptReturnResponse)
	err := c.cc.Invoke(ctx, Returns_AcceptReturn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *returnsClient) ListReturns(ctx context.Context, in *ListReturnsRequest, opts ...grpc.CallOption) (*ListReturnsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReturnsResponse)
	err := c.cc.Invoke(ctx, Returns_ListReturns_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReturnsServer is the server API for Returns service.
// All implementations must embed UnimplementedReturnsServer
// for forward compatibility
type ReturnsServer interface {
	// AcceptReturn принимает возврат от клиента
	AcceptReturn(context.Context, *AcceptReturnRequest) (*AcceptReturnResponse, error)
	// ListReturns получает список возвратов
	ListReturns(context.Context, *ListReturnsRequest) (*ListReturnsResponse, error)
	mustEmbedUnimplementedReturnsServer()
}

// UnimplementedReturnsServer must be embedded to have forward compatible implementations.
type UnimplementedReturnsServer struct {
}

func (UnimplementedReturnsServer) AcceptReturn(context.Context, *AcceptReturnRequest) (*AcceptReturnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptReturn not implemented")
}
func (UnimplementedReturnsServer) ListReturns(context.Context, *ListReturnsRequest) (*ListReturnsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReturns not implemented")
}
func (UnimplementedReturnsServer) mustEmbedUnimplementedReturnsServer() {}

// UnsafeReturnsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReturnsServer will
// result in compilation errors.
type UnsafeReturnsServer interface {
	mustEmbedUnimplementedReturnsServer()
}

func RegisterReturnsServer(s grpc.ServiceRegistrar, srv ReturnsServer) {
	s.RegisterService(&Returns_ServiceDesc, srv)
}

func _Returns_AcceptReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptReturnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReturnsServer).AcceptReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Returns_AcceptReturn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReturnsServer).AcceptReturn(ctx, req.(*AcceptReturnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Returns_ListReturns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReturnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReturnsServer).ListReturns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Returns_ListReturns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReturnsServer).ListReturns(ctx, req.(*ListReturnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Returns_ServiceDesc is the grpc.ServiceDesc for Returns service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Returns_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "returns.Returns",
	HandlerType: (*ReturnsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptReturn",
			Handler:    _Returns_AcceptReturn_Handler,
		},
		{
			MethodName: "ListReturns",
			Handler:    _Returns_ListReturns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "returns.proto",
}
