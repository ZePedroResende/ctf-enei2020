// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flag.proto

package flag

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FlagRequest struct {
	Pass                 string   `protobuf:"bytes,1,opt,name=pass,proto3" json:"pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagRequest) Reset()         { *m = FlagRequest{} }
func (m *FlagRequest) String() string { return proto.CompactTextString(m) }
func (*FlagRequest) ProtoMessage()    {}
func (*FlagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdf51d06af45bb, []int{0}
}

func (m *FlagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagRequest.Unmarshal(m, b)
}
func (m *FlagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagRequest.Marshal(b, m, deterministic)
}
func (m *FlagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagRequest.Merge(m, src)
}
func (m *FlagRequest) XXX_Size() int {
	return xxx_messageInfo_FlagRequest.Size(m)
}
func (m *FlagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlagRequest proto.InternalMessageInfo

func (m *FlagRequest) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type FlagReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagReply) Reset()         { *m = FlagReply{} }
func (m *FlagReply) String() string { return proto.CompactTextString(m) }
func (*FlagReply) ProtoMessage()    {}
func (*FlagReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdf51d06af45bb, []int{1}
}

func (m *FlagReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagReply.Unmarshal(m, b)
}
func (m *FlagReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagReply.Marshal(b, m, deterministic)
}
func (m *FlagReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagReply.Merge(m, src)
}
func (m *FlagReply) XXX_Size() int {
	return xxx_messageInfo_FlagReply.Size(m)
}
func (m *FlagReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagReply.DiscardUnknown(m)
}

var xxx_messageInfo_FlagReply proto.InternalMessageInfo

func (m *FlagReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FlagRequest)(nil), "flag.FlagRequest")
	proto.RegisterType((*FlagReply)(nil), "flag.FlagReply")
}

func init() { proto.RegisterFile("flag.proto", fileDescriptor_01fdf51d06af45bb) }

var fileDescriptor_01fdf51d06af45bb = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0x49, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0xb8, 0xdd, 0x72,
	0x12, 0xd3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x0a, 0x12, 0x8b,
	0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x55, 0x2e, 0x4e, 0x88, 0x92,
	0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xa8, 0x1a,
	0x18, 0xd7, 0xc8, 0x94, 0x8b, 0x05, 0xa4, 0x4c, 0x48, 0x97, 0x8b, 0x3d, 0x3d, 0xb5, 0x04, 0xcc,
	0x14, 0xd4, 0x03, 0xdb, 0x87, 0x64, 0x81, 0x14, 0x3f, 0xb2, 0x50, 0x41, 0x4e, 0x65, 0x12, 0x1b,
	0xd8, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x9f, 0x48, 0x46, 0x9b, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FlagClient is the client API for Flag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlagClient interface {
	GetFlag(ctx context.Context, in *FlagRequest, opts ...grpc.CallOption) (*FlagReply, error)
}

type flagClient struct {
	cc grpc.ClientConnInterface
}

func NewFlagClient(cc grpc.ClientConnInterface) FlagClient {
	return &flagClient{cc}
}

func (c *flagClient) GetFlag(ctx context.Context, in *FlagRequest, opts ...grpc.CallOption) (*FlagReply, error) {
	out := new(FlagReply)
	err := c.cc.Invoke(ctx, "/flag.Flag/getFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlagServer is the server API for Flag service.
type FlagServer interface {
	GetFlag(context.Context, *FlagRequest) (*FlagReply, error)
}

// UnimplementedFlagServer can be embedded to have forward compatible implementations.
type UnimplementedFlagServer struct {
}

func (*UnimplementedFlagServer) GetFlag(ctx context.Context, req *FlagRequest) (*FlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}

func RegisterFlagServer(s *grpc.Server, srv FlagServer) {
	s.RegisterService(&_Flag_serviceDesc, srv)
}

func _Flag_GetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlagServer).GetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flag.Flag/GetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlagServer).GetFlag(ctx, req.(*FlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Flag_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flag.Flag",
	HandlerType: (*FlagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getFlag",
			Handler:    _Flag_GetFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flag.proto",
}
