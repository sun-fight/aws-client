// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: exp.proto

package exp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnumExpUpdate int32

const (
	EnumExpUpdate_UpdateNone   EnumExpUpdate = 0
	EnumExpUpdate_UpdateAdd    EnumExpUpdate = 1
	EnumExpUpdate_UpdateSet    EnumExpUpdate = 2
	EnumExpUpdate_UpdateRemove EnumExpUpdate = 3
	EnumExpUpdate_UpdateDelete EnumExpUpdate = 4
)

// Enum value maps for EnumExpUpdate.
var (
	EnumExpUpdate_name = map[int32]string{
		0: "UpdateNone",
		1: "UpdateAdd",
		2: "UpdateSet",
		3: "UpdateRemove",
		4: "UpdateDelete",
	}
	EnumExpUpdate_value = map[string]int32{
		"UpdateNone":   0,
		"UpdateAdd":    1,
		"UpdateSet":    2,
		"UpdateRemove": 3,
		"UpdateDelete": 4,
	}
)

func (x EnumExpUpdate) Enum() *EnumExpUpdate {
	p := new(EnumExpUpdate)
	*p = x
	return p
}

func (x EnumExpUpdate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumExpUpdate) Descriptor() protoreflect.EnumDescriptor {
	return file_exp_proto_enumTypes[0].Descriptor()
}

func (EnumExpUpdate) Type() protoreflect.EnumType {
	return &file_exp_proto_enumTypes[0]
}

func (x EnumExpUpdate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumExpUpdate.Descriptor instead.
func (EnumExpUpdate) EnumDescriptor() ([]byte, []int) {
	return file_exp_proto_rawDescGZIP(), []int{0}
}

// setValueMode https://github.com/aws/aws-sdk-go-v2/blob/main/feature/dynamodb/expression/operand.go
// set value use special way
type EnumExpUpdateSet int32

const (
	EnumExpUpdateSet_UpdateSetNone        EnumExpUpdateSet = 0
	EnumExpUpdateSet_UpdateSetPlus        EnumExpUpdateSet = 1
	EnumExpUpdateSet_UpdateSetMinus       EnumExpUpdateSet = 2
	EnumExpUpdateSet_UpdateSetListAppend  EnumExpUpdateSet = 3
	EnumExpUpdateSet_UpdateSetIfNotExists EnumExpUpdateSet = 4
)

// Enum value maps for EnumExpUpdateSet.
var (
	EnumExpUpdateSet_name = map[int32]string{
		0: "UpdateSetNone",
		1: "UpdateSetPlus",
		2: "UpdateSetMinus",
		3: "UpdateSetListAppend",
		4: "UpdateSetIfNotExists",
	}
	EnumExpUpdateSet_value = map[string]int32{
		"UpdateSetNone":        0,
		"UpdateSetPlus":        1,
		"UpdateSetMinus":       2,
		"UpdateSetListAppend":  3,
		"UpdateSetIfNotExists": 4,
	}
)

func (x EnumExpUpdateSet) Enum() *EnumExpUpdateSet {
	p := new(EnumExpUpdateSet)
	*p = x
	return p
}

func (x EnumExpUpdateSet) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumExpUpdateSet) Descriptor() protoreflect.EnumDescriptor {
	return file_exp_proto_enumTypes[1].Descriptor()
}

func (EnumExpUpdateSet) Type() protoreflect.EnumType {
	return &file_exp_proto_enumTypes[1]
}

func (x EnumExpUpdateSet) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumExpUpdateSet.Descriptor instead.
func (EnumExpUpdateSet) EnumDescriptor() ([]byte, []int) {
	return file_exp_proto_rawDescGZIP(), []int{1}
}

type ExpUpdateSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string           `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	EnumExpUpdateSet EnumExpUpdateSet `protobuf:"varint,2,opt,name=EnumExpUpdateSet,proto3,enum=pb.EnumExpUpdateSet" json:"EnumExpUpdateSet,omitempty"`
}

func (x *ExpUpdateSet) Reset() {
	*x = ExpUpdateSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpUpdateSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpUpdateSet) ProtoMessage() {}

func (x *ExpUpdateSet) ProtoReflect() protoreflect.Message {
	mi := &file_exp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpUpdateSet.ProtoReflect.Descriptor instead.
func (*ExpUpdateSet) Descriptor() ([]byte, []int) {
	return file_exp_proto_rawDescGZIP(), []int{0}
}

func (x *ExpUpdateSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExpUpdateSet) GetEnumExpUpdateSet() EnumExpUpdateSet {
	if x != nil {
		return x.EnumExpUpdateSet
	}
	return EnumExpUpdateSet_UpdateSetNone
}

type ExpUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnumExpUpdate EnumExpUpdate   `protobuf:"varint,1,opt,name=EnumExpUpdate,proto3,enum=pb.EnumExpUpdate" json:"EnumExpUpdate,omitempty"`
	ExpUpdateSets []*ExpUpdateSet `protobuf:"bytes,2,rep,name=ExpUpdateSets,proto3" json:"ExpUpdateSets,omitempty"`
}

func (x *ExpUpdate) Reset() {
	*x = ExpUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpUpdate) ProtoMessage() {}

func (x *ExpUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_exp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpUpdate.ProtoReflect.Descriptor instead.
func (*ExpUpdate) Descriptor() ([]byte, []int) {
	return file_exp_proto_rawDescGZIP(), []int{1}
}

func (x *ExpUpdate) GetEnumExpUpdate() EnumExpUpdate {
	if x != nil {
		return x.EnumExpUpdate
	}
	return EnumExpUpdate_UpdateNone
}

func (x *ExpUpdate) GetExpUpdateSets() []*ExpUpdateSet {
	if x != nil {
		return x.ExpUpdateSets
	}
	return nil
}

var File_exp_proto protoreflect.FileDescriptor

var file_exp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x64, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x22, 0x7c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x45, 0x6e,
	0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x45,
	0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x52, 0x0d, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x74, 0x73, 0x2a, 0x61, 0x0a, 0x0d, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x10, 0x04, 0x2a, 0x7f, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x73, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x4d, 0x69, 0x6e,
	0x75, 0x73, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x18, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x04, 0x42, 0x0f, 0x5a, 0x0d, 0x6d, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x64, 0x62, 0x2f, 0x65, 0x78, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exp_proto_rawDescOnce sync.Once
	file_exp_proto_rawDescData = file_exp_proto_rawDesc
)

func file_exp_proto_rawDescGZIP() []byte {
	file_exp_proto_rawDescOnce.Do(func() {
		file_exp_proto_rawDescData = protoimpl.X.CompressGZIP(file_exp_proto_rawDescData)
	})
	return file_exp_proto_rawDescData
}

var file_exp_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_exp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exp_proto_goTypes = []interface{}{
	(EnumExpUpdate)(0),    // 0: pb.EnumExpUpdate
	(EnumExpUpdateSet)(0), // 1: pb.EnumExpUpdateSet
	(*ExpUpdateSet)(nil),  // 2: pb.ExpUpdateSet
	(*ExpUpdate)(nil),     // 3: pb.ExpUpdate
}
var file_exp_proto_depIdxs = []int32{
	1, // 0: pb.ExpUpdateSet.EnumExpUpdateSet:type_name -> pb.EnumExpUpdateSet
	0, // 1: pb.ExpUpdate.EnumExpUpdate:type_name -> pb.EnumExpUpdate
	2, // 2: pb.ExpUpdate.ExpUpdateSets:type_name -> pb.ExpUpdateSet
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_exp_proto_init() }
func file_exp_proto_init() {
	if File_exp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpUpdateSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exp_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exp_proto_goTypes,
		DependencyIndexes: file_exp_proto_depIdxs,
		EnumInfos:         file_exp_proto_enumTypes,
		MessageInfos:      file_exp_proto_msgTypes,
	}.Build()
	File_exp_proto = out.File
	file_exp_proto_rawDesc = nil
	file_exp_proto_goTypes = nil
	file_exp_proto_depIdxs = nil
}
