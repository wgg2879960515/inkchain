// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/configuration.proto

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes" json:"preferred_max_bytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount uint64 `protobuf:"varint,1,opt,name=max_count,json=maxCount" json:"max_count,omitempty"`
}

func (m *ChannelRestrictions) Reset()                    { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string            { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()               {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4f, 0xc2, 0x40,
	0x14, 0x84, 0x53, 0x21, 0x22, 0x1b, 0x89, 0xb2, 0x5c, 0x9a, 0x70, 0x21, 0xd5, 0x03, 0x51, 0x53,
	0x12, 0xfd, 0x07, 0xe5, 0x68, 0xb8, 0x14, 0x4e, 0x5e, 0xc8, 0xb6, 0x3c, 0xda, 0x0d, 0x74, 0x5f,
	0xf3, 0x76, 0x37, 0x29, 0xfe, 0x0f, 0xff, 0xaf, 0xd9, 0x6d, 0x2b, 0xde, 0xe6, 0xcd, 0x7c, 0x9b,
	0xcc, 0x0e, 0x9b, 0x23, 0x1d, 0x80, 0x80, 0x56, 0x39, 0xaa, 0xa3, 0x2c, 0x2c, 0x09, 0x23, 0x51,
	0xc5, 0x35, 0xa1, 0x41, 0x3e, 0xea, 0xc2, 0xe8, 0x89, 0x4d, 0xd6, 0xa8, 0x34, 0x28, 0x6d, 0xf5,
	0xee, 0x52, 0x03, 0xe7, 0x6c, 0x68, 0x2e, 0x35, 0x84, 0xc1, 0x22, 0x58, 0x8e, 0x53, 0xaf, 0xa3,
	0x9f, 0x80, 0x8d, 0x13, 0x61, 0xf2, 0x72, 0x2b, 0xbf, 0x81, 0xbf, 0xb0, 0x69, 0x25, 0x9a, 0x7d,
	0x05, 0x5a, 0x8b, 0x02, 0xf6, 0x39, 0x5a, 0x65, 0x3c, 0x3e, 0x49, 0x1f, 0x2a, 0xd1, 0x6c, 0x5a,
	0x7f, 0xed, 0x6c, 0xfe, 0xc6, 0xb8, 0xc8, 0x34, 0x9e, 0xad, 0x81, 0xbd, 0x7b, 0x94, 0x5d, 0x0c,
	0xe8, 0xf0, 0xc6, 0xc3, 0x8f, 0x7d, 0xb2, 0x11, 0x4d, 0xe2, 0x7c, 0x1e, 0xb3, 0x59, 0x4d, 0x70,
	0x04, 0x22, 0x38, 0xfc, 0xc3, 0x07, 0x1e, 0x9f, 0xfe, 0x45, 0x3d, 0x1f, 0x2d, 0xd9, 0xbd, 0xaf,
	0xb5, 0x93, 0x15, 0xa0, 0x35, 0x3c, 0x64, 0x23, 0xd3, 0xca, 0xae, 0x7e, 0x7f, 0x3a, 0xf2, 0x53,
	0x1c, 0x4f, 0x22, 0x21, 0x3c, 0x01, 0x69, 0x47, 0x66, 0xad, 0x0c, 0x83, 0xc5, 0xc0, 0x91, 0xdd,
	0x19, 0xbd, 0xb3, 0xd9, 0xba, 0x14, 0x4a, 0xc1, 0x39, 0x05, 0x6d, 0x48, 0xe6, 0x6e, 0x35, 0xcd,
	0xe7, 0x6c, 0xec, 0x0a, 0x5d, 0x3f, 0x3b, 0x4c, 0xef, 0x2a, 0xd1, 0xf8, 0x5f, 0x26, 0x5b, 0xf6,
	0x8c, 0x54, 0xc4, 0x52, 0x9d, 0xf2, 0x52, 0x48, 0x75, 0x15, 0x7e, 0x6c, 0x1d, 0x77, 0x63, 0x7f,
	0xbd, 0x16, 0xd2, 0x94, 0x36, 0x8b, 0x73, 0xac, 0x56, 0x3d, 0x73, 0x15, 0x2d, 0xbc, 0xea, 0xe0,
	0xec, 0xd6, 0xdf, 0x1f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0x81, 0x29, 0x0c, 0xc8, 0x01,
	0x00, 0x00,
}
