// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post/post.proto

package post

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

type OneReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneReq) Reset()         { *m = OneReq{} }
func (m *OneReq) String() string { return proto.CompactTextString(m) }
func (*OneReq) ProtoMessage()    {}
func (*OneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e4a9952ba66d64, []int{0}
}

func (m *OneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneReq.Unmarshal(m, b)
}
func (m *OneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneReq.Marshal(b, m, deterministic)
}
func (m *OneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneReq.Merge(m, src)
}
func (m *OneReq) XXX_Size() int {
	return xxx_messageInfo_OneReq.Size(m)
}
func (m *OneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OneReq.DiscardUnknown(m)
}

var xxx_messageInfo_OneReq proto.InternalMessageInfo

func (m *OneReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OneRes struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneRes) Reset()         { *m = OneRes{} }
func (m *OneRes) String() string { return proto.CompactTextString(m) }
func (*OneRes) ProtoMessage()    {}
func (*OneRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e4a9952ba66d64, []int{1}
}

func (m *OneRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneRes.Unmarshal(m, b)
}
func (m *OneRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneRes.Marshal(b, m, deterministic)
}
func (m *OneRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneRes.Merge(m, src)
}
func (m *OneRes) XXX_Size() int {
	return xxx_messageInfo_OneRes.Size(m)
}
func (m *OneRes) XXX_DiscardUnknown() {
	xxx_messageInfo_OneRes.DiscardUnknown(m)
}

var xxx_messageInfo_OneRes proto.InternalMessageInfo

func (m *OneRes) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *OneRes) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*OneReq)(nil), "OneReq")
	proto.RegisterType((*OneRes)(nil), "OneRes")
}

func init() {
	proto.RegisterFile("post/post.proto", fileDescriptor_e9e4a9952ba66d64)
}

var fileDescriptor_e9e4a9952ba66d64 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc8, 0x2f, 0x2e,
	0xd1, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x12, 0x5c, 0x6c, 0xfe, 0x79, 0xa9,
	0x41, 0xa9, 0x85, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x4c, 0x99, 0x29, 0x4a, 0x46, 0x50, 0x99, 0x62, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x9c,
	0x54, 0xa8, 0x24, 0x84, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x94, 0x9f, 0x52, 0x29, 0xc1, 0x04, 0x16,
	0x04, 0xb3, 0x8d, 0x14, 0xb9, 0x58, 0x02, 0xf2, 0x8b, 0x4b, 0x84, 0x24, 0xb9, 0x98, 0xdd, 0x53,
	0x4b, 0x84, 0xd8, 0xf5, 0x20, 0x66, 0x4b, 0x41, 0x19, 0xc5, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x7b,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xed, 0x7c, 0xa3, 0x8a, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostClient interface {
	Get(ctx context.Context, in *OneReq, opts ...grpc.CallOption) (*OneRes, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) Get(ctx context.Context, in *OneReq, opts ...grpc.CallOption) (*OneRes, error) {
	out := new(OneRes)
	err := c.cc.Invoke(ctx, "/Post/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
type PostServer interface {
	Get(context.Context, *OneReq) (*OneRes, error)
}

// UnimplementedPostServer can be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (*UnimplementedPostServer) Get(ctx context.Context, req *OneReq) (*OneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterPostServer(s *grpc.Server, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

func _Post_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Post/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Get(ctx, req.(*OneReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Post_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/post.proto",
}