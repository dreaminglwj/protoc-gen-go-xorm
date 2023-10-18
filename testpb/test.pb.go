// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v4.22.2
// source: test.proto

package testpb

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

type Status_Def int32

const (
	Status_Unspecified Status_Def = 0
	Status_FooStatus   Status_Def = 1
	Status_BarStatus   Status_Def = 2
)

// Enum value maps for Status_Def.
var (
	Status_Def_name = map[int32]string{
		0: "Unspecified",
		1: "FooStatus",
		2: "BarStatus",
	}
	Status_Def_value = map[string]int32{
		"Unspecified": 0,
		"FooStatus":   1,
		"BarStatus":   2,
	}
)

func (x Status_Def) Enum() *Status_Def {
	p := new(Status_Def)
	*p = x
	return p
}

func (x Status_Def) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Def) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (Status_Def) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Status_Def) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Status_Def) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Status_Def(num)
	return nil
}

// Deprecated: Use Status_Def.Descriptor instead.
func (Status_Def) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0}
}

type Status1_Def int32

const (
	Status1_Unspecified Status1_Def = 0
	Status1_FooStatus   Status1_Def = 1
	Status1_BarStatus   Status1_Def = 2
)

// Enum value maps for Status1_Def.
var (
	Status1_Def_name = map[int32]string{
		0: "Unspecified",
		1: "FooStatus",
		2: "BarStatus",
	}
	Status1_Def_value = map[string]int32{
		"Unspecified": 0,
		"FooStatus":   1,
		"BarStatus":   2,
	}
)

func (x Status1_Def) Enum() *Status1_Def {
	p := new(Status1_Def)
	*p = x
	return p
}

func (x Status1_Def) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status1_Def) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[1].Descriptor()
}

func (Status1_Def) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[1]
}

func (x Status1_Def) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Status1_Def) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Status1_Def(num)
	return nil
}

// Deprecated: Use Status1_Def.Descriptor instead.
func (Status1_Def) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1, 0}
}

type Status2_Def int32

const (
	Status2_Unspecified Status2_Def = 0
	Status2_FooStatus   Status2_Def = 1
	Status2_BarStatus   Status2_Def = 2
)

// Enum value maps for Status2_Def.
var (
	Status2_Def_name = map[int32]string{
		0: "Unspecified",
		1: "FooStatus",
		2: "BarStatus",
	}
	Status2_Def_value = map[string]int32{
		"Unspecified": 0,
		"FooStatus":   1,
		"BarStatus":   2,
	}
)

func (x Status2_Def) Enum() *Status2_Def {
	p := new(Status2_Def)
	*p = x
	return p
}

func (x Status2_Def) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status2_Def) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[2].Descriptor()
}

func (Status2_Def) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[2]
}

func (x Status2_Def) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Status2_Def) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Status2_Def(num)
	return nil
}

// Deprecated: Use Status2_Def.Descriptor instead.
func (Status2_Def) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2, 0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

type Status1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status1) Reset() {
	*x = Status1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status1) ProtoMessage() {}

func (x *Status1) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status1.ProtoReflect.Descriptor instead.
func (*Status1) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

type Status2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status2) Reset() {
	*x = Status2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status2) ProtoMessage() {}

func (x *Status2) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status2.ProtoReflect.Descriptor instead.
func (*Status2) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x34,
	0x0a, 0x03, 0x44, 0x65, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x6f, 0x6f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x10, 0x02, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x31, 0x22,
	0x34, 0x0a, 0x03, 0x44, 0x65, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x6f, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x10, 0x02, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x22, 0x34, 0x0a, 0x03, 0x44, 0x65, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x6f, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x10, 0x02, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x6c, 0x77, 0x6a,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x78,
	0x6f, 0x72, 0x6d, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_proto_goTypes = []interface{}{
	(Status_Def)(0),  // 0: testpb.Status.Def
	(Status1_Def)(0), // 1: testpb.Status1.Def
	(Status2_Def)(0), // 2: testpb.Status2.Def
	(*Status)(nil),   // 3: testpb.Status
	(*Status1)(nil),  // 4: testpb.Status1
	(*Status2)(nil),  // 5: testpb.Status2
}
var file_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status1); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status2); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
