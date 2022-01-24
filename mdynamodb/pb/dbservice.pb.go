// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: dbservice.proto

package pb

import (
	context "context"
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

type ReqGetUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *ReqGetUserInfo) Reset() {
	*x = ReqGetUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetUserInfo) ProtoMessage() {}

func (x *ReqGetUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dbservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetUserInfo.ProtoReflect.Descriptor instead.
func (*ReqGetUserInfo) Descriptor() ([]byte, []int) {
	return file_dbservice_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGetUserInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Grid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X  int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y  int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	ID int32 `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Grid) Reset() {
	*x = Grid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grid) ProtoMessage() {}

func (x *Grid) ProtoReflect() protoreflect.Message {
	mi := &file_dbservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grid.ProtoReflect.Descriptor instead.
func (*Grid) Descriptor() ([]byte, []int) {
	return file_dbservice_proto_rawDescGZIP(), []int{1}
}

func (x *Grid) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Grid) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Grid) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int64            `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CreatedAt   int32            `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	DeletedAt   int32            `protobuf:"varint,4,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	LastLoginAt int32            `protobuf:"varint,5,opt,name=LastLoginAt,proto3" json:"LastLoginAt,omitempty"`
	Plunder     map[string]int32 `protobuf:"bytes,6,rep,name=Plunder,proto3" json:"Plunder,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Maps        map[string]*Grid `protobuf:"bytes,7,rep,name=Maps,proto3" json:"Maps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dbservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_dbservice_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserInfo) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserInfo) GetDeletedAt() int32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *UserInfo) GetLastLoginAt() int32 {
	if x != nil {
		return x.LastLoginAt
	}
	return 0
}

func (x *UserInfo) GetPlunder() map[string]int32 {
	if x != nil {
		return x.Plunder
	}
	return nil
}

func (x *UserInfo) GetMaps() map[string]*Grid {
	if x != nil {
		return x.Maps
	}
	return nil
}

type ResGetUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
}

func (x *ResGetUserInfo) Reset() {
	*x = ResGetUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetUserInfo) ProtoMessage() {}

func (x *ResGetUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dbservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetUserInfo.ProtoReflect.Descriptor instead.
func (*ResGetUserInfo) Descriptor() ([]byte, []int) {
	return file_dbservice_proto_rawDescGZIP(), []int{3}
}

func (x *ResGetUserInfo) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_dbservice_proto protoreflect.FileDescriptor

var file_dbservice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x32, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x59, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x22, 0xe0, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x50, 0x6c, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x50, 0x6c, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x4d,
	0x61, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x4d, 0x61, 0x70, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x6e, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x69, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0x44, 0x0a, 0x09, 0x44, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbservice_proto_rawDescOnce sync.Once
	file_dbservice_proto_rawDescData = file_dbservice_proto_rawDesc
)

func file_dbservice_proto_rawDescGZIP() []byte {
	file_dbservice_proto_rawDescOnce.Do(func() {
		file_dbservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbservice_proto_rawDescData)
	})
	return file_dbservice_proto_rawDescData
}

var file_dbservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dbservice_proto_goTypes = []interface{}{
	(*ReqGetUserInfo)(nil), // 0: pb.ReqGetUserInfo
	(*Grid)(nil),           // 1: pb.Grid
	(*UserInfo)(nil),       // 2: pb.UserInfo
	(*ResGetUserInfo)(nil), // 3: pb.ResGetUserInfo
	nil,                    // 4: pb.UserInfo.PlunderEntry
	nil,                    // 5: pb.UserInfo.MapsEntry
}
var file_dbservice_proto_depIdxs = []int32{
	4, // 0: pb.UserInfo.Plunder:type_name -> pb.UserInfo.PlunderEntry
	5, // 1: pb.UserInfo.Maps:type_name -> pb.UserInfo.MapsEntry
	2, // 2: pb.ResGetUserInfo.UserInfo:type_name -> pb.UserInfo
	1, // 3: pb.UserInfo.MapsEntry.value:type_name -> pb.Grid
	0, // 4: pb.DbService.GetUserInfo:input_type -> pb.ReqGetUserInfo
	3, // 5: pb.DbService.GetUserInfo:output_type -> pb.ResGetUserInfo
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dbservice_proto_init() }
func file_dbservice_proto_init() {
	if File_dbservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetUserInfo); i {
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
		file_dbservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grid); i {
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
		file_dbservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_dbservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetUserInfo); i {
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
			RawDescriptor: file_dbservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbservice_proto_goTypes,
		DependencyIndexes: file_dbservice_proto_depIdxs,
		MessageInfos:      file_dbservice_proto_msgTypes,
	}.Build()
	File_dbservice_proto = out.File
	file_dbservice_proto_rawDesc = nil
	file_dbservice_proto_goTypes = nil
	file_dbservice_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.1.2. DO NOT EDIT.

type DbService interface {
	GetUserInfo(ctx context.Context, req *ReqGetUserInfo) (res *ResGetUserInfo, err error)
}
