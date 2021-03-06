// Code generated by protoc-gen-go.
// source: webpush.proto
// DO NOT EDIT!

/*
Package webpush is a generated protocol buffer package.

It is generated from these files:
	webpush.proto

It has these top-level messages:
	SubscribeRequest
	SubscribeResponse
	PushRequest
	PushResponse
	MonitorRequest
	Message
	AckRequest
	AckResponse
	ReceiptRequest
	Receipt
*/
package webpush

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

//
type SubscribeRequest struct {
	// A UA should group subscriptions in a set. First request from a
	// UA will not include a set - it is typically a subscription associated with
	// the UA itself.
	PushSet string `protobuf:"bytes,1,opt,name=push_set" json:"push_set,omitempty"`
	// Included as Crypto-Key: p256ecdsa parameter.
	// Corresponds to the applicationServerKey parameter in the PushSubscriptionOptions in
	// the W3C API
	SenderVapid string `protobuf:"bytes,2,opt,name=sender_vapid" json:"sender_vapid,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Subscribe response includes the elements in the spec.
type SubscribeResponse struct {
	// Returned as Link: rel="urn:ietf:params:push"
	// Spec examples use a full path ( /push/xxxx1 )
	// TODO: clarify if it can be a full URL
	Push string `protobuf:"bytes,1,opt,name=push" json:"push,omitempty"`
	// Optional response: it
	// returned as Link: rel=urn:ietf:params:push:set
	// Spec examples use a full path ( /subscription-set/xxxx2 ).
	// TODO: clarify it can be a full URL, like subscription
	PushSet string `protobuf:"bytes,2,opt,name=push_set" json:"push_set,omitempty"`
	// Push subscription resource. This is the full URL where the UA will use to
	// receive the messages, using the PUSH promise http2 frame.
	//
	//
	// Returned as Location header in the spec
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
}

func (m *SubscribeResponse) Reset()                    { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()               {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PushRequest struct {
	// The value returned in the SubscribeResponse push, without the hostname.
	Push    string `protobuf:"bytes,1,opt,name=push" json:"push,omitempty"`
	Ttl     int32  `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Urgency string `protobuf:"bytes,4,opt,name=urgency" json:"urgency,omitempty"`
	// Prefer header indicating delivery receipt request.
	RespondAsync    bool   `protobuf:"varint,5,opt,name=respond_async" json:"respond_async,omitempty"`
	Topic           string `protobuf:"bytes,6,opt,name=topic" json:"topic,omitempty"`
	ContentEncoding string `protobuf:"bytes,7,opt,name=content_encoding" json:"content_encoding,omitempty"`
	Salt            string `protobuf:"bytes,8,opt,name=salt" json:"salt,omitempty"`
	Dh              string `protobuf:"bytes,9,opt,name=dh" json:"dh,omitempty"`
}

