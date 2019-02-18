// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/agent.proto

package agent

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
	return fileDescriptor_253ddf9934d56132, []int{0}
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
	return fileDescriptor_253ddf9934d56132, []int{1}
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
	return fileDescriptor_253ddf9934d56132, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
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
	return fileDescriptor_253ddf9934d56132, []int{1}
}

func (m *DialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialRequest.Unmarshal(m, b)
}
func (m *DialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialRequest.Marshal(b, m, deterministic)
}
func (m *DialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialRequest.Merge(m, src)
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
	return fileDescriptor_253ddf9934d56132, []int{2}
}

func (m *DialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DialResponse.Unmarshal(m, b)
}
func (m *DialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DialResponse.Marshal(b, m, deterministic)
}
func (m *DialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialResponse.Merge(m, src)
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
	return fileDescriptor_253ddf9934d56132, []int{3}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
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
	proto.RegisterEnum("PacketType", PacketType_name, PacketType_value)
	proto.RegisterEnum("Error", Error_name, Error_value)
	proto.RegisterType((*Packet)(nil), "Packet")
	proto.RegisterType((*DialRequest)(nil), "DialRequest")
	proto.RegisterType((*DialResponse)(nil), "DialResponse")
	proto.RegisterType((*Data)(nil), "Data")
}

func init() { proto.RegisterFile("agent/agent.proto", fileDescriptor_253ddf9934d56132) }

var fileDescriptor_253ddf9934d56132 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x65, 0x0a, 0x94, 0x72, 0x41, 0x83, 0x37, 0xc6, 0x90, 0xd5, 0x64, 0x37, 0x3c, 0x35, 0xfb,
	0xc0, 0x6e, 0xba, 0x5f, 0x80, 0x85, 0xa6, 0x8d, 0x8d, 0xad, 0xd3, 0x3e, 0xe9, 0x83, 0x19, 0x61,
	0x62, 0x1a, 0x2b, 0x83, 0xc3, 0x68, 0xe4, 0xcf, 0xfc, 0x3c, 0xd3, 0x81, 0x16, 0x6a, 0xe2, 0x0b,
	0xe1, 0x9c, 0xb9, 0xe7, 0x9c, 0x7b, 0xef, 0x0c, 0xbc, 0x60, 0x5f, 0x79, 0xa9, 0x1e, 0xf4, 0x37,
	0xae, 0xa4, 0x50, 0x22, 0xfa, 0x43, 0x60, 0xbc, 0x65, 0xf9, 0x37, 0xae, 0xf0, 0x16, 0x2c, 0xd5,
	0x54, 0x3c, 0x24, 0x77, 0x64, 0xfa, 0x7c, 0xe6, 0xc5, 0x2d, 0xbd, 0x6f, 0x2a, 0x4e, 0xf5, 0x01,
	0x3e, 0x82, 0x57, 0x1c, 0xd8, 0x91, 0xf2, 0x1f, 0x3f, 0x79, 0xad, 0xc2, 0xd1, 0x1d, 0x99, 0x7a,
	0x33, 0x3f, 0x4e, 0x7b, 0x6e, 0x69, 0xd0, 0x61, 0x09, 0x3e, 0x81, 0xdf, 0xc2, 0xba, 0x12, 0x65,
	0xcd, 0x43, 0x53, 0x4b, 0x9e, 0x75, 0x92, 0x96, 0x5c, 0x1a, 0xf4, 0xaa, 0x08, 0x5f, 0x83, 0x55,
	0x30, 0xc5, 0x42, 0x4b, 0x17, 0xdb, 0x71, 0xca, 0x14, 0x5b, 0x1a, 0x54, 0x93, 0x6f, 0x5d, 0x70,
	0x2a, 0xd6, 0x1c, 0x05, 0x2b, 0xa2, 0x4f, 0xe0, 0x0d, 0xa2, 0xf1, 0x06, 0x26, 0x7a, 0xa4, 0x5c,
	0x1c, 0xf5, 0x08, 0x2e, 0xbd, 0x60, 0x0c, 0xc1, 0x61, 0x45, 0x21, 0x79, 0x5d, 0xeb, 0xae, 0x5d,
	0x7a, 0x86, 0xf8, 0x0a, 0xc6, 0x92, 0x95, 0x85, 0xf8, 0xae, 0x7b, 0x33, 0x69, 0x87, 0xa2, 0x8f,
	0xe0, 0x0f, 0x9b, 0xc4, 0x97, 0x60, 0x73, 0x29, 0x85, 0xec, 0xac, 0x5b, 0x80, 0x6f, 0xc0, 0xcd,
	0x45, 0x59, 0xf2, 0x5c, 0xad, 0x52, 0xed, 0x6c, 0xd2, 0x9e, 0xf8, 0xaf, 0xf7, 0x7b, 0xb0, 0x4e,
	0x33, 0x5d, 0xab, 0xc9, 0xbf, 0xea, 0x4b, 0xe2, 0x68, 0x98, 0x88, 0xdd, 0x72, 0x4e, 0x8e, 0x7e,
	0xbb, 0x93, 0xfb, 0x05, 0x40, 0x7f, 0x57, 0xe8, 0xc3, 0x24, 0x5d, 0x25, 0xeb, 0xcf, 0x34, 0xfb,
	0x10, 0x18, 0x3d, 0xda, 0x6d, 0x03, 0x82, 0x0e, 0x98, 0xc9, 0xfc, 0x5d, 0x30, 0x42, 0x17, 0xec,
	0xf9, 0x7a, 0xb3, 0xcb, 0x02, 0x13, 0x27, 0x60, 0xa5, 0xc9, 0x3e, 0x09, 0xac, 0xfb, 0x00, 0xec,
	0x4c, 0x87, 0x38, 0x60, 0x66, 0x9b, 0x45, 0x60, 0xcc, 0x1e, 0xc0, 0xdf, 0x4a, 0xf1, 0xbb, 0xd9,
	0x71, 0xf9, 0xeb, 0x90, 0x73, 0xbc, 0x05, 0x5b, 0x63, 0x74, 0xba, 0xd7, 0x71, 0x73, 0xfe, 0x89,
	0x8c, 0x29, 0x79, 0x24, 0x5f, 0xc6, 0x7a, 0xe5, 0x4f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x52, 0x9d, 0xb9, 0x6a, 0x02, 0x00, 0x00,
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