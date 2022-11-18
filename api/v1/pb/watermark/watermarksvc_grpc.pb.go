// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: watermarksvc.proto

package watermark

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

// WatermarkClient is the client API for Watermark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatermarkClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type watermarkClient struct {
	cc grpc.ClientConnInterface
}

func NewWatermarkClient(cc grpc.ClientConnInterface) WatermarkClient {
	return &watermarkClient{cc}
}

func (c *watermarkClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error) {
	out := new(WatermarkReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Watermark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentReply, error) {
	out := new(AddDocumentReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/AddDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatermarkServer is the server API for Watermark service.
// All implementations must embed UnimplementedWatermarkServer
// for forward compatibility
type WatermarkServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
	mustEmbedUnimplementedWatermarkServer()
}

// UnimplementedWatermarkServer must be embedded to have forward compatible implementations.
type UnimplementedWatermarkServer struct {
}

func (UnimplementedWatermarkServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWatermarkServer) Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watermark not implemented")
}
func (UnimplementedWatermarkServer) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedWatermarkServer) AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocument not implemented")
}
func (UnimplementedWatermarkServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedWatermarkServer) mustEmbedUnimplementedWatermarkServer() {}

// UnsafeWatermarkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatermarkServer will
// result in compilation errors.
type UnsafeWatermarkServer interface {
	mustEmbedUnimplementedWatermarkServer()
}

func RegisterWatermarkServer(s grpc.ServiceRegistrar, srv WatermarkServer) {
	s.RegisterService(&Watermark_ServiceDesc, srv)
}

func _Watermark_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Watermark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Watermark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Watermark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Watermark(ctx, req.(*WatermarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_AddDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).AddDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/AddDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).AddDocument(ctx, req.(*AddDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Watermark_ServiceDesc is the grpc.ServiceDesc for Watermark service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watermark_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Watermark",
	HandlerType: (*WatermarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Watermark_Get_Handler,
		},
		{
			MethodName: "Watermark",
			Handler:    _Watermark_Watermark_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Watermark_Status_Handler,
		},
		{
			MethodName: "AddDocument",
			Handler:    _Watermark_AddDocument_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Watermark_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watermarksvc.proto",
}