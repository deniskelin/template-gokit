// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: mtt/apistatus/service.proto

package apistatus

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mtt_apistatus_service_proto protoreflect.FileDescriptor

var file_mtt_apistatus_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x74, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x1b, 0x6d, 0x74,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x74, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa0, 0x01, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x53, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x2e, 0x6d, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x2e, 0x6d, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x74,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4c, 0x0a, 0x11, 0x6d, 0x74,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x2e, 0x72, 0x6e, 0x64, 0x2e, 0x6d, 0x74, 0x74, 0x2f, 0x6d, 0x74, 0x74,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x74, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x61,
	0x70, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mtt_apistatus_service_proto_goTypes = []interface{}{
	(*APIVersionRequest)(nil),  // 0: mtt.apistatus.APIVersionRequest
	(*InfoRequest)(nil),        // 1: mtt.apistatus.InfoRequest
	(*APIVersionResponse)(nil), // 2: mtt.apistatus.APIVersionResponse
	(*InfoResponse)(nil),       // 3: mtt.apistatus.InfoResponse
}
var file_mtt_apistatus_service_proto_depIdxs = []int32{
	0, // 0: mtt.apistatus.System.APIVersion:input_type -> mtt.apistatus.APIVersionRequest
	1, // 1: mtt.apistatus.System.Info:input_type -> mtt.apistatus.InfoRequest
	2, // 2: mtt.apistatus.System.APIVersion:output_type -> mtt.apistatus.APIVersionResponse
	3, // 3: mtt.apistatus.System.Info:output_type -> mtt.apistatus.InfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mtt_apistatus_service_proto_init() }
func file_mtt_apistatus_service_proto_init() {
	if File_mtt_apistatus_service_proto != nil {
		return
	}
	file_mtt_apistatus_request_proto_init()
	file_mtt_apistatus_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mtt_apistatus_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mtt_apistatus_service_proto_goTypes,
		DependencyIndexes: file_mtt_apistatus_service_proto_depIdxs,
	}.Build()
	File_mtt_apistatus_service_proto = out.File
	file_mtt_apistatus_service_proto_rawDesc = nil
	file_mtt_apistatus_service_proto_goTypes = nil
	file_mtt_apistatus_service_proto_depIdxs = nil
}