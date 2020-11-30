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
// source: sogou/speech/tts/v1/tts.proto

package tts

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

// Configuration to set up audio encoder. The encoding determines the output
// audio format that we'd like.
type AudioConfig_AudioEncoding int32

const (
	// Not specified. Will return result 400.
	AudioConfig_ENCODING_UNSPECIFIED AudioConfig_AudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	// Audio content returned as LINEAR16 also contains a WAV header.
	// sample rate = 16000
	AudioConfig_LINEAR16 AudioConfig_AudioEncoding = 1
	// MP3 audio. currently sample rate = 16000
	AudioConfig_MP3 AudioConfig_AudioEncoding = 2
	// Adaptive Multi-Rate Narrowband codec. sample rate = 8000.
	AudioConfig_AMR AudioConfig_AudioEncoding = 3
)

// Enum value maps for AudioConfig_AudioEncoding.
var (
	AudioConfig_AudioEncoding_name = map[int32]string{
		0: "ENCODING_UNSPECIFIED",
		1: "LINEAR16",
		2: "MP3",
		3: "AMR",
	}
	AudioConfig_AudioEncoding_value = map[string]int32{
		"ENCODING_UNSPECIFIED": 0,
		"LINEAR16":             1,
		"MP3":                  2,
		"AMR":                  3,
	}
)

func (x AudioConfig_AudioEncoding) Enum() *AudioConfig_AudioEncoding {
	p := new(AudioConfig_AudioEncoding)
	*p = x
	return p
}

func (x AudioConfig_AudioEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioConfig_AudioEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_sogou_speech_tts_v1_tts_proto_enumTypes[0].Descriptor()
}

func (AudioConfig_AudioEncoding) Type() protoreflect.EnumType {
	return &file_sogou_speech_tts_v1_tts_proto_enumTypes[0]
}

func (x AudioConfig_AudioEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioConfig_AudioEncoding.Descriptor instead.
func (AudioConfig_AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{4, 0}
}

type SynthesizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required*. The Synthesizer requires either plain text as input.
	Input *SynthesisInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// *Required*. The configuration of the synthesize.
	Config *SynthesizeConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SynthesizeRequest) Reset() {
	*x = SynthesizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeRequest) ProtoMessage() {}

func (x *SynthesizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeRequest.ProtoReflect.Descriptor instead.
func (*SynthesizeRequest) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{0}
}

func (x *SynthesizeRequest) GetInput() *SynthesisInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *SynthesizeRequest) GetConfig() *SynthesizeConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type SynthesizeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required*. The desired voice of the synthesized audio.
	VoiceConfig *VoiceConfig `protobuf:"bytes,1,opt,name=voice_config,json=voiceConfig,proto3" json:"voice_config,omitempty"`
	// *Required*. The configuration of the synthesized audio.
	AudioConfig *AudioConfig `protobuf:"bytes,2,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"`
}

func (x *SynthesizeConfig) Reset() {
	*x = SynthesizeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeConfig) ProtoMessage() {}

func (x *SynthesizeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeConfig.ProtoReflect.Descriptor instead.
func (*SynthesizeConfig) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{1}
}

func (x *SynthesizeConfig) GetVoiceConfig() *VoiceConfig {
	if x != nil {
		return x.VoiceConfig
	}
	return nil
}

func (x *SynthesizeConfig) GetAudioConfig() *AudioConfig {
	if x != nil {
		return x.AudioConfig
	}
	return nil
}

type SynthesisInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required*. The raw text to be synthesized.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SynthesisInput) Reset() {
	*x = SynthesisInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesisInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesisInput) ProtoMessage() {}

func (x *SynthesisInput) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesisInput.ProtoReflect.Descriptor instead.
func (*SynthesisInput) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{2}
}

func (x *SynthesisInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type VoiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required*. The language code of the voice.
	// Example: "zh-cmn-Hans-CN".
	// See [Language Support](https://zhiyin.sogou.com/doc/?url=/docs/content/tts/concepts/languages/) for a list of the currently supported language codes.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// *Optional*. The name of the speaker. If not set, the service will choose a
	// voice based on the other parameters such as language_code.
	// Current availabe speakers: [male, female]
	Speaker string `protobuf:"bytes,2,opt,name=speaker,proto3" json:"speaker,omitempty"`
}