func (m *PushRequest) Reset()                    { *m = PushRequest{} }
func (m *PushRequest) String() string            { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()               {}
func (*PushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PushResponse struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id" json:"message_id,omitempty"`
	// If request includes the respond_async parameter.
	//
	PushReceipt string `protobuf:"bytes,2,opt,name=push_receipt" json:"push_receipt,omitempty"`
}

func (m *PushResponse) Reset()                    { *m = PushResponse{} }
func (m *PushResponse) String() string            { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()               {}
func (*PushResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type MonitorRequest struct {
	// This is the push or push_set in the subscribe response.
	PushSet string `protobuf:"bytes,1,opt,name=push_set" json:"push_set,omitempty"`
	// JWT token, signed with key
	Authorization string `protobuf:"bytes,2,opt,name=authorization" json:"authorization,omitempty"`
	// Public key used for signing, identifies sender/receiver
	Key string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
}

func (m *MonitorRequest) Reset()                    { *m = MonitorRequest{} }
func (m *MonitorRequest) String() string            { return proto.CompactTextString(m) }
func (*MonitorRequest) ProtoMessage()               {}
func (*MonitorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Message is returned as PUSH PROMISE frames in the spec. The gRPC interface defines
// it as a stream, returned in normal DATA frames.
type Message struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id" json:"message_id,omitempty"`
	// Maps to the SubscribeResponse push parameter, returned as Link rel="urn:ietf:params:push"
	// in the push promise.
	Push            string `protobuf:"bytes,2,opt,name=push" json:"push,omitempty"`
	Data            []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	SenderVapid     string `protobuf:"bytes,4,opt,name=sender_vapid" json:"sender_vapid,omitempty"`
	ContentEncoding string `protobuf:"bytes,7,opt,name=content_encoding" json:"content_encoding,omitempty"`
	Salt            string `protobuf:"bytes,8,opt,name=salt" json:"salt,omitempty"`
	Dh              string `protobuf:"bytes,9,opt,name=dh" json:"dh,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AckRequest struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id" json:"message_id,omitempty"`
}

func (m *AckRequest) Reset()                    { *m = AckRequest{} }
func (m *AckRequest) String() string            { return proto.CompactTextString(m) }
func (*AckRequest) ProtoMessage()               {}
func (*AckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type AckResponse struct {
}

func (m *AckResponse) Reset()                    { *m = AckResponse{} }
func (m *AckResponse) String() string            { return proto.CompactTextString(m) }
func (*AckResponse) ProtoMessage()               {}
func (*AckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ReceiptRequest struct {
	ReceiptSubscription string `protobuf:"bytes,1,opt,name=receipt_subscription" json:"receipt_subscription,omitempty"`
}

func (m *ReceiptRequest) Reset()                    { *m = ReceiptRequest{} }
func (m *ReceiptRequest) String() string            { return proto.CompactTextString(m) }
func (*ReceiptRequest) ProtoMessage()               {}
func (*ReceiptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Receipt struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id" json:"message_id,omitempty"`
}

func (m *Receipt) Reset()                    { *m = Receipt{} }
func (m *Receipt) String() string            { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()               {}
func (*Receipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "webpush.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "webpush.SubscribeResponse")
	proto.RegisterType((*PushRequest)(nil), "webpush.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "webpush.PushResponse")
	proto.RegisterType((*MonitorRequest)(nil), "webpush.MonitorRequest")
	proto.RegisterType((*Message)(nil), "webpush.Message")
	proto.RegisterType((*AckRequest)(nil), "webpush.AckRequest")
	proto.RegisterType((*AckResponse)(nil), "webpush.AckResponse")
	proto.RegisterType((*ReceiptRequest)(nil), "webpush.ReceiptRequest")
	proto.RegisterType((*Receipt)(nil), "webpush.Receipt")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Webpush service

type WebpushClient interface {
	// Subscribe maps the the webpush subscribe request
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// Monitor allows a UA to receive push messages from the push service
	// Replaced push promises with a stream of Message objects.
	Monitor(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (Webpush_MonitorClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error)
	// Push allows an application server to send messages to UA, using the push service.
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
	// Monitor allows an AS to receive push messages receipts from the push service
	// Replaced push promises with a stream of Message objects.
	Receipts(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (Webpush_ReceiptsClient, error)
}

type webpushClient struct {
	cc *grpc.ClientConn
}

func NewWebpushClient(cc *grpc.ClientConn) WebpushClient {
	return &webpushClient{cc}
}

func (c *webpushClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := grpc.Invoke(ctx, "/webpush.Webpush/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webpushClient) Monitor(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (Webpush_MonitorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Webpush_serviceDesc.Streams[0], c.cc, "/webpush.Webpush/Monitor", opts...)
	if err != nil {
		return nil, err
	}
	x := &webpushMonitorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Webpush_MonitorClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type webpushMonitorClient struct {
	grpc.ClientStream
}

func (x *webpushMonitorClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *webpushClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error) {
	out := new(AckResponse)
	err := grpc.Invoke(ctx, "/webpush.Webpush/Ack", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webpushClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := grpc.Invoke(ctx, "/webpush.Webpush/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webpushClient) Receipts(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (Webpush_ReceiptsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Webpush_serviceDesc.Streams[1], c.cc, "/webpush.Webpush/Receipts", opts...)
	if err != nil {
		return nil, err
	}
	x := &webpushReceiptsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Webpush_ReceiptsClient interface {
	Recv() (*Receipt, error)
	grpc.ClientStream
}

type webpushReceiptsClient struct {
	grpc.ClientStream
}

func (x *webpushReceiptsClient) Recv() (*Receipt, error) {
	m := new(Receipt)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Webpush service

type WebpushServer interface {
	// Subscribe maps the the webpush subscribe request
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// Monitor allows a UA to receive push messages from the push service
	// Replaced push promises with a stream of Message objects.
	Monitor(*MonitorRequest, Webpush_MonitorServer) error
	Ack(context.Context, *AckRequest) (*AckResponse, error)
	// Push allows an application server to send messages to UA, using the push service.
	Push(context.Context, *PushRequest) (*PushResponse, error)
	// Monitor allows an AS to receive push messages receipts from the push service
	// Replaced push promises with a stream of Message objects.
	Receipts(*ReceiptRequest, Webpush_ReceiptsServer) error
}

func RegisterWebpushServer(s *grpc.Server, srv WebpushServer) {
	s.RegisterService(&_Webpush_serviceDesc, srv)
}

func _Webpush_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebpushServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpush.Webpush/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebpushServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webpush_Monitor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WebpushServer).Monitor(m, &webpushMonitorServer{stream})
}

type Webpush_MonitorServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type webpushMonitorServer struct {
	grpc.ServerStream
}

func (x *webpushMonitorServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Webpush_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebpushServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpush.Webpush/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebpushServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webpush_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebpushServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webpush.Webpush/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebpushServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webpush_Receipts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiptRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WebpushServer).Receipts(m, &webpushReceiptsServer{stream})
}

type Webpush_ReceiptsServer interface {
	Send(*Receipt) error
	grpc.ServerStream
}

type webpushReceiptsServer struct {
	grpc.ServerStream
}

func (x *webpushReceiptsServer) Send(m *Receipt) error {
	return x.ServerStream.SendMsg(m)
}

var _Webpush_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webpush.Webpush",
	HandlerType: (*WebpushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Webpush_Subscribe_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _Webpush_Ack_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Webpush_Push_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Monitor",
			Handler:       _Webpush_Monitor_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receipts",
			Handler:       _Webpush_Receipts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("webpush.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xdd, 0xa4, 0xed, 0xa6, 0x9d, 0xb6, 0x4b, 0x31, 0x5d, 0x61, 0x22, 0x90, 0x56, 0x39, 0x71,
	0xaa, 0xd0, 0x22, 0xa4, 0x85, 0x1b, 0x12, 0x70, 0x5b, 0x09, 0x2d, 0x07, 0x8e, 0x51, 0xea, 0x58,
	0xad, 0xb5, 0xc5, 0x0e, 0xb6, 0x03, 0x2a, 0x57, 0x7e, 0x09, 0xbf, 0x80, 0xbf, 0x88, 0xe3, 0x8f,
	0x34, 0xad, 0xc2, 0x65, 0x8f, 0x99, 0x37, 0x6f, 0xe6, 0xcd, 0xf3, 0x0b, 0xcc, 0x7f, 0xd2, 0x75,
	0x55, 0xab, 0xed, 0xaa, 0x92, 0x42, 0x0b, 0x94, 0xf8, 0xcf, 0xec, 0x1d, 0x2c, 0xbe, 0xd4, 0x6b,
	0x45, 0x24, 0x5b, 0xd3, 0x3b, 0xfa, 0xbd, 0xa6, 0x4a, 0xa3, 0x05, 0x8c, 0x1b, 0x2c, 0x57, 0x54,
	0xe3, 0xe8, 0x2a, 0x7a, 0x39, 0x41, 0x4b, 0x98, 0x29, 0xca, 0x4b, 0x2a, 0xf3, 0x1f, 0x45, 0xc5,
	0x4a, 0x1c, 0x37, 0xd5, 0xec, 0x23, 0x3c, 0xee, 0x70, 0x55, 0x25, 0xb8, 0xa2, 0x68, 0x06, 0xc3,
	0x86, 0xec, 0x89, 0xdd, 0x51, 0x71, 0xa8, 0xec, 0x04, 0x29, 0x34, 0x13, 0x1c, 0x0f, 0xec, 0x98,
	0x3f, 0x11, 0x4c, 0x3f, 0x9b, 0xa6, 0xb0, 0xfe, 0x78, 0xc2, 0x14, 0x06, 0x5a, 0xef, 0x2c, 0x79,
	0xd4, 0x40, 0x65, 0xa1, 0x0b, 0x4b, 0x9c, 0xa1, 0x47, 0x90, 0xd4, 0x72, 0x43, 0x39, 0xd9, 0xe3,
	0xa1, 0xed, 0xbd, 0x84, 0xb9, 0xb4, 0x3a, 0xca, 0xbc, 0x50, 0x7b, 0x4e, 0xf0, 0xc8, 0x94, 0xc7,
	0x68, 0x0e, 0x23, 0x2d, 0x2a, 0x46, 0xf0, 0xb9, 0xed, 0xc2, 0xb0, 0x20, 0x82, 0x6b, 0xca, 0x75,
	0x6e, 0xb8, 0xa2, 0x64, 0x7c, 0x83, 0x13, 0x8b, 0x98, 0xf1, 0xaa, 0xd8, 0x69, 0x3c, 0xb6, 0x5f,
	0x00, 0x71, 0xb9, 0xc5, 0x13, 0xab, 0xf1, 0x06, 0x66, 0x4e, 0xa2, 0xbf, 0xd2, 0x80, 0xdf, 0xa8,
	0x52, 0xc5, 0x86, 0xe6, 0xc6, 0x8e, 0xd6, 0x24, 0x7b, 0xab, 0xa4, 0x84, 0xb2, 0xca, 0xdf, 0x9b,
	0x7d, 0x82, 0x8b, 0x5b, 0xc1, 0x99, 0x16, 0xf2, 0xff, 0xf6, 0x1a, 0xdd, 0x45, 0xad, 0xb7, 0x42,
	0xb2, 0x5f, 0xce, 0x98, 0x38, 0x9c, 0x7e, 0x4f, 0xf7, 0xde, 0xa5, 0xdf, 0x11, 0x24, 0xb7, 0x6e,
	0x65, 0xef, 0xf6, 0xe0, 0x5a, 0x1c, 0xbe, 0x3a, 0x46, 0x9d, 0x3e, 0xdf, 0xf0, 0xc1, 0x3e, 0x5c,
	0x01, 0xbc, 0x27, 0xf7, 0xe1, 0x92, 0x1e, 0x1d, 0xd9, 0x1c, 0xa6, 0xb6, 0xc3, 0x19, 0x95, 0xad,
	0xe0, 0xe2, 0xce, 0xf9, 0x11, 0x48, 0xcf, 0x61, 0xe9, 0x1d, 0xca, 0x95, 0x4b, 0x4f, 0x65, 0x6f,
	0x76, 0xf4, 0x17, 0x90, 0xf8, 0xfe, 0xbe, 0xe9, 0xd7, 0x7f, 0x63, 0x48, 0xbe, 0xba, 0xe8, 0xa2,
	0x0f, 0x30, 0x69, 0xe3, 0x87, 0x9e, 0xad, 0x42, 0xc0, 0x4f, 0xe3, 0x9c, 0xa6, 0x7d, 0x90, 0x97,
	0x77, 0x86, 0x6e, 0x8c, 0xad, 0xee, 0x7d, 0xd0, 0xd3, 0xb6, 0xf1, 0xf8, 0xc5, 0xd2, 0xc5, 0x01,
	0x70, 0x82, 0xb2, 0xb3, 0x57, 0x11, 0xba, 0x86, 0x81, 0xb9, 0x14, 0x3d, 0x69, 0xc1, 0x83, 0x33,
	0xe9, 0xf2, 0xb8, 0xd8, 0x6e, 0x7b, 0x03, 0xc3, 0x26, 0x47, 0xe8, 0x80, 0x77, 0x92, 0x9f, 0x5e,
	0x9e, 0x54, 0x5b, 0xda, 0x5b, 0x18, 0x7b, 0x57, 0x54, 0x47, 0xe5, 0xb1, 0xb1, 0x1d, 0x95, 0x1e,
	0x68, 0x54, 0xae, 0xcf, 0xed, 0x0f, 0xff, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x3f,
	0xbb, 0xea, 0x01, 0x04, 0x00, 0x00,
}
