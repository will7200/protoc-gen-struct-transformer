// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: options/annotations.proto

// Package transformer contains extend options for protobuf files, messages and
// fields.
// Options are used for customizing transformation process.

package options

import (
	fmt "fmt"
	math "math"

	"github.com/golang/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_GoModelsFilePath = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5201,
	Name:          "transformer.go_models_file_path",
	Tag:           "bytes,5201,opt,name=go_models_file_path",
	Filename:      "options/annotations.proto",
}

var E_GoRepoPackage = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5202,
	Name:          "transformer.go_repo_package",
	Tag:           "bytes,5202,opt,name=go_repo_package",
	Filename:      "options/annotations.proto",
}

var E_GoProtobufPackage = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5203,
	Name:          "transformer.go_protobuf_package",
	Tag:           "bytes,5203,opt,name=go_protobuf_package",
	Filename:      "options/annotations.proto",
}

var E_GoStruct = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5100,
	Name:          "transformer.go_struct",
	Tag:           "bytes,5100,opt,name=go_struct",
	Filename:      "options/annotations.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         5300,
	Name:          "transformer.embed",
	Tag:           "varint,5300,opt,name=embed",
	Filename:      "options/annotations.proto",
}

var E_Skip = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         5301,
	Name:          "transformer.skip",
	Tag:           "varint,5301,opt,name=skip",
	Filename:      "options/annotations.proto",
}

var E_MapTo = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5303,
	Name:          "transformer.map_to",
	Tag:           "bytes,5303,opt,name=map_to",
	Filename:      "options/annotations.proto",
}

var E_MapAs = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         5304,
	Name:          "transformer.map_as",
	Tag:           "bytes,5304,opt,name=map_as",
	Filename:      "options/annotations.proto",
}

var E_Custom = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         5305,
	Name:          "transformer.custom",
	Tag:           "varint,5305,opt,name=custom",
	Filename:      "options/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_GoModelsFilePath)
	proto.RegisterExtension(E_GoRepoPackage)
	proto.RegisterExtension(E_GoProtobufPackage)
	proto.RegisterExtension(E_GoStruct)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Skip)
	proto.RegisterExtension(E_MapTo)
	proto.RegisterExtension(E_MapAs)
	proto.RegisterExtension(E_Custom)
}

func init() { proto.RegisterFile("options/annotations.proto", fileDescriptor_5df765dc541320cc) }

var fileDescriptor_5df765dc541320cc = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc0, 0xf1, 0x46, 0xfa, 0xb5, 0xbf, 0xd6, 0x08, 0x01, 0x61, 0x01, 0x04, 0xa1, 0x4c, 0xb4,
	0x4b, 0x2a, 0xf1, 0x6f, 0x88, 0xc4, 0x00, 0x12, 0x4c, 0x54, 0x54, 0x85, 0x89, 0xc5, 0x72, 0x93,
	0xab, 0x1b, 0x35, 0xc9, 0x59, 0xb6, 0xfb, 0x1e, 0x3c, 0x0c, 0x08, 0x78, 0x03, 0xc6, 0x02, 0x0b,
	0x23, 0x6a, 0x57, 0x1e, 0x02, 0x61, 0x37, 0x30, 0x80, 0x14, 0xb6, 0x48, 0x77, 0x9f, 0xaf, 0x6f,
	0x08, 0x59, 0x45, 0xa1, 0x63, 0xcc, 0x54, 0x8b, 0x65, 0x19, 0x6a, 0x66, 0xbe, 0x7d, 0x21, 0x51,
	0xa3, 0x3b, 0xa7, 0x25, 0xcb, 0x54, 0x1f, 0x65, 0x0a, 0x72, 0xad, 0xce, 0x11, 0x79, 0x02, 0x2d,
	0x33, 0xea, 0x8d, 0xfa, 0xad, 0x08, 0x54, 0x28, 0x63, 0xa1, 0x51, 0xda, 0xf5, 0xe0, 0x8c, 0x2c,
	0x73, 0xa4, 0x29, 0x46, 0x90, 0x28, 0xda, 0x8f, 0x13, 0xa0, 0x82, 0xe9, 0x81, 0xbb, 0xee, 0x5b,
	0xe9, 0xe7, 0xd2, 0x3f, 0x8d, 0x13, 0x38, 0xb7, 0xaf, 0xae, 0x3c, 0x35, 0xea, 0x4e, 0xa3, 0xd6,
	0x5d, 0xe4, 0xd8, 0x36, 0xf0, 0x73, 0xd6, 0x61, 0x7a, 0x10, 0x9c, 0x90, 0x05, 0x8e, 0x54, 0x82,
	0x40, 0x2a, 0x58, 0x38, 0x64, 0x1c, 0x0a, 0x4a, 0xcf, 0xb6, 0x34, 0xcf, 0xb1, 0x0b, 0x02, 0x3b,
	0xd6, 0x04, 0x6d, 0x73, 0x54, 0x0e, 0xfe, 0x98, 0x7a, 0xb1, 0xa9, 0x25, 0x8e, 0x9d, 0xd9, 0x38,
	0xcf, 0x1d, 0x92, 0x1a, 0x47, 0xaa, 0xb4, 0x1c, 0x85, 0xda, 0xdd, 0xfc, 0x11, 0x69, 0x83, 0x52,
	0x8c, 0x7f, 0x75, 0xde, 0xb7, 0x4d, 0xa7, 0xca, 0xf1, 0xc2, 0x88, 0x60, 0x8f, 0x94, 0x21, 0xed,
	0x41, 0xe4, 0x6e, 0xfc, 0xf2, 0x3e, 0x24, 0x51, 0x0e, 0x6f, 0x9a, 0x75, 0xa7, 0x51, 0xed, 0xda,
	0xe5, 0x60, 0x87, 0xfc, 0x53, 0xc3, 0x58, 0x14, 0xa1, 0x5b, 0x8b, 0xcc, 0x6e, 0xb0, 0x4f, 0x2a,
	0x29, 0x13, 0x54, 0x63, 0x91, 0xba, 0x6b, 0x9a, 0x1b, 0xcb, 0x29, 0x13, 0x97, 0x98, 0x33, 0xa6,
	0x8a, 0xd8, 0xfd, 0x37, 0x3b, 0x52, 0xc1, 0x01, 0xa9, 0x84, 0x23, 0xa5, 0x31, 0x2d, 0x62, 0x0f,
	0xf6, 0xc6, 0xd9, 0xf6, 0xf1, 0xd6, 0xe3, 0xc4, 0x73, 0xc6, 0x13, 0xcf, 0x79, 0x9b, 0x78, 0xce,
	0xf5, 0xd4, 0x2b, 0x8d, 0xa7, 0x5e, 0xe9, 0x75, 0xea, 0x95, 0xae, 0xfe, 0xcf, 0x7e, 0xcb, 0x5e,
	0xc5, 0x84, 0x76, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x4b, 0xc5, 0xea, 0xa8, 0x02, 0x00,
	0x00,
}
