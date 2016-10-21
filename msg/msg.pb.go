// Code generated by protoc-gen-gogo.
// source: msg.proto
// DO NOT EDIT!

/*
	Package msg is a generated protocol buffer package.

	It is generated from these files:
		msg.proto

	It has these top-level messages:
		Message
		MessageList
*/
package msg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_Subscribe   MessageType = 0
	MessageType_Publish     MessageType = 1
	MessageType_UnSubscribe MessageType = 2
	MessageType_Ack         MessageType = 3
	MessageType_Heartbeat   MessageType = 4
	MessageType_Error       MessageType = 5
	MessageType_Close       MessageType = 6
)

var MessageType_name = map[int32]string{
	0: "Subscribe",
	1: "Publish",
	2: "UnSubscribe",
	3: "Ack",
	4: "Heartbeat",
	5: "Error",
	6: "Close",
}
var MessageType_value = map[string]int32{
	"Subscribe":   0,
	"Publish":     1,
	"UnSubscribe": 2,
	"Ack":         3,
	"Heartbeat":   4,
	"Error":       5,
	"Close":       6,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

type PublishType int32

const (
	PublishType_DirectType PublishType = 0
	PublishType_FanoutType PublishType = 1
)

var PublishType_name = map[int32]string{
	0: "DirectType",
	1: "FanoutType",
}
var PublishType_value = map[string]int32{
	"DirectType": 0,
	"FanoutType": 1,
}

func (x PublishType) Enum() *PublishType {
	p := new(PublishType)
	*p = x
	return p
}
func (x PublishType) String() string {
	return proto.EnumName(PublishType_name, int32(x))
}
func (x *PublishType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PublishType_value, data, "PublishType")
	if err != nil {
		return err
	}
	*x = PublishType(value)
	return nil
}
func (PublishType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

type Message struct {
	MsgId     uint64      `protobuf:"varint,1,opt,name=msg_id" json:"msg_id"`
	Type      MessageType `protobuf:"varint,2,opt,name=type,enum=MessageType" json:"type"`
	Topic     string      `protobuf:"bytes,3,opt,name=topic" json:"topic"`
	Filter    string      `protobuf:"bytes,4,opt,name=filter" json:"filter"`
	Body      []byte      `protobuf:"bytes,5,opt,name=body" json:"body,omitempty"`
	Timestamp int64       `protobuf:"varint,6,opt,name=timestamp" json:"timestamp"`
	PubType   PublishType `protobuf:"varint,7,opt,name=pubType,enum=PublishType" json:"pubType"`
	Count     int64       `protobuf:"varint,8,opt,name=count" json:"count"`
	Persist   bool        `protobuf:"varint,9,opt,name=persist" json:"persist"`
	NeedAck   bool        `protobuf:"varint,10,opt,name=need_ack" json:"need_ack"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

func (m *Message) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_Subscribe
}

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Message) GetPubType() PublishType {
	if m != nil {
		return m.PubType
	}
	return PublishType_DirectType
}

func (m *Message) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Message) GetPersist() bool {
	if m != nil {
		return m.Persist
	}
	return false
}

func (m *Message) GetNeedAck() bool {
	if m != nil {
		return m.NeedAck
	}
	return false
}

type MessageList struct {
	Msgs []*Message `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *MessageList) Reset()                    { *m = MessageList{} }
func (m *MessageList) String() string            { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()               {}
func (*MessageList) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

func (m *MessageList) GetMsgs() []*Message {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*MessageList)(nil), "MessageList")
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("PublishType", PublishType_name, PublishType_value)
}
func (m *Message) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Message) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMsg(data, i, uint64(m.MsgId))
	data[i] = 0x10
	i++
	i = encodeVarintMsg(data, i, uint64(m.Type))
	data[i] = 0x1a
	i++
	i = encodeVarintMsg(data, i, uint64(len(m.Topic)))
	i += copy(data[i:], m.Topic)
	data[i] = 0x22
	i++
	i = encodeVarintMsg(data, i, uint64(len(m.Filter)))
	i += copy(data[i:], m.Filter)
	if m.Body != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintMsg(data, i, uint64(len(m.Body)))
		i += copy(data[i:], m.Body)
	}
	data[i] = 0x30
	i++
	i = encodeVarintMsg(data, i, uint64(m.Timestamp))
	data[i] = 0x38
	i++
	i = encodeVarintMsg(data, i, uint64(m.PubType))
	data[i] = 0x40
	i++
	i = encodeVarintMsg(data, i, uint64(m.Count))
	data[i] = 0x48
	i++
	if m.Persist {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x50
	i++
	if m.NeedAck {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *MessageList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MessageList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, msg := range m.Msgs {
			data[i] = 0xa
			i++
			i = encodeVarintMsg(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Msg(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Msg(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMsg(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMsg(uint64(m.MsgId))
	n += 1 + sovMsg(uint64(m.Type))
	l = len(m.Topic)
	n += 1 + l + sovMsg(uint64(l))
	l = len(m.Filter)
	n += 1 + l + sovMsg(uint64(l))
	if m.Body != nil {
		l = len(m.Body)
		n += 1 + l + sovMsg(uint64(l))
	}
	n += 1 + sovMsg(uint64(m.Timestamp))
	n += 1 + sovMsg(uint64(m.PubType))
	n += 1 + sovMsg(uint64(m.Count))
	n += 2
	n += 2
	return n
}

func (m *MessageList) Size() (n int) {
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovMsg(uint64(l))
		}
	}
	return n
}

func sovMsg(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgId", wireType)
			}
			m.MsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MsgId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filter = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], data[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubType", wireType)
			}
			m.PubType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PubType |= (PublishType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Persist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Persist = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedAck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NeedAck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MessageList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &Message{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsg(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMsg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsg
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMsg(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMsg = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msg.proto", fileDescriptorMsg) }

var fileDescriptorMsg = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x8e, 0xda, 0x30,
	0x14, 0x85, 0x09, 0x49, 0x48, 0x72, 0x43, 0x69, 0xe4, 0xfe, 0x59, 0x5d, 0x50, 0x44, 0x55, 0x09,
	0x21, 0x01, 0x12, 0x6f, 0x50, 0xfa, 0xa3, 0x2e, 0x5a, 0xa9, 0x52, 0xdb, 0x35, 0x8a, 0x83, 0x09,
	0x16, 0x24, 0x8e, 0x6c, 0x67, 0xc1, 0xb6, 0x4f, 0xd0, 0xc7, 0x62, 0xd9, 0x27, 0xa8, 0xaa, 0x99,
	0x17, 0x99, 0x9b, 0xe0, 0x19, 0x98, 0x85, 0x25, 0x9f, 0xcf, 0xe7, 0xf8, 0xf8, 0x1a, 0xa2, 0x42,
	0xe7, 0xf3, 0x4a, 0x49, 0x23, 0x5f, 0xcf, 0x72, 0x61, 0x76, 0x35, 0x9b, 0x67, 0xb2, 0x58, 0xe4,
	0x32, 0x97, 0x8b, 0x16, 0xb3, 0x7a, 0xdb, 0xaa, 0x56, 0xb4, 0xbb, 0xb3, 0x7d, 0xfc, 0xbb, 0x0b,
	0xc1, 0x37, 0xae, 0x75, 0x9a, 0x73, 0xf2, 0x1c, 0x7a, 0x78, 0xcf, 0x5a, 0x6c, 0xa8, 0x33, 0x72,
	0x26, 0xde, 0xca, 0x3b, 0xfd, 0x7b, 0xd3, 0x21, 0x23, 0xf0, 0xcc, 0xb1, 0xe2, 0xb4, 0x8b, 0x6c,
	0xb0, 0xec, 0xcf, 0xad, 0xfb, 0x27, 0x32, 0xeb, 0x78, 0x06, 0xbe, 0x91, 0x95, 0xc8, 0xa8, 0x8b,
	0x96, 0xc8, 0x42, 0xbc, 0x6c, 0x2b, 0x0e, 0x86, 0x2b, 0xea, 0x5d, 0x51, 0x02, 0x1e, 0x93, 0x9b,
	0x23, 0xf5, 0x91, 0xf5, 0x5b, 0xe6, 0x90, 0x57, 0x10, 0x19, 0x51, 0x70, 0x6d, 0xd2, 0xa2, 0xa2,
	0x3d, 0x3c, 0x70, 0xad, 0xf9, 0x2d, 0x04, 0x55, 0xcd, 0x9a, 0x22, 0x1a, 0xd8, 0xf2, 0xef, 0x35,
	0x3b, 0x08, 0xbd, 0x7b, 0x5c, 0x9e, 0xc9, 0xba, 0x34, 0x34, 0xbc, 0x4a, 0xbe, 0xc0, 0x24, 0x57,
	0x5a, 0x68, 0x43, 0x23, 0xc4, 0xa1, 0xc5, 0x2f, 0x21, 0x2c, 0x39, 0xdf, 0xac, 0xd3, 0x6c, 0x4f,
	0xe1, 0xc2, 0xc7, 0xef, 0x20, 0xb6, 0x53, 0x7d, 0xc5, 0x08, 0xda, 0x3c, 0xfc, 0x07, 0x8d, 0xbf,
	0xe0, 0x4e, 0xe2, 0x65, 0x78, 0x3f, 0xf1, 0x74, 0xf7, 0x60, 0x6b, 0xfa, 0xc9, 0x13, 0x88, 0x7e,
	0xd4, 0x4c, 0x67, 0x4a, 0x30, 0x9e, 0x74, 0x48, 0x0c, 0x81, 0x7d, 0x5d, 0xe2, 0x90, 0xa7, 0x10,
	0xff, 0x2a, 0x2f, 0xa7, 0x5d, 0x12, 0x80, 0xfb, 0x3e, 0xdb, 0x27, 0x6e, 0x93, 0xfa, 0xc2, 0x53,
	0x65, 0x18, 0x4f, 0x4d, 0xe2, 0x91, 0x08, 0xfc, 0x4f, 0x4a, 0x49, 0x95, 0xf8, 0xcd, 0xf6, 0xc3,
	0x41, 0x6a, 0x9e, 0xf4, 0xa6, 0x33, 0x88, 0xaf, 0x26, 0x25, 0x03, 0x80, 0x8f, 0x42, 0xf1, 0xcc,
	0x34, 0x0a, 0xab, 0x50, 0x7f, 0x4e, 0x4b, 0x59, 0x9f, 0xb5, 0xb3, 0x4a, 0x4e, 0x37, 0x43, 0xe7,
	0x2f, 0xae, 0xff, 0xb8, 0xfe, 0xdc, 0x0e, 0x3b, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0xac,
	0xa3, 0xf5, 0x11, 0x02, 0x00, 0x00,
}
