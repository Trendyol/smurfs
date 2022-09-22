// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: protos/backend.proto

package protos

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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Usage       string         `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
	Example     string         `protobuf:"bytes,4,opt,name=example,proto3" json:"example,omitempty"`
	Aliases     []string       `protobuf:"bytes,5,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Flags       []*CommandFlag `protobuf:"bytes,6,rep,name=flags,proto3" json:"flags,omitempty"`
	PluginRef   string         `protobuf:"bytes,7,opt,name=pluginRef,proto3" json:"pluginRef,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_protos_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Command) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Command) GetUsage() string {
	if x != nil {
		return x.Usage
	}
	return ""
}

func (x *Command) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *Command) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Command) GetFlags() []*CommandFlag {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Command) GetPluginRef() string {
	if x != nil {
		return x.PluginRef
	}
	return ""
}

type CommandFlag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DefaultValue string `protobuf:"bytes,3,opt,name=defaultValue,proto3" json:"defaultValue,omitempty"`
	Required     bool   `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Repeated     bool   `protobuf:"varint,5,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Short        string `protobuf:"bytes,6,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *CommandFlag) Reset() {
	*x = CommandFlag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandFlag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandFlag) ProtoMessage() {}

func (x *CommandFlag) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandFlag.ProtoReflect.Descriptor instead.
func (*CommandFlag) Descriptor() ([]byte, []int) {
	return file_protos_backend_proto_rawDescGZIP(), []int{1}
}

func (x *CommandFlag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandFlag) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommandFlag) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *CommandFlag) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *CommandFlag) GetRepeated() bool {
	if x != nil {
		return x.Repeated
	}
	return false
}

func (x *CommandFlag) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Arch   string `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Os     string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_protos_backend_proto_rawDescGZIP(), []int{2}
}

func (x *CommandRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CommandRequest) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *CommandRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_protos_backend_proto_rawDescGZIP(), []int{3}
}

func (x *CommandResponse) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_protos_backend_proto protoreflect.FileDescriptor

var file_protos_backend_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x66, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x4c, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x22, 0x37, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x32, 0x44, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x79, 0x6f, 0x6c,
	0x2f, 0x73, 0x6d, 0x75, 0x72, 0x66, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_backend_proto_rawDescOnce sync.Once
	file_protos_backend_proto_rawDescData = file_protos_backend_proto_rawDesc
)

func file_protos_backend_proto_rawDescGZIP() []byte {
	file_protos_backend_proto_rawDescOnce.Do(func() {
		file_protos_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_backend_proto_rawDescData)
	})
	return file_protos_backend_proto_rawDescData
}

var file_protos_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_backend_proto_goTypes = []interface{}{
	(*Command)(nil),         // 0: Command
	(*CommandFlag)(nil),     // 1: CommandFlag
	(*CommandRequest)(nil),  // 2: CommandRequest
	(*CommandResponse)(nil), // 3: CommandResponse
}
var file_protos_backend_proto_depIdxs = []int32{
	1, // 0: Command.flags:type_name -> CommandFlag
	0, // 1: CommandResponse.commands:type_name -> Command
	2, // 2: CommandService.GetCommands:input_type -> CommandRequest
	3, // 3: CommandService.GetCommands:output_type -> CommandResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_backend_proto_init() }
func file_protos_backend_proto_init() {
	if File_protos_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_protos_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandFlag); i {
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
		file_protos_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_protos_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
			RawDescriptor: file_protos_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_backend_proto_goTypes,
		DependencyIndexes: file_protos_backend_proto_depIdxs,
		MessageInfos:      file_protos_backend_proto_msgTypes,
	}.Build()
	File_protos_backend_proto = out.File
	file_protos_backend_proto_rawDesc = nil
	file_protos_backend_proto_goTypes = nil
	file_protos_backend_proto_depIdxs = nil
}
