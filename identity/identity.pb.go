// Code generated by protoc-gen-go.
// source: identity.proto
// DO NOT EDIT!

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	identity.proto

It has these top-level messages:
	Empty
	ListUserResponse
	CreateUserRequest
	User
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListUserResponse struct {
	List []*User `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ListUserResponse) Reset()                    { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()               {}
func (*ListUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListUserResponse) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

type CreateUserRequest struct {
	Username    string            `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	FirstName   string            `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName    string            `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email       string            `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	MobilePhone string            `protobuf:"bytes,5,opt,name=mobile_phone,json=mobilePhone" json:"mobile_phone,omitempty"`
	Password    string            `protobuf:"bytes,6,opt,name=password" json:"password,omitempty"`
	Group       string            `protobuf:"bytes,7,opt,name=group" json:"group,omitempty"`
	Fields      map[string]string `protobuf:"bytes,8,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateUserRequest) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type User struct {
	UserId      int32             `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Username    string            `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	FirstName   string            `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName    string            `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email       string            `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	MobilePhone string            `protobuf:"bytes,6,opt,name=mobile_phone,json=mobilePhone" json:"mobile_phone,omitempty"`
	Group       string            `protobuf:"bytes,7,opt,name=group" json:"group,omitempty"`
	Fields      map[string]string `protobuf:"bytes,8,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *User) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "identity.Empty")
	proto.RegisterType((*ListUserResponse)(nil), "identity.ListUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "identity.CreateUserRequest")
	proto.RegisterType((*User)(nil), "identity.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserManager service

type UserManagerClient interface {
	ListUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
}

type userManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerClient(cc *grpc.ClientConn) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) ListUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := grpc.Invoke(ctx, "/identity.UserManager/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/identity.UserManager/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/identity.UserManager/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/identity.UserManager/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserManager service

type UserManagerServer interface {
	ListUsers(context.Context, *Empty) (*ListUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*Empty, error)
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.UserManager/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).ListUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.UserManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.UserManager/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.UserManager/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identity.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserManager_ListUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserManager_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManager_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManager_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0x2f, 0xb9, 0x4b, 0xee, 0x6e, 0x22, 0xb5, 0x2e, 0x42, 0x43, 0x8a, 0x50, 0x17, 0xc4,
	0x3e, 0x45, 0x38, 0x41, 0xb4, 0x3e, 0x08, 0xd5, 0x0a, 0x82, 0x4a, 0x09, 0xf4, 0xf9, 0xd8, 0x9a,
	0x69, 0xbb, 0x74, 0x93, 0x8d, 0xbb, 0x1b, 0x4b, 0x3e, 0x98, 0x5f, 0xc4, 0x2f, 0xe1, 0xd7, 0x90,
	0xdd, 0x34, 0xb9, 0x36, 0x6d, 0xf3, 0xe4, 0x5b, 0xfe, 0xf3, 0x9f, 0x99, 0xec, 0xfe, 0x76, 0x06,
	0xb6, 0x78, 0x8e, 0xa5, 0xe1, 0xa6, 0x49, 0x2b, 0x25, 0x8d, 0x24, 0x8b, 0x4e, 0xd3, 0x39, 0x04,
	0x47, 0x45, 0x65, 0x1a, 0xfa, 0x06, 0xb6, 0xbf, 0x72, 0x6d, 0x4e, 0x34, 0xaa, 0x0c, 0x75, 0x25,
	0x4b, 0x8d, 0x84, 0xc2, 0x4c, 0x70, 0x6d, 0x62, 0x6f, 0x6f, 0xba, 0x1f, 0xad, 0xb6, 0xd2, 0xbe,
	0x8b, 0xcb, 0x72, 0x1e, 0xfd, 0xe3, 0xc3, 0x93, 0x8f, 0x0a, 0x99, 0xc1, 0xb6, 0xf4, 0x67, 0x8d,
	0xda, 0x90, 0x04, 0x16, 0xb5, 0x46, 0x55, 0xb2, 0x02, 0x63, 0x6f, 0xcf, 0xdb, 0x5f, 0x66, 0xbd,
	0x26, 0xcf, 0x00, 0xce, 0xb8, 0xd2, 0x66, 0xed, 0x5c, 0xdf, 0xb9, 0x4b, 0x17, 0xf9, 0x6e, 0xed,
	0x5d, 0x58, 0x0a, 0xd6, 0xb9, 0xd3, 0xb6, 0xd6, 0x06, 0x9c, 0xf9, 0x14, 0x02, 0x2c, 0x18, 0x17,
	0xf1, 0xcc, 0x19, 0xad, 0x20, 0xcf, 0xe1, 0x51, 0x21, 0x4f, 0xb9, 0xc0, 0x75, 0x75, 0x21, 0x4b,
	0x8c, 0x03, 0x67, 0x46, 0x6d, 0xec, 0xd8, 0x86, 0xec, 0x81, 0x2a, 0xa6, 0xf5, 0x95, 0x54, 0x79,
	0x1c, 0xb6, 0x4d, 0x3b, 0x6d, 0x9b, 0x9e, 0x2b, 0x59, 0x57, 0xf1, 0xbc, 0x6d, 0xea, 0x04, 0xf9,
	0x00, 0xe1, 0x19, 0x47, 0x91, 0xeb, 0x78, 0xe1, 0xae, 0xff, 0x72, 0x73, 0xfd, 0x3b, 0xf7, 0x4d,
	0x3f, 0xbb, 0xcc, 0xa3, 0xd2, 0xa8, 0x26, 0xbb, 0x2e, 0x4b, 0xde, 0x41, 0x74, 0x23, 0x4c, 0xb6,
	0x61, 0x7a, 0x89, 0xcd, 0x35, 0x0d, 0xfb, 0x69, 0xff, 0xfb, 0x8b, 0x89, 0xba, 0x63, 0xd0, 0x8a,
	0x03, 0xff, 0xad, 0x47, 0x7f, 0xfb, 0x30, 0xb3, 0xed, 0xc9, 0x0e, 0xcc, 0x2d, 0xb7, 0x35, 0xcf,
	0x5d, 0x61, 0x90, 0x85, 0x56, 0x7e, 0xc9, 0x6f, 0x01, 0xf6, 0x47, 0x01, 0x4f, 0x47, 0x01, 0xcf,
	0x1e, 0x02, 0x1c, 0x8c, 0x01, 0x0e, 0xef, 0x02, 0xbe, 0x1f, 0xe2, 0x6a, 0x00, 0x31, 0xb9, 0x3d,
	0x43, 0xff, 0x99, 0xdb, 0xea, 0xaf, 0x07, 0x91, 0xed, 0xfb, 0x8d, 0x95, 0xec, 0x1c, 0x15, 0x39,
	0x80, 0x65, 0x37, 0xd4, 0x9a, 0x3c, 0xde, 0xfc, 0xdb, 0x8d, 0x7c, 0x72, 0xe3, 0x30, 0xc3, 0xd1,
	0xa7, 0x13, 0xf2, 0x1e, 0x60, 0xf3, 0xce, 0x64, 0x77, 0xe4, 0xf5, 0x93, 0xc1, 0x66, 0xd0, 0x09,
	0x49, 0x01, 0x4e, 0xaa, 0xbc, 0x2b, 0x1e, 0xf8, 0xf7, 0xe4, 0xbf, 0x02, 0xf8, 0x84, 0x02, 0x1f,
	0xc8, 0x1f, 0x9e, 0x9c, 0x4e, 0x0e, 0x5f, 0xc0, 0xce, 0x0f, 0x59, 0xa4, 0x05, 0xea, 0x8b, 0x2b,
	0x26, 0x2e, 0x51, 0xf5, 0x29, 0x87, 0xfd, 0x72, 0x1f, 0x7b, 0xa7, 0xa1, 0xdb, 0xf7, 0xd7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x21, 0xa9, 0x48, 0x7a, 0x01, 0x04, 0x00, 0x00,
}
