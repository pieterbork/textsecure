// Code generated by protoc-gen-go.
// source: WhisperTextProtocol.proto
// DO NOT EDIT!

/*
Package textsecure is a generated protocol buffer package.

It is generated from these files:
	WhisperTextProtocol.proto
	LocalStorageProtocol.proto
	FingerprintProtocol.proto

It has these top-level messages:
	SignalMessage
	PreKeySignalMessage
	KeyExchangeMessage
	SenderKeyMessage
	SenderKeyDistributionMessage
	DeviceConsistencyCodeMessage
	SessionStructure
	RecordStructure
	PreKeyRecordStructure
	SignedPreKeyRecordStructure
	IdentityKeyPairStructure
	SenderKeyStateStructure
	SenderKeyRecordStructure
	LogicalFingerprint
	CombinedFingerprints
*/
package textsecure

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignalMessage struct {
	RatchetKey       []byte  `protobuf:"bytes,1,opt,name=ratchetKey" json:"ratchetKey,omitempty"`
	Counter          *uint32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	PreviousCounter  *uint32 `protobuf:"varint,3,opt,name=previousCounter" json:"previousCounter,omitempty"`
	Ciphertext       []byte  `protobuf:"bytes,4,opt,name=ciphertext" json:"ciphertext,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignalMessage) Reset()                    { *m = SignalMessage{} }
func (m *SignalMessage) String() string            { return proto.CompactTextString(m) }
func (*SignalMessage) ProtoMessage()               {}
func (*SignalMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignalMessage) GetRatchetKey() []byte {
	if m != nil {
		return m.RatchetKey
	}
	return nil
}

func (m *SignalMessage) GetCounter() uint32 {
	if m != nil && m.Counter != nil {
		return *m.Counter
	}
	return 0
}

func (m *SignalMessage) GetPreviousCounter() uint32 {
	if m != nil && m.PreviousCounter != nil {
		return *m.PreviousCounter
	}
	return 0
}

func (m *SignalMessage) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type PreKeySignalMessage struct {
	RegistrationId   *uint32 `protobuf:"varint,5,opt,name=registrationId" json:"registrationId,omitempty"`
	PreKeyId         *uint32 `protobuf:"varint,1,opt,name=preKeyId" json:"preKeyId,omitempty"`
	SignedPreKeyId   *uint32 `protobuf:"varint,6,opt,name=signedPreKeyId" json:"signedPreKeyId,omitempty"`
	BaseKey          []byte  `protobuf:"bytes,2,opt,name=baseKey" json:"baseKey,omitempty"`
	IdentityKey      []byte  `protobuf:"bytes,3,opt,name=identityKey" json:"identityKey,omitempty"`
	Message          []byte  `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PreKeySignalMessage) Reset()                    { *m = PreKeySignalMessage{} }
func (m *PreKeySignalMessage) String() string            { return proto.CompactTextString(m) }
func (*PreKeySignalMessage) ProtoMessage()               {}
func (*PreKeySignalMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PreKeySignalMessage) GetRegistrationId() uint32 {
	if m != nil && m.RegistrationId != nil {
		return *m.RegistrationId
	}
	return 0
}

func (m *PreKeySignalMessage) GetPreKeyId() uint32 {
	if m != nil && m.PreKeyId != nil {
		return *m.PreKeyId
	}
	return 0
}

func (m *PreKeySignalMessage) GetSignedPreKeyId() uint32 {
	if m != nil && m.SignedPreKeyId != nil {
		return *m.SignedPreKeyId
	}
	return 0
}

func (m *PreKeySignalMessage) GetBaseKey() []byte {
	if m != nil {
		return m.BaseKey
	}
	return nil
}

func (m *PreKeySignalMessage) GetIdentityKey() []byte {
	if m != nil {
		return m.IdentityKey
	}
	return nil
}

