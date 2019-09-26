// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/criterion_category_channel_availability_mode.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enum containing the possible CriterionCategoryChannelAvailabilityMode.
type CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode int32

const (
	// Not specified.
	CriterionCategoryChannelAvailabilityModeEnum_UNSPECIFIED CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionCategoryChannelAvailabilityModeEnum_UNKNOWN CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode = 1
	// The category is available to campaigns of all channel types and subtypes.
	CriterionCategoryChannelAvailabilityModeEnum_ALL_CHANNELS CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode = 2
	// The category is available to campaigns of a specific channel type,
	// including all subtypes under it.
	CriterionCategoryChannelAvailabilityModeEnum_CHANNEL_TYPE_AND_ALL_SUBTYPES CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode = 3
	// The category is available to campaigns of a specific channel type and
	// subtype(s).
	CriterionCategoryChannelAvailabilityModeEnum_CHANNEL_TYPE_AND_SUBSET_SUBTYPES CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode = 4
)

var CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL_CHANNELS",
	3: "CHANNEL_TYPE_AND_ALL_SUBTYPES",
	4: "CHANNEL_TYPE_AND_SUBSET_SUBTYPES",
}
var CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"ALL_CHANNELS":                     2,
	"CHANNEL_TYPE_AND_ALL_SUBTYPES":    3,
	"CHANNEL_TYPE_AND_SUBSET_SUBTYPES": 4,
}

func (x CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode) String() string {
	return proto.EnumName(CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode_name, int32(x))
}
func (CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_channel_availability_mode_fa60aba1511b3de1, []int{0, 0}
}

// Describes channel availability mode for a criterion availability - whether
// the availability is meant to include all advertising channels, or a
// particular channel with all its channel subtypes, or a channel with a certain
// subset of channel subtypes.
type CriterionCategoryChannelAvailabilityModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionCategoryChannelAvailabilityModeEnum) Reset() {
	*m = CriterionCategoryChannelAvailabilityModeEnum{}
}
func (m *CriterionCategoryChannelAvailabilityModeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*CriterionCategoryChannelAvailabilityModeEnum) ProtoMessage() {}
func (*CriterionCategoryChannelAvailabilityModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_channel_availability_mode_fa60aba1511b3de1, []int{0}
}
func (m *CriterionCategoryChannelAvailabilityModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum.Unmarshal(m, b)
}
func (m *CriterionCategoryChannelAvailabilityModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum.Marshal(b, m, deterministic)
}
func (dst *CriterionCategoryChannelAvailabilityModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum.Merge(dst, src)
}
func (m *CriterionCategoryChannelAvailabilityModeEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum.Size(m)
}
func (m *CriterionCategoryChannelAvailabilityModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryChannelAvailabilityModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CriterionCategoryChannelAvailabilityModeEnum)(nil), "google.ads.googleads.v0.enums.CriterionCategoryChannelAvailabilityModeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode", CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode_name, CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/criterion_category_channel_availability_mode.proto", fileDescriptor_criterion_category_channel_availability_mode_fa60aba1511b3de1)
}

