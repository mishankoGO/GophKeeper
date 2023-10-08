// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: text.proto

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
	Texts_Insert_FullMethodName = "/api.Texts/Insert"
	Texts_Get_FullMethodName    = "/api.Texts/Get"
	Texts_Update_FullMethodName = "/api.Texts/Update"
	Texts_Delete_FullMethodName = "/api.Texts/Delete"
	Texts_List_FullMethodName   = "/api.Texts/List"
)

// TextsClient is the client API for Texts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextsClient interface {
	Insert(ctx context.Context, in *InsertTextRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetTextResponse, error)
	Update(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error)
	Delete(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListTextRequest, opts ...grpc.CallOption) (*ListTextResponse, error)
}

type textsClient struct {
	cc grpc.ClientConnInterface
}

func NewTextsClient(cc grpc.ClientConnInterface) TextsClient {
	return &textsClient{cc}
}

func (c *textsClient) Insert(ctx context.Context, in *InsertTextRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, Texts_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetTextResponse, error) {
	out := new(GetTextResponse)
	err := c.cc.Invoke(ctx, Texts_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textsClient) Update(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error) {
	out := new(UpdateTextResponse)
	err := c.cc.Invoke(ctx, Texts_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textsClient) Delete(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Texts_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textsClient) List(ctx context.Context, in *ListTextRequest, opts ...grpc.CallOption) (*ListTextResponse, error) {
	out := new(ListTextResponse)
	err := c.cc.Invoke(ctx, Texts_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextsServer is the server API for Texts service.
// All implementations must embed UnimplementedTextsServer
// for forward compatibility
type TextsServer interface {
	Insert(context.Context, *InsertTextRequest) (*InsertResponse, error)
	Get(context.Context, *GetRequest) (*GetTextResponse, error)
	Update(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error)
	Delete(context.Context, *DeleteTextRequest) (*DeleteResponse, error)
	List(context.Context, *ListTextRequest) (*ListTextResponse, error)
	mustEmbedUnimplementedTextsServer()
}

// UnimplementedTextsServer must be embedded to have forward compatible implementations.
type UnimplementedTextsServer struct {
}

func (UnimplementedTextsServer) Insert(context.Context, *InsertTextRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTextsServer) Get(context.Context, *GetRequest) (*GetTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTextsServer) Update(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTextsServer) Delete(context.Context, *DeleteTextRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTextsServer) List(context.Context, *ListTextRequest) (*ListTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTextsServer) mustEmbedUnimplementedTextsServer() {}

// UnsafeTextsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextsServer will
// result in compilation errors.
type UnsafeTextsServer interface {
	mustEmbedUnimplementedTextsServer()
}

func RegisterTextsServer(s grpc.ServiceRegistrar, srv TextsServer) {
	s.RegisterService(&Texts_ServiceDesc, srv)
}

func _Texts_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Texts_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextsServer).Insert(ctx, req.(*InsertTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Texts_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Texts_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Texts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Texts_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextsServer).Update(ctx, req.(*UpdateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Texts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Texts_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextsServer).Delete(ctx, req.(*DeleteTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Texts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Texts_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextsServer).List(ctx, req.(*ListTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Texts_ServiceDesc is the grpc.ServiceDesc for Texts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Texts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Texts",
	HandlerType: (*TextsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Texts_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Texts_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Texts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Texts_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Texts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "text.proto",
}
