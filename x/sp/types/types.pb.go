// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Status is the status of a storage provider.
type Status int32

const (
	STATUS_IN_SERVICE       Status = 0
	STATUS_IN_JAILED        Status = 1
	STATUS_GRACEFUL_EXITING Status = 2
	STATUS_OUT_OF_SERVICE   Status = 3
)

var Status_name = map[int32]string{
	0: "STATUS_IN_SERVICE",
	1: "STATUS_IN_JAILED",
	2: "STATUS_GRACEFUL_EXITING",
	3: "STATUS_OUT_OF_SERVICE",
}

var Status_value = map[string]int32{
	"STATUS_IN_SERVICE":       0,
	"STATUS_IN_JAILED":        1,
	"STATUS_GRACEFUL_EXITING": 2,
	"STATUS_OUT_OF_SERVICE":   3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a9af9b5be8c2eeb, []int{0}
}

// Description defines a storage provider description.
type Description struct {
	// moniker defines a human-readable name for the storage provider
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// identity defines an optional identity signature (ex. UPort or Keybase).
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	// website defines an optional website link.
	Website string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	// security_contact defines an optional email for security contact.
	SecurityContact string `protobuf:"bytes,4,opt,name=security_contact,json=securityContact,proto3" json:"security_contact,omitempty"`
	// details define other optional details.
	Details string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *Description) Reset()      { *m = Description{} }
func (*Description) ProtoMessage() {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a9af9b5be8c2eeb, []int{0}
}
func (m *Description) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Description.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(m, src)
}
func (m *Description) XXX_Size() int {
	return m.Size()
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

func (m *Description) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Description) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Description) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Description) GetSecurityContact() string {
	if m != nil {
		return m.SecurityContact
	}
	return ""
}

func (m *Description) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

// StorageProvider defines the meta info of storage provider
// TODO: add endpoint for RPC/HTTP/Websocket and p2p identity
type StorageProvider struct {
	// operator_address defines the address of the sp's operator; It also is the unqiue index key of sp.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// fund_address is the account address of the storage provider for deposit, remuneration.
	FundingAddress string `protobuf:"bytes,2,opt,name=funding_address,json=fundingAddress,proto3" json:"funding_address,omitempty"`
	// seal_address is the account address of the storage provider for sealObject
	SealAddress string `protobuf:"bytes,3,opt,name=seal_address,json=sealAddress,proto3" json:"seal_address,omitempty"`
	// approval_address is the account address of the storage provider for ack CreateBuclet/Object.
	ApprovalAddress string `protobuf:"bytes,4,opt,name=approval_address,json=approvalAddress,proto3" json:"approval_address,omitempty"`
	// total_deposit define the deposit token
	TotalDeposit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_deposit,json=totalDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_deposit"`
	// status is the status of sp, which can be (in_service/read_only_service/graceful_exiting/out_of_service)
	Status Status `protobuf:"varint,6,opt,name=status,proto3,enum=bnbchain.greenfield.sp.Status" json:"status,omitempty"`
	// description defines the description terms for the validator.
	Description Description `protobuf:"bytes,7,opt,name=description,proto3" json:"description"`
}

func (m *StorageProvider) Reset()      { *m = StorageProvider{} }
func (*StorageProvider) ProtoMessage() {}
func (*StorageProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a9af9b5be8c2eeb, []int{1}
}
func (m *StorageProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StorageProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StorageProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StorageProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageProvider.Merge(m, src)
}
func (m *StorageProvider) XXX_Size() int {
	return m.Size()
}
func (m *StorageProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageProvider.DiscardUnknown(m)
}

var xxx_messageInfo_StorageProvider proto.InternalMessageInfo

func (m *StorageProvider) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *StorageProvider) GetFundingAddress() string {
	if m != nil {
		return m.FundingAddress
	}
	return ""
}

func (m *StorageProvider) GetSealAddress() string {
	if m != nil {
		return m.SealAddress
	}
	return ""
}

func (m *StorageProvider) GetApprovalAddress() string {
	if m != nil {
		return m.ApprovalAddress
	}
	return ""
}

func (m *StorageProvider) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return STATUS_IN_SERVICE
}

func (m *StorageProvider) GetDescription() Description {
	if m != nil {
		return m.Description
	}
	return Description{}
}

type RewardInfo struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount  types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *RewardInfo) Reset()         { *m = RewardInfo{} }
func (m *RewardInfo) String() string { return proto.CompactTextString(m) }
func (*RewardInfo) ProtoMessage()    {}
func (*RewardInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a9af9b5be8c2eeb, []int{2}
}
func (m *RewardInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardInfo.Merge(m, src)
}
func (m *RewardInfo) XXX_Size() int {
	return m.Size()
}
func (m *RewardInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RewardInfo proto.InternalMessageInfo

func (m *RewardInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RewardInfo) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.sp.Status", Status_name, Status_value)
	proto.RegisterType((*Description)(nil), "bnbchain.greenfield.sp.Description")
	proto.RegisterType((*StorageProvider)(nil), "bnbchain.greenfield.sp.StorageProvider")
	proto.RegisterType((*RewardInfo)(nil), "bnbchain.greenfield.sp.RewardInfo")
}

