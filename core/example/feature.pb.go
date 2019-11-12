// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feature.proto

package example // import "github.com/galeone/tfgo/core/example"

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

// Containers to hold repeated fundamental values.
type BytesList struct {
	Value                [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesList) Reset()         { *m = BytesList{} }
func (m *BytesList) String() string { return proto.CompactTextString(m) }
func (*BytesList) ProtoMessage()    {}
func (*BytesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{0}
}
func (m *BytesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesList.Unmarshal(m, b)
}
func (m *BytesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesList.Marshal(b, m, deterministic)
}
func (dst *BytesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesList.Merge(dst, src)
}
func (m *BytesList) XXX_Size() int {
	return xxx_messageInfo_BytesList.Size(m)
}
func (m *BytesList) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesList.DiscardUnknown(m)
}

var xxx_messageInfo_BytesList proto.InternalMessageInfo

func (m *BytesList) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FloatList struct {
	Value                []float32 `protobuf:"fixed32,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FloatList) Reset()         { *m = FloatList{} }
func (m *FloatList) String() string { return proto.CompactTextString(m) }
func (*FloatList) ProtoMessage()    {}
func (*FloatList) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{1}
}
func (m *FloatList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatList.Unmarshal(m, b)
}
func (m *FloatList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatList.Marshal(b, m, deterministic)
}
func (dst *FloatList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatList.Merge(dst, src)
}
func (m *FloatList) XXX_Size() int {
	return xxx_messageInfo_FloatList.Size(m)
}
func (m *FloatList) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatList.DiscardUnknown(m)
}

var xxx_messageInfo_FloatList proto.InternalMessageInfo

func (m *FloatList) GetValue() []float32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Int64List struct {
	Value                []int64  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64List) Reset()         { *m = Int64List{} }
func (m *Int64List) String() string { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()    {}
func (*Int64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{2}
}
func (m *Int64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64List.Unmarshal(m, b)
}
func (m *Int64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64List.Marshal(b, m, deterministic)
}
func (dst *Int64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64List.Merge(dst, src)
}
func (m *Int64List) XXX_Size() int {
	return xxx_messageInfo_Int64List.Size(m)
}
func (m *Int64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64List.DiscardUnknown(m)
}

var xxx_messageInfo_Int64List proto.InternalMessageInfo

func (m *Int64List) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

// Containers for non-sequential data.
type Feature struct {
	// Each feature can be exactly one kind.
	//
	// Types that are valid to be assigned to Kind:
	//	*Feature_BytesList
	//	*Feature_FloatList
	//	*Feature_Int64List
	Kind                 isFeature_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{3}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

type isFeature_Kind interface {
	isFeature_Kind()
}

type Feature_BytesList struct {
	BytesList *BytesList `protobuf:"bytes,1,opt,name=bytes_list,json=bytesList,proto3,oneof"`
}
type Feature_FloatList struct {
	FloatList *FloatList `protobuf:"bytes,2,opt,name=float_list,json=floatList,proto3,oneof"`
}
type Feature_Int64List struct {
	Int64List *Int64List `protobuf:"bytes,3,opt,name=int64_list,json=int64List,proto3,oneof"`
}

func (*Feature_BytesList) isFeature_Kind() {}
func (*Feature_FloatList) isFeature_Kind() {}
func (*Feature_Int64List) isFeature_Kind() {}

func (m *Feature) GetKind() isFeature_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Feature) GetBytesList() *BytesList {
	if x, ok := m.GetKind().(*Feature_BytesList); ok {
		return x.BytesList
	}
	return nil
}

func (m *Feature) GetFloatList() *FloatList {
	if x, ok := m.GetKind().(*Feature_FloatList); ok {
		return x.FloatList
	}
	return nil
}

func (m *Feature) GetInt64List() *Int64List {
	if x, ok := m.GetKind().(*Feature_Int64List); ok {
		return x.Int64List
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Feature) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Feature_OneofMarshaler, _Feature_OneofUnmarshaler, _Feature_OneofSizer, []interface{}{
		(*Feature_BytesList)(nil),
		(*Feature_FloatList)(nil),
		(*Feature_Int64List)(nil),
	}
}

func _Feature_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Feature)
	// kind
	switch x := m.Kind.(type) {
	case *Feature_BytesList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BytesList); err != nil {
			return err
		}
	case *Feature_FloatList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FloatList); err != nil {
			return err
		}
	case *Feature_Int64List:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int64List); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Feature.Kind has unexpected type %T", x)
	}
	return nil
}

func _Feature_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Feature)
	switch tag {
	case 1: // kind.bytes_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BytesList)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_BytesList{msg}
		return true, err
	case 2: // kind.float_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FloatList)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_FloatList{msg}
		return true, err
	case 3: // kind.int64_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int64List)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_Int64List{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Feature_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Feature)
	// kind
	switch x := m.Kind.(type) {
	case *Feature_BytesList:
		s := proto.Size(x.BytesList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Feature_FloatList:
		s := proto.Size(x.FloatList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Feature_Int64List:
		s := proto.Size(x.Int64List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Features struct {
	// Map from feature name to feature.
	Feature              map[string]*Feature `protobuf:"bytes,1,rep,name=feature,proto3" json:"feature,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Features) Reset()         { *m = Features{} }
func (m *Features) String() string { return proto.CompactTextString(m) }
func (*Features) ProtoMessage()    {}
func (*Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{4}
}
func (m *Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Features.Unmarshal(m, b)
}
func (m *Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Features.Marshal(b, m, deterministic)
}
func (dst *Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Features.Merge(dst, src)
}
func (m *Features) XXX_Size() int {
	return xxx_messageInfo_Features.Size(m)
}
func (m *Features) XXX_DiscardUnknown() {
	xxx_messageInfo_Features.DiscardUnknown(m)
}

var xxx_messageInfo_Features proto.InternalMessageInfo

func (m *Features) GetFeature() map[string]*Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

// Containers for sequential data.
//
// A FeatureList contains lists of Features.  These may hold zero or more
// Feature values.
//
// FeatureLists are organized into categories by name.  The FeatureLists message
// contains the mapping from name to FeatureList.
//
type FeatureList struct {
	Feature              []*Feature `protobuf:"bytes,1,rep,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FeatureList) Reset()         { *m = FeatureList{} }
func (m *FeatureList) String() string { return proto.CompactTextString(m) }
func (*FeatureList) ProtoMessage()    {}
func (*FeatureList) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{5}
}
func (m *FeatureList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureList.Unmarshal(m, b)
}
func (m *FeatureList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureList.Marshal(b, m, deterministic)
}
func (dst *FeatureList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureList.Merge(dst, src)
}
func (m *FeatureList) XXX_Size() int {
	return xxx_messageInfo_FeatureList.Size(m)
}
func (m *FeatureList) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureList.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureList proto.InternalMessageInfo

func (m *FeatureList) GetFeature() []*Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type FeatureLists struct {
	// Map from feature name to feature list.
	FeatureList          map[string]*FeatureList `protobuf:"bytes,1,rep,name=feature_list,json=featureList,proto3" json:"feature_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FeatureLists) Reset()         { *m = FeatureLists{} }
func (m *FeatureLists) String() string { return proto.CompactTextString(m) }
func (*FeatureLists) ProtoMessage()    {}
func (*FeatureLists) Descriptor() ([]byte, []int) {
	return fileDescriptor_feature_d7594a49dd86f264, []int{6}
}
func (m *FeatureLists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureLists.Unmarshal(m, b)
}
func (m *FeatureLists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureLists.Marshal(b, m, deterministic)
}
func (dst *FeatureLists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureLists.Merge(dst, src)
}
func (m *FeatureLists) XXX_Size() int {
	return xxx_messageInfo_FeatureLists.Size(m)
}
func (m *FeatureLists) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureLists.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureLists proto.InternalMessageInfo

func (m *FeatureLists) GetFeatureList() map[string]*FeatureList {
	if m != nil {
		return m.FeatureList
	}
	return nil
}

func init() {
	proto.RegisterType((*BytesList)(nil), "tensorflow.BytesList")
	proto.RegisterType((*FloatList)(nil), "tensorflow.FloatList")
	proto.RegisterType((*Int64List)(nil), "tensorflow.Int64List")
	proto.RegisterType((*Feature)(nil), "tensorflow.Feature")
	proto.RegisterType((*Features)(nil), "tensorflow.Features")
	proto.RegisterMapType((map[string]*Feature)(nil), "tensorflow.Features.FeatureEntry")
	proto.RegisterType((*FeatureList)(nil), "tensorflow.FeatureList")
	proto.RegisterType((*FeatureLists)(nil), "tensorflow.FeatureLists")
	proto.RegisterMapType((map[string]*FeatureList)(nil), "tensorflow.FeatureLists.FeatureListEntry")
}

func init() { proto.RegisterFile("feature.proto", fileDescriptor_feature_d7594a49dd86f264) }

var fileDescriptor_feature_d7594a49dd86f264 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x4e, 0xab, 0x40,
	0x14, 0x86, 0xef, 0xc0, 0xbd, 0xed, 0xe5, 0xd0, 0x26, 0x0d, 0xf7, 0xaa, 0x4d, 0x57, 0x2d, 0xd1,
	0xa4, 0x5d, 0x14, 0x92, 0x6a, 0x88, 0x51, 0x57, 0x24, 0x36, 0x9a, 0x34, 0xb1, 0x61, 0x63, 0xe2,
	0xc6, 0x40, 0x1d, 0x90, 0x94, 0x32, 0x0d, 0x4c, 0xd5, 0xbe, 0x89, 0x2f, 0xe2, 0xc2, 0x37, 0x73,
	0x69, 0x06, 0x06, 0x3a, 0x6d, 0x71, 0x37, 0x67, 0xe6, 0xff, 0xcf, 0x7c, 0xff, 0x81, 0x81, 0xa6,
	0x8f, 0x5d, 0xba, 0x4a, 0xb0, 0xb1, 0x4c, 0x08, 0x25, 0x1a, 0x50, 0x1c, 0xa7, 0x24, 0xf1, 0x23,
	0xf2, 0xaa, 0xf7, 0x40, 0xb1, 0xd7, 0x14, 0xa7, 0x93, 0x30, 0xa5, 0xda, 0x7f, 0xf8, 0xf3, 0xe2,
	0x46, 0x2b, 0xdc, 0x46, 0x5d, 0xb9, 0xdf, 0x70, 0xf2, 0x42, 0x3f, 0x01, 0x65, 0x1c, 0x11, 0x97,
	0x66, 0x92, 0xb6, 0x28, 0x91, 0x6c, 0xa9, 0x85, 0x04, 0xd9, 0x6d, 0x4c, 0xad, 0xb3, 0x7d, 0x99,
	0x2c, 0xca, 0x3e, 0x11, 0xd4, 0xc7, 0x39, 0x8e, 0x66, 0x01, 0x78, 0xec, 0xf2, 0xc7, 0x28, 0x4c,
	0x69, 0x1b, 0x75, 0x51, 0x5f, 0x1d, 0x1d, 0x18, 0x1b, 0x3a, 0xa3, 0x44, 0xbb, 0xf9, 0xe5, 0x28,
	0x5e, 0xc9, 0x69, 0x01, 0xf8, 0x8c, 0x28, 0xf7, 0x49, 0xfb, 0xbe, 0x92, 0x97, 0xf9, 0xfc, 0x12,
	0xde, 0x02, 0x08, 0x19, 0x62, 0xee, 0x93, 0xf7, 0x7d, 0x65, 0x00, 0xe6, 0x0b, 0x8b, 0xc2, 0xae,
	0xc1, 0xef, 0x79, 0x18, 0x3f, 0xe9, 0xef, 0x08, 0xfe, 0x72, 0xf6, 0x54, 0xbb, 0x84, 0x3a, 0x1f,
	0x6b, 0x16, 0x52, 0x1d, 0xf5, 0xb6, 0x08, 0xb8, 0xac, 0x58, 0x5c, 0xc7, 0x34, 0x59, 0x3b, 0x85,
	0xa3, 0x73, 0x07, 0x0d, 0xf1, 0x40, 0x6b, 0x81, 0x3c, 0xc7, 0xeb, 0x6c, 0x04, 0x8a, 0xc3, 0x96,
	0xda, 0xa0, 0x98, 0x60, 0x1e, 0xef, 0x5f, 0x45, 0x73, 0x3e, 0xd2, 0x0b, 0xe9, 0x1c, 0xe9, 0x57,
	0xa0, 0xf2, 0xdd, 0x2c, 0xe9, 0x70, 0x17, 0xae, 0xd2, 0x5f, 0x68, 0xf4, 0x0f, 0x54, 0xf2, 0x30,
	0x7b, 0xaa, 0x4d, 0xa0, 0xc1, 0xcf, 0x8a, 0x6f, 0xc3, 0x9a, 0x0c, 0x2a, 0x9a, 0x64, 0x7a, 0xb1,
	0xc8, 0x93, 0xaa, 0xfe, 0x66, 0xa7, 0x73, 0x0f, 0xad, 0x5d, 0x41, 0x45, 0xe2, 0xe1, 0x76, 0xe2,
	0xa3, 0x1f, 0x2e, 0x13, 0x52, 0xdb, 0x0e, 0x1c, 0x92, 0x24, 0x10, 0x85, 0xf8, 0xcd, 0x5d, 0x2c,
	0x23, 0x6c, 0x37, 0xb9, 0x63, 0xca, 0xfe, 0xf8, 0x74, 0x8a, 0x1e, 0x8e, 0x83, 0x90, 0x3e, 0xaf,
	0x3c, 0x63, 0x46, 0x16, 0x66, 0xe0, 0x46, 0x98, 0xc4, 0xd8, 0xa4, 0x7e, 0x40, 0xcc, 0x19, 0x49,
	0xb0, 0xc9, 0x6d, 0x5f, 0x08, 0x79, 0xb5, 0xec, 0x91, 0x9c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0x14, 0x70, 0x0e, 0x35, 0x03, 0x00, 0x00,
}