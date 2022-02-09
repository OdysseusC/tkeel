// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	// create a tenant.
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	// get a tenant.
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	// list tenant.
	ListTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTenantResponse, error)
	// delete a tenant.
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// create a user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	//  get user.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// list user.
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	// delete a user.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// update user
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	// add a plugin
	AddTenantPlugin(ctx context.Context, in *AddTenantPluginRequest, opts ...grpc.CallOption) (*AddTenantPluginResponse, error)
	// list plugin
	ListTenantPlugin(ctx context.Context, in *ListTenantPluginRequest, opts ...grpc.CallOption) (*ListTenantPluginResponse, error)
	// delete plugin
	DeleteTenantPlugin(ctx context.Context, in *DeleteTenantPluginRequest, opts ...grpc.CallOption) (*DeleteTenantPluginResponse, error)
	TenantPluginPermissible(ctx context.Context, in *PluginPermissibleRequest, opts ...grpc.CallOption) (*PluginPermissibleResponse, error)
	GetResetPasswordKey(ctx context.Context, in *GetResetPasswordKeyRequest, opts ...grpc.CallOption) (*GetResetPasswordKeyResponse, error)
	ResetPasswordKeyInfo(ctx context.Context, in *RPKInfoRequest, opts ...grpc.CallOption) (*RPKInfoResponse, error)
}

type tenantClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantClient(cc grpc.ClientConnInterface) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) ListTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTenantResponse, error) {
	out := new(ListTenantResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/ListTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) AddTenantPlugin(ctx context.Context, in *AddTenantPluginRequest, opts ...grpc.CallOption) (*AddTenantPluginResponse, error) {
	out := new(AddTenantPluginResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/AddTenantPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) ListTenantPlugin(ctx context.Context, in *ListTenantPluginRequest, opts ...grpc.CallOption) (*ListTenantPluginResponse, error) {
	out := new(ListTenantPluginResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/ListTenantPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) DeleteTenantPlugin(ctx context.Context, in *DeleteTenantPluginRequest, opts ...grpc.CallOption) (*DeleteTenantPluginResponse, error) {
	out := new(DeleteTenantPluginResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/DeleteTenantPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) TenantPluginPermissible(ctx context.Context, in *PluginPermissibleRequest, opts ...grpc.CallOption) (*PluginPermissibleResponse, error) {
	out := new(PluginPermissibleResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/TenantPluginPermissible", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetResetPasswordKey(ctx context.Context, in *GetResetPasswordKeyRequest, opts ...grpc.CallOption) (*GetResetPasswordKeyResponse, error) {
	out := new(GetResetPasswordKeyResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/GetResetPasswordKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) ResetPasswordKeyInfo(ctx context.Context, in *RPKInfoRequest, opts ...grpc.CallOption) (*RPKInfoResponse, error) {
	out := new(RPKInfoResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.security.api.tenant.v1.Tenant/ResetPasswordKeyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	// create a tenant.
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	// get a tenant.
	GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	// list tenant.
	ListTenant(context.Context, *emptypb.Empty) (*ListTenantResponse, error)
	// delete a tenant.
	DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error)
	// create a user.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	//  get user.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// list user.
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	// delete a user.
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	// update user
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// add a plugin
	AddTenantPlugin(context.Context, *AddTenantPluginRequest) (*AddTenantPluginResponse, error)
	// list plugin
	ListTenantPlugin(context.Context, *ListTenantPluginRequest) (*ListTenantPluginResponse, error)
	// delete plugin
	DeleteTenantPlugin(context.Context, *DeleteTenantPluginRequest) (*DeleteTenantPluginResponse, error)
	TenantPluginPermissible(context.Context, *PluginPermissibleRequest) (*PluginPermissibleResponse, error)
	GetResetPasswordKey(context.Context, *GetResetPasswordKeyRequest) (*GetResetPasswordKeyResponse, error)
	ResetPasswordKeyInfo(context.Context, *RPKInfoRequest) (*RPKInfoResponse, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
}

func (UnimplementedTenantServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantServer) GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantServer) ListTenant(context.Context, *emptypb.Empty) (*ListTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenant not implemented")
}
func (UnimplementedTenantServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTenantServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedTenantServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedTenantServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedTenantServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedTenantServer) AddTenantPlugin(context.Context, *AddTenantPluginRequest) (*AddTenantPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTenantPlugin not implemented")
}
func (UnimplementedTenantServer) ListTenantPlugin(context.Context, *ListTenantPluginRequest) (*ListTenantPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenantPlugin not implemented")
}
func (UnimplementedTenantServer) DeleteTenantPlugin(context.Context, *DeleteTenantPluginRequest) (*DeleteTenantPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenantPlugin not implemented")
}
func (UnimplementedTenantServer) TenantPluginPermissible(context.Context, *PluginPermissibleRequest) (*PluginPermissibleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantPluginPermissible not implemented")
}
func (UnimplementedTenantServer) GetResetPasswordKey(context.Context, *GetResetPasswordKeyRequest) (*GetResetPasswordKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResetPasswordKey not implemented")
}
func (UnimplementedTenantServer) ResetPasswordKeyInfo(context.Context, *RPKInfoRequest) (*RPKInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordKeyInfo not implemented")
}
func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_ListTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).ListTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/ListTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).ListTenant(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_AddTenantPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTenantPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).AddTenantPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/AddTenantPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).AddTenantPlugin(ctx, req.(*AddTenantPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_ListTenantPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).ListTenantPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/ListTenantPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).ListTenantPlugin(ctx, req.(*ListTenantPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_DeleteTenantPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).DeleteTenantPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/DeleteTenantPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).DeleteTenantPlugin(ctx, req.(*DeleteTenantPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_TenantPluginPermissible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginPermissibleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).TenantPluginPermissible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/TenantPluginPermissible",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).TenantPluginPermissible(ctx, req.(*PluginPermissibleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetResetPasswordKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResetPasswordKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetResetPasswordKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/GetResetPasswordKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetResetPasswordKey(ctx, req.(*GetResetPasswordKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_ResetPasswordKeyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPKInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).ResetPasswordKeyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.security.api.tenant.v1.Tenant/ResetPasswordKeyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).ResetPasswordKeyInfo(ctx, req.(*RPKInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.tkeel.security.api.tenant.v1.Tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _Tenant_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _Tenant_GetTenant_Handler,
		},
		{
			MethodName: "ListTenant",
			Handler:    _Tenant_ListTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _Tenant_DeleteTenant_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Tenant_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Tenant_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _Tenant_ListUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Tenant_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Tenant_UpdateUser_Handler,
		},
		{
			MethodName: "AddTenantPlugin",
			Handler:    _Tenant_AddTenantPlugin_Handler,
		},
		{
			MethodName: "ListTenantPlugin",
			Handler:    _Tenant_ListTenantPlugin_Handler,
		},
		{
			MethodName: "DeleteTenantPlugin",
			Handler:    _Tenant_DeleteTenantPlugin_Handler,
		},
		{
			MethodName: "TenantPluginPermissible",
			Handler:    _Tenant_TenantPluginPermissible_Handler,
		},
		{
			MethodName: "GetResetPasswordKey",
			Handler:    _Tenant_GetResetPasswordKey_Handler,
		},
		{
			MethodName: "ResetPasswordKeyInfo",
			Handler:    _Tenant_ResetPasswordKeyInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tenant/v1/tenant.proto",
}
