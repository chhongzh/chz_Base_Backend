// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: response.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HasPermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasPermission bool                   `protobuf:"varint,1,opt,name=hasPermission,proto3" json:"hasPermission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HasPermissionResponse) Reset() {
	*x = HasPermissionResponse{}
	mi := &file_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasPermissionResponse) ProtoMessage() {}

func (x *HasPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasPermissionResponse.ProtoReflect.Descriptor instead.
func (*HasPermissionResponse) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

func (x *HasPermissionResponse) GetHasPermission() bool {
	if x != nil {
		return x.HasPermission
	}
	return false
}

var File_response_proto protoreflect.FileDescriptor

const file_response_proto_rawDesc = "" +
	"\n" +
	"\x0eresponse.proto\x12\x02pb\"=\n" +
	"\x15HasPermissionResponse\x12$\n" +
	"\rhasPermission\x18\x01 \x01(\bR\rhasPermissionB-Z+github.com/chhongzh/chz_Base_Backend/pkg/pbb\x06proto3"

var (
	file_response_proto_rawDescOnce sync.Once
	file_response_proto_rawDescData []byte
)

func file_response_proto_rawDescGZIP() []byte {
	file_response_proto_rawDescOnce.Do(func() {
		file_response_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_response_proto_rawDesc), len(file_response_proto_rawDesc)))
	})
	return file_response_proto_rawDescData
}

var file_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_proto_goTypes = []any{
	(*HasPermissionResponse)(nil), // 0: pb.HasPermissionResponse
}
var file_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_proto_init() }
func file_response_proto_init() {
	if File_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_response_proto_rawDesc), len(file_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_proto_goTypes,
		DependencyIndexes: file_response_proto_depIdxs,
		MessageInfos:      file_response_proto_msgTypes,
	}.Build()
	File_response_proto = out.File
	file_response_proto_goTypes = nil
	file_response_proto_depIdxs = nil
}
