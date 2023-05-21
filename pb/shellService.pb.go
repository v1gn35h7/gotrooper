// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.5
// source: pb/shellService.proto

package pb

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

type ShellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId         string `protobuf:"bytes,1,opt,name=AgentId,proto3" json:"AgentId,omitempty"`
	Platform        string `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	OperatingSystem string `protobuf:"bytes,3,opt,name=OperatingSystem,proto3" json:"OperatingSystem,omitempty"`
	Architecture    string `protobuf:"bytes,4,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
}

func (x *ShellRequest) Reset() {
	*x = ShellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shellService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellRequest) ProtoMessage() {}

func (x *ShellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shellService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellRequest.ProtoReflect.Descriptor instead.
func (*ShellRequest) Descriptor() ([]byte, []int) {
	return file_pb_shellService_proto_rawDescGZIP(), []int{0}
}

func (x *ShellRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *ShellRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ShellRequest) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *ShellRequest) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

type ShellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scripts []*ShellScript `protobuf:"bytes,1,rep,name=Scripts,proto3" json:"Scripts,omitempty"`
}

func (x *ShellResponse) Reset() {
	*x = ShellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shellService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellResponse) ProtoMessage() {}

func (x *ShellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shellService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellResponse.ProtoReflect.Descriptor instead.
func (*ShellResponse) Descriptor() ([]byte, []int) {
	return file_pb_shellService_proto_rawDescGZIP(), []int{1}
}

func (x *ShellResponse) GetScripts() []*ShellScript {
	if x != nil {
		return x.Scripts
	}
	return nil
}

type ShellScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	Args   string `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ShellScript) Reset() {
	*x = ShellScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shellService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellScript) ProtoMessage() {}

func (x *ShellScript) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shellService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellScript.ProtoReflect.Descriptor instead.
func (*ShellScript) Descriptor() ([]byte, []int) {
	return file_pb_shellService_proto_rawDescGZIP(), []int{2}
}

func (x *ShellScript) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *ShellScript) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *ShellScript) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_pb_shellService_proto protoreflect.FileDescriptor

var file_pb_shellService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x28,
	0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x37, 0x0a, 0x0d,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x32, 0x3d, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x73, 0x12, 0x0d, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x31, 0x67, 0x6e, 0x33, 0x35, 0x68, 0x37, 0x2f, 0x67, 0x6f, 0x74, 0x72, 0x6f,
	0x6f, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_shellService_proto_rawDescOnce sync.Once
	file_pb_shellService_proto_rawDescData = file_pb_shellService_proto_rawDesc
)

func file_pb_shellService_proto_rawDescGZIP() []byte {
	file_pb_shellService_proto_rawDescOnce.Do(func() {
		file_pb_shellService_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_shellService_proto_rawDescData)
	})
	return file_pb_shellService_proto_rawDescData
}

var file_pb_shellService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_shellService_proto_goTypes = []interface{}{
	(*ShellRequest)(nil),  // 0: ShellRequest
	(*ShellResponse)(nil), // 1: ShellResponse
	(*ShellScript)(nil),   // 2: ShellScript
}
var file_pb_shellService_proto_depIdxs = []int32{
	2, // 0: ShellResponse.Scripts:type_name -> ShellScript
	0, // 1: ShellService.GetScripts:input_type -> ShellRequest
	1, // 2: ShellService.GetScripts:output_type -> ShellResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_shellService_proto_init() }
func file_pb_shellService_proto_init() {
	if File_pb_shellService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_shellService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellRequest); i {
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
		file_pb_shellService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellResponse); i {
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
		file_pb_shellService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellScript); i {
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
			RawDescriptor: file_pb_shellService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_shellService_proto_goTypes,
		DependencyIndexes: file_pb_shellService_proto_depIdxs,
		MessageInfos:      file_pb_shellService_proto_msgTypes,
	}.Build()
	File_pb_shellService_proto = out.File
	file_pb_shellService_proto_rawDesc = nil
	file_pb_shellService_proto_goTypes = nil
	file_pb_shellService_proto_depIdxs = nil
}
