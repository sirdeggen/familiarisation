// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: arithmetic_service.proto

package pb

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

var File_arithmetic_service_proto protoreflect.FileDescriptor

var file_arithmetic_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xf4, 0x01, 0x0a, 0x11, 0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_arithmetic_service_proto_goTypes = []interface{}{
	(*OperationParameters)(nil), // 0: pb.OperationParameters
	(*Answer)(nil),              // 1: pb.Answer
}
var file_arithmetic_service_proto_depIdxs = []int32{
	0, // 0: pb.ArithmeticService.GetAddition:input_type -> pb.OperationParameters
	0, // 1: pb.ArithmeticService.GetSubtraction:input_type -> pb.OperationParameters
	0, // 2: pb.ArithmeticService.GetMultiplication:input_type -> pb.OperationParameters
	0, // 3: pb.ArithmeticService.GetDivision:input_type -> pb.OperationParameters
	1, // 4: pb.ArithmeticService.GetAddition:output_type -> pb.Answer
	1, // 5: pb.ArithmeticService.GetSubtraction:output_type -> pb.Answer
	1, // 6: pb.ArithmeticService.GetMultiplication:output_type -> pb.Answer
	1, // 7: pb.ArithmeticService.GetDivision:output_type -> pb.Answer
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arithmetic_service_proto_init() }
func file_arithmetic_service_proto_init() {
	if File_arithmetic_service_proto != nil {
		return
	}
	file_number_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_arithmetic_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arithmetic_service_proto_goTypes,
		DependencyIndexes: file_arithmetic_service_proto_depIdxs,
	}.Build()
	File_arithmetic_service_proto = out.File
	file_arithmetic_service_proto_rawDesc = nil
	file_arithmetic_service_proto_goTypes = nil
	file_arithmetic_service_proto_depIdxs = nil
}
