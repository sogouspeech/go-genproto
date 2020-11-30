// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TtsClient is the client API for Tts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TtsClient interface {
	// Synthesizes speech synchronously: receive results after all text input
	// has been processed.
	Synthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (*SynthesizeResponse, error)
	// Performs unidirectional streaming synthesizes
	StreamingSynthesize(ctx context.Context, in *SynthesizeRequest, opts ...grpc.CallOption) (Tts_StreamingSynthesizeClient, error)
}

type ttsClient struct {
	cc grpc.ClientConnInterface
}

func NewTtsClient(cc grpc.ClientConnInterface) TtsClient {
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
// All implementations must embed UnimplementedTtsServer
// for forward compatibility
type TtsServer interface {
	// Synthesizes speech synchronously: receive results after all text input
	// has been processed.
	Synthesize(context.Context, *SynthesizeRequest) (*SynthesizeResponse, error)
	// Performs unidirectional streaming synthesizes
	StreamingSynthesize(*SynthesizeRequest, Tts_StreamingSynthesizeServer) error
	mustEmbedUnimplementedTtsServer()
}

// UnimplementedTtsServer must be embedded to have forward compatible implementations.
type UnimplementedTtsServer struct {
}

func (UnimplementedTtsServer) Synthesize(context.Context, *SynthesizeRequest) (*SynthesizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Synthesize not implemented")
}
func (UnimplementedTtsServer) StreamingSynthesize(*SynthesizeRequest, Tts_StreamingSynthesizeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingSynthesize not implemented")
}
func (UnimplementedTtsServer) mustEmbedUnimplementedTtsServer() {}

// UnsafeTtsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TtsServer will
// result in compilation errors.
type UnsafeTtsServer interface {
	mustEmbedUnimplementedTtsServer()
}

func RegisterTtsServer(s grpc.ServiceRegistrar, srv TtsServer) {
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