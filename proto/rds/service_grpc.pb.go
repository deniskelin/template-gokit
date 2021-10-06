// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rds

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RDSClient is the client API for RDS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RDSClient interface {
	GetBillingByEnvAndID(ctx context.Context, in *GetBillingByEnvAndIDRequest, opts ...grpc.CallOption) (*GetBillingByEnvAndIDResponse, error)
	GetBillingEnvRouteByAccountID(ctx context.Context, in *GetBillingEnvRouteByAccountIDRequest, opts ...grpc.CallOption) (*GetBillingEnvRouteByAccountIDResponse, error)
	SetRouteLabel(ctx context.Context, in *SetRouteLabelRequest, opts ...grpc.CallOption) (*SetRouteLabelResponse, error)
	GetRouteLabel(ctx context.Context, in *GetRouteLabelRequest, opts ...grpc.CallOption) (*GetRouteLabelResponse, error)
}

type rDSClient struct {
	cc grpc.ClientConnInterface
}

func NewRDSClient(cc grpc.ClientConnInterface) RDSClient {
	return &rDSClient{cc}
}

func (c *rDSClient) GetBillingByEnvAndID(ctx context.Context, in *GetBillingByEnvAndIDRequest, opts ...grpc.CallOption) (*GetBillingByEnvAndIDResponse, error) {
	out := new(GetBillingByEnvAndIDResponse)
	err := c.cc.Invoke(ctx, "/mtt.nmshl.rds.RDS/GetBillingByEnvAndID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) GetBillingEnvRouteByAccountID(ctx context.Context, in *GetBillingEnvRouteByAccountIDRequest, opts ...grpc.CallOption) (*GetBillingEnvRouteByAccountIDResponse, error) {
	out := new(GetBillingEnvRouteByAccountIDResponse)
	err := c.cc.Invoke(ctx, "/mtt.nmshl.rds.RDS/GetBillingEnvRouteByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) SetRouteLabel(ctx context.Context, in *SetRouteLabelRequest, opts ...grpc.CallOption) (*SetRouteLabelResponse, error) {
	out := new(SetRouteLabelResponse)
	err := c.cc.Invoke(ctx, "/mtt.nmshl.rds.RDS/SetRouteLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) GetRouteLabel(ctx context.Context, in *GetRouteLabelRequest, opts ...grpc.CallOption) (*GetRouteLabelResponse, error) {
	out := new(GetRouteLabelResponse)
	err := c.cc.Invoke(ctx, "/mtt.nmshl.rds.RDS/GetRouteLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RDSServer is the server API for RDS service.
// All implementations must embed UnimplementedRDSServer
// for forward compatibility
type RDSServer interface {
	GetBillingByEnvAndID(context.Context, *GetBillingByEnvAndIDRequest) (*GetBillingByEnvAndIDResponse, error)
	GetBillingEnvRouteByAccountID(context.Context, *GetBillingEnvRouteByAccountIDRequest) (*GetBillingEnvRouteByAccountIDResponse, error)
	SetRouteLabel(context.Context, *SetRouteLabelRequest) (*SetRouteLabelResponse, error)
	GetRouteLabel(context.Context, *GetRouteLabelRequest) (*GetRouteLabelResponse, error)
	mustEmbedUnimplementedRDSServer()
}

// UnimplementedRDSServer must be embedded to have forward compatible implementations.
type UnimplementedRDSServer struct {
}

func (UnimplementedRDSServer) GetBillingByEnvAndID(context.Context, *GetBillingByEnvAndIDRequest) (*GetBillingByEnvAndIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingByEnvAndID not implemented")
}
func (UnimplementedRDSServer) GetBillingEnvRouteByAccountID(context.Context, *GetBillingEnvRouteByAccountIDRequest) (*GetBillingEnvRouteByAccountIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingEnvRouteByAccountID not implemented")
}
func (UnimplementedRDSServer) SetRouteLabel(context.Context, *SetRouteLabelRequest) (*SetRouteLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRouteLabel not implemented")
}
func (UnimplementedRDSServer) GetRouteLabel(context.Context, *GetRouteLabelRequest) (*GetRouteLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteLabel not implemented")
}
func (UnimplementedRDSServer) mustEmbedUnimplementedRDSServer() {}

// UnsafeRDSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RDSServer will
// result in compilation errors.
type UnsafeRDSServer interface {
	mustEmbedUnimplementedRDSServer()
}

func RegisterRDSServer(s grpc.ServiceRegistrar, srv RDSServer) {
	s.RegisterService(&RDS_ServiceDesc, srv)
}

func _RDS_GetBillingByEnvAndID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingByEnvAndIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).GetBillingByEnvAndID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtt.nmshl.rds.RDS/GetBillingByEnvAndID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).GetBillingByEnvAndID(ctx, req.(*GetBillingByEnvAndIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_GetBillingEnvRouteByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingEnvRouteByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).GetBillingEnvRouteByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtt.nmshl.rds.RDS/GetBillingEnvRouteByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).GetBillingEnvRouteByAccountID(ctx, req.(*GetBillingEnvRouteByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_SetRouteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRouteLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).SetRouteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtt.nmshl.rds.RDS/SetRouteLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).SetRouteLabel(ctx, req.(*SetRouteLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_GetRouteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).GetRouteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtt.nmshl.rds.RDS/GetRouteLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).GetRouteLabel(ctx, req.(*GetRouteLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RDS_ServiceDesc is the grpc.ServiceDesc for RDS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RDS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtt.nmshl.rds.RDS",
	HandlerType: (*RDSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingByEnvAndID",
			Handler:    _RDS_GetBillingByEnvAndID_Handler,
		},
		{
			MethodName: "GetBillingEnvRouteByAccountID",
			Handler:    _RDS_GetBillingEnvRouteByAccountID_Handler,
		},
		{
			MethodName: "SetRouteLabel",
			Handler:    _RDS_SetRouteLabel_Handler,
		},
		{
			MethodName: "GetRouteLabel",
			Handler:    _RDS_GetRouteLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mtt/nmshl/rds/service.proto",
}