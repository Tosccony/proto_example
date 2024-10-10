// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: services.proto

package services_protos

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
	MyService_GetMenu_FullMethodName        = "/services.MyService/GetMenu"
	MyService_PlaceOrder_FullMethodName     = "/services.MyService/PlaceOrder"
	MyService_GetOrderStatus_FullMethodName = "/services.MyService/GetOrderStatus"
)

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Menu], error)
	PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error)
	GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}


func (c *myServiceClient) GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Menu], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], MyService_GetMenu_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MenuRequest, Menu]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil

}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_GetMenuClient = grpc.ServerStreamingClient[Menu]

func (c *myServiceClient) PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, MyService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderStatus)
	err := c.cc.Invoke(ctx, MyService_GetOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility.
type MyServiceServer interface {
	GetMenu(*MenuRequest, grpc.ServerStreamingServer[Menu]) error
	PlaceOrder(context.Context, *Order) (*Receipt, error)
	GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMyServiceServer struct{}

func (UnimplementedMyServiceServer) GetMenu(*MenuRequest, grpc.ServerStreamingServer[Menu]) error {
	return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedMyServiceServer) PlaceOrder(context.Context, *Order) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedMyServiceServer) GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}
func (UnimplementedMyServiceServer) testEmbeddedByValue()                   {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	// If the following call pancis, it indicates UnimplementedMyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_GetMenu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MenuRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyServiceServer).GetMenu(m, &grpc.GenericServerStream[MenuRequest, Menu]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_GetMenuServer = grpc.ServerStreamingServer[Menu]

func _MyService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).PlaceOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_GetOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).GetOrderStatus(ctx, req.(*Receipt))
	}
	return interceptor(ctx, in, info, handler)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _MyService_PlaceOrder_Handler,
		},
		{
			MethodName: "GetOrderStatus",
			Handler:    _MyService_GetOrderStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMenu",
			Handler:       _MyService_GetMenu_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}
