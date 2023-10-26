// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: usercenter.proto

package usercenter

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
	Usercenter_Register_FullMethodName          = "/usercenter.Usercenter/Register"
	Usercenter_Login_FullMethodName             = "/usercenter.Usercenter/Login"
	Usercenter_GetUserInfo_FullMethodName       = "/usercenter.Usercenter/GetUserInfo"
	Usercenter_GetUserByMobile_FullMethodName   = "/usercenter.Usercenter/GetUserByMobile"
	Usercenter_UpdateUser_FullMethodName        = "/usercenter.Usercenter/UpdateUser"
	Usercenter_ExistUserByUserId_FullMethodName = "/usercenter.Usercenter/ExistUserByUserId"
	Usercenter_ExistUserByMobile_FullMethodName = "/usercenter.Usercenter/ExistUserByMobile"
	Usercenter_GenerateToken_FullMethodName     = "/usercenter.Usercenter/GenerateToken"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	GetUserByMobile(ctx context.Context, in *GetUserByMobileReq, opts ...grpc.CallOption) (*GetUserByMobileResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	ExistUserByUserId(ctx context.Context, in *ExistUserByUserIdReq, opts ...grpc.CallOption) (*ExistUserByUserIdResp, error)
	ExistUserByMobile(ctx context.Context, in *ExistUserByMobileReq, opts ...grpc.CallOption) (*ExistUserByMobileResp, error)
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Usercenter_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Usercenter_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, Usercenter_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserByMobile(ctx context.Context, in *GetUserByMobileReq, opts ...grpc.CallOption) (*GetUserByMobileResp, error) {
	out := new(GetUserByMobileResp)
	err := c.cc.Invoke(ctx, Usercenter_GetUserByMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, Usercenter_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) ExistUserByUserId(ctx context.Context, in *ExistUserByUserIdReq, opts ...grpc.CallOption) (*ExistUserByUserIdResp, error) {
	out := new(ExistUserByUserIdResp)
	err := c.cc.Invoke(ctx, Usercenter_ExistUserByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) ExistUserByMobile(ctx context.Context, in *ExistUserByMobileReq, opts ...grpc.CallOption) (*ExistUserByMobileResp, error) {
	out := new(ExistUserByMobileResp)
	err := c.cc.Invoke(ctx, Usercenter_ExistUserByMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, Usercenter_GenerateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	GetUserByMobile(context.Context, *GetUserByMobileReq) (*GetUserByMobileResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	ExistUserByUserId(context.Context, *ExistUserByUserIdReq) (*ExistUserByUserIdResp, error)
	ExistUserByMobile(context.Context, *ExistUserByMobileReq) (*ExistUserByMobileResp, error)
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsercenterServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsercenterServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUsercenterServer) GetUserByMobile(context.Context, *GetUserByMobileReq) (*GetUserByMobileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (UnimplementedUsercenterServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsercenterServer) ExistUserByUserId(context.Context, *ExistUserByUserIdReq) (*ExistUserByUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUserByUserId not implemented")
}
func (UnimplementedUsercenterServer) ExistUserByMobile(context.Context, *ExistUserByMobileReq) (*ExistUserByMobileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUserByMobile not implemented")
}
func (UnimplementedUsercenterServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByMobileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetUserByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserByMobile(ctx, req.(*GetUserByMobileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_ExistUserByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistUserByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).ExistUserByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_ExistUserByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).ExistUserByUserId(ctx, req.(*ExistUserByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_ExistUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistUserByMobileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).ExistUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_ExistUserByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).ExistUserByMobile(ctx, req.(*ExistUserByMobileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GenerateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.Usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Usercenter_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Usercenter_Login_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Usercenter_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _Usercenter_GetUserByMobile_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Usercenter_UpdateUser_Handler,
		},
		{
			MethodName: "ExistUserByUserId",
			Handler:    _Usercenter_ExistUserByUserId_Handler,
		},
		{
			MethodName: "ExistUserByMobile",
			Handler:    _Usercenter_ExistUserByMobile_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _Usercenter_GenerateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}