func (m *PreKeySignalMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type KeyExchangeMessage struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BaseKey          []byte  `protobuf:"bytes,2,opt,name=baseKey" json:"baseKey,omitempty"`
	RatchetKey       []byte  `protobuf:"bytes,3,opt,name=ratchetKey" json:"ratchetKey,omitempty"`
	IdentityKey      []byte  `protobuf:"bytes,4,opt,name=identityKey" json:"identityKey,omitempty"`
	BaseKeySignature []byte  `protobuf:"bytes,5,opt,name=baseKeySignature" json:"baseKeySignature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyExchangeMessage) Reset()                    { *m = KeyExchangeMessage{} }
func (m *KeyExchangeMessage) String() string            { return proto.CompactTextString(m) }
func (*KeyExchangeMessage) ProtoMessage()               {}
func (*KeyExchangeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyExchangeMessage) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *KeyExchangeMessage) GetBaseKey() []byte {
	if m != nil {
		return m.BaseKey
	}
	return nil
}

func (m *KeyExchangeMessage) GetRatchetKey() []byte {
	if m != nil {
		return m.RatchetKey
	}
	return nil
}

func (m *KeyExchangeMessage) GetIdentityKey() []byte {
	if m != nil {
		return m.IdentityKey
	}
	return nil
}

func (m *KeyExchangeMessage) GetBaseKeySignature() []byte {
	if m != nil {
		return m.BaseKeySignature
	}
	return nil
}

type SenderKeyMessage struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Iteration        *uint32 `protobuf:"varint,2,opt,name=iteration" json:"iteration,omitempty"`
	Ciphertext       []byte  `protobuf:"bytes,3,opt,name=ciphertext" json:"ciphertext,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SenderKeyMessage) Reset()                    { *m = SenderKeyMessage{} }
func (m *SenderKeyMessage) String() string            { return proto.CompactTextString(m) }
func (*SenderKeyMessage) ProtoMessage()               {}
func (*SenderKeyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SenderKeyMessage) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SenderKeyMessage) GetIteration() uint32 {
	if m != nil && m.Iteration != nil {
		return *m.Iteration
	}
	return 0
}

func (m *SenderKeyMessage) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type SenderKeyDistributionMessage struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Iteration        *uint32 `protobuf:"varint,2,opt,name=iteration" json:"iteration,omitempty"`
	ChainKey         []byte  `protobuf:"bytes,3,opt,name=chainKey" json:"chainKey,omitempty"`
	SigningKey       []byte  `protobuf:"bytes,4,opt,name=signingKey" json:"signingKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SenderKeyDistributionMessage) Reset()                    { *m = SenderKeyDistributionMessage{} }
func (m *SenderKeyDistributionMessage) String() string            { return proto.CompactTextString(m) }
func (*SenderKeyDistributionMessage) ProtoMessage()               {}
func (*SenderKeyDistributionMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SenderKeyDistributionMessage) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SenderKeyDistributionMessage) GetIteration() uint32 {
	if m != nil && m.Iteration != nil {
		return *m.Iteration
	}
	return 0
}

func (m *SenderKeyDistributionMessage) GetChainKey() []byte {
	if m != nil {
		return m.ChainKey
	}
	return nil
}

func (m *SenderKeyDistributionMessage) GetSigningKey() []byte {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

type DeviceConsistencyCodeMessage struct {
	Generation       *uint32 `protobuf:"varint,1,opt,name=generation" json:"generation,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeviceConsistencyCodeMessage) Reset()                    { *m = DeviceConsistencyCodeMessage{} }
func (m *DeviceConsistencyCodeMessage) String() string            { return proto.CompactTextString(m) }
func (*DeviceConsistencyCodeMessage) ProtoMessage()               {}
func (*DeviceConsistencyCodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeviceConsistencyCodeMessage) GetGeneration() uint32 {
	if m != nil && m.Generation != nil {
		return *m.Generation
	}
	return 0
}

func (m *DeviceConsistencyCodeMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*SignalMessage)(nil), "textsecure.SignalMessage")
	proto.RegisterType((*PreKeySignalMessage)(nil), "textsecure.PreKeySignalMessage")
	proto.RegisterType((*KeyExchangeMessage)(nil), "textsecure.KeyExchangeMessage")
	proto.RegisterType((*SenderKeyMessage)(nil), "textsecure.SenderKeyMessage")
	proto.RegisterType((*SenderKeyDistributionMessage)(nil), "textsecure.SenderKeyDistributionMessage")
	proto.RegisterType((*DeviceConsistencyCodeMessage)(nil), "textsecure.DeviceConsistencyCodeMessage")
}

