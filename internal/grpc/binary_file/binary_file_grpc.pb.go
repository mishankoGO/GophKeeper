// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: binary_file.proto

package api

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

const (
	BinaryFiles_Get_FullMethodName    = "/api.BinaryFiles/Get"
	BinaryFiles_Update_FullMethodName = "/api.BinaryFiles/Update"
	BinaryFiles_Delete_FullMethodName = "/api.BinaryFiles/Delete"
)

// BinaryFilesClient is the client API for BinaryFiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BinaryFilesClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteRequest, error)
}

type binaryFilesClient struct {
	cc grpc.ClientConnInterface
}

func NewBinaryFilesClient(cc grpc.ClientConnInterface) BinaryFilesClient {
	return &binaryFilesClient{cc}
}

func (c *binaryFilesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, BinaryFiles_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryFilesClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, BinaryFiles_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryFilesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteRequest, error) {
	out := new(DeleteRequest)
	err := c.cc.Invoke(ctx, BinaryFiles_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinaryFilesServer is the server API for BinaryFiles service.
// All implementations must embed UnimplementedBinaryFilesServer
// for forward compatibility
type BinaryFilesServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteRequest, error)
	mustEmbedUnimplementedBinaryFilesServer()
}

// UnimplementedBinaryFilesServer must be embedded to have forward compatible implementations.
type UnimplementedBinaryFilesServer struct {
}

func (UnimplementedBinaryFilesServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBinaryFilesServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBinaryFilesServer) Delete(context.Context, *DeleteRequest) (*DeleteRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBinaryFilesServer) mustEmbedUnimplementedBinaryFilesServer() {}

// UnsafeBinaryFilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BinaryFilesServer will
// result in compilation errors.
type UnsafeBinaryFilesServer interface {
	mustEmbedUnimplementedBinaryFilesServer()
}

func RegisterBinaryFilesServer(s grpc.ServiceRegistrar, srv BinaryFilesServer) {
	s.RegisterService(&BinaryFiles_ServiceDesc, srv)
}

func _BinaryFiles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryFilesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinaryFiles_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryFilesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryFiles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryFilesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinaryFiles_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryFilesServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryFiles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryFilesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BinaryFiles_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryFilesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BinaryFiles_ServiceDesc is the grpc.ServiceDesc for BinaryFiles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BinaryFiles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BinaryFiles",
	HandlerType: (*BinaryFilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BinaryFiles_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BinaryFiles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BinaryFiles_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "binary_file.proto",
}