func (x *VoiceConfig) Reset() {
	*x = VoiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceConfig) ProtoMessage() {}

func (x *VoiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceConfig.ProtoReflect.Descriptor instead.
func (*VoiceConfig) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{3}
}

func (x *VoiceConfig) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *VoiceConfig) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

type AudioConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *Required*. The format of the requested audio byte stream.
	AudioEncoding AudioConfig_AudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=sogou.speech.tts.v1.AudioConfig_AudioEncoding" json:"audio_encoding,omitempty"`
	// *Optional* speaking rate/speed, in the range [0.7, 1.3]. 1.0 is the normal
	// native speed supported by the specific voice.
	// If unset(0.0), defaults to the native 1.0 speed. Any
	// other values < 0.7 or > 1.3 will return an error.
	SpeakingRate float64 `protobuf:"fixed64,2,opt,name=speaking_rate,json=speakingRate,proto3" json:"speaking_rate,omitempty"`
	// *Optional* speaking pitch, in the range [0.8, 1.2]. 1.0 is the default pitch.
	// If unset(0.0), defaults to the native 1.0 pitch. Any
	// other values < 0.8 or > 1.2 will return an error.
	Pitch float64 `protobuf:"fixed64,3,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// *Optional* audio volume, in the range [0.7, 1.3]. 1.0 is the default volume.
	// If unset(0.0), defaults to the native 1.0 volume. Any
	// other values < 0.7 or > 1.3 will return an error.
	Volume float64 `protobuf:"fixed64,4,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *AudioConfig) Reset() {
	*x = AudioConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioConfig) ProtoMessage() {}

func (x *AudioConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioConfig.ProtoReflect.Descriptor instead.
func (*AudioConfig) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{4}
}

func (x *AudioConfig) GetAudioEncoding() AudioConfig_AudioEncoding {
	if x != nil {
		return x.AudioEncoding
	}
	return AudioConfig_ENCODING_UNSPECIFIED
}

func (x *AudioConfig) GetSpeakingRate() float64 {
	if x != nil {
		return x.SpeakingRate
	}
	return 0
}

func (x *AudioConfig) GetPitch() float64 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *AudioConfig) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type SynthesizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The audio data bytes encoded as specified in the request, including the
	// header (For LINEAR16 audio, we include the WAV header). Note: as
	// with all bytes fields, protobuffers use a pure binary representation.
	AudioContent []byte `protobuf:"bytes,2,opt,name=audio_content,json=audioContent,proto3" json:"audio_content,omitempty"`
}

func (x *SynthesizeResponse) Reset() {
	*x = SynthesizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeResponse) ProtoMessage() {}

func (x *SynthesizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sogou_speech_tts_v1_tts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeResponse.ProtoReflect.Descriptor instead.
func (*SynthesizeResponse) Descriptor() ([]byte, []int) {
	return file_sogou_speech_tts_v1_tts_proto_rawDescGZIP(), []int{5}
}

func (x *SynthesizeResponse) GetAudioContent() []byte {
	if x != nil {
		return x.AudioContent
	}
	return nil
}

var File_sogou_speech_tts_v1_tts_proto protoreflect.FileDescriptor

var file_sogou_speech_tts_v1_tts_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x74,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x67, 0x6f,
	0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73,
	0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0c, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43,
	0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4c, 0x0a, 0x0b, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x22, 0x82, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x22, 0x49, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x31, 0x36, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x50,
	0x33, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4d, 0x52, 0x10, 0x03, 0x22, 0x39, 0x0a, 0x12,
	0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x03, 0x74, 0x74, 0x73, 0x12,
	0x5d, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x2e,
	0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68,
	0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x6e, 0x74, 0x68,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x89, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x74, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x54, 0x54, 0x53, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x27, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e,
	0x73, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x74, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x53,
	0x50, 0x42, 0xaa, 0x02, 0x13, 0x53, 0x6f, 0x67, 0x6f, 0x75, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x2e, 0x54, 0x54, 0x53, 0x2e, 0x56, 0x31, 0xba, 0x02, 0x03, 0x53, 0x50, 0x42, 0xca, 0x02,
	0x13, 0x53, 0x6f, 0x67, 0x6f, 0x75, 0x5c, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5c, 0x54, 0x54,
	0x53, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sogou_speech_tts_v1_tts_proto_rawDescOnce sync.Once
	file_sogou_speech_tts_v1_tts_proto_rawDescData = file_sogou_speech_tts_v1_tts_proto_rawDesc
)

