// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: friend.proto

package monify

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

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*FriendEmpty, error)
	ListFriend(ctx context.Context, in *FriendEmpty, opts ...grpc.CallOption) (*ListFriendResponse, error)
	InviteFriend(ctx context.Context, in *InviteFriendRequest, opts ...grpc.CallOption) (*InviteFriendResponse, error)
	ListFriendInvitation(ctx context.Context, in *FriendEmpty, opts ...grpc.CallOption) (*ListFriendInvitationResponse, error)
	AcceptInvitation(ctx context.Context, in *AcceptInvitationRequest, opts ...grpc.CallOption) (*FriendEmpty, error)
	RejectInvitation(ctx context.Context, in *RejectInvitationRequest, opts ...grpc.CallOption) (*FriendEmpty, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*FriendEmpty, error) {
	out := new(FriendEmpty)
	err := c.cc.Invoke(ctx, "/FriendService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) ListFriend(ctx context.Context, in *FriendEmpty, opts ...grpc.CallOption) (*ListFriendResponse, error) {
	out := new(ListFriendResponse)
	err := c.cc.Invoke(ctx, "/FriendService/ListFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) InviteFriend(ctx context.Context, in *InviteFriendRequest, opts ...grpc.CallOption) (*InviteFriendResponse, error) {
	out := new(InviteFriendResponse)
	err := c.cc.Invoke(ctx, "/FriendService/InviteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) ListFriendInvitation(ctx context.Context, in *FriendEmpty, opts ...grpc.CallOption) (*ListFriendInvitationResponse, error) {
	out := new(ListFriendInvitationResponse)
	err := c.cc.Invoke(ctx, "/FriendService/ListFriendInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) AcceptInvitation(ctx context.Context, in *AcceptInvitationRequest, opts ...grpc.CallOption) (*FriendEmpty, error) {
	out := new(FriendEmpty)
	err := c.cc.Invoke(ctx, "/FriendService/AcceptInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) RejectInvitation(ctx context.Context, in *RejectInvitationRequest, opts ...grpc.CallOption) (*FriendEmpty, error) {
	out := new(FriendEmpty)
	err := c.cc.Invoke(ctx, "/FriendService/RejectInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations must embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	DeleteFriend(context.Context, *DeleteFriendRequest) (*FriendEmpty, error)
	ListFriend(context.Context, *FriendEmpty) (*ListFriendResponse, error)
	InviteFriend(context.Context, *InviteFriendRequest) (*InviteFriendResponse, error)
	ListFriendInvitation(context.Context, *FriendEmpty) (*ListFriendInvitationResponse, error)
	AcceptInvitation(context.Context, *AcceptInvitationRequest) (*FriendEmpty, error)
	RejectInvitation(context.Context, *RejectInvitationRequest) (*FriendEmpty, error)
	mustEmbedUnimplementedFriendServiceServer()
}

// UnimplementedFriendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*FriendEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendServiceServer) ListFriend(context.Context, *FriendEmpty) (*ListFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriend not implemented")
}
func (UnimplementedFriendServiceServer) InviteFriend(context.Context, *InviteFriendRequest) (*InviteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteFriend not implemented")
}
func (UnimplementedFriendServiceServer) ListFriendInvitation(context.Context, *FriendEmpty) (*ListFriendInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriendInvitation not implemented")
}
func (UnimplementedFriendServiceServer) AcceptInvitation(context.Context, *AcceptInvitationRequest) (*FriendEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvitation not implemented")
}
func (UnimplementedFriendServiceServer) RejectInvitation(context.Context, *RejectInvitationRequest) (*FriendEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectInvitation not implemented")
}
func (UnimplementedFriendServiceServer) mustEmbedUnimplementedFriendServiceServer() {}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_ListFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).ListFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/ListFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).ListFriend(ctx, req.(*FriendEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_InviteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).InviteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/InviteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).InviteFriend(ctx, req.(*InviteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_ListFriendInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).ListFriendInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/ListFriendInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).ListFriendInvitation(ctx, req.(*FriendEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_AcceptInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).AcceptInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/AcceptInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).AcceptInvitation(ctx, req.(*AcceptInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_RejectInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).RejectInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/RejectInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).RejectInvitation(ctx, req.(*RejectInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FriendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFriend",
			Handler:    _FriendService_DeleteFriend_Handler,
		},
		{
			MethodName: "ListFriend",
			Handler:    _FriendService_ListFriend_Handler,
		},
		{
			MethodName: "InviteFriend",
			Handler:    _FriendService_InviteFriend_Handler,
		},
		{
			MethodName: "ListFriendInvitation",
			Handler:    _FriendService_ListFriendInvitation_Handler,
		},
		{
			MethodName: "AcceptInvitation",
			Handler:    _FriendService_AcceptInvitation_Handler,
		},
		{
			MethodName: "RejectInvitation",
			Handler:    _FriendService_RejectInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend.proto",
}
