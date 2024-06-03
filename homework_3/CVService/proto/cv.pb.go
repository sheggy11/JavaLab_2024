package proto

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

type CVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skills []string `protobuf:"bytes,1,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *CVRequest) Reset() {
	*x = CVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVRequest) ProtoMessage() {}

func (x *CVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVRequest.ProtoReflect.Descriptor instead.
func (*CVRequest) Descriptor() ([]byte, []int) {
	return file_proto_cv_proto_rawDescGZIP(), []int{0}
}

func (x *CVRequest) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

type CV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age    uint32   `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Skills []string `protobuf:"bytes,4,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *CV) Reset() {
	*x = CV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CV) ProtoMessage() {}

func (x *CV) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CV.ProtoReflect.Descriptor instead.
func (*CV) Descriptor() ([]byte, []int) {
	return file_proto_cv_proto_rawDescGZIP(), []int{1}
}

func (x *CV) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CV) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CV) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CV) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

type CVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvList []*CV `protobuf:"bytes,1,rep,name=cvList,proto3" json:"cvList,omitempty"`
}

func (x *CVResponse) Reset() {
	*x = CVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVResponse) ProtoMessage() {}

func (x *CVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVResponse.ProtoReflect.Descriptor instead.
func (*CVResponse) Descriptor() ([]byte, []int) {
	return file_proto_cv_proto_rawDescGZIP(), []int{2}
}

func (x *CVResponse) GetCvList() []*CV {
	if x != nil {
		return x.CvList
	}
	return nil
}

var File_proto_cv_proto protoreflect.FileDescriptor

var file_proto_cv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x23, 0x0a, 0x09, 0x43, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x52, 0x0a, 0x02, 0x43, 0x56, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x41, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x29, 0x0a, 0x0a, 0x43, 0x56, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x76, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x03, 0x2e, 0x43, 0x56, 0x52, 0x06, 0x63, 0x76,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0x32, 0x0a, 0x09, 0x43, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x56, 0x12, 0x0a, 0x2e,
	0x43, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x43, 0x56, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cv_proto_rawDescOnce sync.Once
	file_proto_cv_proto_rawDescData = file_proto_cv_proto_rawDesc
)

func file_proto_cv_proto_rawDescGZIP() []byte {
	file_proto_cv_proto_rawDescOnce.Do(func() {
		file_proto_cv_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cv_proto_rawDescData)
	})
	return file_proto_cv_proto_rawDescData
}

var file_proto_cv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_cv_proto_goTypes = []interface{}{
	(*CVRequest)(nil),  // 0: CVRequest
	(*CV)(nil),         // 1: CV
	(*CVResponse)(nil), // 2: CVResponse
}
var file_proto_cv_proto_depIdxs = []int32{
	1, // 0: CVResponse.cvList:type_name -> CV
	0, // 1: CVService.GetAllCV:input_type -> CVRequest
	2, // 2: CVService.GetAllCV:output_type -> CVResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cv_proto_init() }
func file_proto_cv_proto_init() {
	if File_proto_cv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVRequest); i {
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
		file_proto_cv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CV); i {
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
		file_proto_cv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVResponse); i {
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
			RawDescriptor: file_proto_cv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cv_proto_goTypes,
		DependencyIndexes: file_proto_cv_proto_depIdxs,
		MessageInfos:      file_proto_cv_proto_msgTypes,
	}.Build()
	File_proto_cv_proto = out.File
	file_proto_cv_proto_rawDesc = nil
	file_proto_cv_proto_goTypes = nil
	file_proto_cv_proto_depIdxs = nil
}