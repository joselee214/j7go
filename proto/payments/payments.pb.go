// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payments.proto

package payments // import "j7go/proto/payments"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/joselee214/j7f/proto/common"

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

// 创建方式
type WayCreation int32

const (
	// 初始值无意义
	WayCreation_WAY_CREATION_MOD WayCreation = 0
	// 系统默认
	WayCreation_SYSTEM_DEFAULT WayCreation = 1
	// 自定义添加
	WayCreation_CUSTOMIZE WayCreation = 2
)

var WayCreation_name = map[int32]string{
	0: "WAY_CREATION_MOD",
	1: "SYSTEM_DEFAULT",
	2: "CUSTOMIZE",
}
var WayCreation_value = map[string]int32{
	"WAY_CREATION_MOD": 0,
	"SYSTEM_DEFAULT":   1,
	"CUSTOMIZE":        2,
}

func (x WayCreation) String() string {
	return proto.EnumName(WayCreation_name, int32(x))
}
func (WayCreation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{0}
}

// 是否支持在线支付
type OnlinePayment int32

const (
	// 初始值无意义
	OnlinePayment_ONLINE_PAY_STATUS OnlinePayment = 0
	// 在线
	OnlinePayment_ONLINE_STATUS OnlinePayment = 1
	// 不支持子痫
	OnlinePayment_NOT_ONLINE_STATUS OnlinePayment = 2
)

var OnlinePayment_name = map[int32]string{
	0: "ONLINE_PAY_STATUS",
	1: "ONLINE_STATUS",
	2: "NOT_ONLINE_STATUS",
}
var OnlinePayment_value = map[string]int32{
	"ONLINE_PAY_STATUS": 0,
	"ONLINE_STATUS":     1,
	"NOT_ONLINE_STATUS": 2,
}

func (x OnlinePayment) String() string {
	return proto.EnumName(OnlinePayment_name, int32(x))
}
func (OnlinePayment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{1}
}

// 支付方式列表请求
type PaymentsRequest struct {
	// 公共请求头
	Header               *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PaymentsRequest) Reset()         { *m = PaymentsRequest{} }
