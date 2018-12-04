// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sogou/speech/tts/v1/tts.proto

package tts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

var AudioConfig_AudioEncoding_name = map[int32]string{
	0: "ENCODING_UNSPECIFIED",
	1: "LINEAR16",
	2: "MP3",
	3: "AMR",
}

var AudioConfig_AudioEncoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"MP3":                  2,
	"AMR":                  3,
}

func (x AudioConfig_AudioEncoding) String() string {
	return proto.EnumName(AudioConfig_AudioEncoding_name, int32(x))
}

func (AudioConfig_AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{4, 0}
}

type SynthesizeRequest struct {
	// *Required*. The Synthesizer requires either plain text as input.
	Input *SynthesisInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// *Required*. The configuration of the synthesize.
	Config               *SynthesizeConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SynthesizeRequest) Reset()         { *m = SynthesizeRequest{} }
func (m *SynthesizeRequest) String() string { return proto.CompactTextString(m) }
func (*SynthesizeRequest) ProtoMessage()    {}
func (*SynthesizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{0}
}

func (m *SynthesizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeRequest.Unmarshal(m, b)
}
func (m *SynthesizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeRequest.Marshal(b, m, deterministic)
}
func (m *SynthesizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeRequest.Merge(m, src)
}
func (m *SynthesizeRequest) XXX_Size() int {
	return xxx_messageInfo_SynthesizeRequest.Size(m)
}
func (m *SynthesizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeRequest proto.InternalMessageInfo

func (m *SynthesizeRequest) GetInput() *SynthesisInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *SynthesizeRequest) GetConfig() *SynthesizeConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type SynthesizeConfig struct {
	// *Required*. The desired voice of the synthesized audio.
	VoiceConfig *VoiceConfig `protobuf:"bytes,1,opt,name=voice_config,json=voiceConfig,proto3" json:"voice_config,omitempty"`
	// *Required*. The configuration of the synthesized audio.
	AudioConfig          *AudioConfig `protobuf:"bytes,2,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SynthesizeConfig) Reset()         { *m = SynthesizeConfig{} }
func (m *SynthesizeConfig) String() string { return proto.CompactTextString(m) }
func (*SynthesizeConfig) ProtoMessage()    {}
func (*SynthesizeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{1}
}

func (m *SynthesizeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeConfig.Unmarshal(m, b)
}
func (m *SynthesizeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeConfig.Marshal(b, m, deterministic)
}
func (m *SynthesizeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeConfig.Merge(m, src)
}
func (m *SynthesizeConfig) XXX_Size() int {
	return xxx_messageInfo_SynthesizeConfig.Size(m)
}
func (m *SynthesizeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeConfig proto.InternalMessageInfo

func (m *SynthesizeConfig) GetVoiceConfig() *VoiceConfig {
	if m != nil {
		return m.VoiceConfig
	}
	return nil
}

func (m *SynthesizeConfig) GetAudioConfig() *AudioConfig {
	if m != nil {
		return m.AudioConfig
	}
	return nil
}

type SynthesisInput struct {
	// *Required*. The raw text to be synthesized.
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynthesisInput) Reset()         { *m = SynthesisInput{} }
func (m *SynthesisInput) String() string { return proto.CompactTextString(m) }
func (*SynthesisInput) ProtoMessage()    {}
func (*SynthesisInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{2}
}

func (m *SynthesisInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesisInput.Unmarshal(m, b)
}
func (m *SynthesisInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesisInput.Marshal(b, m, deterministic)
}
func (m *SynthesisInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesisInput.Merge(m, src)
}
func (m *SynthesisInput) XXX_Size() int {
	return xxx_messageInfo_SynthesisInput.Size(m)
}
func (m *SynthesisInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesisInput.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesisInput proto.InternalMessageInfo

func (m *SynthesisInput) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type VoiceConfig struct {
	// *Required*. The language code of the voice.
	// Example: "zh-cmn-Hans-CN".
	// See [Language Support](https://docs.speech.sogou.com/docs/tts/resources/languages) for a list of the currently supported language codes.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// *Optional*. The name of the speaker. If not set, the service will choose a
	// voice based on the other parameters such as language_code.
	// Current availabe speakers: [male, female]
	Speaker              string   `protobuf:"bytes,2,opt,name=speaker,proto3" json:"speaker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoiceConfig) Reset()         { *m = VoiceConfig{} }
func (m *VoiceConfig) String() string { return proto.CompactTextString(m) }
func (*VoiceConfig) ProtoMessage()    {}
func (*VoiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{3}
}

func (m *VoiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoiceConfig.Unmarshal(m, b)
}
func (m *VoiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoiceConfig.Marshal(b, m, deterministic)
}
func (m *VoiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoiceConfig.Merge(m, src)
}
func (m *VoiceConfig) XXX_Size() int {
	return xxx_messageInfo_VoiceConfig.Size(m)
}
func (m *VoiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VoiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VoiceConfig proto.InternalMessageInfo

func (m *VoiceConfig) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *VoiceConfig) GetSpeaker() string {
	if m != nil {
		return m.Speaker
	}
	return ""
}

type AudioConfig struct {
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
	Volume               float64  `protobuf:"fixed64,4,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AudioConfig) Reset()         { *m = AudioConfig{} }
func (m *AudioConfig) String() string { return proto.CompactTextString(m) }
func (*AudioConfig) ProtoMessage()    {}
func (*AudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{4}
}

func (m *AudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioConfig.Unmarshal(m, b)
}
func (m *AudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioConfig.Marshal(b, m, deterministic)
}
func (m *AudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioConfig.Merge(m, src)
}
func (m *AudioConfig) XXX_Size() int {
	return xxx_messageInfo_AudioConfig.Size(m)
}
func (m *AudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AudioConfig proto.InternalMessageInfo

func (m *AudioConfig) GetAudioEncoding() AudioConfig_AudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return AudioConfig_ENCODING_UNSPECIFIED
}

func (m *AudioConfig) GetSpeakingRate() float64 {
	if m != nil {
		return m.SpeakingRate
	}
	return 0
}

func (m *AudioConfig) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *AudioConfig) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type SynthesizeResponse struct {
	// The audio data bytes encoded as specified in the request, including the
	// header (For LINEAR16 audio, we include the WAV header). Note: as
	// with all bytes fields, protobuffers use a pure binary representation,
	// whereas JSON representations use base64.
	AudioContent         []byte   `protobuf:"bytes,2,opt,name=audio_content,json=audioContent,proto3" json:"audio_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynthesizeResponse) Reset()         { *m = SynthesizeResponse{} }
func (m *SynthesizeResponse) String() string { return proto.CompactTextString(m) }
func (*SynthesizeResponse) ProtoMessage()    {}
func (*SynthesizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28cfaee149f7a2, []int{5}
}

func (m *SynthesizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeResponse.Unmarshal(m, b)
}
func (m *SynthesizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeResponse.Marshal(b, m, deterministic)
}
func (m *SynthesizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeResponse.Merge(m, src)
}
func (m *SynthesizeResponse) XXX_Size() int {
	return xxx_messageInfo_SynthesizeResponse.Size(m)
}
func (m *SynthesizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeResponse proto.InternalMessageInfo

func (m *SynthesizeResponse) GetAudioContent() []byte {
	if m != nil {
		return m.AudioContent
	}
	return nil
}

func init() {
	proto.RegisterEnum("sogou.speech.tts.v1.AudioConfig_AudioEncoding", AudioConfig_AudioEncoding_name, AudioConfig_AudioEncoding_value)
	proto.RegisterType((*SynthesizeRequest)(nil), "sogou.speech.tts.v1.SynthesizeRequest")
	proto.RegisterType((*SynthesizeConfig)(nil), "sogou.speech.tts.v1.SynthesizeConfig")
	proto.RegisterType((*SynthesisInput)(nil), "sogou.speech.tts.v1.SynthesisInput")
	proto.RegisterType((*VoiceConfig)(nil), "sogou.speech.tts.v1.VoiceConfig")
	proto.RegisterType((*AudioConfig)(nil), "sogou.speech.tts.v1.AudioConfig")
	proto.RegisterType((*SynthesizeResponse)(nil), "sogou.speech.tts.v1.SynthesizeResponse")
}

func init() { proto.RegisterFile("sogou/speech/tts/v1/tts.proto", fileDescriptor_8e28cfaee149f7a2) }

var fileDescriptor_8e28cfaee149f7a2 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x8e, 0xd2, 0x40,
	0x18, 0x77, 0x60, 0xff, 0x7e, 0x14, 0x82, 0xc3, 0x46, 0x1b, 0x13, 0x93, 0x4d, 0x57, 0x5d, 0x4f,
	0x45, 0xd8, 0xc4, 0x64, 0x63, 0x3c, 0x00, 0x8b, 0xa6, 0xc9, 0x2e, 0x92, 0x69, 0x97, 0x83, 0xc1,
	0x90, 0x5a, 0xc6, 0xd2, 0xb8, 0x74, 0x2a, 0x33, 0x25, 0xba, 0x47, 0x6f, 0x1e, 0x7c, 0x03, 0x4f,
	0x1e, 0x7d, 0x04, 0x1f, 0xc1, 0x83, 0xcf, 0xe3, 0xd1, 0xcc, 0x4c, 0x2b, 0x60, 0x88, 0x9c, 0x3c,
	0x75, 0xbe, 0xef, 0xf7, 0x67, 0x7e, 0x5f, 0x3b, 0x1d, 0xb8, 0xcb, 0x59, 0xc8, 0xd2, 0x3a, 0x4f,
	0x28, 0x0d, 0x26, 0x75, 0x21, 0x78, 0x7d, 0xde, 0x90, 0x0f, 0x3b, 0x99, 0x31, 0xc1, 0x70, 0x4d,
	0xc1, 0xb6, 0x86, 0x6d, 0xd9, 0x9f, 0x37, 0xac, 0xcf, 0x08, 0x6e, 0xba, 0x1f, 0x62, 0x31, 0xa1,
	0x3c, 0xba, 0xa6, 0x84, 0xbe, 0x4b, 0x29, 0x17, 0xf8, 0x14, 0xb6, 0xa3, 0x38, 0x49, 0x85, 0x89,
	0x0e, 0xd1, 0xc3, 0x52, 0xf3, 0xc8, 0x5e, 0x23, 0xb5, 0x73, 0x19, 0x77, 0x24, 0x95, 0x68, 0x05,
	0x7e, 0x0a, 0x3b, 0x01, 0x8b, 0xdf, 0x44, 0xa1, 0x59, 0x50, 0xda, 0xfb, 0xff, 0xd4, 0x5e, 0xd3,
	0x8e, 0x22, 0x93, 0x4c, 0x64, 0x7d, 0x41, 0x50, 0xfd, 0x1b, 0xc4, 0x1d, 0x30, 0xe6, 0x2c, 0x0a,
	0xe8, 0x28, 0x73, 0xd6, 0xa9, 0x0e, 0xd7, 0x3a, 0x0f, 0x24, 0x31, 0x33, 0x2d, 0xcd, 0x17, 0x85,
	0x34, 0xf1, 0xd3, 0x71, 0xc4, 0x46, 0x2b, 0xf1, 0xd6, 0x9b, 0xb4, 0x24, 0x31, 0x37, 0xf1, 0x17,
	0x85, 0x75, 0x0f, 0x2a, 0xab, 0x63, 0x63, 0x0c, 0x5b, 0x82, 0xbe, 0xd7, 0x6f, 0x6a, 0x9f, 0xa8,
	0xb5, 0x75, 0x0e, 0xa5, 0xa5, 0x18, 0xf8, 0x08, 0xca, 0x57, 0x7e, 0x1c, 0xa6, 0x7e, 0x28, 0x27,
	0x18, 0xd3, 0x8c, 0x6b, 0xe4, 0xcd, 0x0e, 0x1b, 0x53, 0x6c, 0xc2, 0x2e, 0x4f, 0xa8, 0xff, 0x96,
	0xce, 0x54, 0xb2, 0x7d, 0x92, 0x97, 0xd6, 0xc7, 0x02, 0x94, 0x96, 0x02, 0xe1, 0x4b, 0xa8, 0xe8,
	0x41, 0x68, 0x1c, 0xb0, 0x71, 0x14, 0xeb, 0xf7, 0x51, 0x69, 0xda, 0x9b, 0x46, 0xd1, 0xeb, 0x6e,
	0xa6, 0x22, 0x65, 0x7f, 0xb9, 0x94, 0x29, 0xd5, 0x8e, 0x51, 0x1c, 0x8e, 0x66, 0xbe, 0xa0, 0x2a,
	0x06, 0x22, 0x46, 0xde, 0x24, 0xbe, 0xa0, 0xf8, 0x00, 0xb6, 0x93, 0x48, 0x04, 0x13, 0xb3, 0xa8,
	0x40, 0x5d, 0xe0, 0x5b, 0xb0, 0x33, 0x67, 0x57, 0xe9, 0x94, 0x9a, 0x5b, 0xaa, 0x9d, 0x55, 0x96,
	0x03, 0xe5, 0x95, 0x2d, 0xb1, 0x09, 0x07, 0xdd, 0x5e, 0xe7, 0xc5, 0x99, 0xd3, 0x7b, 0x3e, 0xba,
	0xec, 0xb9, 0xfd, 0x6e, 0xc7, 0x79, 0xe6, 0x74, 0xcf, 0xaa, 0x37, 0xb0, 0x01, 0x7b, 0xe7, 0x4e,
	0xaf, 0xdb, 0x22, 0x8d, 0xc7, 0x55, 0x84, 0x77, 0xa1, 0x78, 0xd1, 0x3f, 0xa9, 0x16, 0xe4, 0xa2,
	0x75, 0x41, 0xaa, 0x45, 0xeb, 0x14, 0xf0, 0xf2, 0x31, 0xe5, 0x09, 0x8b, 0x39, 0x95, 0x99, 0xff,
	0x7c, 0x53, 0x41, 0x63, 0xa1, 0x32, 0x1b, 0xc4, 0xc8, 0x3f, 0x99, 0xec, 0x35, 0x7f, 0x22, 0x28,
	0x0a, 0xc1, 0xf1, 0x2b, 0x80, 0x85, 0x05, 0x7e, 0xb0, 0xe1, 0x5c, 0x66, 0xbf, 0xc2, 0x9d, 0xe3,
	0x8d, 0xbc, 0x2c, 0xcb, 0x04, 0x6a, 0xae, 0x98, 0x51, 0x7f, 0x1a, 0xc5, 0xe1, 0x7f, 0xdc, 0xe7,
	0x11, 0x6a, 0x7f, 0x42, 0x70, 0x3b, 0x60, 0xd3, 0x75, 0x82, 0xf6, 0x9e, 0xe7, 0xb9, 0x7d, 0xf9,
	0xbb, 0xf7, 0xd1, 0xcb, 0xe3, 0x90, 0xc9, 0x23, 0x96, 0x33, 0x34, 0x3d, 0x60, 0xd3, 0xba, 0x9f,
	0x44, 0x3c, 0xbb, 0x1f, 0x9e, 0x08, 0xc1, 0x7f, 0x21, 0xf4, 0xb5, 0x50, 0x74, 0xfb, 0xed, 0x6f,
	0x85, 0x9a, 0xab, 0x48, 0xae, 0x56, 0x78, 0x9e, 0x6b, 0x0f, 0x1a, 0xdf, 0x15, 0xf6, 0x23, 0xc3,
	0x86, 0x1a, 0x1b, 0x7a, 0x9e, 0x3b, 0x1c, 0x34, 0x5e, 0xef, 0xa8, 0xbb, 0xe5, 0xe4, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0xb8, 0x7c, 0x83, 0x7c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TtsClient is the client API for Tts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TtsClient interface {
	// Synthesizes speech synchronously: receive results after all text input
	// has been processed.
	Synthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (*SynthesizeResponse, error)
	// Performs unidirectional streaming synthesizes
	StreamingSynthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (Tts_StreamingSynthesizeClient, error)
}

type ttsClient struct {
	cc *grpc.ClientConn
}

func NewTtsClient(cc *grpc.ClientConn) TtsClient {
	return &ttsClient{cc}
}

func (c *ttsClient) Synthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (*SynthesizeResponse, error) {
	out := new(SynthesizeResponse)
	err := c.cc.Invoke(ctx, "/sogou.speech.tts.v1.tts/Synthesize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ttsClient) StreamingSynthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (Tts_StreamingSynthesizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tts_serviceDesc.Streams[0], "/sogou.speech.tts.v1.tts/StreamingSynthesize", opts...)
	if err != nil {
		return nil, err
	}
	x := &ttsStreamingSynthesizeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tts_StreamingSynthesizeClient interface {
	Recv() (*SynthesizeResponse, error)
	grpc.ClientStream
}

type ttsStreamingSynthesizeClient struct {
	grpc.ClientStream
}

func (x *ttsStreamingSynthesizeClient) Recv() (*SynthesizeResponse, error) {
	m := new(SynthesizeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TtsServer is the server API for Tts service.
type TtsServer interface {
	// Synthesizes speech synchronously: receive results after all text input
	// has been processed.
	Synthesize(context.Context, *SynthesizeRequest) (*SynthesizeResponse, error)
	// Performs unidirectional streaming synthesizes
	StreamingSynthesize(*SynthesizeRequest, Tts_StreamingSynthesizeServer) error
}

func RegisterTtsServer(s *grpc.Server, srv TtsServer) {
	s.RegisterService(&_Tts_serviceDesc, srv)
}

func _Tts_Synthesize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynthesizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TtsServer).Synthesize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sogou.speech.tts.v1.tts/Synthesize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TtsServer).Synthesize(ctx, req.(*SynthesizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tts_StreamingSynthesize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SynthesizeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TtsServer).StreamingSynthesize(m, &ttsStreamingSynthesizeServer{stream})
}

type Tts_StreamingSynthesizeServer interface {
	Send(*SynthesizeResponse) error
	grpc.ServerStream
}

type ttsStreamingSynthesizeServer struct {
	grpc.ServerStream
}

func (x *ttsStreamingSynthesizeServer) Send(m *SynthesizeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Tts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sogou.speech.tts.v1.tts",
	HandlerType: (*TtsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Synthesize",
			Handler:    _Tts_Synthesize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingSynthesize",
			Handler:       _Tts_StreamingSynthesize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sogou/speech/tts/v1/tts.proto",
}
