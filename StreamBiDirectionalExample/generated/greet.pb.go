// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{2}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
}

func init() {
	proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579)
}

var fileDescriptor_32c0044392f32579 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xdc, 0xb8, 0x38, 0xdc,
	0x41, 0x8c, 0xcc, 0xbc, 0x74, 0x21, 0x59, 0x2e, 0xae, 0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x88, 0x1f, 0x50, 0x40,
	0x48, 0x9a, 0x8b, 0x33, 0x27, 0x11, 0x26, 0xcb, 0x04, 0x96, 0xe5, 0x00, 0x09, 0x80, 0x24, 0x95,
	0x9c, 0xb9, 0x44, 0xc0, 0xe6, 0xb8, 0x96, 0xa5, 0x16, 0x55, 0xe6, 0xe7, 0xa5, 0x06, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x69, 0x73, 0x71, 0xa4, 0x43, 0xcd, 0x07, 0x9b, 0xc8, 0x6d, 0xc4,
	0xaf, 0x07, 0x71, 0x06, 0xcc, 0xda, 0x20, 0xb8, 0x02, 0x25, 0x7d, 0x2e, 0x51, 0x34, 0x43, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0,
	0xae, 0x82, 0xf2, 0x8c, 0x12, 0xb8, 0x78, 0xc0, 0x1a, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x02, 0xb8, 0x78, 0x51, 0x0c, 0x10, 0x92, 0x46, 0xb6, 0x0c, 0xcd, 0x6d, 0x52, 0x32, 0xd8,
	0x25, 0x21, 0x76, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0x71, 0x46, 0xb1, 0x83, 0x15, 0x15,
	0x24, 0x25, 0xb1, 0x81, 0x03, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xd7, 0x48, 0xdb,
	0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Bi directional streaming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Bi directional streaming
	GreetEveryone(GreetService_GreetEveryoneServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
