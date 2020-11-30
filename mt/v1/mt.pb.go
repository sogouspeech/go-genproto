// Copyright 2018 Sogou Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: sogou/speech/mt/v1/mt.proto

package mt

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TranslateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required* The translate configurations
	Config *TranslateConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// *Required* The text will be translated
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TranslateTextRequest) Reset() {
	*x = TranslateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateTextRequest) ProtoMessage() {}

func (x *TranslateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateTextRequest.ProtoReflect.Descriptor instead.
func (*TranslateTextRequest) Descriptor() ([]byte, []int) {
	return file_sogou_speech_mt_v1_mt_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateTextRequest) GetConfig() *TranslateConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *TranslateTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TranslateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required* The language code of the source text,
	// set to one of the language codes listed in [Language Support](https://zhiyin.sogou.com/doc/?url=/docs/content/mt/concepts/lauguages/).
	// Example: "zh-cmn-Hans-CN".
	SourceLanguageCode string `protobuf:"bytes,1,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// *Required* The language code of the target text
	TargetLanguageCode string `protobuf:"bytes,2,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
}

func (x *TranslateConfig) Reset() {
	*x = TranslateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateConfig) ProtoMessage() {}

func (x *TranslateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateConfig.ProtoReflect.Descriptor instead.
func (*TranslateConfig) Descriptor() ([]byte, []int) {
	return file_sogou_speech_mt_v1_mt_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateConfig) GetSourceLanguageCode() string {
	if x != nil {
		return x.SourceLanguageCode
	}
	return ""
}

func (x *TranslateConfig) GetTargetLanguageCode() string {
	if x != nil {
		return x.TargetLanguageCode
	}
	return ""
}

type TranslateTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source text
	SourceText string `protobuf:"bytes,1,opt,name=source_text,json=sourceText,proto3" json:"source_text,omitempty"`
	// Text translated into the target language.
	TranslatedText string `protobuf:"bytes,2,opt,name=translated_text,json=translatedText,proto3" json:"translated_text,omitempty"`
	// The language code of the source text
	SourceLanguageCode string `protobuf:"bytes,3,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// The language code of the target text
	TargetLanguageCode string `protobuf:"bytes,4,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
}

func (x *TranslateTextResponse) Reset() {
	*x = TranslateTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateTextResponse) ProtoMessage() {}

func (x *TranslateTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_mt_v1_mt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateTextResponse.ProtoReflect.Descriptor instead.
func (*TranslateTextResponse) Descriptor() ([]byte, []int) {
	return file_sogou_speech_mt_v1_mt_proto_rawDescGZIP(), []int{2}
}

func (x *TranslateTextResponse) GetSourceText() string {
	if x != nil {
		return x.SourceText
	}
	return ""
}

func (x *TranslateTextResponse) GetTranslatedText() string {
	if x != nil {
		return x.TranslatedText
	}
	return ""
}

func (x *TranslateTextResponse) GetSourceLanguageCode() string {
	if x != nil {
		return x.SourceLanguageCode
	}
	return ""
}

func (x *TranslateTextResponse) GetTargetLanguageCode() string {
	if x != nil {
		return x.TargetLanguageCode
	}
	return ""
}

var File_sogou_speech_mt_v1_mt_proto protoreflect.FileDescriptor

var file_sogou_speech_mt_v1_mt_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x6d,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73,
	0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0x67, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x67, 0x6f,
	0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x75, 0x0a, 0x0f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a,
	0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xc5, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x6a, 0x0a, 0x02, 0x6d, 0x74, 0x12,
	0x64, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x28, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6f, 0x67,
	0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f,
	0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x07, 0x4d, 0x54, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x74, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x53, 0x50, 0x42, 0xaa, 0x02, 0x12, 0x53, 0x6f,
	0x67, 0x6f, 0x75, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x4d, 0x54, 0x2e, 0x56, 0x31,
	0xba, 0x02, 0x03, 0x53, 0x50, 0x42, 0xca, 0x02, 0x12, 0x53, 0x6f, 0x67, 0x6f, 0x75, 0x5c, 0x53,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x5c, 0x4d, 0x54, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sogou_speech_mt_v1_mt_proto_rawDescOnce sync.Once
	file_sogou_speech_mt_v1_mt_proto_rawDescData = file_sogou_speech_mt_v1_mt_proto_rawDesc
)

func file_sogou_speech_mt_v1_mt_proto_rawDescGZIP() []byte {
	file_sogou_speech_mt_v1_mt_proto_rawDescOnce.Do(func() {
		file_sogou_speech_mt_v1_mt_proto_rawDescData = protoimpl.X.CompressGZIP(file_sogou_speech_mt_v1_mt_proto_rawDescData)
	})
	return file_sogou_speech_mt_v1_mt_proto_rawDescData
}

var file_sogou_speech_mt_v1_mt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sogou_speech_mt_v1_mt_proto_goTypes = []interface{}{
	(*TranslateTextRequest)(nil),  // 0: sogou.speech.mt.v1.TranslateTextRequest
	(*TranslateConfig)(nil),       // 1: sogou.speech.mt.v1.TranslateConfig
	(*TranslateTextResponse)(nil), // 2: sogou.speech.mt.v1.TranslateTextResponse
}
var file_sogou_speech_mt_v1_mt_proto_depIdxs = []int32{
	1, // 0: sogou.speech.mt.v1.TranslateTextRequest.config:type_name -> sogou.speech.mt.v1.TranslateConfig
	0, // 1: sogou.speech.mt.v1.mt.TranslateText:input_type -> sogou.speech.mt.v1.TranslateTextRequest
	2, // 2: sogou.speech.mt.v1.mt.TranslateText:output_type -> sogou.speech.mt.v1.TranslateTextResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sogou_speech_mt_v1_mt_proto_init() }
func file_sogou_speech_mt_v1_mt_proto_init() {
	if File_sogou_speech_mt_v1_mt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sogou_speech_mt_v1_mt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateTextRequest); i {
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
		file_sogou_speech_mt_v1_mt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateConfig); i {
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
		file_sogou_speech_mt_v1_mt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateTextResponse); i {
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
			RawDescriptor: file_sogou_speech_mt_v1_mt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sogou_speech_mt_v1_mt_proto_goTypes,
		DependencyIndexes: file_sogou_speech_mt_v1_mt_proto_depIdxs,
		MessageInfos:      file_sogou_speech_mt_v1_mt_proto_msgTypes,
	}.Build()
	File_sogou_speech_mt_v1_mt_proto = out.File
	file_sogou_speech_mt_v1_mt_proto_rawDesc = nil
	file_sogou_speech_mt_v1_mt_proto_goTypes = nil
	file_sogou_speech_mt_v1_mt_proto_depIdxs = nil
}
