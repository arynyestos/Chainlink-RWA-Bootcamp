// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v4.25.3
// source: shared/ptypes/label.proto

package ptypes

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

// SelectorOp defines the operation to be used in a selector
type SelectorOp int32

const (
	SelectorOp_EQ        SelectorOp = 0
	SelectorOp_NOT_EQ    SelectorOp = 1
	SelectorOp_IN        SelectorOp = 2
	SelectorOp_NOT_IN    SelectorOp = 3
	SelectorOp_EXIST     SelectorOp = 4
	SelectorOp_NOT_EXIST SelectorOp = 5
)

// Enum value maps for SelectorOp.
var (
	SelectorOp_name = map[int32]string{
		0: "EQ",
		1: "NOT_EQ",
		2: "IN",
		3: "NOT_IN",
		4: "EXIST",
		5: "NOT_EXIST",
	}
	SelectorOp_value = map[string]int32{
		"EQ":        0,
		"NOT_EQ":    1,
		"IN":        2,
		"NOT_IN":    3,
		"EXIST":     4,
		"NOT_EXIST": 5,
	}
)

func (x SelectorOp) Enum() *SelectorOp {
	p := new(SelectorOp)
	*p = x
	return p
}

func (x SelectorOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SelectorOp) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_ptypes_label_proto_enumTypes[0].Descriptor()
}

func (SelectorOp) Type() protoreflect.EnumType {
	return &file_shared_ptypes_label_proto_enumTypes[0]
}

func (x SelectorOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SelectorOp.Descriptor instead.
func (SelectorOp) EnumDescriptor() ([]byte, []int) {
	return file_shared_ptypes_label_proto_rawDescGZIP(), []int{0}
}

// Label defines a label as a key value pair
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_ptypes_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_shared_ptypes_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_shared_ptypes_label_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

// Selector defines a selector as a key value pair with an operation
type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op    SelectorOp `protobuf:"varint,2,opt,name=op,proto3,enum=api.label.SelectorOp" json:"op,omitempty"`
	Value *string    `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_ptypes_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_shared_ptypes_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_shared_ptypes_label_proto_rawDescGZIP(), []int{1}
}

func (x *Selector) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Selector) GetOp() SelectorOp {
	if x != nil {
		return x.Op
	}
	return SelectorOp_EQ
}

func (x *Selector) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_shared_ptypes_label_proto protoreflect.FileDescriptor

var file_shared_ptypes_label_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x68, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x19, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0x4e, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x12, 0x06,
	0x0a, 0x02, 0x45, 0x51, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x05,
	0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f,
	0x6a, 0x6f, 0x62, 0x2d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shared_ptypes_label_proto_rawDescOnce sync.Once
	file_shared_ptypes_label_proto_rawDescData = file_shared_ptypes_label_proto_rawDesc
)

func file_shared_ptypes_label_proto_rawDescGZIP() []byte {
	file_shared_ptypes_label_proto_rawDescOnce.Do(func() {
		file_shared_ptypes_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_ptypes_label_proto_rawDescData)
	})
	return file_shared_ptypes_label_proto_rawDescData
}

var file_shared_ptypes_label_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shared_ptypes_label_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_ptypes_label_proto_goTypes = []interface{}{
	(SelectorOp)(0),  // 0: api.label.SelectorOp
	(*Label)(nil),    // 1: api.label.Label
	(*Selector)(nil), // 2: api.label.Selector
}
var file_shared_ptypes_label_proto_depIdxs = []int32{
	0, // 0: api.label.Selector.op:type_name -> api.label.SelectorOp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_ptypes_label_proto_init() }
func file_shared_ptypes_label_proto_init() {
	if File_shared_ptypes_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_ptypes_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_shared_ptypes_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
	file_shared_ptypes_label_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_shared_ptypes_label_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_ptypes_label_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_ptypes_label_proto_goTypes,
		DependencyIndexes: file_shared_ptypes_label_proto_depIdxs,
		EnumInfos:         file_shared_ptypes_label_proto_enumTypes,
		MessageInfos:      file_shared_ptypes_label_proto_msgTypes,
	}.Build()
	File_shared_ptypes_label_proto = out.File
	file_shared_ptypes_label_proto_rawDesc = nil
	file_shared_ptypes_label_proto_goTypes = nil
	file_shared_ptypes_label_proto_depIdxs = nil
}
