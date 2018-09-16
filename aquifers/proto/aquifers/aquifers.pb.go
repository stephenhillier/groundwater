// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/aquifers/aquifers.proto

package go_srv_aquifers

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

type Aquifer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Volume               float64  `protobuf:"fixed64,3,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Aquifer) Reset()         { *m = Aquifer{} }
func (m *Aquifer) String() string { return proto.CompactTextString(m) }
func (*Aquifer) ProtoMessage()    {}
func (*Aquifer) Descriptor() ([]byte, []int) {
	return fileDescriptor_62dde725a48b96e2, []int{0}
}
func (m *Aquifer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aquifer.Unmarshal(m, b)
}
func (m *Aquifer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aquifer.Marshal(b, m, deterministic)
}
func (dst *Aquifer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aquifer.Merge(dst, src)
}
func (m *Aquifer) XXX_Size() int {
	return xxx_messageInfo_Aquifer.Size(m)
}
func (m *Aquifer) XXX_DiscardUnknown() {
	xxx_messageInfo_Aquifer.DiscardUnknown(m)
}

var xxx_messageInfo_Aquifer proto.InternalMessageInfo

func (m *Aquifer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Aquifer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Aquifer) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type SingleAquiferRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleAquiferRequest) Reset()         { *m = SingleAquiferRequest{} }
func (m *SingleAquiferRequest) String() string { return proto.CompactTextString(m) }
func (*SingleAquiferRequest) ProtoMessage()    {}
func (*SingleAquiferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62dde725a48b96e2, []int{1}
}
func (m *SingleAquiferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleAquiferRequest.Unmarshal(m, b)
}
func (m *SingleAquiferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleAquiferRequest.Marshal(b, m, deterministic)
}
func (dst *SingleAquiferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleAquiferRequest.Merge(dst, src)
}
func (m *SingleAquiferRequest) XXX_Size() int {
	return xxx_messageInfo_SingleAquiferRequest.Size(m)
}
func (m *SingleAquiferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleAquiferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleAquiferRequest proto.InternalMessageInfo

func (m *SingleAquiferRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Aquifer)(nil), "go.srv.aquifers.Aquifer")
	proto.RegisterType((*SingleAquiferRequest)(nil), "go.srv.aquifers.SingleAquiferRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AquiferServiceClient is the client API for AquiferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AquiferServiceClient interface {
	GetAquifer(ctx context.Context, in *SingleAquiferRequest, opts ...grpc.CallOption) (*Aquifer, error)
}

type aquiferServiceClient struct {
	cc *grpc.ClientConn
}

func NewAquiferServiceClient(cc *grpc.ClientConn) AquiferServiceClient {
	return &aquiferServiceClient{cc}
}

func (c *aquiferServiceClient) GetAquifer(ctx context.Context, in *SingleAquiferRequest, opts ...grpc.CallOption) (*Aquifer, error) {
	out := new(Aquifer)
	err := c.cc.Invoke(ctx, "/go.srv.aquifers.AquiferService/GetAquifer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AquiferServiceServer is the server API for AquiferService service.
type AquiferServiceServer interface {
	GetAquifer(context.Context, *SingleAquiferRequest) (*Aquifer, error)
}

func RegisterAquiferServiceServer(s *grpc.Server, srv AquiferServiceServer) {
	s.RegisterService(&_AquiferService_serviceDesc, srv)
}

func _AquiferService_GetAquifer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleAquiferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AquiferServiceServer).GetAquifer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.srv.aquifers.AquiferService/GetAquifer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AquiferServiceServer).GetAquifer(ctx, req.(*SingleAquiferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AquiferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.srv.aquifers.AquiferService",
	HandlerType: (*AquiferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAquifer",
			Handler:    _AquiferService_GetAquifer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/aquifers/aquifers.proto",
}

func init() { proto.RegisterFile("proto/aquifers/aquifers.proto", fileDescriptor_62dde725a48b96e2) }

var fileDescriptor_62dde725a48b96e2 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x2c, 0xcd, 0x4c, 0x4b, 0x2d, 0x2a, 0x86, 0x33, 0xf4, 0xc0, 0xe2, 0x42,
	0xfc, 0xe9, 0xf9, 0x7a, 0xc5, 0x45, 0x65, 0x7a, 0x30, 0x61, 0x25, 0x57, 0x2e, 0x76, 0x47, 0x08,
	0x5b, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x29, 0x33,
	0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xcc, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0xcb, 0xcf, 0x29, 0xcd, 0x4d, 0x95, 0x60, 0x56, 0x60, 0xd4,
	0x60, 0x0c, 0x82, 0xf2, 0x94, 0xd4, 0xb8, 0x44, 0x82, 0x33, 0xf3, 0xd2, 0x73, 0x52, 0xa1, 0x86,
	0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0xa0, 0x9b, 0x69, 0x94, 0xc8, 0xc5, 0x07, 0x55, 0x11,
	0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xcf, 0xc5, 0xe5, 0x9e, 0x5a, 0x02, 0x73, 0x83,
	0xaa, 0x1e, 0x9a, 0x03, 0xf5, 0xb0, 0x19, 0x2b, 0x25, 0x81, 0xa1, 0x0c, 0xaa, 0x40, 0x89, 0x21,
	0x89, 0x0d, 0xec, 0x53, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xc4, 0x8c, 0xc2, 0x0a,
	0x01, 0x00, 0x00,
}