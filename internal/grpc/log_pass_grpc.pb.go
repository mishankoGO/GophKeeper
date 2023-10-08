// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: log_pass.proto

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
	LogPasses_Insert_FullMethodName = "/api.LogPasses/Insert"
	LogPasses_Get_FullMethodName    = "/api.LogPasses/Get"
	LogPasses_Update_FullMethodName = "/api.LogPasses/Update"
	LogPasses_Delete_FullMethodName = "/api.LogPasses/Delete"
	LogPasses_List_FullMethodName   = "/api.LogPasses/List"
)

// LogPassesClient is the client API for LogPasses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogPassesClient interface {
	Insert(ctx context.Context, in *InsertLogPassRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLogPassResponse, error)
	Update(ctx context.Context, in *UpdateLogPassRequest, opts ...grpc.CallOption) (*UpdateLogPassResponse, error)
	Delete(ctx context.Context, in *DeleteLogPassRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListLogPassRequest, opts ...grpc.CallOption) (*ListLogPassResponse, error)
}

type logPassesClient struct {
	cc grpc.ClientConnInterface
}

func NewLogPassesClient(cc grpc.ClientConnInterface) LogPassesClient {
	return &logPassesClient{cc}
}

func (c *logPassesClient) Insert(ctx context.Context, in *InsertLogPassRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, LogPasses_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPassesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetLogPassResponse, error) {
	out := new(GetLogPassResponse)
	err := c.cc.Invoke(ctx, LogPasses_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPassesClient) Update(ctx context.Context, in *UpdateLogPassRequest, opts ...grpc.CallOption) (*UpdateLogPassResponse, error) {
	out := new(UpdateLogPassResponse)
	err := c.cc.Invoke(ctx, LogPasses_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPassesClient) Delete(ctx context.Context, in *DeleteLogPassRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, LogPasses_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPassesClient) List(ctx context.Context, in *ListLogPassRequest, opts ...grpc.CallOption) (*ListLogPassResponse, error) {
	out := new(ListLogPassResponse)
	err := c.cc.Invoke(ctx, LogPasses_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogPassesServer is the server API for LogPasses service.
// All implementations must embed UnimplementedLogPassesServer
// for forward compatibility
type LogPassesServer interface {
	Insert(context.Context, *InsertLogPassRequest) (*InsertResponse, error)
	Get(context.Context, *GetRequest) (*GetLogPassResponse, error)
	Update(context.Context, *UpdateLogPassRequest) (*UpdateLogPassResponse, error)
	Delete(context.Context, *DeleteLogPassRequest) (*DeleteResponse, error)
	List(context.Context, *ListLogPassRequest) (*ListLogPassResponse, error)
	mustEmbedUnimplementedLogPassesServer()
}

// UnimplementedLogPassesServer must be embedded to have forward compatible implementations.
type UnimplementedLogPassesServer struct {
}

func (UnimplementedLogPassesServer) Insert(context.Context, *InsertLogPassRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedLogPassesServer) Get(context.Context, *GetRequest) (*GetLogPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLogPassesServer) Update(context.Context, *UpdateLogPassRequest) (*UpdateLogPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLogPassesServer) Delete(context.Context, *DeleteLogPassRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLogPassesServer) List(context.Context, *ListLogPassRequest) (*ListLogPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLogPassesServer) mustEmbedUnimplementedLogPassesServer() {}

// UnsafeLogPassesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogPassesServer will
// result in compilation errors.
type UnsafeLogPassesServer interface {
	mustEmbedUnimplementedLogPassesServer()
}

func RegisterLogPassesServer(s grpc.ServiceRegistrar, srv LogPassesServer) {
	s.RegisterService(&LogPasses_ServiceDesc, srv)
}

func _LogPasses_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertLogPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPassesServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogPasses_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPassesServer).Insert(ctx, req.(*InsertLogPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPasses_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPassesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogPasses_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPassesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPasses_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLogPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPassesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogPasses_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPassesServer).Update(ctx, req.(*UpdateLogPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPasses_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPassesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogPasses_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPassesServer).Delete(ctx, req.(*DeleteLogPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPasses_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPassesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogPasses_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPassesServer).List(ctx, req.(*ListLogPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogPasses_ServiceDesc is the grpc.ServiceDesc for LogPasses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogPasses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.LogPasses",
	HandlerType: (*LogPassesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _LogPasses_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LogPasses_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LogPasses_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LogPasses_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LogPasses_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log_pass.proto",
}
