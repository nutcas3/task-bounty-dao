// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: taskbounty/v1/query.proto

package types

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
	Query_Tasks_FullMethodName           = "/taskbounty.v1.Query/Tasks"
	Query_Task_FullMethodName            = "/taskbounty.v1.Query/Task"
	Query_TasksByStatus_FullMethodName   = "/taskbounty.v1.Query/TasksByStatus"
	Query_TasksByCreator_FullMethodName  = "/taskbounty.v1.Query/TasksByCreator"
	Query_TasksByClaimant_FullMethodName = "/taskbounty.v1.Query/TasksByClaimant"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service
type QueryClient interface {
	// Tasks returns all tasks
	Tasks(ctx context.Context, in *QueryTasksRequest, opts ...grpc.CallOption) (*QueryTasksResponse, error)
	// Task returns task details by ID
	Task(ctx context.Context, in *QueryTaskRequest, opts ...grpc.CallOption) (*QueryTaskResponse, error)
	// TasksByStatus returns tasks filtered by status
	TasksByStatus(ctx context.Context, in *QueryTasksByStatusRequest, opts ...grpc.CallOption) (*QueryTasksByStatusResponse, error)
	// TasksByCreator returns tasks created by an address
	TasksByCreator(ctx context.Context, in *QueryTasksByCreatorRequest, opts ...grpc.CallOption) (*QueryTasksByCreatorResponse, error)
	// TasksByClaimant returns tasks claimed by an address
	TasksByClaimant(ctx context.Context, in *QueryTasksByClaimantRequest, opts ...grpc.CallOption) (*QueryTasksByClaimantResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Tasks(ctx context.Context, in *QueryTasksRequest, opts ...grpc.CallOption) (*QueryTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTasksResponse)
	err := c.cc.Invoke(ctx, Query_Tasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Task(ctx context.Context, in *QueryTaskRequest, opts ...grpc.CallOption) (*QueryTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTaskResponse)
	err := c.cc.Invoke(ctx, Query_Task_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TasksByStatus(ctx context.Context, in *QueryTasksByStatusRequest, opts ...grpc.CallOption) (*QueryTasksByStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTasksByStatusResponse)
	err := c.cc.Invoke(ctx, Query_TasksByStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TasksByCreator(ctx context.Context, in *QueryTasksByCreatorRequest, opts ...grpc.CallOption) (*QueryTasksByCreatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTasksByCreatorResponse)
	err := c.cc.Invoke(ctx, Query_TasksByCreator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TasksByClaimant(ctx context.Context, in *QueryTasksByClaimantRequest, opts ...grpc.CallOption) (*QueryTasksByClaimantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTasksByClaimantResponse)
	err := c.cc.Invoke(ctx, Query_TasksByClaimant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service
type QueryServer interface {
	// Tasks returns all tasks
	Tasks(context.Context, *QueryTasksRequest) (*QueryTasksResponse, error)
	// Task returns task details by ID
	Task(context.Context, *QueryTaskRequest) (*QueryTaskResponse, error)
	// TasksByStatus returns tasks filtered by status
	TasksByStatus(context.Context, *QueryTasksByStatusRequest) (*QueryTasksByStatusResponse, error)
	// TasksByCreator returns tasks created by an address
	TasksByCreator(context.Context, *QueryTasksByCreatorRequest) (*QueryTasksByCreatorResponse, error)
	// TasksByClaimant returns tasks claimed by an address
	TasksByClaimant(context.Context, *QueryTasksByClaimantRequest) (*QueryTasksByClaimantResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Tasks(context.Context, *QueryTasksRequest) (*QueryTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tasks not implemented")
}
func (UnimplementedQueryServer) Task(context.Context, *QueryTaskRequest) (*QueryTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Task not implemented")
}
func (UnimplementedQueryServer) TasksByStatus(context.Context, *QueryTasksByStatusRequest) (*QueryTasksByStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TasksByStatus not implemented")
}
func (UnimplementedQueryServer) TasksByCreator(context.Context, *QueryTasksByCreatorRequest) (*QueryTasksByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TasksByCreator not implemented")
}
func (UnimplementedQueryServer) TasksByClaimant(context.Context, *QueryTasksByClaimantRequest) (*QueryTasksByClaimantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TasksByClaimant not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Tasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Tasks(ctx, req.(*QueryTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Task_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Task(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Task_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Task(ctx, req.(*QueryTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TasksByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTasksByStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TasksByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TasksByStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TasksByStatus(ctx, req.(*QueryTasksByStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TasksByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTasksByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TasksByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TasksByCreator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TasksByCreator(ctx, req.(*QueryTasksByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TasksByClaimant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTasksByClaimantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TasksByClaimant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TasksByClaimant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TasksByClaimant(ctx, req.(*QueryTasksByClaimantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taskbounty.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tasks",
			Handler:    _Query_Tasks_Handler,
		},
		{
			MethodName: "Task",
			Handler:    _Query_Task_Handler,
		},
		{
			MethodName: "TasksByStatus",
			Handler:    _Query_TasksByStatus_Handler,
		},
		{
			MethodName: "TasksByCreator",
			Handler:    _Query_TasksByCreator_Handler,
		},
		{
			MethodName: "TasksByClaimant",
			Handler:    _Query_TasksByClaimant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskbounty/v1/query.proto",
}
