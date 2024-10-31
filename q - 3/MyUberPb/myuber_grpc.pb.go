// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: myuber.proto

package __

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
	RideSharingService_RequestRide_FullMethodName   = "/ridesharing.RideSharingService/RequestRide"
	RideSharingService_GetRideStatus_FullMethodName = "/ridesharing.RideSharingService/GetRideStatus"
	RideSharingService_AcceptRide_FullMethodName    = "/ridesharing.RideSharingService/AcceptRide"
	RideSharingService_RejectRide_FullMethodName    = "/ridesharing.RideSharingService/RejectRide"
	RideSharingService_CompleteRide_FullMethodName  = "/ridesharing.RideSharingService/CompleteRide"
)

// RideSharingServiceClient is the client API for RideSharingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service Definition
type RideSharingServiceClient interface {
	// Rider Methods
	RequestRide(ctx context.Context, in *RideRequest, opts ...grpc.CallOption) (*RideResponse, error)
	GetRideStatus(ctx context.Context, in *RideStatusRequest, opts ...grpc.CallOption) (*RideStatusResponse, error)
	// Driver Methods
	AcceptRide(ctx context.Context, in *AcceptRideRequest, opts ...grpc.CallOption) (*AcceptRideResponse, error)
	RejectRide(ctx context.Context, in *RejectRideRequest, opts ...grpc.CallOption) (*RejectRideResponse, error)
	CompleteRide(ctx context.Context, in *RideCompletionRequest, opts ...grpc.CallOption) (*RideCompletionResponse, error)
}

type rideSharingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRideSharingServiceClient(cc grpc.ClientConnInterface) RideSharingServiceClient {
	return &rideSharingServiceClient{cc}
}

func (c *rideSharingServiceClient) RequestRide(ctx context.Context, in *RideRequest, opts ...grpc.CallOption) (*RideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideResponse)
	err := c.cc.Invoke(ctx, RideSharingService_RequestRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rideSharingServiceClient) GetRideStatus(ctx context.Context, in *RideStatusRequest, opts ...grpc.CallOption) (*RideStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideStatusResponse)
	err := c.cc.Invoke(ctx, RideSharingService_GetRideStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rideSharingServiceClient) AcceptRide(ctx context.Context, in *AcceptRideRequest, opts ...grpc.CallOption) (*AcceptRideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptRideResponse)
	err := c.cc.Invoke(ctx, RideSharingService_AcceptRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rideSharingServiceClient) RejectRide(ctx context.Context, in *RejectRideRequest, opts ...grpc.CallOption) (*RejectRideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RejectRideResponse)
	err := c.cc.Invoke(ctx, RideSharingService_RejectRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rideSharingServiceClient) CompleteRide(ctx context.Context, in *RideCompletionRequest, opts ...grpc.CallOption) (*RideCompletionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideCompletionResponse)
	err := c.cc.Invoke(ctx, RideSharingService_CompleteRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RideSharingServiceServer is the server API for RideSharingService service.
// All implementations must embed UnimplementedRideSharingServiceServer
// for forward compatibility.
//
// Service Definition
type RideSharingServiceServer interface {
	// Rider Methods
	RequestRide(context.Context, *RideRequest) (*RideResponse, error)
	GetRideStatus(context.Context, *RideStatusRequest) (*RideStatusResponse, error)
	// Driver Methods
	AcceptRide(context.Context, *AcceptRideRequest) (*AcceptRideResponse, error)
	RejectRide(context.Context, *RejectRideRequest) (*RejectRideResponse, error)
	CompleteRide(context.Context, *RideCompletionRequest) (*RideCompletionResponse, error)
	mustEmbedUnimplementedRideSharingServiceServer()
}

// UnimplementedRideSharingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRideSharingServiceServer struct{}

func (UnimplementedRideSharingServiceServer) RequestRide(context.Context, *RideRequest) (*RideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRide not implemented")
}
func (UnimplementedRideSharingServiceServer) GetRideStatus(context.Context, *RideStatusRequest) (*RideStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRideStatus not implemented")
}
func (UnimplementedRideSharingServiceServer) AcceptRide(context.Context, *AcceptRideRequest) (*AcceptRideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRide not implemented")
}
func (UnimplementedRideSharingServiceServer) RejectRide(context.Context, *RejectRideRequest) (*RejectRideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectRide not implemented")
}
func (UnimplementedRideSharingServiceServer) CompleteRide(context.Context, *RideCompletionRequest) (*RideCompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteRide not implemented")
}
func (UnimplementedRideSharingServiceServer) mustEmbedUnimplementedRideSharingServiceServer() {}
func (UnimplementedRideSharingServiceServer) testEmbeddedByValue()                            {}

// UnsafeRideSharingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RideSharingServiceServer will
// result in compilation errors.
type UnsafeRideSharingServiceServer interface {
	mustEmbedUnimplementedRideSharingServiceServer()
}

func RegisterRideSharingServiceServer(s grpc.ServiceRegistrar, srv RideSharingServiceServer) {
	// If the following call pancis, it indicates UnimplementedRideSharingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RideSharingService_ServiceDesc, srv)
}

func _RideSharingService_RequestRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideSharingServiceServer).RequestRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideSharingService_RequestRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideSharingServiceServer).RequestRide(ctx, req.(*RideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RideSharingService_GetRideStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RideStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideSharingServiceServer).GetRideStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideSharingService_GetRideStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideSharingServiceServer).GetRideStatus(ctx, req.(*RideStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RideSharingService_AcceptRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideSharingServiceServer).AcceptRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideSharingService_AcceptRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideSharingServiceServer).AcceptRide(ctx, req.(*AcceptRideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RideSharingService_RejectRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectRideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideSharingServiceServer).RejectRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideSharingService_RejectRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideSharingServiceServer).RejectRide(ctx, req.(*RejectRideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RideSharingService_CompleteRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RideCompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RideSharingServiceServer).CompleteRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RideSharingService_CompleteRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RideSharingServiceServer).CompleteRide(ctx, req.(*RideCompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RideSharingService_ServiceDesc is the grpc.ServiceDesc for RideSharingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RideSharingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ridesharing.RideSharingService",
	HandlerType: (*RideSharingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRide",
			Handler:    _RideSharingService_RequestRide_Handler,
		},
		{
			MethodName: "GetRideStatus",
			Handler:    _RideSharingService_GetRideStatus_Handler,
		},
		{
			MethodName: "AcceptRide",
			Handler:    _RideSharingService_AcceptRide_Handler,
		},
		{
			MethodName: "RejectRide",
			Handler:    _RideSharingService_RejectRide_Handler,
		},
		{
			MethodName: "CompleteRide",
			Handler:    _RideSharingService_CompleteRide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myuber.proto",
}