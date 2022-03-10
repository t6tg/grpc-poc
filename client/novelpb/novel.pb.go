package novelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NovelIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NovelIdRequest) Reset() {
	*x = NovelIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NovelIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NovelIdRequest) ProtoMessage() {}

func (x *NovelIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_novel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NovelIdRequest.ProtoReflect.Descriptor instead.
func (*NovelIdRequest) Descriptor() ([]byte, []int) {
	return file_novel_proto_rawDescGZIP(), []int{0}
}

func (x *NovelIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NovelIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NovelIdResponse) Reset() {
	*x = NovelIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NovelIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NovelIdResponse) ProtoMessage() {}

func (x *NovelIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_novel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NovelIdResponse.ProtoReflect.Descriptor instead.
func (*NovelIdResponse) Descriptor() ([]byte, []int) {
	return file_novel_proto_rawDescGZIP(), []int{1}
}

func (x *NovelIdResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NovelIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NovelIdResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_novel_proto protoreflect.FileDescriptor

var file_novel_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e,
	0x6f, 0x76, 0x65, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0f, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x48, 0x0a, 0x05, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x6e, 0x6f, 0x76,
	0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_novel_proto_rawDescOnce sync.Once
	file_novel_proto_rawDescData = file_novel_proto_rawDesc
)

func file_novel_proto_rawDescGZIP() []byte {
	file_novel_proto_rawDescOnce.Do(func() {
		file_novel_proto_rawDescData = protoimpl.X.CompressGZIP(file_novel_proto_rawDescData)
	})
	return file_novel_proto_rawDescData
}

var file_novel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_novel_proto_goTypes = []interface{}{
	(*NovelIdRequest)(nil),  // 0: novel.NovelIdRequest
	(*NovelIdResponse)(nil), // 1: novel.NovelIdResponse
}
var file_novel_proto_depIdxs = []int32{
	0, // 0: novel.Novel.GetNovelById:input_type -> novel.NovelIdRequest
	1, // 1: novel.Novel.GetNovelById:output_type -> novel.NovelIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_novel_proto_init() }
func file_novel_proto_init() {
	if File_novel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_novel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NovelIdRequest); i {
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
		file_novel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NovelIdResponse); i {
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
			RawDescriptor: file_novel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_novel_proto_goTypes,
		DependencyIndexes: file_novel_proto_depIdxs,
		MessageInfos:      file_novel_proto_msgTypes,
	}.Build()
	File_novel_proto = out.File
	file_novel_proto_rawDesc = nil
	file_novel_proto_goTypes = nil
	file_novel_proto_depIdxs = nil
}
