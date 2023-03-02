// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icaauth/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgRegisterAccount defines the request message for MsgRegisterAccount
type MsgRegisterAccount struct {
	// owner represents the owner of the interchain account
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// connectionId represents the IBC `connectionId` of the host chain
	ConnectionId string `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty" yaml:"connection_id"`
	// version represents the version of the ICA channel
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *MsgRegisterAccount) Reset()         { *m = MsgRegisterAccount{} }
func (m *MsgRegisterAccount) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccount) ProtoMessage()    {}
func (*MsgRegisterAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf995a0c8f8b8c, []int{0}
}
func (m *MsgRegisterAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccount.Merge(m, src)
}
func (m *MsgRegisterAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccount proto.InternalMessageInfo

func (m *MsgRegisterAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgRegisterAccount) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *MsgRegisterAccount) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// MsgRegisterAccountResponse defines the response message for MsgRegisterAccount
type MsgRegisterAccountResponse struct {
}

func (m *MsgRegisterAccountResponse) Reset()         { *m = MsgRegisterAccountResponse{} }
func (m *MsgRegisterAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccountResponse) ProtoMessage()    {}
func (*MsgRegisterAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf995a0c8f8b8c, []int{1}
}
func (m *MsgRegisterAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccountResponse.Merge(m, src)
}
func (m *MsgRegisterAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccountResponse proto.InternalMessageInfo

// MsgSubmitTx defines the request message for MsgSubmitTx
type MsgSubmitTx struct {
	// owner represents the owner of the interchain account
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// connectionId represents the IBC `connectionId` of the host chain
	ConnectionId string `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty" yaml:"connection_id"`
	// msgs represents the transactions to be submitted to the host chain
	Msgs []*types.Any `protobuf:"bytes,3,rep,name=msgs,proto3" json:"msgs,omitempty"`
	// timeoutDuration represents the timeout duration for the IBC packet from last block
	TimeoutDuration *time.Duration `protobuf:"bytes,4,opt,name=timeoutDuration,proto3,stdduration" json:"timeoutDuration,omitempty"`
}

func (m *MsgSubmitTx) Reset()         { *m = MsgSubmitTx{} }
func (m *MsgSubmitTx) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTx) ProtoMessage()    {}
func (*MsgSubmitTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf995a0c8f8b8c, []int{2}
}
func (m *MsgSubmitTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTx.Merge(m, src)
}
func (m *MsgSubmitTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTx proto.InternalMessageInfo

func (m *MsgSubmitTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgSubmitTx) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *MsgSubmitTx) GetMsgs() []*types.Any {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func (m *MsgSubmitTx) GetTimeoutDuration() *time.Duration {
	if m != nil {
		return m.TimeoutDuration
	}
	return nil
}

// MsgSubmitTxResponse defines the response message for MsgSubmitTx
type MsgSubmitTxResponse struct {
}

func (m *MsgSubmitTxResponse) Reset()         { *m = MsgSubmitTxResponse{} }
func (m *MsgSubmitTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTxResponse) ProtoMessage()    {}
func (*MsgSubmitTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf995a0c8f8b8c, []int{3}
}
func (m *MsgSubmitTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTxResponse.Merge(m, src)
}
func (m *MsgSubmitTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterAccount)(nil), "chainmain.icaauth.v1.MsgRegisterAccount")
	proto.RegisterType((*MsgRegisterAccountResponse)(nil), "chainmain.icaauth.v1.MsgRegisterAccountResponse")
	proto.RegisterType((*MsgSubmitTx)(nil), "chainmain.icaauth.v1.MsgSubmitTx")
	proto.RegisterType((*MsgSubmitTxResponse)(nil), "chainmain.icaauth.v1.MsgSubmitTxResponse")
}

func init() { proto.RegisterFile("icaauth/v1/tx.proto", fileDescriptor_09bf995a0c8f8b8c) }

