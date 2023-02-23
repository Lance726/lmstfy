// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: engine/redis_v2/model/job.proto

package model

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

type JobData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JobData) Reset() {
	*x = JobData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_redis_v2_model_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobData) ProtoMessage() {}

func (x *JobData) ProtoReflect() protoreflect.Message {
	mi := &file_engine_redis_v2_model_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobData.ProtoReflect.Descriptor instead.
func (*JobData) Descriptor() ([]byte, []int) {
	return file_engine_redis_v2_model_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *JobData) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_engine_redis_v2_model_job_proto protoreflect.FileDescriptor

var file_engine_redis_v2_model_job_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x76,
	0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x4a, 0x6f, 0x62,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6d, 0x73, 0x74, 0x66, 0x79, 0x2f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_redis_v2_model_job_proto_rawDescOnce sync.Once
	file_engine_redis_v2_model_job_proto_rawDescData = file_engine_redis_v2_model_job_proto_rawDesc
)

func file_engine_redis_v2_model_job_proto_rawDescGZIP() []byte {
	file_engine_redis_v2_model_job_proto_rawDescOnce.Do(func() {
		file_engine_redis_v2_model_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_redis_v2_model_job_proto_rawDescData)
	})
	return file_engine_redis_v2_model_job_proto_rawDescData
}

var file_engine_redis_v2_model_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_engine_redis_v2_model_job_proto_goTypes = []interface{}{
	(*JobData)(nil), // 0: model.JobData
	nil,             // 1: model.JobData.AttributesEntry
}
var file_engine_redis_v2_model_job_proto_depIdxs = []int32{
	1, // 0: model.JobData.attributes:type_name -> model.JobData.AttributesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_engine_redis_v2_model_job_proto_init() }
func file_engine_redis_v2_model_job_proto_init() {
	if File_engine_redis_v2_model_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_redis_v2_model_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobData); i {
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
			RawDescriptor: file_engine_redis_v2_model_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_redis_v2_model_job_proto_goTypes,
		DependencyIndexes: file_engine_redis_v2_model_job_proto_depIdxs,
		MessageInfos:      file_engine_redis_v2_model_job_proto_msgTypes,
	}.Build()
	File_engine_redis_v2_model_job_proto = out.File
	file_engine_redis_v2_model_job_proto_rawDesc = nil
	file_engine_redis_v2_model_job_proto_goTypes = nil
	file_engine_redis_v2_model_job_proto_depIdxs = nil
}