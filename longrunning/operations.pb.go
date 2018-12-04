// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sogou/speech/longrunning/operations.proto

package longrunning

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	// originally returns it.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *any.Any `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If true, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result               isOperation_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	Response *any.Any `protobuf:"bytes,5,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetError() *status.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *any.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

// The request message for [Operations.GetOperation][sogou.speech.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{1}
}

func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(m, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.ListOperations][sogou.speech.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsRequest) Reset()         { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()    {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{2}
}

func (m *ListOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsRequest.Unmarshal(m, b)
}
func (m *ListOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsRequest.Marshal(b, m, deterministic)
}
func (m *ListOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsRequest.Merge(m, src)
}
func (m *ListOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListOperationsRequest.Size(m)
}
func (m *ListOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsRequest proto.InternalMessageInfo

func (m *ListOperationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListOperationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListOperationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Operations.ListOperations][sogou.speech.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that matches the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsResponse) Reset()         { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()    {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{3}
}

func (m *ListOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsResponse.Unmarshal(m, b)
}
func (m *ListOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsResponse.Marshal(b, m, deterministic)
}
func (m *ListOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsResponse.Merge(m, src)
}
func (m *ListOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListOperationsResponse.Size(m)
}
func (m *ListOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsResponse proto.InternalMessageInfo

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Operations.CancelOperation][sogou.speech.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOperationRequest) Reset()         { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()    {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{4}
}

func (m *CancelOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOperationRequest.Unmarshal(m, b)
}
func (m *CancelOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOperationRequest.Marshal(b, m, deterministic)
}
func (m *CancelOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOperationRequest.Merge(m, src)
}
func (m *CancelOperationRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOperationRequest.Size(m)
}
func (m *CancelOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOperationRequest proto.InternalMessageInfo

func (m *CancelOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.DeleteOperation][sogou.speech.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationRequest) Reset()         { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()    {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721b8ae654c0fce, []int{5}
}

func (m *DeleteOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationRequest.Unmarshal(m, b)
}
func (m *DeleteOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationRequest.Merge(m, src)
}
func (m *DeleteOperationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationRequest.Size(m)
}
func (m *DeleteOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationRequest proto.InternalMessageInfo

func (m *DeleteOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Operation)(nil), "sogou.speech.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "sogou.speech.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "sogou.speech.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "sogou.speech.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "sogou.speech.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "sogou.speech.longrunning.DeleteOperationRequest")
}

func init() {
	proto.RegisterFile("sogou/speech/longrunning/operations.proto", fileDescriptor_6721b8ae654c0fce)
}

var fileDescriptor_6721b8ae654c0fce = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0xe6, 0x4b, 0xc9, 0x14, 0x88, 0xb4, 0x40, 0x30, 0x29, 0x48, 0x51, 0x90, 0x50, 0x8a,
	0x8a, 0x5d, 0x05, 0x6e, 0x9c, 0x48, 0x41, 0xf4, 0x50, 0x89, 0xc8, 0xe1, 0x04, 0x95, 0xaa, 0x8d,
	0x3b, 0x35, 0x16, 0xce, 0xae, 0xd9, 0x5d, 0x4b, 0xb4, 0x17, 0x4e, 0xfc, 0x0e, 0x38, 0xf3, 0x53,
	0xb8, 0xf0, 0x77, 0x38, 0x22, 0xaf, 0x9d, 0xd8, 0x18, 0xbb, 0xca, 0x6d, 0xb3, 0xef, 0xcd, 0x9b,
	0x79, 0x6f, 0x27, 0x86, 0x7d, 0x25, 0x7c, 0x11, 0x3b, 0x2a, 0x42, 0xf4, 0x3e, 0x3a, 0xa1, 0xe0,
	0xbe, 0x8c, 0x39, 0x0f, 0xb8, 0xef, 0x88, 0x08, 0x25, 0xd3, 0x81, 0xe0, 0xca, 0x8e, 0xa4, 0xd0,
	0x82, 0x5a, 0x86, 0x6a, 0xa7, 0x54, 0xbb, 0x40, 0x1d, 0xde, 0xf7, 0x85, 0xf0, 0x43, 0x74, 0x0c,
	0x6f, 0x19, 0x5f, 0x38, 0x8c, 0x5f, 0xa6, 0x45, 0xc3, 0xbd, 0x32, 0x84, 0xab, 0x48, 0xaf, 0xc1,
	0x7b, 0x19, 0x28, 0x23, 0xcf, 0x51, 0x9a, 0xe9, 0x38, 0x6b, 0x35, 0xfe, 0x4d, 0xa0, 0xf7, 0x76,
	0xdd, 0x9f, 0x52, 0x68, 0x71, 0xb6, 0x42, 0x8b, 0x8c, 0xc8, 0xa4, 0xe7, 0x9a, 0x33, 0x3d, 0x84,
	0xee, 0x0a, 0x35, 0x3b, 0x67, 0x9a, 0x59, 0x8d, 0x11, 0x99, 0xec, 0x4e, 0xef, 0xd8, 0xa9, 0x9a,
	0xbd, 0x6e, 0x65, 0xbf, 0xe4, 0x97, 0xee, 0x86, 0x95, 0xa8, 0x9c, 0x0b, 0x8e, 0x56, 0x73, 0x44,
	0x26, 0x5d, 0xd7, 0x9c, 0xe9, 0x13, 0x68, 0xa3, 0x94, 0x42, 0x5a, 0x2d, 0x23, 0x41, 0xd7, 0x12,
	0x32, 0xf2, 0xec, 0x85, 0x19, 0xe8, 0x78, 0xc7, 0x4d, 0x29, 0x74, 0x0a, 0x5d, 0x89, 0x2a, 0x12,
	0x5c, 0xa1, 0xd5, 0xae, 0xef, 0x78, 0xbc, 0xe3, 0x6e, 0x78, 0xb3, 0x2e, 0x74, 0x24, 0xaa, 0x38,
	0xd4, 0xe3, 0x7d, 0xb8, 0xfd, 0x06, 0xf5, 0xc6, 0x93, 0x8b, 0x9f, 0x63, 0x54, 0xba, 0xca, 0xda,
	0xf8, 0x2b, 0xdc, 0x3d, 0x09, 0x54, 0xce, 0x55, 0x65, 0x72, 0xab, 0x90, 0xc3, 0x00, 0x3a, 0x17,
	0x41, 0xa8, 0x51, 0x66, 0x12, 0xd9, 0x2f, 0xba, 0x07, 0xbd, 0x88, 0xf9, 0x78, 0xa6, 0x82, 0x2b,
	0x34, 0x01, 0xb5, 0xdd, 0x6e, 0x72, 0xb1, 0x08, 0xae, 0x90, 0x3e, 0x04, 0x30, 0xa0, 0x16, 0x9f,
	0x90, 0x9b, 0x40, 0x7a, 0xae, 0xa1, 0xbf, 0x4b, 0x2e, 0xc6, 0xdf, 0x08, 0x0c, 0xca, 0x13, 0xa4,
	0x86, 0xe8, 0x11, 0x40, 0xbe, 0x17, 0x16, 0x19, 0x35, 0x27, 0xbb, 0xd3, 0x47, 0x76, 0xdd, 0x62,
	0xd8, 0xb9, 0xdf, 0x42, 0x19, 0x7d, 0x0c, 0x7d, 0x8e, 0x5f, 0xf4, 0x59, 0x61, 0x86, 0x86, 0x99,
	0xe1, 0x66, 0x72, 0x3d, 0xdf, 0xcc, 0x71, 0x00, 0x83, 0x23, 0xc6, 0x3d, 0x0c, 0xb7, 0x8a, 0xed,
	0x00, 0x06, 0xaf, 0x30, 0x44, 0x8d, 0xdb, 0xb0, 0xa7, 0xdf, 0x9b, 0x00, 0xb9, 0x3f, 0xaa, 0xe0,
	0xd6, 0xbf, 0x8e, 0xa9, 0x53, 0xef, 0xaa, 0xf2, 0x75, 0x86, 0x87, 0xdb, 0x17, 0x64, 0x61, 0x2e,
	0xe1, 0x46, 0x71, 0x27, 0xe8, 0xd3, 0x7a, 0x85, 0x8a, 0xdd, 0x19, 0x6e, 0x93, 0x3b, 0xfd, 0x00,
	0xfd, 0x52, 0x2a, 0xf4, 0x9a, 0x41, 0xab, 0x03, 0x1c, 0x0e, 0xfe, 0x5b, 0xf4, 0xd7, 0xc9, 0xbf,
	0x38, 0x11, 0x2f, 0x3d, 0xd0, 0x75, 0xe2, 0xd5, 0x6f, 0x59, 0x27, 0x3e, 0xfb, 0x41, 0xe0, 0x81,
	0x27, 0x56, 0xb5, 0x7a, 0xb3, 0x7e, 0x1e, 0xe9, 0x3c, 0x29, 0x9d, 0x93, 0xf7, 0xcf, 0x7d, 0x11,
	0x32, 0xee, 0xaf, 0xd9, 0x69, 0xa9, 0x27, 0x56, 0x0e, 0x8b, 0x02, 0x55, 0xfc, 0xba, 0xbd, 0x28,
	0x9c, 0xff, 0x10, 0xf2, 0xb3, 0x61, 0x2d, 0x0c, 0x77, 0x91, 0x16, 0x9e, 0x08, 0xee, 0xbb, 0x29,
	0xfa, 0x2b, 0x83, 0x4e, 0x53, 0xe8, 0xb4, 0x00, 0x2d, 0x3b, 0x66, 0xe4, 0x67, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x2c, 0x15, 0x44, 0x45, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationsClient is the client API for Operations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationsClient interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][sogou.speech.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][sogou.speech.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := c.cc.Invoke(ctx, "/sogou.speech.longrunning.Operations/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/sogou.speech.longrunning.Operations/GetOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/sogou.speech.longrunning.Operations/DeleteOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/sogou.speech.longrunning.Operations/CancelOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServer is the server API for Operations service.
type OperationsServer interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(context.Context, *GetOperationRequest) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(context.Context, *DeleteOperationRequest) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][sogou.speech.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][sogou.speech.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(context.Context, *CancelOperationRequest) (*empty.Empty, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sogou.speech.longrunning.Operations/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sogou.speech.longrunning.Operations/GetOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetOperation(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sogou.speech.longrunning.Operations/DeleteOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sogou.speech.longrunning.Operations/CancelOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sogou.speech.longrunning.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOperations",
			Handler:    _Operations_ListOperations_Handler,
		},
		{
			MethodName: "GetOperation",
			Handler:    _Operations_GetOperation_Handler,
		},
		{
			MethodName: "DeleteOperation",
			Handler:    _Operations_DeleteOperation_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _Operations_CancelOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sogou/speech/longrunning/operations.proto",
}
