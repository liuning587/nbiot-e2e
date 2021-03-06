// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/telenordigital/nbiot-e2e/server/pb/nanopb/generator/proto"
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

type Message struct {
	// Types that are valid to be assigned to Message:
	//	*Message_PingMessage
	Message              isMessage_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Message interface {
	isMessage_Message()
}

type Message_PingMessage struct {
	PingMessage *PingMessage `protobuf:"bytes,1,opt,name=ping_message,json=pingMessage,proto3,oneof"`
}

func (*Message_PingMessage) isMessage_Message() {}

func (m *Message) GetMessage() isMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Message) GetPingMessage() *PingMessage {
	if x, ok := m.GetMessage().(*Message_PingMessage); ok {
		return x.PingMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_PingMessage)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// message
	switch x := m.Message.(type) {
	case *Message_PingMessage:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PingMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Message has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // message.ping_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PingMessage)
		err := b.DecodeMessage(msg)
		m.Message = &Message_PingMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// message
	switch x := m.Message.(type) {
	case *Message_PingMessage:
		s := proto.Size(x.PingMessage)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PingMessage struct {
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	PrevRssi             float32  `protobuf:"fixed32,3,opt,name=prev_rssi,json=prevRssi,proto3" json:"prev_rssi,omitempty"`
	NbiotLibHash         uint32   `protobuf:"varint,4,opt,name=nbiot_lib_hash,json=nbiotLibHash,proto3" json:"nbiot_lib_hash,omitempty"`
	E2EHash              uint32   `protobuf:"varint,5,opt,name=e2e_hash,json=e2eHash,proto3" json:"e2e_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PingMessage) GetPrevRssi() float32 {
	if m != nil {
		return m.PrevRssi
	}
	return 0
}

func (m *PingMessage) GetNbiotLibHash() uint32 {
	if m != nil {
		return m.NbiotLibHash
	}
	return 0
}

func (m *PingMessage) GetE2EHash() uint32 {
	if m != nil {
		return m.E2EHash
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "nbiot_e2e.Message")
	proto.RegisterType((*PingMessage)(nil), "nbiot_e2e.PingMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x49, 0x5b, 0x48, 0xe2, 0xb4, 0x08, 0x79, 0x40, 0xa1, 0x2c, 0x55, 0x61, 0xe8, 0xd2,
	0x58, 0x0a, 0x23, 0x5b, 0xa7, 0x0a, 0x81, 0x04, 0x19, 0x59, 0x22, 0xbb, 0x3c, 0x39, 0x96, 0x52,
	0xdb, 0xf8, 0xb9, 0xfd, 0x0b, 0xfe, 0x19, 0xd5, 0x8e, 0x0a, 0xe3, 0xbb, 0xf7, 0xdc, 0x63, 0xc9,
	0x64, 0xb6, 0x07, 0x44, 0x2e, 0xa1, 0xb2, 0xce, 0x78, 0x43, 0x73, 0x2d, 0x94, 0xf1, 0x2d, 0xd4,
	0x30, 0x7f, 0xd0, 0x5c, 0x1b, 0x2b, 0x98, 0x04, 0x0d, 0x8e, 0x7b, 0xe3, 0x58, 0x40, 0x58, 0x8c,
	0x23, 0xbf, 0xfc, 0x20, 0xe9, 0x5b, 0x14, 0xd0, 0x67, 0x32, 0xb5, 0x4a, 0xcb, 0x76, 0x10, 0x96,
	0xc9, 0x22, 0x59, 0x15, 0xf5, 0x6d, 0x75, 0x36, 0x56, 0xef, 0x4a, 0xcb, 0x81, 0xde, 0x5e, 0x34,
	0x85, 0xfd, 0x3b, 0x37, 0x39, 0x49, 0x87, 0xdd, 0xf2, 0x27, 0x21, 0xc5, 0x3f, 0x92, 0xce, 0x49,
	0x86, 0xf0, 0x7d, 0x00, 0xbd, 0x8b, 0xce, 0x59, 0x73, 0xbe, 0xe9, 0x3d, 0xc9, 0xad, 0x83, 0x63,
	0xeb, 0x10, 0x55, 0x39, 0x5e, 0x24, 0xab, 0x51, 0x93, 0x9d, 0x82, 0x06, 0x51, 0xd1, 0x47, 0x72,
	0x1d, 0xdf, 0xee, 0x95, 0x68, 0x3b, 0x8e, 0x5d, 0x39, 0x09, 0xf3, 0x69, 0x48, 0x5f, 0x95, 0xd8,
	0x72, 0xec, 0xe8, 0x1d, 0xc9, 0xa0, 0x86, 0xd8, 0x5f, 0x86, 0x3e, 0x85, 0x1a, 0x4e, 0xd5, 0xcb,
	0x24, 0x1b, 0xdd, 0x8c, 0x37, 0xec, 0x73, 0x2d, 0x95, 0xef, 0x0e, 0xa2, 0xda, 0x99, 0x3d, 0xf3,
	0xd0, 0x83, 0x36, 0xee, 0x4b, 0x49, 0xe5, 0x79, 0xcf, 0x82, 0x6a, 0x0d, 0x35, 0x30, 0x04, 0x77,
	0x04, 0xc7, 0xac, 0x10, 0x57, 0xe1, 0x6b, 0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61, 0x5e,
	0x0a, 0x00, 0x5b, 0x01, 0x00, 0x00,
}
