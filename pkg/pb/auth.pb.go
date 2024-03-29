// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthRequest_Entity int32

const (
	AuthRequest_SYSTEM   AuthRequest_Entity = 0
	AuthRequest_EMPLOYEE AuthRequest_Entity = 1
)

var AuthRequest_Entity_name = map[int32]string{
	0: "SYSTEM",
	1: "EMPLOYEE",
}

var AuthRequest_Entity_value = map[string]int32{
	"SYSTEM":   0,
	"EMPLOYEE": 1,
}

func (x AuthRequest_Entity) String() string {
	return proto.EnumName(AuthRequest_Entity_name, int32(x))
}

func (AuthRequest_Entity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0, 0}
}

type AuthRequest struct {
	Entity               AuthRequest_Entity `protobuf:"varint,1,opt,name=entity,proto3,enum=pb.AuthRequest_Entity" json:"entity,omitempty"`
	Login                string             `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password             string             `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Key                  string             `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetEntity() AuthRequest_Entity {
	if m != nil {
		return m.Entity
	}
	return AuthRequest_SYSTEM
}

func (m *AuthRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.AuthRequest_Entity", AuthRequest_Entity_name, AuthRequest_Entity_value)
	proto.RegisterType((*AuthRequest)(nil), "pb.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "pb.AuthResponse")
	proto.RegisterType((*ValidateRequest)(nil), "pb.ValidateRequest")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0x5a, 0x0d, 0xed, 0x58, 0x6c, 0x58, 0x45, 0x42, 0xf0, 0x50, 0x16, 0xc1, 0x9e, 0x52,
	0x68, 0xbe, 0xc0, 0x43, 0x6e, 0x16, 0x25, 0x11, 0xa1, 0xc7, 0x4d, 0x3b, 0xa4, 0x4b, 0xc3, 0xee,
	0x9a, 0x9d, 0x58, 0xfa, 0x41, 0xfe, 0xa7, 0x24, 0x6b, 0x55, 0x4a, 0x6f, 0xfb, 0xe6, 0xbd, 0x9d,
	0xf7, 0xde, 0x00, 0x88, 0x86, 0xb6, 0xb1, 0xa9, 0x35, 0x69, 0xd6, 0x37, 0x45, 0x74, 0x5f, 0x6a,
	0x5d, 0x56, 0x38, 0x17, 0x46, 0xce, 0x85, 0x52, 0x9a, 0x04, 0x49, 0xad, 0xac, 0x53, 0xf0, 0x2f,
	0x0f, 0xae, 0x9e, 0x1a, 0xda, 0x66, 0xf8, 0xd1, 0xa0, 0x25, 0x16, 0x83, 0x8f, 0x8a, 0x24, 0x1d,
	0x42, 0x6f, 0xea, 0xcd, 0xae, 0x17, 0x77, 0xb1, 0x29, 0xe2, 0x7f, 0x82, 0x38, 0xed, 0xd8, 0xec,
	0x47, 0xc5, 0x6e, 0xe1, 0xb2, 0xd2, 0xa5, 0x54, 0x61, 0x7f, 0xea, 0xcd, 0x46, 0x99, 0x03, 0x2c,
	0x82, 0xa1, 0x11, 0xd6, 0xee, 0x75, 0xbd, 0x09, 0x07, 0x1d, 0xf1, 0x8b, 0x59, 0x00, 0x83, 0x1d,
	0x1e, 0xc2, 0x8b, 0x6e, 0xdc, 0x3e, 0x39, 0x07, 0xdf, 0x6d, 0x65, 0x00, 0x7e, 0xbe, 0xca, 0xdf,
	0xd2, 0x65, 0xd0, 0x63, 0x63, 0x18, 0xa6, 0xcb, 0xd7, 0xe7, 0x97, 0x55, 0x9a, 0x06, 0x1e, 0x7f,
	0x80, 0xb1, 0x4b, 0x61, 0x8d, 0x56, 0x16, 0x5b, 0x5f, 0xd2, 0x3b, 0x54, 0x5d, 0xcc, 0x51, 0xe6,
	0x00, 0x7f, 0x84, 0xc9, 0xbb, 0xa8, 0xe4, 0x46, 0x10, 0x1e, 0x0b, 0x9d, 0x15, 0x2e, 0xf6, 0xae,
	0x75, 0x8e, 0xf5, 0xa7, 0x5c, 0x23, 0x4b, 0xdc, 0xf6, 0xb6, 0xd3, 0x5a, 0x10, 0xb2, 0xc9, 0x49,
	0xeb, 0x28, 0xf8, 0x1b, 0xb8, 0x00, 0xbc, 0xc7, 0x12, 0x18, 0x1e, 0xcd, 0xd8, 0x4d, 0xcb, 0x9f,
	0x58, 0x9f, 0xfb, 0x54, 0xf8, 0xdd, 0xd9, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x03,
	0xdd, 0xf9, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
	Validate(context.Context, *ValidateRequest) (*AuthResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Authenticate(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedAuthServiceServer) Validate(ctx context.Context, req *ValidateRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
