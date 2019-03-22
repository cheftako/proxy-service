// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/agent.proto

package agent

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

type PacketType int32

const (
	PacketType_DIAL_REQ PacketType = 0
	PacketType_DIAL_RSP PacketType = 1
	PacketType_ACK      PacketType = 2
	PacketType_CLOSE    PacketType = 3
	PacketType_DATA     PacketType = 4
)

var PacketType_name = map[int32]string{
	0: "DIAL_REQ",
	1: "DIAL_RSP",
	2: "ACK",
	3: "CLOSE",
	4: "DATA",
}
var PacketType_value = map[string]int32{
	"DIAL_REQ": 0,
	"DIAL_RSP": 1,
	"ACK":      2,
	"CLOSE":    3,
	"DATA":     4,
}

func (x PacketType) String() string {
	return proto.EnumName(PacketType_name, int32(x))
}
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{0}
}

type Error int32

const (
	Error_EOF Error = 0
)

var Error_name = map[int32]string{
	0: "EOF",
}
var Error_value = map[string]int32{
	"EOF": 0,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{1}
}

type Packet struct {
	Type PacketType `protobuf:"varint,1,opt,name=type,proto3,enum=PacketType" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Packet_DialRequest
	//	*Packet_DialResponse
	//	*Packet_Data
	Payload              isPacket_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (dst *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(dst, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetType() PacketType {
	if m != nil {
		return m.Type
	}
	return PacketType_DIAL_REQ
}

type isPacket_Payload interface {
	isPacket_Payload()
}

type Packet_DialRequest struct {
	DialRequest *DialRequest `protobuf:"bytes,2,opt,name=dialRequest,proto3,oneof"`
}

type Packet_DialResponse struct {
	DialResponse *DialResponse `protobuf:"bytes,3,opt,name=dialResponse,proto3,oneof"`
}

type Packet_Data struct {
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

func (*Packet_DialRequest) isPacket_Payload() {}

func (*Packet_DialResponse) isPacket_Payload() {}

func (*Packet_Data) isPacket_Payload() {}

func (m *Packet) GetPayload() isPacket_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Packet) GetDialRequest() *DialRequest {
	if x, ok := m.GetPayload().(*Packet_DialRequest); ok {
		return x.DialRequest
	}
	return nil
}

func (m *Packet) GetDialResponse() *DialResponse {
	if x, ok := m.GetPayload().(*Packet_DialResponse); ok {
		return x.DialResponse
	}
	return nil
}

func (m *Packet) GetData() *Data {
	if x, ok := m.GetPayload().(*Packet_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Packet) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Packet_OneofMarshaler, _Packet_OneofUnmarshaler, _Packet_OneofSizer, []interface{}{
		(*Packet_DialRequest)(nil),
		(*Packet_DialResponse)(nil),
		(*Packet_Data)(nil),
	}
}

func _Packet_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Packet)
	// payload
	switch x := m.Payload.(type) {
	case *Packet_DialRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DialRequest); err != nil {
			return err
		}
	case *Packet_DialResponse:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DialResponse); err != nil {
			return err
		}
	case *Packet_Data:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Packet.Payload has unexpected type %T", x)
	}
	return nil
}

func _Packet_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Packet)
	switch tag {
	case 2: // payload.dialRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DialRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Packet_DialRequest{msg}
		return true, err
	case 3: // payload.dialResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DialResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &Packet_DialResponse{msg}
		return true, err
	case 4: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Data)
		err := b.DecodeMessage(msg)
		m.Payload = &Packet_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Packet_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Packet)
	// payload
	switch x := m.Payload.(type) {
	case *Packet_DialRequest:
		s := proto.Size(x.DialRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_DialResponse:
		s := proto.Size(x.DialResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_Data:
		s := proto.Size(x.Data)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DialRequest struct {
	// tcp or udp?
	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// node:port
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// random id for client, maybe should be longer
	Random               int64    `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DialRequest) Reset()         { *m = DialRequest{} }