var fileDescriptor_09bf995a0c8f8b8c = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbb, 0x8e, 0xd3, 0x40,
	0x14, 0xcd, 0x90, 0xc0, 0xc2, 0x04, 0x69, 0x25, 0xaf, 0x91, 0x1c, 0x0b, 0x79, 0x83, 0xab, 0x50,
	0x64, 0x86, 0x0d, 0x82, 0x02, 0xd1, 0x6c, 0x44, 0xb3, 0x45, 0x24, 0x64, 0x28, 0x10, 0xcd, 0xca,
	0x99, 0x0c, 0x93, 0x11, 0x3b, 0x73, 0x23, 0xcf, 0x38, 0xc4, 0x2d, 0x5f, 0x40, 0xc9, 0x87, 0xf0,
	0x11, 0x88, 0x02, 0x6d, 0x49, 0x05, 0x28, 0xf9, 0x03, 0xbe, 0x00, 0xf9, 0x31, 0x3c, 0x12, 0x10,
	0x34, 0xdb, 0x58, 0x3e, 0x3e, 0xe7, 0x9e, 0x7b, 0xef, 0xf1, 0x0c, 0x3e, 0x90, 0x2c, 0x4d, 0x73,
	0x3b, 0xa7, 0xcb, 0x23, 0x6a, 0x57, 0x64, 0x91, 0x81, 0x05, 0xcf, 0x67, 0xf3, 0x54, 0x6a, 0x95,
	0x4a, 0x4d, 0x1a, 0x9a, 0x2c, 0x8f, 0xc2, 0x1e, 0x03, 0xa3, 0xc0, 0x9c, 0x56, 0x1a, 0x5a, 0x83,
	0xba, 0x20, 0xec, 0x09, 0x00, 0x71, 0xc6, 0x69, 0x85, 0xa6, 0xf9, 0x0b, 0x9a, 0xea, 0xa2, 0xa1,
	0xa2, 0x6d, 0x6a, 0x96, 0x67, 0xa9, 0x95, 0xa0, 0x1b, 0xde, 0x17, 0x20, 0xa0, 0xb6, 0x2c, 0xdf,
	0xea, 0xaf, 0xf1, 0x6b, 0x84, 0xbd, 0x89, 0x11, 0x09, 0x17, 0xd2, 0x58, 0x9e, 0x1d, 0x33, 0x06,
	0xb9, 0xb6, 0x9e, 0x8f, 0x2f, 0xc3, 0x2b, 0xcd, 0xb3, 0x00, 0xf5, 0xd1, 0xe0, 0x5a, 0x52, 0x03,
	0xef, 0x21, 0xbe, 0xce, 0x40, 0x6b, 0xce, 0x4a, 0xdb, 0x93, 0x59, 0x70, 0xa9, 0x24, 0xc7, 0xc1,
	0xb7, 0xcf, 0x87, 0x7e, 0x91, 0xaa, 0xb3, 0x07, 0xf1, 0x4f, 0xf6, 0x54, 0xce, 0xe2, 0xe4, 0x37,
	0xb5, 0x17, 0xe0, 0xbd, 0x25, 0xcf, 0x8c, 0x04, 0x1d, 0xb4, 0x2b, 0x57, 0x07, 0xe3, 0x9b, 0x38,
	0xdc, 0x9d, 0x21, 0xe1, 0x66, 0x01, 0xda, 0xf0, 0x78, 0x83, 0x70, 0x77, 0x62, 0xc4, 0x93, 0x7c,
	0xaa, 0xa4, 0x7d, 0xba, 0xba, 0x90, 0xd9, 0xee, 0xe1, 0x8e, 0x32, 0xc2, 0x04, 0xed, 0x7e, 0x7b,
	0xd0, 0x1d, 0xf9, 0xa4, 0xce, 0x92, 0xb8, 0x2c, 0xc9, 0xb1, 0x2e, 0xc6, 0xdd, 0x0f, 0xef, 0x86,
	0x7b, 0x66, 0xf6, 0x92, 0x94, 0xb3, 0x56, 0x72, 0xef, 0x04, 0xef, 0x5b, 0xa9, 0x38, 0xe4, 0xf6,
	0x51, 0x13, 0x76, 0xd0, 0xe9, 0xa3, 0x41, 0x77, 0xd4, 0xdb, 0x71, 0x70, 0x82, 0x71, 0xe7, 0xed,
	0x97, 0x43, 0x94, 0x6c, 0xd7, 0xc5, 0x37, 0xf0, 0xc1, 0x2f, 0x4b, 0xba, 0xe5, 0x47, 0x1f, 0x11,
	0x6e, 0x4f, 0x8c, 0xf0, 0x14, 0xde, 0xdf, 0xfe, 0x47, 0x03, 0xf2, 0xa7, 0xd3, 0x43, 0x76, 0x93,
	0x0c, 0xef, 0xfc, 0xaf, 0xd2, 0xb5, 0xf5, 0x9e, 0xe1, 0xab, 0x3f, 0xf2, 0xbe, 0xf5, 0xd7, 0x6a,
	0x27, 0x09, 0x6f, 0xff, 0x53, 0xe2, 0x9c, 0xc7, 0x8f, 0xdf, 0xaf, 0x23, 0x74, 0xbe, 0x8e, 0xd0,
	0xd7, 0x75, 0x84, 0xde, 0x6c, 0xa2, 0xd6, 0xf9, 0x26, 0x6a, 0x7d, 0xda, 0x44, 0xad, 0xe7, 0xf7,
	0x85, 0xb4, 0xf3, 0x7c, 0x4a, 0x18, 0x28, 0xca, 0xb2, 0x62, 0x61, 0x61, 0x08, 0x99, 0x18, 0x56,
	0xce, 0xb4, 0x7a, 0x0e, 0xcb, 0x06, 0x74, 0x45, 0xdd, 0x55, 0xb2, 0xc5, 0x82, 0x9b, 0xe9, 0x95,
	0x2a, 0xe3, 0xbb, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x2e, 0xcf, 0x3f, 0x62, 0x03, 0x00,
	0x00,
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
	// RegisterAccount registers an interchain account on host chain with given `connectionId`
	RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error)
	// SubmitTx submits a transaction to the host chain on behalf of interchain account
	SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error) {
	out := new(MsgRegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/chainmain.icaauth.v1.Msg/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error) {
	out := new(MsgSubmitTxResponse)
	err := c.cc.Invoke(ctx, "/chainmain.icaauth.v1.Msg/SubmitTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterAccount registers an interchain account on host chain with given `connectionId`
	RegisterAccount(context.Context, *MsgRegisterAccount) (*MsgRegisterAccountResponse, error)
	// SubmitTx submits a transaction to the host chain on behalf of interchain account
	SubmitTx(context.Context, *MsgSubmitTx) (*MsgSubmitTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterAccount(ctx context.Context, req *MsgRegisterAccount) (*MsgRegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (*UnimplementedMsgServer) SubmitTx(ctx context.Context, req *MsgSubmitTx) (*MsgSubmitTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.icaauth.v1.Msg/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAccount(ctx, req.(*MsgRegisterAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.icaauth.v1.Msg/SubmitTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitTx(ctx, req.(*MsgSubmitTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainmain.icaauth.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _Msg_RegisterAccount_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Msg_SubmitTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "icaauth/v1/tx.proto",
}

func (m *MsgRegisterAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutDuration != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.TimeoutDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.TimeoutDuration):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTx(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.TimeoutDuration != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.TimeoutDuration)
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &types.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimeoutDuration == nil {
				m.TimeoutDuration = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.TimeoutDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
