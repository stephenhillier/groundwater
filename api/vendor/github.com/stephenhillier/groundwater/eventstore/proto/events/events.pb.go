// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/events/events.proto

package go_srv_events

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

type Event struct {
	AggregateId          string   `protobuf:"bytes,1,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType        string   `protobuf:"bytes,2,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventId              string   `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType            string   `protobuf:"bytes,4,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EventData            string   `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_06e82083c0e33784, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_06e82083c0e33784, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Event)(nil), "go.srv.events.Event")
	proto.RegisterType((*Response)(nil), "go.srv.events.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.srv.events.EventService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	CreateEvent(context.Context, *Event) (*Response, error)
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.srv.events.EventService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.srv.events.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventService_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/events/events.proto",
}

func init() { proto.RegisterFile("proto/events/events.proto", fileDescriptor_06e82083c0e33784) }

var fileDescriptor_06e82083c0e33784 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0x87, 0x30,
	0x10, 0xc5, 0x45, 0x45, 0xe0, 0x00, 0x87, 0xc6, 0xc4, 0x62, 0x62, 0xa2, 0x44, 0x13, 0xa7, 0x9a,
	0xe8, 0xee, 0xa2, 0x0e, 0x2c, 0x0e, 0xe8, 0x6e, 0x2a, 0xbd, 0x34, 0x2c, 0xb4, 0x69, 0x2b, 0x09,
	0x5f, 0xc9, 0x4f, 0x69, 0x38, 0x04, 0xf5, 0x3f, 0xb5, 0xf7, 0x7b, 0xef, 0xae, 0xbd, 0x07, 0x95,
	0x75, 0x26, 0x98, 0x5b, 0x1c, 0x71, 0x08, 0xfe, 0xe7, 0x10, 0xc4, 0x58, 0xa9, 0x8d, 0xf0, 0x6e,
	0x14, 0x0b, 0xac, 0xbf, 0x22, 0x88, 0x9f, 0xe7, 0x2b, 0xbb, 0x84, 0x42, 0x6a, 0xed, 0x50, 0xcb,
	0x80, 0xef, 0xbd, 0xe2, 0xd1, 0x45, 0x74, 0x93, 0xb5, 0xf9, 0xc6, 0x1a, 0xc5, 0xae, 0xe1, 0xf8,
	0xd7, 0x12, 0x26, 0x8b, 0x7c, 0x9f, 0x4c, 0xe5, 0x46, 0xdf, 0x26, 0x8b, 0xac, 0x82, 0x94, 0xa6,
	0xcf, 0x53, 0x0e, 0xc8, 0x90, 0x50, 0xdd, 0x28, 0x76, 0x0e, 0xb0, 0x48, 0xd4, 0x7d, 0x48, 0x62,
	0x46, 0x84, 0x3a, 0x37, 0x59, 0xc9, 0x20, 0x79, 0xfc, 0x47, 0x7e, 0x92, 0x41, 0xd6, 0x57, 0x90,
	0xb6, 0xe8, 0xad, 0x19, 0x3c, 0x32, 0x0e, 0x89, 0xff, 0xec, 0x3a, 0xf4, 0x9e, 0x7e, 0x9a, 0xb6,
	0x6b, 0x79, 0xf7, 0x02, 0x05, 0x6d, 0xf4, 0x8a, 0x6e, 0xec, 0x3b, 0x64, 0x0f, 0x90, 0x3f, 0x3a,
	0x94, 0x01, 0x97, 0x3d, 0x4f, 0xc4, 0xbf, 0x04, 0x04, 0xd1, 0xb3, 0xd3, 0x1d, 0xba, 0xbe, 0x53,
	0xef, 0x7d, 0x1c, 0x51, 0x70, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x60, 0x6a, 0xe9,
	0x55, 0x01, 0x00, 0x00,
}