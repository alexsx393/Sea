// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: api/openapi-spec/date/date.proto

package date

import (
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	types "github.com/kaydxh/sea/api/openapi-spec/types"
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

type NowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=RequestId,proto3" json:"request_id,omitempty"` //请求ID
}

func (x *NowRequest) Reset() {
	*x = NowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_date_date_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowRequest) ProtoMessage() {}

func (x *NowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_date_date_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowRequest.ProtoReflect.Descriptor instead.
func (*NowRequest) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_date_date_proto_rawDescGZIP(), []int{0}
}

func (x *NowRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type NowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string       `protobuf:"bytes,1,opt,name=request_id,json=RequestId,proto3" json:"request_id,omitempty"` // 请求ID
	Date      string       `protobuf:"bytes,2,opt,name=date,json=Date,proto3" json:"date,omitempty"`                  //当前时间
	Error     *types.Error `protobuf:"bytes,1000,opt,name=error,json=Error,proto3" json:"error,omitempty"`
}

func (x *NowResponse) Reset() {
	*x = NowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_date_date_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowResponse) ProtoMessage() {}

func (x *NowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_date_date_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowResponse.ProtoReflect.Descriptor instead.
func (*NowResponse) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_date_date_proto_rawDescGZIP(), []int{1}
}

func (x *NowResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *NowResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *NowResponse) GetError() *types.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type NowErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=RequestId,proto3" json:"request_id,omitempty"` //请求ID
}

func (x *NowErrorRequest) Reset() {
	*x = NowErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_date_date_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowErrorRequest) ProtoMessage() {}

func (x *NowErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_date_date_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowErrorRequest.ProtoReflect.Descriptor instead.
func (*NowErrorRequest) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_date_date_proto_rawDescGZIP(), []int{2}
}

func (x *NowErrorRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type NowErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string       `protobuf:"bytes,1,opt,name=request_id,json=RequestId,proto3" json:"request_id,omitempty"` // 请求ID
	Date      string       `protobuf:"bytes,2,opt,name=date,json=Date,proto3" json:"date,omitempty"`                  //当前时间
	Error     *types.Error `protobuf:"bytes,1000,opt,name=error,json=Error,proto3" json:"error,omitempty"`
}

func (x *NowErrorResponse) Reset() {
	*x = NowErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_date_date_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowErrorResponse) ProtoMessage() {}

func (x *NowErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_date_date_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowErrorResponse.ProtoReflect.Descriptor instead.
func (*NowErrorResponse) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_date_date_proto_rawDescGZIP(), []int{3}
}

func (x *NowErrorResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *NowErrorResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *NowErrorResponse) GetError() *types.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_api_openapi_spec_date_date_proto protoreflect.FileDescriptor

var file_api_openapi_spec_date_date_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x70, 0x65, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0a, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x0b, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0xe8,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x30, 0x0a, 0x0f, 0x4e, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x10, 0x4e, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x98, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x4e, 0x6f, 0x77, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4e,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x4e, 0x6f, 0x77, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x4e, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x4e, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x79, 0x64, 0x78, 0x68, 0x2f, 0x73, 0x65, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x3b, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_openapi_spec_date_date_proto_rawDescOnce sync.Once
	file_api_openapi_spec_date_date_proto_rawDescData = file_api_openapi_spec_date_date_proto_rawDesc
)

func file_api_openapi_spec_date_date_proto_rawDescGZIP() []byte {
	file_api_openapi_spec_date_date_proto_rawDescOnce.Do(func() {
		file_api_openapi_spec_date_date_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_openapi_spec_date_date_proto_rawDescData)
	})
	return file_api_openapi_spec_date_date_proto_rawDescData
}

var file_api_openapi_spec_date_date_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_openapi_spec_date_date_proto_goTypes = []interface{}{
	(*NowRequest)(nil),       // 0: sea.api.date.NowRequest
	(*NowResponse)(nil),      // 1: sea.api.date.NowResponse
	(*NowErrorRequest)(nil),  // 2: sea.api.date.NowErrorRequest
	(*NowErrorResponse)(nil), // 3: sea.api.date.NowErrorResponse
	(*types.Error)(nil),      // 4: sea.api.types.Error
}
var file_api_openapi_spec_date_date_proto_depIdxs = []int32{
	4, // 0: sea.api.date.NowResponse.error:type_name -> sea.api.types.Error
	4, // 1: sea.api.date.NowErrorResponse.error:type_name -> sea.api.types.Error
	0, // 2: sea.api.date.DateService.Now:input_type -> sea.api.date.NowRequest
	2, // 3: sea.api.date.DateService.NowError:input_type -> sea.api.date.NowErrorRequest
	1, // 4: sea.api.date.DateService.Now:output_type -> sea.api.date.NowResponse
	3, // 5: sea.api.date.DateService.NowError:output_type -> sea.api.date.NowErrorResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_openapi_spec_date_date_proto_init() }
func file_api_openapi_spec_date_date_proto_init() {
	if File_api_openapi_spec_date_date_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_openapi_spec_date_date_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowRequest); i {
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
		file_api_openapi_spec_date_date_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowResponse); i {
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
		file_api_openapi_spec_date_date_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowErrorRequest); i {
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
		file_api_openapi_spec_date_date_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowErrorResponse); i {
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
			RawDescriptor: file_api_openapi_spec_date_date_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_openapi_spec_date_date_proto_goTypes,
		DependencyIndexes: file_api_openapi_spec_date_date_proto_depIdxs,
		MessageInfos:      file_api_openapi_spec_date_date_proto_msgTypes,
	}.Build()
	File_api_openapi_spec_date_date_proto = out.File
	file_api_openapi_spec_date_date_proto_rawDesc = nil
	file_api_openapi_spec_date_date_proto_goTypes = nil
	file_api_openapi_spec_date_date_proto_depIdxs = nil
}
