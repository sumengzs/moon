// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: prom/v1/rule.proto

package v1

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

// RuleClient is the client API for Rule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleClient interface {
	CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleReply, error)
	UpdateRule(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*UpdateRuleReply, error)
	DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*DeleteRuleReply, error)
	GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*GetRuleReply, error)
	ListRule(ctx context.Context, in *ListRuleRequest, opts ...grpc.CallOption) (*ListRuleReply, error)
}

type ruleClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleClient(cc grpc.ClientConnInterface) RuleClient {
	return &ruleClient{cc}
}

func (c *ruleClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleReply, error) {
	out := new(CreateRuleReply)
	err := c.cc.Invoke(ctx, "/api.prom.Rule/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleClient) UpdateRule(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*UpdateRuleReply, error) {
	out := new(UpdateRuleReply)
	err := c.cc.Invoke(ctx, "/api.prom.Rule/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleClient) DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*DeleteRuleReply, error) {
	out := new(DeleteRuleReply)
	err := c.cc.Invoke(ctx, "/api.prom.Rule/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleClient) GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*GetRuleReply, error) {
	out := new(GetRuleReply)
	err := c.cc.Invoke(ctx, "/api.prom.Rule/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleClient) ListRule(ctx context.Context, in *ListRuleRequest, opts ...grpc.CallOption) (*ListRuleReply, error) {
	out := new(ListRuleReply)
	err := c.cc.Invoke(ctx, "/api.prom.Rule/ListRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleServer is the server API for Rule service.
// All implementations must embed UnimplementedRuleServer
// for forward compatibility
type RuleServer interface {
	CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleReply, error)
	UpdateRule(context.Context, *UpdateRuleRequest) (*UpdateRuleReply, error)
	DeleteRule(context.Context, *DeleteRuleRequest) (*DeleteRuleReply, error)
	GetRule(context.Context, *GetRuleRequest) (*GetRuleReply, error)
	ListRule(context.Context, *ListRuleRequest) (*ListRuleReply, error)
	mustEmbedUnimplementedRuleServer()
}

// UnimplementedRuleServer must be embedded to have forward compatible implementations.
type UnimplementedRuleServer struct {
}

func (UnimplementedRuleServer) CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRule not implemented")
}
func (UnimplementedRuleServer) UpdateRule(context.Context, *UpdateRuleRequest) (*UpdateRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (UnimplementedRuleServer) DeleteRule(context.Context, *DeleteRuleRequest) (*DeleteRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedRuleServer) GetRule(context.Context, *GetRuleRequest) (*GetRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (UnimplementedRuleServer) ListRule(context.Context, *ListRuleRequest) (*ListRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRule not implemented")
}
func (UnimplementedRuleServer) mustEmbedUnimplementedRuleServer() {}

// UnsafeRuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleServer will
// result in compilation errors.
type UnsafeRuleServer interface {
	mustEmbedUnimplementedRuleServer()
}

func RegisterRuleServer(s grpc.ServiceRegistrar, srv RuleServer) {
	s.RegisterService(&Rule_ServiceDesc, srv)
}

func _Rule_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.Rule/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServer).CreateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rule_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.Rule/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServer).UpdateRule(ctx, req.(*UpdateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rule_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.Rule/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServer).DeleteRule(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rule_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.Rule/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServer).GetRule(ctx, req.(*GetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rule_ListRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServer).ListRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.Rule/ListRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServer).ListRule(ctx, req.(*ListRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rule_ServiceDesc is the grpc.ServiceDesc for Rule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.prom.Rule",
	HandlerType: (*RuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _Rule_CreateRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Rule_UpdateRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Rule_DeleteRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Rule_GetRule_Handler,
		},
		{
			MethodName: "ListRule",
			Handler:    _Rule_ListRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prom/v1/rule.proto",
}