var fileDescriptor_criterion_category_channel_availability_mode_fa60aba1511b3de1 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x6f, 0xd2, 0xcb, 0xbd, 0x30, 0xbd, 0x70, 0x43, 0xf6, 0x05, 0x5b, 0x5c, 0xb8, 0xd0,
	0x49, 0xc0, 0xdd, 0xb8, 0x9a, 0xa4, 0xb1, 0x16, 0x6b, 0x0c, 0xa4, 0xa9, 0x28, 0x81, 0x61, 0x9a,
	0x0c, 0x63, 0x20, 0x99, 0x29, 0x99, 0xb4, 0xd0, 0xad, 0x8f, 0xa2, 0x3b, 0x1f, 0xc5, 0xb5, 0x4f,
	0xe1, 0x53, 0x48, 0x92, 0xb6, 0x0a, 0xa2, 0x74, 0x33, 0xfc, 0x67, 0xce, 0x3f, 0x1f, 0x73, 0xfe,
	0x03, 0x02, 0x2e, 0x25, 0xcf, 0x99, 0x45, 0x53, 0x65, 0xb5, 0xb2, 0x56, 0x2b, 0xdb, 0x62, 0x62,
	0x59, 0x28, 0x2b, 0x29, 0xb3, 0x8a, 0x95, 0x99, 0x14, 0x24, 0xa1, 0x15, 0xe3, 0xb2, 0x5c, 0x93,
	0xe4, 0x9e, 0x0a, 0xc1, 0x72, 0x42, 0x57, 0x34, 0xcb, 0xe9, 0x3c, 0xcb, 0xb3, 0x6a, 0x4d, 0x0a,
	0x99, 0x32, 0xb8, 0x28, 0x65, 0x25, 0xcd, 0x5e, 0x8b, 0x81, 0x34, 0x55, 0x70, 0x47, 0x84, 0x2b,
	0x1b, 0x36, 0xc4, 0xc1, 0xab, 0x06, 0x8e, 0xdd, 0x2d, 0xd5, 0xdd, 0x40, 0xdd, 0x96, 0x89, 0x3f,
	0x21, 0xaf, 0x64, 0xca, 0x3c, 0xb1, 0x2c, 0x06, 0x4f, 0x1a, 0x38, 0xda, 0xf7, 0x81, 0xf9, 0x1f,
	0x74, 0x23, 0x3f, 0x0c, 0x3c, 0x77, 0x7c, 0x3e, 0xf6, 0x86, 0xc6, 0x2f, 0xb3, 0x0b, 0xfe, 0x46,
	0xfe, 0xa5, 0x7f, 0x7d, 0xe3, 0x1b, 0x9a, 0x69, 0x80, 0x7f, 0x78, 0x32, 0x21, 0xee, 0x05, 0xf6,
	0x7d, 0x6f, 0x12, 0x1a, 0xba, 0xd9, 0x07, 0xbd, 0x4d, 0x45, 0xa6, 0xb7, 0x81, 0x47, 0xb0, 0x3f,
	0x24, 0xb5, 0x25, 0x8c, 0x9c, 0xba, 0x0e, 0x8d, 0x8e, 0x79, 0x08, 0x0e, 0xbe, 0x58, 0xc2, 0xc8,
	0x09, 0xbd, 0xe9, 0x87, 0xeb, 0xb7, 0xf3, 0xa0, 0x83, 0x7e, 0x22, 0x0b, 0xf8, 0xe3, 0xf0, 0xce,
	0xc9, 0xbe, 0x83, 0x04, 0x75, 0x94, 0x81, 0x76, 0xe7, 0x6c, 0x78, 0x5c, 0xe6, 0x54, 0x70, 0x28,
	0x4b, 0x6e, 0x71, 0x26, 0x9a, 0xa0, 0xb7, 0xeb, 0x5a, 0x64, 0xea, 0x9b, 0xed, 0x9d, 0x35, 0xe7,
	0xa3, 0xde, 0x19, 0x61, 0xfc, 0xac, 0xf7, 0x46, 0x2d, 0x0a, 0xa7, 0x0a, 0xb6, 0xb2, 0x56, 0x33,
	0x1b, 0xd6, 0x29, 0xab, 0x97, 0x6d, 0x3f, 0xc6, 0xa9, 0x8a, 0x77, 0xfd, 0x78, 0x66, 0xc7, 0x4d,
	0xff, 0x4d, 0xef, 0xb7, 0x97, 0x08, 0xe1, 0x54, 0x21, 0xb4, 0x73, 0x20, 0x34, 0xb3, 0x11, 0x6a,
	0x3c, 0xf3, 0x3f, 0xcd, 0xc7, 0x4e, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc6, 0xdf, 0x2c,
	0x55, 0x02, 0x00, 0x00,
}