func init() { proto.RegisterFile("WhisperTextProtocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0x05, 0x96, 0xa1, 0x5d, 0x56, 0xe1, 0x12, 0x56, 0x15, 0xaa, 0x22, 0x81, 0x2a,
	0x0e, 0xbd, 0xf1, 0x02, 0xdb, 0xe5, 0xb0, 0xaa, 0x90, 0xaa, 0x2c, 0x12, 0x17, 0x0e, 0xa4, 0xce,
	0xc8, 0x19, 0xa9, 0x6b, 0x47, 0xb6, 0x53, 0x9a, 0x37, 0xe0, 0xce, 0x6b, 0xf0, 0x3e, 0xbc, 0x0e,
	0xb2, 0x9d, 0xbf, 0x26, 0x82, 0x03, 0xb7, 0xcc, 0x97, 0xb1, 0xbf, 0x6f, 0xbe, 0xcf, 0x03, 0xaf,
	0xbf, 0xe4, 0xa4, 0x0b, 0x54, 0x9f, 0xf1, 0x64, 0x76, 0x4a, 0x1a, 0xc9, 0xe4, 0x61, 0x5d, 0xd8,
	0x8f, 0x10, 0x0c, 0x9e, 0x8c, 0x46, 0x56, 0x2a, 0x8c, 0x7f, 0x06, 0x30, 0x7f, 0x20, 0x2e, 0xd2,
	0xc3, 0x27, 0xd4, 0x3a, 0xe5, 0x18, 0xbe, 0x01, 0x50, 0xa9, 0x61, 0x39, 0x9a, 0x2d, 0x56, 0x51,
	0xb0, 0x0c, 0x56, 0xb3, 0xa4, 0x87, 0x84, 0x11, 0x3c, 0x63, 0xb2, 0x14, 0x06, 0x55, 0x34, 0x59,
	0x06, 0xab, 0x79, 0xd2, 0x94, 0xe1, 0x0a, 0x5e, 0x16, 0x0a, 0x8f, 0x24, 0x4b, 0xbd, 0xa9, 0x3b,
	0xa6, 0xae, 0x63, 0x08, 0x5b, 0x0e, 0x46, 0x45, 0x8e, 0xca, 0x2a, 0x89, 0x2e, 0x3c, 0x47, 0x87,
	0xc4, 0xbf, 0x03, 0x78, 0xb5, 0x53, 0xb8, 0xc5, 0xea, 0x5c, 0xdb, 0x3b, 0xb8, 0x52, 0xc8, 0x49,
	0x1b, 0x95, 0x1a, 0x92, 0xe2, 0x3e, 0x8b, 0x9e, 0x38, 0x82, 0x01, 0x1a, 0xde, 0xc0, 0x65, 0xe1,
	0x8e, 0xdf, 0x67, 0x6e, 0x82, 0x79, 0xd2, 0xd6, 0xf6, 0x0e, 0x4d, 0x5c, 0x60, 0xb6, 0x6b, 0x3a,
	0x9e, 0xfa, 0x3b, 0xce, 0x51, 0x3b, 0xe7, 0x3e, 0xd5, 0xb6, 0x70, 0x73, 0xce, 0x92, 0xa6, 0x0c,
	0x97, 0xf0, 0x82, 0x32, 0x14, 0x86, 0x4c, 0x65, 0xff, 0x4e, 0xdd, 0xdf, 0x3e, 0x64, 0xcf, 0x3e,
	0x7a, 0xc9, 0xf5, 0x70, 0x4d, 0x19, 0xff, 0x0a, 0x20, 0xdc, 0x62, 0xf5, 0xf1, 0xc4, 0xf2, 0x54,
	0x70, 0x6c, 0x06, 0xbb, 0x82, 0x09, 0x35, 0x52, 0x27, 0xf4, 0x2f, 0xf2, 0xf3, 0x78, 0xa6, 0xa3,
	0x78, 0x06, 0xe2, 0x2e, 0xc6, 0xe2, 0xde, 0xc3, 0x75, 0x7d, 0x99, 0x33, 0xd7, 0x94, 0x0a, 0x9d,
	0x8d, 0xb3, 0x64, 0x84, 0xc7, 0xdf, 0xe0, 0xfa, 0x01, 0x45, 0x86, 0x6a, 0x8b, 0xd5, 0xdf, 0xb4,
	0x2e, 0xe0, 0x39, 0x19, 0xf4, 0xde, 0xd7, 0x4f, 0xa2, 0x03, 0x06, 0x51, 0x4f, 0x47, 0x51, 0xff,
	0x08, 0x60, 0xd1, 0x52, 0xdc, 0xd9, 0x10, 0x69, 0x5f, 0xda, 0x93, 0xff, 0x47, 0x77, 0x03, 0x97,
	0x2c, 0x4f, 0x49, 0x74, 0xe6, 0xb4, 0xb5, 0x95, 0x62, 0x33, 0x26, 0xc1, 0x3b, 0x67, 0x7a, 0x48,
	0xfc, 0x15, 0x16, 0x77, 0x78, 0x24, 0x86, 0x1b, 0x29, 0x34, 0x69, 0x83, 0x82, 0x55, 0x1b, 0x99,
	0x61, 0x6f, 0x33, 0x38, 0x8a, 0x86, 0xda, 0x2b, 0xea, 0x21, 0x56, 0x99, 0x6e, 0x1d, 0xf5, 0xb1,
	0x75, 0xc0, 0xed, 0x07, 0x78, 0x2b, 0x15, 0x5f, 0x7f, 0xf7, 0x6b, 0xa9, 0x2b, 0x6d, 0xf0, 0x51,
	0xaf, 0x0f, 0xb4, 0x77, 0x2d, 0xf5, 0x6e, 0x32, 0x79, 0xb8, 0x9d, 0xf9, 0x37, 0xef, 0x96, 0x56,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x32, 0x36, 0x08, 0x6d, 0xc8, 0x03, 0x00, 0x00,
}
