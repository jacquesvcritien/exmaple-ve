// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examplevess/examplevess/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("examplevess/examplevess/tx.proto", fileDescriptor_65564e56f1e4083b) }

var fileDescriptor_65564e56f1e4083b = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0x4b, 0x2d, 0x2e, 0xd6, 0x47, 0x66, 0x97, 0x54, 0xe8, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x89, 0x23, 0x89, 0xea, 0x21, 0xb1, 0x8d, 0x58, 0xb9, 0x98, 0x7d, 0x8b, 0xd3,
	0x9d, 0x2c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x1e, 0xd9, 0xbc,
	0x0a, 0x54, 0xd3, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x36, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x1d, 0xca, 0xc5, 0xe8, 0x85, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examplevess.examplevess.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "examplevess/examplevess/tx.proto",
}
