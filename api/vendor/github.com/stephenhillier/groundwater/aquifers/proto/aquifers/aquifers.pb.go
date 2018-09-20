// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aquifers/proto/aquifers/aquifers.proto

package go_srv_aquifers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wells "github.com/stephenhillier/groundwater/wells/proto/wells"

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
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Volume               float64       `protobuf:"fixed64,3,opt,name=volume,proto3" json:"volume,omitempty"`
	Wells                []*wells.Well `protobuf:"bytes,4,rep,name=wells,proto3" json:"wells,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Aquifer) Reset()         { *m = Aquifer{} }
func (m *Aquifer) String() string { return proto.CompactTextString(m) }
func (*Aquifer) ProtoMessage()    {}
func (*Aquifer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ad465a33996f31d, []int{0}
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

func (m *Aquifer) GetWells() []*wells.Well {
	if m != nil {
		return m.Wells
	}
	return nil
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
	return fileDescriptor_6ad465a33996f31d, []int{1}
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
	Metadata: "aquifers/proto/aquifers/aquifers.proto",
}

func init() {
	proto.RegisterFile("aquifers/proto/aquifers/aquifers.proto", fileDescriptor_6ad465a33996f31d)
}

var fileDescriptor_6ad465a33996f31d = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x6a, 0xf3, 0x30,
	0x10, 0xc4, 0x3f, 0x39, 0x7f, 0x3e, 0xba, 0x85, 0x14, 0x44, 0x29, 0x22, 0x27, 0x63, 0x68, 0xd0,
	0x49, 0x82, 0xf4, 0x09, 0x4a, 0x0f, 0x3d, 0x16, 0x94, 0x43, 0xcf, 0x4e, 0xbc, 0xb5, 0x05, 0xb2,
	0x95, 0xc8, 0x92, 0xf3, 0xfa, 0x05, 0x49, 0xee, 0x21, 0xed, 0x65, 0x99, 0x5d, 0x86, 0xdf, 0x2c,
	0x03, 0xbb, 0xfa, 0x12, 0xf4, 0x17, 0xba, 0x51, 0x9e, 0x9d, 0xf5, 0x56, 0xfe, 0xac, 0xb3, 0x10,
	0xf1, 0x4e, 0x1f, 0x5a, 0x2b, 0x46, 0x37, 0x89, 0xf9, 0xbc, 0x7d, 0x6b, 0xb5, 0xef, 0xc2, 0x51,
	0x9c, 0x6c, 0x2f, 0x47, 0x8f, 0xe7, 0x0e, 0x87, 0x4e, 0x1b, 0xa3, 0xd1, 0xc9, 0xd6, 0xd9, 0x30,
	0x34, 0xd7, 0xda, 0xa3, 0x93, 0x57, 0x34, 0x66, 0x66, 0x27, 0x1d, 0x67, 0xa2, 0x56, 0x16, 0xfe,
	0xbf, 0x26, 0x20, 0xdd, 0x40, 0xa1, 0x1b, 0x46, 0x4a, 0xc2, 0x57, 0xaa, 0xd0, 0x0d, 0xa5, 0xb0,
	0x1c, 0xea, 0x1e, 0x59, 0x51, 0x12, 0x7e, 0xa7, 0xa2, 0xa6, 0x4f, 0xb0, 0x9e, 0xac, 0x09, 0x3d,
	0xb2, 0x45, 0x49, 0x38, 0x51, 0x79, 0xa3, 0x1c, 0x56, 0x91, 0xca, 0x96, 0xe5, 0x82, 0xdf, 0xef,
	0xa9, 0xc8, 0xcf, 0xa6, 0xa8, 0x4f, 0x34, 0x46, 0x25, 0x43, 0xb5, 0x83, 0xc7, 0x83, 0x1e, 0x5a,
	0x83, 0x39, 0x56, 0xe1, 0x25, 0xe0, 0xe8, 0x6f, 0xd3, 0xf7, 0x35, 0x6c, 0xb2, 0xe3, 0x80, 0x6e,
	0xd2, 0x27, 0xa4, 0x1f, 0x00, 0xef, 0xe8, 0xe7, 0x6f, 0x9f, 0xc5, 0x4d, 0x1f, 0xe2, 0x2f, 0xec,
	0x96, 0xfd, 0xb2, 0x65, 0x43, 0xf5, 0xef, 0xb8, 0x8e, 0x15, 0xbc, 0x7c, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0x66, 0x1b, 0x8b, 0x82, 0x01, 0x00, 0x00,
}