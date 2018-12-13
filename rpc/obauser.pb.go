// Code generated by protoc-gen-go. DO NOT EDIT.
// source: obauser.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9cbb1645250020, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateUserRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9cbb1645250020, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type User struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9cbb1645250020, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "niixo.obauser.v1.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "niixo.obauser.v1.CreateUserRequest")
	proto.RegisterType((*User)(nil), "niixo.obauser.v1.User")
}

func init() { proto.RegisterFile("obauser.proto", fileDescriptor_da9cbb1645250020) }

var fileDescriptor_da9cbb1645250020 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4f, 0x4a, 0x2c,
	0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0xcb, 0xcc, 0xac, 0xc8,
	0xd7, 0x83, 0x09, 0x96, 0x19, 0x2a, 0x69, 0x72, 0xf1, 0xb9, 0xa7, 0x96, 0x84, 0x16, 0xa7, 0x16,
	0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x73, 0xb1, 0x83, 0x24, 0xe3, 0x33, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x25, 0x03, 0x2e, 0x41,
	0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0x64, 0xd5, 0xd2, 0x5c, 0x9c, 0x60, 0xd5, 0x79, 0x89, 0xb9,
	0xa9, 0x60, 0xf5, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x25, 0x1b, 0x2e, 0x16,
	0x90, 0x5a, 0x74, 0x23, 0x79, 0x61, 0x46, 0xa2, 0xea, 0x66, 0x42, 0xd5, 0x6d, 0x34, 0x97, 0x91,
	0x8b, 0x1d, 0xa4, 0xdd, 0xb1, 0x20, 0x53, 0xc8, 0x95, 0x8b, 0x1d, 0xea, 0x4c, 0x21, 0x05, 0x3d,
	0x74, 0x4f, 0xe8, 0xa1, 0xfa, 0x40, 0x4a, 0x0c, 0x53, 0x05, 0x48, 0x5a, 0x89, 0x41, 0xc8, 0x9b,
	0x8b, 0x0b, 0xe1, 0x05, 0x21, 0x65, 0x4c, 0x75, 0x18, 0x1e, 0xc4, 0x6d, 0x98, 0x93, 0x5c, 0x94,
	0x4c, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x58, 0x95, 0x3e, 0x54,
	0x95, 0x7e, 0x51, 0x41, 0x72, 0x12, 0x1b, 0x38, 0xcc, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x78, 0xa9, 0xb7, 0x47, 0x84, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserApiClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userApiClient struct {
	cc *grpc.ClientConn
}

func NewUserApiClient(cc *grpc.ClientConn) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/niixo.obauser.v1.UserApi/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/niixo.obauser.v1.UserApi/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
type UserApiServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
}

func RegisterUserApiServer(s *grpc.Server, srv UserApiServer) {
	s.RegisterService(&_UserApi_serviceDesc, srv)
}

func _UserApi_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/niixo.obauser.v1.UserApi/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/niixo.obauser.v1.UserApi/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "niixo.obauser.v1.UserApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserApi_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserApi_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "obauser.proto",
}
