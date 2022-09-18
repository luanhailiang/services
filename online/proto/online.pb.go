// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: online.proto

package online

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

// Call Heart 心跳空包，活跃检查
type CmdHeart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty" bson:"time"`
	Plat int32 `protobuf:"varint,2,opt,name=plat,proto3" json:"plat,omitempty" bson:"plat"`
}

func (x *CmdHeart) Reset() {
	*x = CmdHeart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_online_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdHeart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdHeart) ProtoMessage() {}

func (x *CmdHeart) ProtoReflect() protoreflect.Message {
	mi := &file_online_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdHeart.ProtoReflect.Descriptor instead.
func (*CmdHeart) Descriptor() ([]byte, []int) {
	return file_online_proto_rawDescGZIP(), []int{0}
}

func (x *CmdHeart) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *CmdHeart) GetPlat() int32 {
	if x != nil {
		return x.Plat
	}
	return 0
}

// 在线玩家列表
type CmdOnlineList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space string   `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty" bson:"space"`
	List  []string `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty" bson:"list"`
}

func (x *CmdOnlineList) Reset() {
	*x = CmdOnlineList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_online_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdOnlineList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdOnlineList) ProtoMessage() {}

func (x *CmdOnlineList) ProtoReflect() protoreflect.Message {
	mi := &file_online_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdOnlineList.ProtoReflect.Descriptor instead.
func (*CmdOnlineList) Descriptor() ([]byte, []int) {
	return file_online_proto_rawDescGZIP(), []int{1}
}

func (x *CmdOnlineList) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *CmdOnlineList) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

// 在线玩家人数
type CmdOnlineCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space  uint32 `protobuf:"varint,1,opt,name=space,proto3" json:"space,omitempty" bson:"space"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty" bson:"amount"`
}

func (x *CmdOnlineCount) Reset() {
	*x = CmdOnlineCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_online_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdOnlineCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdOnlineCount) ProtoMessage() {}

func (x *CmdOnlineCount) ProtoReflect() protoreflect.Message {
	mi := &file_online_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdOnlineCount.ProtoReflect.Descriptor instead.
func (*CmdOnlineCount) Descriptor() ([]byte, []int) {
	return file_online_proto_rawDescGZIP(), []int{2}
}

func (x *CmdOnlineCount) GetSpace() uint32 {
	if x != nil {
		return x.Space
	}
	return 0
}

func (x *CmdOnlineCount) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SpaceCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space    uint32 `protobuf:"varint,1,opt,name=space,proto3" json:"space,omitempty" bson:"space"`
	IosCount uint32 `protobuf:"varint,2,opt,name=ios_count,json=iosCount,proto3" json:"ios_count,omitempty" bson:"ios_count"`
	AndCount uint32 `protobuf:"varint,3,opt,name=and_count,json=andCount,proto3" json:"and_count,omitempty" bson:"and_count"`
}

func (x *SpaceCount) Reset() {
	*x = SpaceCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_online_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceCount) ProtoMessage() {}

func (x *SpaceCount) ProtoReflect() protoreflect.Message {
	mi := &file_online_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceCount.ProtoReflect.Descriptor instead.
func (*SpaceCount) Descriptor() ([]byte, []int) {
	return file_online_proto_rawDescGZIP(), []int{3}
}

func (x *SpaceCount) GetSpace() uint32 {
	if x != nil {
		return x.Space
	}
	return 0
}

func (x *SpaceCount) GetIosCount() uint32 {
	if x != nil {
		return x.IosCount
	}
	return 0
}

func (x *SpaceCount) GetAndCount() uint32 {
	if x != nil {
		return x.AndCount
	}
	return 0
}

// 在线玩家人数
type CmdOnlineAllCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*SpaceCount `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" bson:"list"`
}

func (x *CmdOnlineAllCount) Reset() {
	*x = CmdOnlineAllCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_online_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdOnlineAllCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdOnlineAllCount) ProtoMessage() {}

func (x *CmdOnlineAllCount) ProtoReflect() protoreflect.Message {
	mi := &file_online_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdOnlineAllCount.ProtoReflect.Descriptor instead.
func (*CmdOnlineAllCount) Descriptor() ([]byte, []int) {
	return file_online_proto_rawDescGZIP(), []int{4}
}

func (x *CmdOnlineAllCount) GetList() []*SpaceCount {
	if x != nil {
		return x.List
	}
	return nil
}

var File_online_proto protoreflect.FileDescriptor

var file_online_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x43, 0x6d, 0x64, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x6d,
	0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x6d, 0x64, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6f, 0x73,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6f,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x6e, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x6d, 0x64, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_online_proto_rawDescOnce sync.Once
	file_online_proto_rawDescData = file_online_proto_rawDesc
)

func file_online_proto_rawDescGZIP() []byte {
	file_online_proto_rawDescOnce.Do(func() {
		file_online_proto_rawDescData = protoimpl.X.CompressGZIP(file_online_proto_rawDescData)
	})
	return file_online_proto_rawDescData
}

var file_online_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_online_proto_goTypes = []interface{}{
	(*CmdHeart)(nil),          // 0: online.CmdHeart
	(*CmdOnlineList)(nil),     // 1: online.CmdOnlineList
	(*CmdOnlineCount)(nil),    // 2: online.CmdOnlineCount
	(*SpaceCount)(nil),        // 3: online.SpaceCount
	(*CmdOnlineAllCount)(nil), // 4: online.CmdOnlineAllCount
}
var file_online_proto_depIdxs = []int32{
	3, // 0: online.CmdOnlineAllCount.list:type_name -> online.SpaceCount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_online_proto_init() }
func file_online_proto_init() {
	if File_online_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_online_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdHeart); i {
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
		file_online_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdOnlineList); i {
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
		file_online_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdOnlineCount); i {
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
		file_online_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceCount); i {
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
		file_online_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdOnlineAllCount); i {
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
			RawDescriptor: file_online_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_online_proto_goTypes,
		DependencyIndexes: file_online_proto_depIdxs,
		MessageInfos:      file_online_proto_msgTypes,
	}.Build()
	File_online_proto = out.File
	file_online_proto_rawDesc = nil
	file_online_proto_goTypes = nil
	file_online_proto_depIdxs = nil
}
