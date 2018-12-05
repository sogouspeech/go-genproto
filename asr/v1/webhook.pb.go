// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sogou/speech/asr/v1/webhook.proto

package asr // import "golang.speech.sogou.com/apis/asr/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config a webhook when speech recoginition reaches a final result.
type RecognitionResultHookConfig struct {
	// this hook name. when hook call complete, the extra field of SpeechRecognitionResult or StreamingRecognitionResult will
	// be filled by a key-value pair, the content is hook result, and key will be set to this name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the hook uri, the SpeechRecognitionResult will be jsonize and post to this uri
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// *Optional* Headers will be set when post
	Headers              map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RecognitionResultHookConfig) Reset()         { *m = RecognitionResultHookConfig{} }
func (m *RecognitionResultHookConfig) String() string { return proto.CompactTextString(m) }
func (*RecognitionResultHookConfig) ProtoMessage()    {}
func (*RecognitionResultHookConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhook_273ff0cbeabf2a88, []int{0}
}
func (m *RecognitionResultHookConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognitionResultHookConfig.Unmarshal(m, b)
}
func (m *RecognitionResultHookConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognitionResultHookConfig.Marshal(b, m, deterministic)
}
func (dst *RecognitionResultHookConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognitionResultHookConfig.Merge(dst, src)
}
func (m *RecognitionResultHookConfig) XXX_Size() int {
	return xxx_messageInfo_RecognitionResultHookConfig.Size(m)
}
func (m *RecognitionResultHookConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognitionResultHookConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RecognitionResultHookConfig proto.InternalMessageInfo

func (m *RecognitionResultHookConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecognitionResultHookConfig) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *RecognitionResultHookConfig) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*RecognitionResultHookConfig)(nil), "sogou.speech.asr.v1.RecognitionResultHookConfig")
	proto.RegisterMapType((map[string]string)(nil), "sogou.speech.asr.v1.RecognitionResultHookConfig.HeadersEntry")
}

func init() {
	proto.RegisterFile("sogou/speech/asr/v1/webhook.proto", fileDescriptor_webhook_273ff0cbeabf2a88)
}

var fileDescriptor_webhook_273ff0cbeabf2a88 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0xce, 0x4f, 0xcf,
	0x2f, 0xd5, 0x2f, 0x2e, 0x48, 0x4d, 0x4d, 0xce, 0xd0, 0x4f, 0x2c, 0x2e, 0xd2, 0x2f, 0x33, 0xd4,
	0x2f, 0x4f, 0x4d, 0xca, 0xc8, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06,
	0x2b, 0xd1, 0x83, 0x28, 0xd1, 0x4b, 0x2c, 0x2e, 0xd2, 0x2b, 0x33, 0x54, 0xba, 0xc1, 0xc8, 0x25,
	0x1d, 0x94, 0x9a, 0x9c, 0x9f, 0x9e, 0x97, 0x59, 0x92, 0x99, 0x9f, 0x17, 0x94, 0x5a, 0x5c, 0x9a,
	0x53, 0xe2, 0x91, 0x9f, 0x9f, 0xed, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x2e, 0x24, 0xc4, 0xc5, 0x92,
	0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x09, 0x70, 0x31,
	0x97, 0x16, 0x65, 0x4a, 0x30, 0x81, 0x85, 0x40, 0x4c, 0xa1, 0x70, 0x2e, 0xf6, 0x8c, 0xd4, 0xc4,
	0x94, 0xd4, 0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x5b, 0x3d, 0x2c, 0x96, 0xe9,
	0xe1, 0xb1, 0x48, 0xcf, 0x03, 0xa2, 0xdf, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0x66, 0x9a, 0x94,
	0x15, 0x17, 0x0f, 0xb2, 0x04, 0xc8, 0xea, 0xec, 0xd4, 0x4a, 0xa8, 0x6b, 0x40, 0x4c, 0x21, 0x11,
	0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0xa8, 0x73, 0x20, 0x1c, 0x2b, 0x26, 0x0b, 0x46, 0xa7,
	0x5e, 0x46, 0x2e, 0xf1, 0xe4, 0xfc, 0x5c, 0x6c, 0x2e, 0x71, 0xe2, 0x09, 0x4f, 0x4d, 0x02, 0x59,
	0x1e, 0x00, 0x0a, 0x99, 0x00, 0xc6, 0x28, 0xf5, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x98, 0x2a,
	0x88, 0x96, 0xe4, 0xfc, 0x5c, 0xfd, 0xc4, 0x82, 0xcc, 0x62, 0x68, 0x70, 0x5a, 0x27, 0x16, 0x17,
	0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0x1c, 0x1c, 0xe0, 0xb4, 0x8a, 0x49, 0x38, 0x18, 0xac, 0x28,
	0x18, 0xa2, 0xc3, 0xb1, 0xb8, 0x48, 0x2f, 0xcc, 0x70, 0x17, 0x58, 0xee, 0x14, 0x54, 0x2e, 0x06,
	0x22, 0x17, 0xe3, 0x58, 0x5c, 0x14, 0x13, 0x66, 0x98, 0xc4, 0x06, 0x8e, 0x06, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0xa1, 0xbf, 0xad, 0xab, 0x01, 0x00, 0x00,
}