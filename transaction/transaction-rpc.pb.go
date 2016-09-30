// Code generated by protoc-gen-go.
// source: transaction/transaction-rpc.proto
// DO NOT EDIT!

/*
Package transaction is a generated protocol buffer package.

It is generated from these files:
	transaction/transaction-rpc.proto

It has these top-level messages:
	TxnStepReq
	TxnStepResp
*/
package transaction

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxnStepReq struct {
	StepFunc string `protobuf:"bytes,1,opt,name=StepFunc,json=stepFunc" json:"StepFunc,omitempty"`
	Context  []byte `protobuf:"bytes,2,opt,name=Context,json=context,proto3" json:"Context,omitempty"`
}

func (m *TxnStepReq) Reset()                    { *m = TxnStepReq{} }
func (m *TxnStepReq) String() string            { return proto.CompactTextString(m) }
func (*TxnStepReq) ProtoMessage()               {}
func (*TxnStepReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TxnStepResp struct {
	Error string `protobuf:"bytes,1,opt,name=Error,json=error" json:"Error,omitempty"`
	Resp  []byte `protobuf:"bytes,2,opt,name=Resp,json=resp,proto3" json:"Resp,omitempty"`
}

func (m *TxnStepResp) Reset()                    { *m = TxnStepResp{} }
func (m *TxnStepResp) String() string            { return proto.CompactTextString(m) }
func (*TxnStepResp) ProtoMessage()               {}
func (*TxnStepResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TxnStepReq)(nil), "transaction.TxnStepReq")
	proto.RegisterType((*TxnStepResp)(nil), "transaction.TxnStepResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for TxnSvc service

type TxnSvcClient interface {
	RunStep(ctx context.Context, in *TxnStepReq, opts ...grpc.CallOption) (*TxnStepResp, error)
}

type txnSvcClient struct {
	cc *grpc.ClientConn
}

func NewTxnSvcClient(cc *grpc.ClientConn) TxnSvcClient {
	return &txnSvcClient{cc}
}

func (c *txnSvcClient) RunStep(ctx context.Context, in *TxnStepReq, opts ...grpc.CallOption) (*TxnStepResp, error) {
	out := new(TxnStepResp)
	err := grpc.Invoke(ctx, "/transaction.TxnSvc/RunStep", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TxnSvc service

type TxnSvcServer interface {
	RunStep(context.Context, *TxnStepReq) (*TxnStepResp, error)
}

func RegisterTxnSvcServer(s *grpc.Server, srv TxnSvcServer) {
	s.RegisterService(&_TxnSvc_serviceDesc, srv)
}

func _TxnSvc_RunStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnStepReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxnSvcServer).RunStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TxnSvc/RunStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxnSvcServer).RunStep(ctx, req.(*TxnStepReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TxnSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TxnSvc",
	HandlerType: (*TxnSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunStep",
			Handler:    _TxnSvc_RunStep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("transaction/transaction-rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x47, 0x62, 0xeb, 0x16, 0x15, 0x24, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x23, 0x09, 0x2b, 0x39, 0x71, 0x71, 0x85, 0x54, 0xe4, 0x05,
	0x97, 0xa4, 0x16, 0x04, 0xa5, 0x16, 0x0a, 0x49, 0x71, 0x71, 0x80, 0x98, 0x6e, 0xa5, 0x79, 0xc9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0xc5, 0x50, 0xbe, 0x90, 0x04, 0x17, 0xbb, 0x73,
	0x7e, 0x5e, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x13, 0x50, 0x8a, 0x27, 0x88, 0x3d, 0x19, 0xc2, 0x55,
	0x32, 0xe7, 0xe2, 0x86, 0x9b, 0x51, 0x5c, 0x20, 0x24, 0xc2, 0xc5, 0xea, 0x5a, 0x54, 0x94, 0x5f,
	0x04, 0x35, 0x81, 0x35, 0x15, 0xc4, 0x11, 0x12, 0xe2, 0x62, 0x01, 0xc9, 0x42, 0xf5, 0xb2, 0x14,
	0x01, 0xd9, 0x46, 0x1e, 0x5c, 0x6c, 0x20, 0x8d, 0x65, 0xc9, 0x42, 0x76, 0x5c, 0xec, 0x41, 0xa5,
	0x60, 0x23, 0x84, 0xc4, 0xf5, 0x90, 0xdc, 0xa7, 0x87, 0x70, 0x9c, 0x94, 0x04, 0x76, 0x89, 0xe2,
	0x02, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xd7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xcf,
	0x8a, 0xfd, 0xff, 0x00, 0x00, 0x00,
}