func init() { proto.RegisterFile("greenfield/sp/types.proto", fileDescriptor_7a9af9b5be8c2eeb) }

var fileDescriptor_7a9af9b5be8c2eeb = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xde, 0x85, 0xfe, 0x0a, 0xbf, 0x29, 0xd2, 0x3a, 0x01, 0xdd, 0xd6, 0x64, 0x4b, 0x30, 0x31,
	0x48, 0xd2, 0xdd, 0x50, 0x13, 0x4d, 0xd4, 0x4b, 0xff, 0x41, 0x56, 0x09, 0x98, 0x6d, 0x31, 0xc6,
	0xcb, 0x66, 0x76, 0x77, 0x58, 0x26, 0xb4, 0x33, 0x9b, 0x99, 0x29, 0xd8, 0x6f, 0xe0, 0xc5, 0xc4,
	0xa3, 0x47, 0x12, 0xaf, 0x1e, 0xf9, 0x10, 0x1c, 0x09, 0x27, 0xe3, 0x81, 0x18, 0xb8, 0xf8, 0x31,
	0x4c, 0x77, 0x67, 0x4b, 0x0f, 0x1a, 0x3c, 0xb5, 0xef, 0xf3, 0x3e, 0xcf, 0xb3, 0xf3, 0xe7, 0x79,
	0x07, 0x94, 0x23, 0x8e, 0x31, 0xdd, 0x27, 0xb8, 0x1f, 0xda, 0x22, 0xb6, 0xe5, 0x28, 0xc6, 0xc2,
	0x8a, 0x39, 0x93, 0x0c, 0xde, 0xf3, 0xa9, 0x1f, 0x1c, 0x20, 0x42, 0xad, 0x1b, 0x8e, 0x25, 0xe2,
	0x8a, 0x19, 0x30, 0x31, 0x60, 0xc2, 0xf6, 0x91, 0xc0, 0xf6, 0xd1, 0x86, 0x8f, 0x25, 0xda, 0xb0,
	0x03, 0x46, 0x68, 0xaa, 0xab, 0x94, 0xd3, 0xbe, 0x97, 0x54, 0x76, 0x5a, 0xa8, 0xd6, 0x52, 0xc4,
	0x22, 0x96, 0xe2, 0xe3, 0x7f, 0x29, 0xba, 0xfa, 0x4d, 0x07, 0x85, 0x36, 0x16, 0x01, 0x27, 0xb1,
	0x24, 0x8c, 0x42, 0x03, 0xcc, 0x0d, 0x18, 0x25, 0x87, 0x98, 0x1b, 0xfa, 0x8a, 0xbe, 0xf6, 0xbf,
	0x9b, 0x95, 0xb0, 0x02, 0xe6, 0x49, 0x88, 0xa9, 0x24, 0x72, 0x64, 0xcc, 0x24, 0xad, 0x49, 0x3d,
	0x56, 0x1d, 0x63, 0x5f, 0x10, 0x89, 0x8d, 0xd9, 0x54, 0xa5, 0x4a, 0xf8, 0x18, 0x94, 0x04, 0x0e,
	0x86, 0x9c, 0xc8, 0x91, 0x17, 0x30, 0x2a, 0x51, 0x20, 0x8d, 0x5c, 0x42, 0x29, 0x66, 0x78, 0x2b,
	0x85, 0xc7, 0x26, 0x21, 0x96, 0x88, 0xf4, 0x85, 0xf1, 0x5f, 0x6a, 0xa2, 0xca, 0xe7, 0xf3, 0x5f,
	0x4e, 0xaa, 0xda, 0xaf, 0x93, 0xaa, 0xbe, 0xfa, 0x29, 0x07, 0x8a, 0x5d, 0xc9, 0x38, 0x8a, 0xf0,
	0x1b, 0xce, 0x8e, 0x48, 0x88, 0x39, 0x6c, 0x81, 0x12, 0x8b, 0x31, 0x47, 0x92, 0x71, 0x0f, 0x85,
	0x21, 0xc7, 0x42, 0xa4, 0x6b, 0x6f, 0x1a, 0x17, 0xa7, 0xb5, 0x25, 0x75, 0x08, 0x8d, 0xb4, 0xd3,
	0x95, 0x9c, 0xd0, 0xc8, 0x2d, 0x66, 0x0a, 0x05, 0xc3, 0x06, 0x28, 0xee, 0x0f, 0x69, 0x48, 0x68,
	0x34, 0xf1, 0x98, 0xb9, 0xc5, 0x63, 0x51, 0x09, 0x32, 0x8b, 0x17, 0x60, 0x41, 0x60, 0xd4, 0x9f,
	0xe8, 0x67, 0x6f, 0xd1, 0x17, 0xc6, 0xec, 0x4c, 0xdc, 0x02, 0x25, 0x14, 0xc7, 0x9c, 0x1d, 0x4d,
	0x19, 0xe4, 0x6e, 0xdb, 0x44, 0xa6, 0xc8, 0x4c, 0x10, 0xb8, 0x23, 0x99, 0x44, 0x7d, 0x2f, 0xc4,
	0x31, 0x13, 0x44, 0xa6, 0xe7, 0xd8, 0x7c, 0x79, 0x76, 0x59, 0xd5, 0x7e, 0x5c, 0x56, 0x1f, 0x45,
	0x44, 0x1e, 0x0c, 0x7d, 0x2b, 0x60, 0x03, 0x15, 0x0d, 0xf5, 0x53, 0x13, 0xe1, 0xa1, 0x8a, 0x9f,
	0x43, 0xe5, 0xc5, 0x69, 0x0d, 0xa8, 0xef, 0x39, 0x54, 0xba, 0x0b, 0x89, 0x65, 0x3b, 0x75, 0x84,
	0x4f, 0x41, 0x5e, 0x48, 0x24, 0x87, 0xc2, 0xc8, 0xaf, 0xe8, 0x6b, 0x8b, 0x75, 0xd3, 0xfa, 0x73,
	0x52, 0xad, 0x6e, 0xc2, 0x72, 0x15, 0x1b, 0xbe, 0x06, 0x85, 0xf0, 0x26, 0x66, 0xc6, 0xdc, 0x8a,
	0xbe, 0x56, 0xa8, 0x3f, 0xfc, 0x9b, 0x78, 0x2a, 0x91, 0xcd, 0xdc, 0x78, 0xf5, 0xee, 0xb4, 0x7a,
	0x92, 0x07, 0x6d, 0x75, 0x04, 0x80, 0x8b, 0x8f, 0x11, 0x0f, 0x1d, 0xba, 0xcf, 0x60, 0x1d, 0xcc,
	0xfd, 0x6b, 0x00, 0x32, 0x22, 0x7c, 0x06, 0xf2, 0x68, 0xc0, 0x86, 0x54, 0x26, 0xf7, 0x5d, 0xa8,
	0x97, 0x2d, 0xc5, 0x1f, 0x8f, 0x98, 0xa5, 0x46, 0xcc, 0x6a, 0x31, 0x92, 0xad, 0x44, 0xd1, 0xd7,
	0x05, 0xc8, 0xa7, 0x7b, 0x84, 0xcb, 0xe0, 0x6e, 0xb7, 0xd7, 0xe8, 0xed, 0x75, 0x3d, 0x67, 0xc7,
	0xeb, 0x76, 0xdc, 0xb7, 0x4e, 0xab, 0x53, 0xd2, 0xe0, 0x12, 0x28, 0xdd, 0xc0, 0xaf, 0x1a, 0xce,
	0x76, 0xa7, 0x5d, 0xd2, 0xe1, 0x03, 0x70, 0x5f, 0xa1, 0x5b, 0x6e, 0xa3, 0xd5, 0xd9, 0xdc, 0xdb,
	0xf6, 0x3a, 0xef, 0x9c, 0x9e, 0xb3, 0xb3, 0x55, 0x9a, 0x81, 0x65, 0xb0, 0xac, 0x9a, 0xbb, 0x7b,
	0x3d, 0x6f, 0x77, 0x73, 0xe2, 0x36, 0x5b, 0xc9, 0x7d, 0xfc, 0x6a, 0x6a, 0xcd, 0xf6, 0xd9, 0x95,
	0xa9, 0x9f, 0x5f, 0x99, 0xfa, 0xcf, 0x2b, 0x53, 0xff, 0x7c, 0x6d, 0x6a, 0xe7, 0xd7, 0xa6, 0xf6,
	0xfd, 0xda, 0xd4, 0xde, 0xaf, 0x4f, 0x5d, 0xae, 0x4f, 0xfd, 0x5a, 0x72, 0xac, 0xf6, 0xd4, 0x0b,
	0xf3, 0x61, 0xf2, 0xc6, 0xf8, 0xf9, 0x64, 0xf6, 0x9f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x88,
	0xf2, 0x5b, 0x73, 0x81, 0x04, 0x00, 0x00,
}

func (this *Description) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Description)
	if !ok {
		that2, ok := that.(Description)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Identity != that1.Identity {
		return false
	}
	if this.Website != that1.Website {
		return false
	}
	if this.SecurityContact != that1.SecurityContact {
		return false
	}
	if this.Details != that1.Details {
		return false
	}
	return true
}
func (m *Description) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Description) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Description) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StorageProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StorageProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.TotalDeposit.Size()
		i -= size
		if _, err := m.TotalDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ApprovalAddress) > 0 {
		i -= len(m.ApprovalAddress)
		copy(dAtA[i:], m.ApprovalAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ApprovalAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SealAddress) > 0 {
		i -= len(m.SealAddress)
		copy(dAtA[i:], m.SealAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SealAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FundingAddress) > 0 {
		i -= len(m.FundingAddress)
		copy(dAtA[i:], m.FundingAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FundingAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RewardInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Description) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *StorageProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.FundingAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SealAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ApprovalAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.TotalDeposit.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	l = m.Description.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *RewardInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Description) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Description: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Description: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *StorageProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StorageProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FundingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SealAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SealAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RewardInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RewardInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)