func file_sogou_speech_tts_v1_tts_proto_rawDescGZIP() []byte {
	file_sogou_speech_tts_v1_tts_proto_rawDescOnce.Do(func() {
		file_sogou_speech_tts_v1_tts_proto_rawDescData = protoimpl.X.CompressGZIP(file_sogou_speech_tts_v1_tts_proto_rawDescData)
	})
	return file_sogou_speech_tts_v1_tts_proto_rawDescData
}

var file_sogou_speech_tts_v1_tts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sogou_speech_tts_v1_tts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sogou_speech_tts_v1_tts_proto_goTypes = []interface{}{
	(AudioConfig_AudioEncoding)(0), // 0: sogou.speech.tts.v1.AudioConfig.AudioEncoding
	(*SynthesizeRequest)(nil),      // 1: sogou.speech.tts.v1.SynthesizeRequest
	(*SynthesizeConfig)(nil),       // 2: sogou.speech.tts.v1.SynthesizeConfig
	(*SynthesisInput)(nil),         // 3: sogou.speech.tts.v1.SynthesisInput
	(*VoiceConfig)(nil),            // 4: sogou.speech.tts.v1.VoiceConfig
	(*AudioConfig)(nil),            // 5: sogou.speech.tts.v1.AudioConfig
	(*SynthesizeResponse)(nil),     // 6: sogou.speech.tts.v1.SynthesizeResponse
}
var file_sogou_speech_tts_v1_tts_proto_depIdxs = []int32{
	3, // 0: sogou.speech.tts.v1.SynthesizeRequest.input:type_name -> sogou.speech.tts.v1.SynthesisInput
	2, // 1: sogou.speech.tts.v1.SynthesizeRequest.config:type_name -> sogou.speech.tts.v1.SynthesizeConfig
	4, // 2: sogou.speech.tts.v1.SynthesizeConfig.voice_config:type_name -> sogou.speech.tts.v1.VoiceConfig
	5, // 3: sogou.speech.tts.v1.SynthesizeConfig.audio_config:type_name -> sogou.speech.tts.v1.AudioConfig
	0, // 4: sogou.speech.tts.v1.AudioConfig.audio_encoding:type_name -> sogou.speech.tts.v1.AudioConfig.AudioEncoding
	1, // 5: sogou.speech.tts.v1.tts.Synthesize:input_type -> sogou.speech.tts.v1.SynthesizeRequest
	1, // 6: sogou.speech.tts.v1.tts.StreamingSynthesize:input_type -> sogou.speech.tts.v1.SynthesizeRequest
	6, // 7: sogou.speech.tts.v1.tts.Synthesize:output_type -> sogou.speech.tts.v1.SynthesizeResponse
	6, // 8: sogou.speech.tts.v1.tts.StreamingSynthesize:output_type -> sogou.speech.tts.v1.SynthesizeResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sogou_speech_tts_v1_tts_proto_init() }
func file_sogou_speech_tts_v1_tts_proto_init() {
	if File_sogou_speech_tts_v1_tts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sogou_speech_tts_v1_tts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeRequest); i {
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
		file_sogou_speech_tts_v1_tts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeConfig); i {
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
		file_sogou_speech_tts_v1_tts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesisInput); i {
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
		file_sogou_speech_tts_v1_tts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceConfig); i {
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
		file_sogou_speech_tts_v1_tts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioConfig); i {
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
		file_sogou_speech_tts_v1_tts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeResponse); i {
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
			RawDescriptor: file_sogou_speech_tts_v1_tts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sogou_speech_tts_v1_tts_proto_goTypes,
		DependencyIndexes: file_sogou_speech_tts_v1_tts_proto_depIdxs,
		EnumInfos:         file_sogou_speech_tts_v1_tts_proto_enumTypes,
		MessageInfos:      file_sogou_speech_tts_v1_tts_proto_msgTypes,
	}.Build()
	File_sogou_speech_tts_v1_tts_proto = out.File
	file_sogou_speech_tts_v1_tts_proto_rawDesc = nil
	file_sogou_speech_tts_v1_tts_proto_goTypes = nil
	file_sogou_speech_tts_v1_tts_proto_depIdxs = nil
}
