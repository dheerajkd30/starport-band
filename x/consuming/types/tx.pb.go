// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consuming/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgCoinRatesData struct {
	Creator        string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OracleScriptID uint64                                   `protobuf:"varint,2,opt,name=oracle_script_id,json=oracleScriptId,proto3" json:"oracle_script_id,omitempty" yaml:"oracle_script_id"`
	SourceChannel  string                                   `protobuf:"bytes,3,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	Calldata       *CoinRatesCallData                       `protobuf:"bytes,4,opt,name=calldata,proto3" json:"calldata,omitempty"`
	AskCount       uint64                                   `protobuf:"varint,5,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	MinCount       uint64                                   `protobuf:"varint,6,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	FeeLimit       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
	RequestKey     string                                   `protobuf:"bytes,8,opt,name=request_key,json=requestKey,proto3" json:"request_key,omitempty"`
	PrepareGas     uint64                                   `protobuf:"varint,9,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty"`
	ExecuteGas     uint64                                   `protobuf:"varint,10,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty"`
	ClientID       string                                   `protobuf:"bytes,11,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *MsgCoinRatesData) Reset()         { *m = MsgCoinRatesData{} }
func (m *MsgCoinRatesData) String() string { return proto.CompactTextString(m) }
func (*MsgCoinRatesData) ProtoMessage()    {}
func (*MsgCoinRatesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5bd027f3d974bb4, []int{0}
}
func (m *MsgCoinRatesData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinRatesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinRatesData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinRatesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinRatesData.Merge(m, src)
}
func (m *MsgCoinRatesData) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinRatesData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinRatesData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinRatesData proto.InternalMessageInfo

func (m *MsgCoinRatesData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCoinRatesData) GetOracleScriptID() uint64 {
	if m != nil {
		return m.OracleScriptID
	}
	return 0
}

func (m *MsgCoinRatesData) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgCoinRatesData) GetCalldata() *CoinRatesCallData {
	if m != nil {
		return m.Calldata
	}
	return nil
}

func (m *MsgCoinRatesData) GetAskCount() uint64 {
	if m != nil {
		return m.AskCount
	}
	return 0
}

func (m *MsgCoinRatesData) GetMinCount() uint64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *MsgCoinRatesData) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func (m *MsgCoinRatesData) GetRequestKey() string {
	if m != nil {
		return m.RequestKey
	}
	return ""
}

func (m *MsgCoinRatesData) GetPrepareGas() uint64 {
	if m != nil {
		return m.PrepareGas
	}
	return 0
}

func (m *MsgCoinRatesData) GetExecuteGas() uint64 {
	if m != nil {
		return m.ExecuteGas
	}
	return 0
}

func (m *MsgCoinRatesData) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type MsgCoinRatesDataResponse struct {
}

func (m *MsgCoinRatesDataResponse) Reset()         { *m = MsgCoinRatesDataResponse{} }
func (m *MsgCoinRatesDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCoinRatesDataResponse) ProtoMessage()    {}
func (*MsgCoinRatesDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5bd027f3d974bb4, []int{1}
}
func (m *MsgCoinRatesDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinRatesDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinRatesDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinRatesDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinRatesDataResponse.Merge(m, src)
}
func (m *MsgCoinRatesDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinRatesDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinRatesDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinRatesDataResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCoinRatesData)(nil), "cosmonaut.oracle.consuming.MsgCoinRatesData")
	proto.RegisterType((*MsgCoinRatesDataResponse)(nil), "cosmonaut.oracle.consuming.MsgCoinRatesDataResponse")
}

func init() { proto.RegisterFile("consuming/tx.proto", fileDescriptor_f5bd027f3d974bb4) }

var fileDescriptor_f5bd027f3d974bb4 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x8a, 0xd3, 0x40,
	0x1c, 0x6e, 0xdc, 0x75, 0x37, 0x9d, 0xba, 0x65, 0x09, 0x82, 0x31, 0x42, 0x52, 0x0a, 0x42, 0x05,
	0x9b, 0xb8, 0xd5, 0x93, 0xc7, 0xb6, 0xb0, 0x14, 0x5d, 0x84, 0x78, 0x10, 0xbc, 0x84, 0xe9, 0xf4,
	0xb7, 0xe9, 0xd0, 0x64, 0x26, 0x66, 0x26, 0xd2, 0xfa, 0x14, 0xbe, 0x86, 0x3e, 0xc9, 0x1e, 0xf7,
	0xe8, 0xa9, 0x4a, 0xfb, 0x06, 0x3e, 0x81, 0xcc, 0x4c, 0xb6, 0xae, 0x05, 0x05, 0x4f, 0xc9, 0x7c,
	0xdf, 0xf7, 0xfb, 0xfb, 0xcd, 0x20, 0x87, 0x70, 0x26, 0xaa, 0x9c, 0xb2, 0x34, 0x92, 0xcb, 0xb0,
	0x28, 0xb9, 0xe4, 0x8e, 0x47, 0xb8, 0xc8, 0x39, 0xc3, 0x95, 0x0c, 0x79, 0x89, 0x49, 0x06, 0xe1,
	0x4e, 0xe4, 0xdd, 0x4f, 0x79, 0xca, 0xb5, 0x2c, 0x52, 0x7f, 0x26, 0xc2, 0xf3, 0x75, 0x84, 0x88,
	0xa6, 0x58, 0x40, 0xf4, 0xf1, 0x6c, 0x0a, 0x12, 0x9f, 0x45, 0x84, 0x53, 0x56, 0xf3, 0xde, 0xef,
	0x2a, 0x0a, 0x4d, 0x4a, 0x2c, 0x41, 0x18, 0xae, 0xfb, 0xe5, 0x10, 0x9d, 0x5e, 0x88, 0x74, 0xc4,
	0x29, 0x8b, 0x15, 0x3c, 0xc6, 0x12, 0x3b, 0x2e, 0x3a, 0x26, 0x25, 0x60, 0xc9, 0x4b, 0xd7, 0xea,
	0x58, 0xbd, 0x66, 0x7c, 0x73, 0x74, 0xde, 0xa1, 0x53, 0xd3, 0x54, 0x22, 0x48, 0x49, 0x0b, 0x99,
	0xd0, 0x99, 0x7b, 0xa7, 0x63, 0xf5, 0x0e, 0x87, 0xfd, 0xcd, 0x3a, 0x68, 0xbf, 0xd1, 0xdc, 0x5b,
	0x4d, 0x4d, 0xc6, 0x3f, 0xd7, 0xc1, 0x83, 0x15, 0xce, 0xb3, 0x97, 0xdd, 0xfd, 0x98, 0x6e, 0xdc,
	0xe6, 0xb7, 0xa5, 0x33, 0xe7, 0x31, 0x6a, 0x0b, 0x5e, 0x95, 0x04, 0x12, 0x32, 0xc7, 0x8c, 0x41,
	0xe6, 0x1e, 0xe8, 0xca, 0x27, 0x06, 0x1d, 0x19, 0xd0, 0x99, 0x20, 0x9b, 0xe0, 0x2c, 0x9b, 0x61,
	0x89, 0xdd, 0xc3, 0x8e, 0xd5, 0x6b, 0x0d, 0xfa, 0xe1, 0xdf, 0xf7, 0x15, 0xee, 0xc6, 0x1a, 0xe1,
	0x2c, 0x53, 0xa3, 0xc5, 0xbb, 0x70, 0xe7, 0x11, 0x6a, 0x62, 0xb1, 0x48, 0x08, 0xaf, 0x98, 0x74,
	0xef, 0xaa, 0x19, 0x62, 0x1b, 0x8b, 0xc5, 0x48, 0x9d, 0x15, 0x99, 0x53, 0x56, 0x93, 0x47, 0x86,
	0xcc, 0x29, 0x33, 0xe4, 0x1c, 0x35, 0x2f, 0x01, 0x92, 0x8c, 0xe6, 0x54, 0xba, 0xc7, 0x9d, 0x83,
	0x5e, 0x6b, 0xf0, 0xd0, 0x74, 0x21, 0x42, 0xe5, 0x41, 0x58, 0x7b, 0xa0, 0xcb, 0x0f, 0x9f, 0x5d,
	0xad, 0x83, 0xc6, 0xd7, 0xef, 0x41, 0x2f, 0xa5, 0x72, 0x5e, 0x4d, 0x43, 0xc2, 0xf3, 0xa8, 0x36,
	0xcc, 0x7c, 0xfa, 0x62, 0xb6, 0x88, 0xe4, 0xaa, 0x00, 0xa1, 0x03, 0x44, 0x6c, 0x5f, 0x02, 0xbc,
	0x56, 0xc9, 0x9d, 0x00, 0xb5, 0x4a, 0xf8, 0x50, 0x81, 0x90, 0xc9, 0x02, 0x56, 0xae, 0xad, 0x57,
	0x82, 0x6a, 0xe8, 0x15, 0xac, 0x94, 0xa0, 0x28, 0xa1, 0xc0, 0x25, 0x24, 0x29, 0x16, 0x6e, 0x53,
	0x77, 0x8a, 0x6a, 0xe8, 0x1c, 0x0b, 0x25, 0x80, 0x25, 0x90, 0x4a, 0x1a, 0x01, 0x32, 0x82, 0x1a,
	0x52, 0x82, 0x27, 0xa8, 0x49, 0x32, 0x0a, 0x4c, 0x5b, 0xd9, 0x52, 0x05, 0x86, 0xf7, 0x36, 0xeb,
	0xc0, 0x1e, 0x69, 0x70, 0x32, 0x8e, 0x6d, 0x43, 0x4f, 0x66, 0x5d, 0x0f, 0xb9, 0xfb, 0x57, 0x25,
	0x06, 0x51, 0x70, 0x26, 0x60, 0xf0, 0x09, 0x1d, 0x5c, 0x88, 0xd4, 0x11, 0xe8, 0xe4, 0xcf, 0xab,
	0xf4, 0xf4, 0x5f, 0xf6, 0xec, 0x67, 0xf3, 0x5e, 0xfc, 0x8f, 0xfa, 0xa6, 0xf6, 0xf0, 0xfc, 0x6a,
	0xe3, 0x5b, 0xd7, 0x1b, 0xdf, 0xfa, 0xb1, 0xf1, 0xad, 0xcf, 0x5b, 0xbf, 0x71, 0xbd, 0xf5, 0x1b,
	0xdf, 0xb6, 0x7e, 0xe3, 0x7d, 0x7f, 0x7f, 0xe7, 0x2a, 0x73, 0x64, 0x32, 0x47, 0xcb, 0xe8, 0xd6,
	0xeb, 0x53, 0xeb, 0x9f, 0x1e, 0xe9, 0x37, 0xf1, 0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x6f, 0xd8, 0xef, 0x97, 0x03, 0x00, 0x00,
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
	CoinRatesData(ctx context.Context, in *MsgCoinRatesData, opts ...grpc.CallOption) (*MsgCoinRatesDataResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CoinRatesData(ctx context.Context, in *MsgCoinRatesData, opts ...grpc.CallOption) (*MsgCoinRatesDataResponse, error) {
	out := new(MsgCoinRatesDataResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.consuming.Msg/CoinRatesData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CoinRatesData(context.Context, *MsgCoinRatesData) (*MsgCoinRatesDataResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CoinRatesData(ctx context.Context, req *MsgCoinRatesData) (*MsgCoinRatesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinRatesData not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CoinRatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCoinRatesData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CoinRatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.consuming.Msg/CoinRatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CoinRatesData(ctx, req.(*MsgCoinRatesData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.oracle.consuming.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoinRatesData",
			Handler:    _Msg_CoinRatesData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consuming/tx.proto",
}

func (m *MsgCoinRatesData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinRatesData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinRatesData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x5a
	}
	if m.ExecuteGas != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ExecuteGas))
		i--
		dAtA[i] = 0x50
	}
	if m.PrepareGas != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PrepareGas))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RequestKey) > 0 {
		i -= len(m.RequestKey)
		copy(dAtA[i:], m.RequestKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RequestKey)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.MinCount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x30
	}
	if m.AskCount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x28
	}
	if m.Calldata != nil {
		{
			size, err := m.Calldata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OracleScriptID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.OracleScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCoinRatesDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinRatesDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinRatesDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCoinRatesData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.OracleScriptID != 0 {
		n += 1 + sovTx(uint64(m.OracleScriptID))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Calldata != nil {
		l = m.Calldata.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AskCount != 0 {
		n += 1 + sovTx(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovTx(uint64(m.MinCount))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.RequestKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PrepareGas != 0 {
		n += 1 + sovTx(uint64(m.PrepareGas))
	}
	if m.ExecuteGas != 0 {
		n += 1 + sovTx(uint64(m.ExecuteGas))
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCoinRatesDataResponse) Size() (n int) {
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
func (m *MsgCoinRatesData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCoinRatesData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinRatesData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptID", wireType)
			}
			m.OracleScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldata", wireType)
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
			if m.Calldata == nil {
				m.Calldata = &CoinRatesCallData{}
			}
			if err := m.Calldata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
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
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestKey", wireType)
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
			m.RequestKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGas", wireType)
			}
			m.PrepareGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGas", wireType)
			}
			m.ExecuteGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
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
			m.ClientID = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCoinRatesDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCoinRatesDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinRatesDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