func (m *PaymentsRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentsRequest) ProtoMessage()    {}
func (*PaymentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{0}
}
func (m *PaymentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentsRequest.Unmarshal(m, b)
}
func (m *PaymentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentsRequest.Marshal(b, m, deterministic)
}
func (dst *PaymentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentsRequest.Merge(dst, src)
}
func (m *PaymentsRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentsRequest.Size(m)
}
func (m *PaymentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentsRequest proto.InternalMessageInfo

func (m *PaymentsRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 支付方式列表返回
type PaymentsResponse struct {
	// 公共响应状态
	Status               *common.BusinessStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PayMethodsSource     []*PaymentResult       `protobuf:"bytes,2,rep,name=pay_methods_source,json=payMethodsSource,proto3" json:"pay_methods_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PaymentsResponse) Reset()         { *m = PaymentsResponse{} }
func (m *PaymentsResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentsResponse) ProtoMessage()    {}
func (*PaymentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{1}
}
func (m *PaymentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentsResponse.Unmarshal(m, b)
}
func (m *PaymentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentsResponse.Marshal(b, m, deterministic)
}
func (dst *PaymentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentsResponse.Merge(dst, src)
}
func (m *PaymentsResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentsResponse.Size(m)
}
func (m *PaymentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentsResponse proto.InternalMessageInfo

func (m *PaymentsResponse) GetStatus() *common.BusinessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PaymentsResponse) GetPayMethodsSource() []*PaymentResult {
	if m != nil {
		return m.PayMethodsSource
	}
	return nil
}

// 编辑支付方式请求
type EditPayRequest struct {
	// 公共请求头
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 支付方式id
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 是否启用
	IsDel                common.DelStatus `protobuf:"varint,3,opt,name=is_del,json=isDel,proto3,enum=common.DelStatus" json:"is_del,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EditPayRequest) Reset()         { *m = EditPayRequest{} }
func (m *EditPayRequest) String() string { return proto.CompactTextString(m) }
func (*EditPayRequest) ProtoMessage()    {}
func (*EditPayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{2}
}
func (m *EditPayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditPayRequest.Unmarshal(m, b)
}
func (m *EditPayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditPayRequest.Marshal(b, m, deterministic)
}
func (dst *EditPayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPayRequest.Merge(dst, src)
}
func (m *EditPayRequest) XXX_Size() int {
	return xxx_messageInfo_EditPayRequest.Size(m)
}
func (m *EditPayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EditPayRequest proto.InternalMessageInfo

func (m *EditPayRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EditPayRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EditPayRequest) GetIsDel() common.DelStatus {
	if m != nil {
		return m.IsDel
	}
	return common.DelStatus_NOT_DEL
}

// 编辑支付方式响应
type EditPayResponse struct {
	Status               *common.BusinessStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EditPayResponse) Reset()         { *m = EditPayResponse{} }
func (m *EditPayResponse) String() string { return proto.CompactTextString(m) }
func (*EditPayResponse) ProtoMessage()    {}
func (*EditPayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{3}
}
func (m *EditPayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditPayResponse.Unmarshal(m, b)
}
func (m *EditPayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditPayResponse.Marshal(b, m, deterministic)
}
func (dst *EditPayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPayResponse.Merge(dst, src)
}
func (m *EditPayResponse) XXX_Size() int {
	return xxx_messageInfo_EditPayResponse.Size(m)
}
func (m *EditPayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EditPayResponse proto.InternalMessageInfo

func (m *EditPayResponse) GetStatus() *common.BusinessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PaymentResult struct {
	// 主键id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 支付名称
	PaymentName string `protobuf:"bytes,2,opt,name=payment_name,json=paymentName,proto3" json:"payment_name,omitempty"`
	// 创建方式 1-系统默认 2-自定义添加
	CreatingMode WayCreation `protobuf:"varint,3,opt,name=creating_mode,json=creatingMode,proto3,enum=payments.WayCreation" json:"creating_mode,omitempty"`
	// 是否在线支付 1-在线 2-离线
	IsOnlinePay OnlinePayment `protobuf:"varint,4,opt,name=is_online_pay,json=isOnlinePay,proto3,enum=payments.OnlinePayment" json:"is_online_pay,omitempty"`
	// 是否删除
	IsDel common.DelStatus `protobuf:"varint,5,opt,name=is_del,json=isDel,proto3,enum=common.DelStatus" json:"is_del,omitempty"`
	// 创建时间
	CreatedTime uint32 `protobuf:"varint,6,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// 更新时间
	UpdatedTime          uint32   `protobuf:"varint,7,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentResult) Reset()         { *m = PaymentResult{} }
func (m *PaymentResult) String() string { return proto.CompactTextString(m) }
func (*PaymentResult) ProtoMessage()    {}
func (*PaymentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_payments_90de37742759c1a5, []int{4}
}
func (m *PaymentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentResult.Unmarshal(m, b)
}
func (m *PaymentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentResult.Marshal(b, m, deterministic)
}
func (dst *PaymentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentResult.Merge(dst, src)
}
func (m *PaymentResult) XXX_Size() int {
	return xxx_messageInfo_PaymentResult.Size(m)
}
func (m *PaymentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentResult.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentResult proto.InternalMessageInfo

func (m *PaymentResult) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PaymentResult) GetPaymentName() string {
	if m != nil {
		return m.PaymentName
	}
	return ""
}

func (m *PaymentResult) GetCreatingMode() WayCreation {
	if m != nil {
		return m.CreatingMode
	}
	return WayCreation_WAY_CREATION_MOD
}

func (m *PaymentResult) GetIsOnlinePay() OnlinePayment {
	if m != nil {
		return m.IsOnlinePay
	}
	return OnlinePayment_ONLINE_PAY_STATUS
}

func (m *PaymentResult) GetIsDel() common.DelStatus {
	if m != nil {
		return m.IsDel
	}
	return common.DelStatus_NOT_DEL
}

func (m *PaymentResult) GetCreatedTime() uint32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PaymentResult) GetUpdatedTime() uint32 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*PaymentsRequest)(nil), "payments.PaymentsRequest")
	proto.RegisterType((*PaymentsResponse)(nil), "payments.PaymentsResponse")
	proto.RegisterType((*EditPayRequest)(nil), "payments.EditPayRequest")
	proto.RegisterType((*EditPayResponse)(nil), "payments.EditPayResponse")
	proto.RegisterType((*PaymentResult)(nil), "payments.PaymentResult")
	proto.RegisterEnum("payments.WayCreation", WayCreation_name, WayCreation_value)
	proto.RegisterEnum("payments.OnlinePayment", OnlinePayment_name, OnlinePayment_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentsServerClient is the client API for PaymentsServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsServerClient interface {
	// 支付方式列表
	GetPayments(ctx context.Context, opts ...grpc.CallOption) (PaymentsServer_GetPaymentsClient, error)
	// 编辑支付方式
	GetEditPay(ctx context.Context, opts ...grpc.CallOption) (PaymentsServer_GetEditPayClient, error)
}

type paymentsServerClient struct {
	cc *grpc.ClientConn
}

func NewPaymentsServerClient(cc *grpc.ClientConn) PaymentsServerClient {
	return &paymentsServerClient{cc}
}

func (c *paymentsServerClient) GetPayments(ctx context.Context, opts ...grpc.CallOption) (PaymentsServer_GetPaymentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PaymentsServer_serviceDesc.Streams[0], "/payments.PaymentsServer/GetPayments", opts...)
	if err != nil {
		return nil, err
	}
	x := &paymentsServerGetPaymentsClient{stream}
	return x, nil
}

type PaymentsServer_GetPaymentsClient interface {
	Send(*PaymentsRequest) error
	Recv() (*PaymentsResponse, error)
	grpc.ClientStream
}

type paymentsServerGetPaymentsClient struct {
	grpc.ClientStream
}

func (x *paymentsServerGetPaymentsClient) Send(m *PaymentsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *paymentsServerGetPaymentsClient) Recv() (*PaymentsResponse, error) {
	m := new(PaymentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *paymentsServerClient) GetEditPay(ctx context.Context, opts ...grpc.CallOption) (PaymentsServer_GetEditPayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PaymentsServer_serviceDesc.Streams[1], "/payments.PaymentsServer/GetEditPay", opts...)
	if err != nil {
		return nil, err
	}
	x := &paymentsServerGetEditPayClient{stream}
	return x, nil
}

type PaymentsServer_GetEditPayClient interface {
	Send(*EditPayRequest) error
	Recv() (*EditPayResponse, error)
	grpc.ClientStream
}

type paymentsServerGetEditPayClient struct {
	grpc.ClientStream
}

func (x *paymentsServerGetEditPayClient) Send(m *EditPayRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *paymentsServerGetEditPayClient) Recv() (*EditPayResponse, error) {
	m := new(EditPayResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PaymentsServerServer is the server API for PaymentsServer service.
type PaymentsServerServer interface {
	// 支付方式列表
	GetPayments(PaymentsServer_GetPaymentsServer) error
	// 编辑支付方式
	GetEditPay(PaymentsServer_GetEditPayServer) error
}

func RegisterPaymentsServerServer(s *grpc.Server, srv PaymentsServerServer) {
	s.RegisterService(&_PaymentsServer_serviceDesc, srv)
}

func _PaymentsServer_GetPayments_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaymentsServerServer).GetPayments(&paymentsServerGetPaymentsServer{stream})
}

type PaymentsServer_GetPaymentsServer interface {
	Send(*PaymentsResponse) error
	Recv() (*PaymentsRequest, error)
	grpc.ServerStream
}

type paymentsServerGetPaymentsServer struct {
	grpc.ServerStream
}

func (x *paymentsServerGetPaymentsServer) Send(m *PaymentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *paymentsServerGetPaymentsServer) Recv() (*PaymentsRequest, error) {
	m := new(PaymentsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PaymentsServer_GetEditPay_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaymentsServerServer).GetEditPay(&paymentsServerGetEditPayServer{stream})
}

type PaymentsServer_GetEditPayServer interface {
	Send(*EditPayResponse) error
	Recv() (*EditPayRequest, error)
	grpc.ServerStream
}

type paymentsServerGetEditPayServer struct {
	grpc.ServerStream
}

func (x *paymentsServerGetEditPayServer) Send(m *EditPayResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *paymentsServerGetEditPayServer) Recv() (*EditPayRequest, error) {
	m := new(EditPayRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PaymentsServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payments.PaymentsServer",
	HandlerType: (*PaymentsServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPayments",
			Handler:       _PaymentsServer_GetPayments_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetEditPay",
			Handler:       _PaymentsServer_GetEditPay_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "payments.proto",
}

func init() { proto.RegisterFile("payments.proto", fileDescriptor_payments_90de37742759c1a5) }

var fileDescriptor_payments_90de37742759c1a5 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0x5d, 0x1a, 0xe8, 0x75, 0x9d, 0x3a, 0xa3, 0x16, 0xdc, 0xac, 0x42, 0x56, 0x51, 0x05,
	0x09, 0x0a, 0x62, 0x43, 0x17, 0x28, 0x4d, 0xdc, 0x87, 0x68, 0x1e, 0xb2, 0x1d, 0x55, 0xad, 0x90,
	0x46, 0x6e, 0x7c, 0xd5, 0x8e, 0x88, 0x3d, 0x26, 0x33, 0x46, 0xf8, 0x13, 0xf8, 0x0c, 0xd6, 0xfc,
	0x09, 0x5f, 0x85, 0xe2, 0x8c, 0xf3, 0x10, 0x15, 0x8b, 0x2e, 0xa2, 0xd8, 0xe7, 0x9c, 0x7b, 0xe7,
	0x9e, 0xa3, 0x3b, 0x86, 0x72, 0x12, 0x64, 0x11, 0xc6, 0x52, 0x34, 0x93, 0x19, 0x97, 0x9c, 0xbc,
	0x28, 0xde, 0xab, 0x7b, 0x13, 0x1e, 0x45, 0x3c, 0x5e, 0xe0, 0xf5, 0x4f, 0xb0, 0x3f, 0x52, 0x8c,
	0x8b, 0xdf, 0x52, 0x14, 0x92, 0xbc, 0x81, 0xd2, 0x03, 0x06, 0x21, 0xce, 0x6c, 0xad, 0xa6, 0x35,
	0x8c, 0xf6, 0x41, 0x53, 0x55, 0x74, 0xf3, 0xbf, 0x8b, 0x9c, 0x73, 0x95, 0xa6, 0xfe, 0x53, 0x03,
	0x6b, 0xd5, 0x41, 0x24, 0x3c, 0x16, 0x48, 0x9a, 0x50, 0x12, 0x32, 0x90, 0xa9, 0x50, 0x2d, 0x5e,
	0x16, 0x2d, 0x4e, 0x53, 0xc1, 0x62, 0x14, 0xc2, 0xcb, 0x59, 0x57, 0xa9, 0x88, 0x03, 0x24, 0x09,
	0x32, 0x1a, 0xa1, 0x7c, 0xe0, 0xa1, 0xa0, 0x82, 0xa7, 0xb3, 0x09, 0xda, 0x7a, 0x6d, 0xbb, 0x61,
	0xb4, 0x5f, 0x35, 0x97, 0x56, 0xd4, 0x39, 0x2e, 0x8a, 0x74, 0x2a, 0x5d, 0x2b, 0x09, 0xb2, 0xfe,
	0xa2, 0xc2, 0xcb, 0x0b, 0xea, 0x3f, 0xa0, 0xec, 0x84, 0x4c, 0x8e, 0x82, 0xec, 0x49, 0x5e, 0x48,
	0x19, 0x74, 0x16, 0xda, 0x7a, 0x4d, 0x6b, 0x98, 0xae, 0xce, 0x42, 0xd2, 0x80, 0x12, 0x13, 0x34,
	0xc4, 0xa9, 0xbd, 0x5d, 0xd3, 0x1a, 0xe5, 0x76, 0xa5, 0xa8, 0xee, 0xe1, 0x54, 0x39, 0xd8, 0x61,
	0xa2, 0x87, 0xd3, 0x7a, 0x07, 0xf6, 0x97, 0x27, 0x3f, 0x2d, 0x83, 0xfa, 0x6f, 0x1d, 0xcc, 0x0d,
	0x83, 0x6a, 0x1c, 0x6d, 0x39, 0xce, 0x6b, 0xd8, 0x53, 0x51, 0xd0, 0x38, 0x88, 0x30, 0x1f, 0x74,
	0xd7, 0x35, 0x14, 0x36, 0x08, 0x22, 0x24, 0x1f, 0xc1, 0x9c, 0xcc, 0x30, 0x90, 0x2c, 0xbe, 0xa7,
	0x11, 0x0f, 0x51, 0x0d, 0x7e, 0xb8, 0xca, 0xf0, 0x3a, 0xc8, 0xba, 0xb9, 0x82, 0xc7, 0xee, 0x5e,
	0xa1, 0xed, 0xf3, 0x10, 0xc9, 0x09, 0x98, 0x4c, 0x50, 0x1e, 0x4f, 0x59, 0x8c, 0x34, 0x09, 0x32,
	0xfb, 0x59, 0x5e, 0xbb, 0x96, 0xff, 0x30, 0xe7, 0x8a, 0x21, 0x0d, 0x26, 0x96, 0xc0, 0x5a, 0x54,
	0x3b, 0xff, 0x8f, 0x6a, 0xee, 0x22, 0x3f, 0x16, 0x43, 0x2a, 0x59, 0x84, 0x76, 0x29, 0xf7, 0x67,
	0x28, 0xcc, 0x67, 0x11, 0xce, 0x25, 0x69, 0x12, 0xae, 0x24, 0xcf, 0x17, 0x12, 0x85, 0xcd, 0x25,
	0xc7, 0x67, 0x60, 0xac, 0x39, 0x21, 0x07, 0x60, 0x5d, 0x77, 0x6e, 0x68, 0xd7, 0x75, 0x3a, 0xfe,
	0xe5, 0x70, 0x40, 0xfb, 0xc3, 0x9e, 0xb5, 0x45, 0x08, 0x94, 0xbd, 0x1b, 0xcf, 0x77, 0xfa, 0xb4,
	0xe7, 0x9c, 0x75, 0xc6, 0x57, 0xbe, 0xa5, 0x11, 0x13, 0x76, 0xbb, 0x63, 0xcf, 0x1f, 0xf6, 0x2f,
	0x6f, 0x1d, 0x4b, 0x3f, 0x1e, 0x81, 0xb9, 0xe1, 0x8a, 0x1c, 0x42, 0x65, 0x38, 0xb8, 0xba, 0x1c,
	0x38, 0x74, 0xd4, 0xb9, 0xa1, 0x9e, 0xdf, 0xf1, 0xc7, 0x9e, 0xb5, 0x45, 0x2a, 0x60, 0x2a, 0x58,
	0x41, 0xda, 0x5c, 0x39, 0x18, 0xfa, 0x74, 0x13, 0xd6, 0xdb, 0xbf, 0x34, 0x28, 0x17, 0x17, 0xc2,
	0xc3, 0xd9, 0x77, 0x9c, 0x91, 0x0b, 0x30, 0xce, 0x51, 0x16, 0x20, 0x39, 0xfa, 0x67, 0xa3, 0x8b,
	0xbb, 0x57, 0xad, 0x3e, 0x46, 0x2d, 0x16, 0xaa, 0xa1, 0xbd, 0xd3, 0x88, 0x03, 0x70, 0x8e, 0x52,
	0xad, 0x1a, 0xb1, 0x57, 0xea, 0xcd, 0xbd, 0xaf, 0x1e, 0x3d, 0xc2, 0xac, 0xda, 0x9c, 0x7e, 0xbe,
	0xfd, 0x70, 0xcf, 0xe4, 0x34, 0xb8, 0x6b, 0x0a, 0x99, 0x85, 0xcd, 0x49, 0xdc, 0x4a, 0x25, 0x9f,
	0xff, 0xde, 0x8a, 0xf0, 0x6b, 0xeb, 0x9e, 0xb7, 0xf2, 0x8f, 0x43, 0xab, 0xe8, 0x71, 0x52, 0x3c,
	0xfc, 0xd1, 0x2b, 0x63, 0xc9, 0x47, 0x73, 0xf2, 0x4b, 0x31, 0xdf, 0x5d, 0x29, 0x17, 0xbf, 0xff,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x67, 0x9c, 0x3f, 0xbe, 0x73, 0x04, 0x00, 0x00,
}