func (m *DialRequest) String() string { return proto.CompactTextString(m) }
func (*DialRequest) ProtoMessage()    {}
func (*DialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{1}
}
func (m *DialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialRequest.Unmarshal(m, b)
}
func (m *DialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialRequest.Marshal(b, m, deterministic)
}
func (dst *DialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialRequest.Merge(dst, src)
}
func (m *DialRequest) XXX_Size() int {
	return xxx_messageInfo_DialRequest.Size(m)
}
func (m *DialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DialRequest proto.InternalMessageInfo

func (m *DialRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DialRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DialRequest) GetRandom() int64 {
	if m != nil {
		return m.Random
	}
	return 0
}

type DialResponse struct {
	// error failed reason; enum?
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// connectID indicates the identifier of the connection
	ConnectID int64 `protobuf:"varint,2,opt,name=connectID,proto3" json:"connectID,omitempty"`
	// random copied from DialRequest
	Random               int64    `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DialResponse) Reset()         { *m = DialResponse{} }
func (m *DialResponse) String() string { return proto.CompactTextString(m) }
func (*DialResponse) ProtoMessage()    {}
func (*DialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{2}
}
func (m *DialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialResponse.Unmarshal(m, b)
}
func (m *DialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialResponse.Marshal(b, m, deterministic)
}
func (dst *DialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialResponse.Merge(dst, src)
}
func (m *DialResponse) XXX_Size() int {
	return xxx_messageInfo_DialResponse.Size(m)
}
func (m *DialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DialResponse proto.InternalMessageInfo

func (m *DialResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DialResponse) GetConnectID() int64 {
	if m != nil {
		return m.ConnectID
	}
	return 0
}

func (m *DialResponse) GetRandom() int64 {
	if m != nil {
		return m.Random
	}
	return 0
}

type Data struct {
	// connectID to connect to
	ConnectID int64 `protobuf:"varint,1,opt,name=connectID,proto3" json:"connectID,omitempty"`
	// error message if error happens
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// stream data
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_e6dff374c951db2a, []int{3}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetConnectID() int64 {
	if m != nil {
		return m.ConnectID
	}
	return 0
}

func (m *Data) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "Packet")
	proto.RegisterType((*DialRequest)(nil), "DialRequest")
	proto.RegisterType((*DialResponse)(nil), "DialResponse")
	proto.RegisterType((*Data)(nil), "Data")
	proto.RegisterEnum("PacketType", PacketType_name, PacketType_value)
	proto.RegisterEnum("Error", Error_name, Error_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyServiceClient interface {
	Proxy(ctx context.Context, opts ...grpc.CallOption) (ProxyService_ProxyClient, error)
}

type proxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewProxyServiceClient(cc *grpc.ClientConn) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) Proxy(ctx context.Context, opts ...grpc.CallOption) (ProxyService_ProxyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProxyService_serviceDesc.Streams[0], "/ProxyService/Proxy", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceProxyClient{stream}
	return x, nil
}

type ProxyService_ProxyClient interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ClientStream
}

type proxyServiceProxyClient struct {
	grpc.ClientStream
}

func (x *proxyServiceProxyClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyServiceProxyClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServiceServer is the server API for ProxyService service.
type ProxyServiceServer interface {
	Proxy(ProxyService_ProxyServer) error
}

func RegisterProxyServiceServer(s *grpc.Server, srv ProxyServiceServer) {
	s.RegisterService(&_ProxyService_serviceDesc, srv)
}

func _ProxyService_Proxy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServiceServer).Proxy(&proxyServiceProxyServer{stream})
}

type ProxyService_ProxyServer interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type proxyServiceProxyServer struct {
	grpc.ServerStream
}

func (x *proxyServiceProxyServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyServiceProxyServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Proxy",
			Handler:       _ProxyService_Proxy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent/agent.proto",
}

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	// Agent Identifier?
	Connect(ctx context.Context, opts ...grpc.CallOption) (AgentService_ConnectClient, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (AgentService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentService_serviceDesc.Streams[0], "/AgentService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceConnectClient{stream}
	return x, nil
}

type AgentService_ConnectClient interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ClientStream
}

type agentServiceConnectClient struct {
	grpc.ClientStream
}

func (x *agentServiceConnectClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceConnectClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	// Agent Identifier?
	Connect(AgentService_ConnectServer) error
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).Connect(&agentServiceConnectServer{stream})
}

type AgentService_ConnectServer interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type agentServiceConnectServer struct {
	grpc.ServerStream
}

func (x *agentServiceConnectServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceConnectServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _AgentService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent/agent.proto",
}

func init() { proto.RegisterFile("agent/agent.proto", fileDescriptor_agent_e6dff374c951db2a) }

var fileDescriptor_agent_e6dff374c951db2a = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x65, 0x0a, 0x94, 0x72, 0x41, 0x83, 0x37, 0xc6, 0x90, 0xd5, 0x64, 0x57, 0x9e, 0x9a, 0x7d,
	0x60, 0x57, 0xf6, 0x0b, 0xb0, 0xb0, 0xe9, 0xc6, 0x8d, 0x5b, 0xa7, 0xfb, 0xa4, 0x0f, 0x66, 0x84,
	0x89, 0x69, 0xac, 0x0c, 0x0e, 0xa3, 0x91, 0x3f, 0xf3, 0xf3, 0x4c, 0x07, 0x5a, 0xe8, 0x26, 0x7d,
	0x21, 0x9c, 0x33, 0xf7, 0xdc, 0x73, 0xcf, 0xcc, 0x85, 0x17, 0xec, 0x3b, 0xaf, 0xd4, 0x95, 0xfe,
	0xc6, 0xb5, 0x14, 0x4a, 0x44, 0xff, 0x08, 0x4c, 0x57, 0xac, 0xf8, 0xc1, 0x15, 0x9e, 0x83, 0xa5,
	0xda, 0x9a, 0x87, 0xe4, 0x82, 0xcc, 0x9f, 0x27, 0x5e, 0xdc, 0xd1, 0x8f, 0x6d, 0xcd, 0xa9, 0x3e,
	0xc0, 0x6b, 0xf0, 0xca, 0x0d, 0xdb, 0x52, 0xfe, 0xeb, 0x37, 0x6f, 0x54, 0x38, 0xb9, 0x20, 0x73,
	0x2f, 0xf1, 0xe3, 0x6c, 0xe0, 0x96, 0x06, 0x1d, 0x97, 0xe0, 0x0d, 0xf8, 0x1d, 0x6c, 0x6a, 0x51,
	0x35, 0x3c, 0x34, 0xb5, 0xe4, 0x59, 0x2f, 0xe9, 0xc8, 0xa5, 0x41, 0x8f, 0x8a, 0xf0, 0x35, 0x58,
	0x25, 0x53, 0x2c, 0xb4, 0x74, 0xb1, 0x1d, 0x67, 0x4c, 0xb1, 0xa5, 0x41, 0x35, 0xf9, 0xde, 0x05,
	0xa7, 0x66, 0xed, 0x56, 0xb0, 0x32, 0xfa, 0x02, 0xde, 0xc8, 0x1a, 0xcf, 0x60, 0xa6, 0x23, 0x15,
	0x62, 0xab, 0x23, 0xb8, 0xf4, 0x80, 0x31, 0x04, 0x87, 0x95, 0xa5, 0xe4, 0x4d, 0xa3, 0xa7, 0x76,
	0xe9, 0x1e, 0xe2, 0x2b, 0x98, 0x4a, 0x56, 0x95, 0xe2, 0xa7, 0x9e, 0xcd, 0xa4, 0x3d, 0x8a, 0x3e,
	0x83, 0x3f, 0x1e, 0x12, 0x5f, 0x82, 0xcd, 0xa5, 0x14, 0xb2, 0x6f, 0xdd, 0x01, 0x7c, 0x03, 0x6e,
	0x21, 0xaa, 0x8a, 0x17, 0xea, 0x2e, 0xd3, 0x9d, 0x4d, 0x3a, 0x10, 0x27, 0x7b, 0x7f, 0x04, 0x6b,
	0x97, 0xe9, 0x58, 0x4d, 0x9e, 0xaa, 0x0f, 0x8e, 0x93, 0xb1, 0x23, 0xf6, 0x97, 0xb3, 0xeb, 0xe8,
	0x77, 0x77, 0x72, 0x79, 0x0b, 0x30, 0xbc, 0x15, 0xfa, 0x30, 0xcb, 0xee, 0xd2, 0xfb, 0xaf, 0x34,
	0xff, 0x14, 0x18, 0x03, 0x5a, 0xaf, 0x02, 0x82, 0x0e, 0x98, 0xe9, 0xe2, 0x43, 0x30, 0x41, 0x17,
	0xec, 0xc5, 0xfd, 0xc3, 0x3a, 0x0f, 0x4c, 0x9c, 0x81, 0x95, 0xa5, 0x8f, 0x69, 0x60, 0x5d, 0x06,
	0x60, 0xe7, 0xda, 0xc4, 0x01, 0x33, 0x7f, 0xb8, 0x0d, 0x8c, 0xe4, 0x0a, 0xfc, 0x95, 0x14, 0x7f,
	0xdb, 0x35, 0x97, 0x7f, 0x36, 0x05, 0xc7, 0x73, 0xb0, 0x35, 0x46, 0xa7, 0xdf, 0x8e, 0xb3, 0xfd,
	0x4f, 0x64, 0xcc, 0xc9, 0x35, 0x49, 0xde, 0x81, 0x9f, 0xee, 0xb6, 0x6b, 0x2f, 0x78, 0x0b, 0xce,
	0xa2, 0x4b, 0x74, 0x4a, 0xf2, 0x6d, 0xaa, 0x5f, 0xe9, 0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x34, 0xb1, 0x6b, 0x60, 0x9d, 0x02, 0x00, 0x00,
}
