// Code generated by protoc-gen-go. DO NOT EDIT.
// source: personal_course_tpl.proto

package product // import "j7go/proto/product"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.7yes.com/j7f/proto/common"

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

// 设置私教课状态请求体
type SetCourseStateRequest struct {
	// 公共请求头
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 课程ID
	CourseId uint32 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	// 操作人ID
	OperatorId           uint32   `protobuf:"varint,3,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetCourseStateRequest) Reset()         { *m = SetCourseStateRequest{} }
func (m *SetCourseStateRequest) String() string { return proto.CompactTextString(m) }
func (*SetCourseStateRequest) ProtoMessage()    {}
func (*SetCourseStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{0}
}
func (m *SetCourseStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetCourseStateRequest.Unmarshal(m, b)
}
func (m *SetCourseStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetCourseStateRequest.Marshal(b, m, deterministic)
}
func (dst *SetCourseStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetCourseStateRequest.Merge(dst, src)
}
func (m *SetCourseStateRequest) XXX_Size() int {
	return xxx_messageInfo_SetCourseStateRequest.Size(m)
}
func (m *SetCourseStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetCourseStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetCourseStateRequest proto.InternalMessageInfo

func (m *SetCourseStateRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetCourseStateRequest) GetCourseId() uint32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

func (m *SetCourseStateRequest) GetOperatorId() uint32 {
	if m != nil {
		return m.OperatorId
	}
	return 0
}

// 新增私教课请求体
type PersonalCourseRequest struct {
	// 公共请求头
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 品牌ID
	BrandId uint32 `protobuf:"varint,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	// 门店ID
	ShopId uint32 `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// 课程ID，新增时传0，更新时传具体ID
	CourseId uint32 `protobuf:"varint,4,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	// 课程名
	CourseName string `protobuf:"bytes,5,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	// 课程类别
	CourseCategory uint32 `protobuf:"varint,6,opt,name=course_category,json=courseCategory,proto3" json:"course_category,omitempty"`
	// 训练目标
	TrainAim []uint32 `protobuf:"varint,7,rep,packed,name=train_aim,json=trainAim,proto3" json:"train_aim,omitempty"`
	// 课程时长
	Duration uint32 `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	// 建议售价
	Price uint32 `protobuf:"varint,9,opt,name=price,proto3" json:"price,omitempty"`
	// 课程图片
	CourseImg string `protobuf:"bytes,10,opt,name=course_img,json=courseImg,proto3" json:"course_img,omitempty"`
	// 课程描述
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	// 单节课有效时长
	EffectiveUnit uint32 `protobuf:"varint,12,opt,name=effective_unit,json=effectiveUnit,proto3" json:"effective_unit,omitempty"`
	// 发布渠道 1品牌 2门店
	PublishChannel       PublishChannel `protobuf:"varint,13,opt,name=publish_channel,json=publishChannel,proto3,enum=product.PublishChannel" json:"publish_channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PersonalCourseRequest) Reset()         { *m = PersonalCourseRequest{} }
func (m *PersonalCourseRequest) String() string { return proto.CompactTextString(m) }
func (*PersonalCourseRequest) ProtoMessage()    {}
func (*PersonalCourseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{1}
}
func (m *PersonalCourseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalCourseRequest.Unmarshal(m, b)
}
func (m *PersonalCourseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalCourseRequest.Marshal(b, m, deterministic)
}
func (dst *PersonalCourseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalCourseRequest.Merge(dst, src)
}
func (m *PersonalCourseRequest) XXX_Size() int {
	return xxx_messageInfo_PersonalCourseRequest.Size(m)
}
func (m *PersonalCourseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalCourseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalCourseRequest proto.InternalMessageInfo

func (m *PersonalCourseRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PersonalCourseRequest) GetBrandId() uint32 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *PersonalCourseRequest) GetShopId() uint32 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *PersonalCourseRequest) GetCourseId() uint32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

func (m *PersonalCourseRequest) GetCourseName() string {
	if m != nil {
		return m.CourseName
	}
	return ""
}

func (m *PersonalCourseRequest) GetCourseCategory() uint32 {
	if m != nil {
		return m.CourseCategory
	}
	return 0
}

func (m *PersonalCourseRequest) GetTrainAim() []uint32 {
	if m != nil {
		return m.TrainAim
	}
	return nil
}

func (m *PersonalCourseRequest) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PersonalCourseRequest) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PersonalCourseRequest) GetCourseImg() string {
	if m != nil {
		return m.CourseImg
	}
	return ""
}

func (m *PersonalCourseRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PersonalCourseRequest) GetEffectiveUnit() uint32 {
	if m != nil {
		return m.EffectiveUnit
	}
	return 0
}

func (m *PersonalCourseRequest) GetPublishChannel() PublishChannel {
	if m != nil {
		return m.PublishChannel
	}
	return PublishChannel_PUBLISH_CHANNEL_INIT
}

// 新增私教课、设置门店、设置售卖价格返回体
type PersonalCourseResponse struct {
	// 公共返回域
	Status *common.BusinessStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 私教课ID
	CourseId             uint32   `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonalCourseResponse) Reset()         { *m = PersonalCourseResponse{} }
func (m *PersonalCourseResponse) String() string { return proto.CompactTextString(m) }
func (*PersonalCourseResponse) ProtoMessage()    {}
func (*PersonalCourseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{2}
}
func (m *PersonalCourseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalCourseResponse.Unmarshal(m, b)
}
func (m *PersonalCourseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalCourseResponse.Marshal(b, m, deterministic)
}
func (dst *PersonalCourseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalCourseResponse.Merge(dst, src)
}
func (m *PersonalCourseResponse) XXX_Size() int {
	return xxx_messageInfo_PersonalCourseResponse.Size(m)
}
func (m *PersonalCourseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalCourseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalCourseResponse proto.InternalMessageInfo

func (m *PersonalCourseResponse) GetStatus() *common.BusinessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PersonalCourseResponse) GetCourseId() uint32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

// 设置课程门店请求体
type SetCourseShopsRequest struct {
	// 公共请求头
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 课程ID
	CourseId uint32 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	// 上课门店 1全部门店 2指定门店
	ShopSetting uint32 `protobuf:"varint,3,opt,name=shop_setting,json=shopSetting,proto3" json:"shop_setting,omitempty"`
	// 门店ID 全部门店为空数组
	ShopIds []uint32 `protobuf:"varint,4,rep,packed,name=shop_ids,json=shopIds,proto3" json:"shop_ids,omitempty"`
	// 课程支持教练
	CoachIds []uint32 `protobuf:"varint,5,rep,packed,name=coach_ids,json=coachIds,proto3" json:"coach_ids,omitempty"`
	// 操作人ID
	OperatorId           uint32   `protobuf:"varint,6,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetCourseShopsRequest) Reset()         { *m = SetCourseShopsRequest{} }
func (m *SetCourseShopsRequest) String() string { return proto.CompactTextString(m) }
func (*SetCourseShopsRequest) ProtoMessage()    {}
func (*SetCourseShopsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{3}
}
func (m *SetCourseShopsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetCourseShopsRequest.Unmarshal(m, b)
}
func (m *SetCourseShopsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetCourseShopsRequest.Marshal(b, m, deterministic)
}
func (dst *SetCourseShopsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetCourseShopsRequest.Merge(dst, src)
}
func (m *SetCourseShopsRequest) XXX_Size() int {
	return xxx_messageInfo_SetCourseShopsRequest.Size(m)
}
func (m *SetCourseShopsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetCourseShopsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetCourseShopsRequest proto.InternalMessageInfo

func (m *SetCourseShopsRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetCourseShopsRequest) GetCourseId() uint32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

func (m *SetCourseShopsRequest) GetShopSetting() uint32 {
	if m != nil {
		return m.ShopSetting
	}
	return 0
}

func (m *SetCourseShopsRequest) GetShopIds() []uint32 {
	if m != nil {
		return m.ShopIds
	}
	return nil
}

func (m *SetCourseShopsRequest) GetCoachIds() []uint32 {
	if m != nil {
		return m.CoachIds
	}
	return nil
}

func (m *SetCourseShopsRequest) GetOperatorId() uint32 {
	if m != nil {
		return m.OperatorId
	}
	return 0
}

// 设置课程售卖价格请求体
type SetSalePriceRequest struct {
	// 公共请求头
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 门店ID，门店定价时传
	ShopId uint32 `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// 课程ID
	CourseId uint32 `protobuf:"varint,3,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	// 课程定价方式 1品牌统一 2门店自定义
	PriceSetting SetPriceType `protobuf:"varint,4,opt,name=price_setting,json=priceSetting,proto3,enum=product.SetPriceType" json:"price_setting,omitempty"`
	// 课程价格 品牌统一定价时为空数组
	CoursePrices         []*CoursePriceSetting `protobuf:"bytes,5,rep,name=course_prices,json=coursePrices,proto3" json:"course_prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SetSalePriceRequest) Reset()         { *m = SetSalePriceRequest{} }
func (m *SetSalePriceRequest) String() string { return proto.CompactTextString(m) }
func (*SetSalePriceRequest) ProtoMessage()    {}
func (*SetSalePriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{4}
}
func (m *SetSalePriceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSalePriceRequest.Unmarshal(m, b)
}
func (m *SetSalePriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSalePriceRequest.Marshal(b, m, deterministic)
}
func (dst *SetSalePriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSalePriceRequest.Merge(dst, src)
}
func (m *SetSalePriceRequest) XXX_Size() int {
	return xxx_messageInfo_SetSalePriceRequest.Size(m)
}
func (m *SetSalePriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSalePriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetSalePriceRequest proto.InternalMessageInfo

func (m *SetSalePriceRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetSalePriceRequest) GetShopId() uint32 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SetSalePriceRequest) GetCourseId() uint32 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

func (m *SetSalePriceRequest) GetPriceSetting() SetPriceType {
	if m != nil {
		return m.PriceSetting
	}
	return SetPriceType_SET_PRICE_TYPE_INIT
}

func (m *SetSalePriceRequest) GetCoursePrices() []*CoursePriceSetting {
	if m != nil {
		return m.CoursePrices
	}
	return nil
}

// 课程价格定义结构体
type CoursePriceSetting struct {
	// 课程价格ID，新增时传0
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 教练级别
	CoachLevel uint32 `protobuf:"varint,2,opt,name=coach_level,json=coachLevel,proto3" json:"coach_level,omitempty"`
	// 梯度最小出售节数
	SaleMin uint32 `protobuf:"varint,3,opt,name=sale_min,json=saleMin,proto3" json:"sale_min,omitempty"`
	// 梯度最大出售节数
	SaleMax uint32 `protobuf:"varint,4,opt,name=sale_max,json=saleMax,proto3" json:"sale_max,omitempty"`
	// 出售价格
	Price uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	// 转让类型 1万分比（注意：数据库存储是以万分比为准，聚合层给参数时要转换成百分比） 2金额
	TransferType uint32 `protobuf:"varint,6,opt,name=transfer_type,json=transferType,proto3" json:"transfer_type,omitempty"`
	// 转让总数，如类型为元、转让数量为100，结合起来为100元
	TransferNum uint32 `protobuf:"varint,7,opt,name=transfer_num,json=transferNum,proto3" json:"transfer_num,omitempty"`
	// 是否支持在线售卖 0不支持 1支持
	IsOnlineSale         uint32   `protobuf:"varint,8,opt,name=is_online_sale,json=isOnlineSale,proto3" json:"is_online_sale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoursePriceSetting) Reset()         { *m = CoursePriceSetting{} }
func (m *CoursePriceSetting) String() string { return proto.CompactTextString(m) }
func (*CoursePriceSetting) ProtoMessage()    {}
func (*CoursePriceSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_personal_course_tpl_a706bec0ed83f88f, []int{5}
}
func (m *CoursePriceSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoursePriceSetting.Unmarshal(m, b)
}
func (m *CoursePriceSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoursePriceSetting.Marshal(b, m, deterministic)
}
func (dst *CoursePriceSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoursePriceSetting.Merge(dst, src)
}
func (m *CoursePriceSetting) XXX_Size() int {
	return xxx_messageInfo_CoursePriceSetting.Size(m)
}
func (m *CoursePriceSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CoursePriceSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CoursePriceSetting proto.InternalMessageInfo

func (m *CoursePriceSetting) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CoursePriceSetting) GetCoachLevel() uint32 {
	if m != nil {
		return m.CoachLevel
	}
	return 0
}

func (m *CoursePriceSetting) GetSaleMin() uint32 {
	if m != nil {
		return m.SaleMin
	}
	return 0
}

func (m *CoursePriceSetting) GetSaleMax() uint32 {
	if m != nil {
		return m.SaleMax
	}
	return 0
}

func (m *CoursePriceSetting) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CoursePriceSetting) GetTransferType() uint32 {
	if m != nil {
		return m.TransferType
	}
	return 0
}

func (m *CoursePriceSetting) GetTransferNum() uint32 {
	if m != nil {
		return m.TransferNum
	}
	return 0
}

func (m *CoursePriceSetting) GetIsOnlineSale() uint32 {
	if m != nil {
		return m.IsOnlineSale
	}
	return 0
}

func init() {
	proto.RegisterType((*SetCourseStateRequest)(nil), "product.SetCourseStateRequest")
	proto.RegisterType((*PersonalCourseRequest)(nil), "product.PersonalCourseRequest")
	proto.RegisterType((*PersonalCourseResponse)(nil), "product.PersonalCourseResponse")
	proto.RegisterType((*SetCourseShopsRequest)(nil), "product.SetCourseShopsRequest")
	proto.RegisterType((*SetSalePriceRequest)(nil), "product.SetSalePriceRequest")
	proto.RegisterType((*CoursePriceSetting)(nil), "product.CoursePriceSetting")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonalCourseSrvClient is the client API for PersonalCourseSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonalCourseSrvClient interface {
	// 新增私教课
	AddCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_AddCourseClient, error)
	// 更新私教课
	EditCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_EditCourseClient, error)
	// 设置课程支持门店
	SetCourseShops(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseShopsClient, error)
	// 设置售卖价格
	SetSalePrice(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetSalePriceClient, error)
	// 删除私教课
	DelCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_DelCourseClient, error)
	// 私教课置为无效
	SetCourseInvalid(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseInvalidClient, error)
	// 私教课置为有效
	SetCourseValid(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseValidClient, error)
}

type personalCourseSrvClient struct {
	cc *grpc.ClientConn
}

func NewPersonalCourseSrvClient(cc *grpc.ClientConn) PersonalCourseSrvClient {
	return &personalCourseSrvClient{cc}
}

func (c *personalCourseSrvClient) AddCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_AddCourseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[0], "/product.PersonalCourseSrv/AddCourse", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvAddCourseClient{stream}
	return x, nil
}

type PersonalCourseSrv_AddCourseClient interface {
	Send(*PersonalCourseRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvAddCourseClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvAddCourseClient) Send(m *PersonalCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvAddCourseClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) EditCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_EditCourseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[1], "/product.PersonalCourseSrv/EditCourse", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvEditCourseClient{stream}
	return x, nil
}

type PersonalCourseSrv_EditCourseClient interface {
	Send(*PersonalCourseRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvEditCourseClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvEditCourseClient) Send(m *PersonalCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvEditCourseClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) SetCourseShops(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseShopsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[2], "/product.PersonalCourseSrv/SetCourseShops", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvSetCourseShopsClient{stream}
	return x, nil
}

type PersonalCourseSrv_SetCourseShopsClient interface {
	Send(*SetCourseShopsRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvSetCourseShopsClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvSetCourseShopsClient) Send(m *SetCourseShopsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseShopsClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) SetSalePrice(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetSalePriceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[3], "/product.PersonalCourseSrv/SetSalePrice", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvSetSalePriceClient{stream}
	return x, nil
}

type PersonalCourseSrv_SetSalePriceClient interface {
	Send(*SetSalePriceRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvSetSalePriceClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvSetSalePriceClient) Send(m *SetSalePriceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvSetSalePriceClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) DelCourse(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_DelCourseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[4], "/product.PersonalCourseSrv/DelCourse", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvDelCourseClient{stream}
	return x, nil
}

type PersonalCourseSrv_DelCourseClient interface {
	Send(*SetCourseStateRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvDelCourseClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvDelCourseClient) Send(m *SetCourseStateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvDelCourseClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) SetCourseInvalid(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseInvalidClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[5], "/product.PersonalCourseSrv/SetCourseInvalid", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvSetCourseInvalidClient{stream}
	return x, nil
}

type PersonalCourseSrv_SetCourseInvalidClient interface {
	Send(*SetCourseStateRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvSetCourseInvalidClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvSetCourseInvalidClient) Send(m *SetCourseStateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseInvalidClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personalCourseSrvClient) SetCourseValid(ctx context.Context, opts ...grpc.CallOption) (PersonalCourseSrv_SetCourseValidClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonalCourseSrv_serviceDesc.Streams[6], "/product.PersonalCourseSrv/SetCourseValid", opts...)
	if err != nil {
		return nil, err
	}
	x := &personalCourseSrvSetCourseValidClient{stream}
	return x, nil
}

type PersonalCourseSrv_SetCourseValidClient interface {
	Send(*SetCourseStateRequest) error
	Recv() (*PersonalCourseResponse, error)
	grpc.ClientStream
}

type personalCourseSrvSetCourseValidClient struct {
	grpc.ClientStream
}

func (x *personalCourseSrvSetCourseValidClient) Send(m *SetCourseStateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseValidClient) Recv() (*PersonalCourseResponse, error) {
	m := new(PersonalCourseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonalCourseSrvServer is the server API for PersonalCourseSrv service.
type PersonalCourseSrvServer interface {
	// 新增私教课
	AddCourse(PersonalCourseSrv_AddCourseServer) error
	// 更新私教课
	EditCourse(PersonalCourseSrv_EditCourseServer) error
	// 设置课程支持门店
	SetCourseShops(PersonalCourseSrv_SetCourseShopsServer) error
	// 设置售卖价格
	SetSalePrice(PersonalCourseSrv_SetSalePriceServer) error
	// 删除私教课
	DelCourse(PersonalCourseSrv_DelCourseServer) error
	// 私教课置为无效
	SetCourseInvalid(PersonalCourseSrv_SetCourseInvalidServer) error
	// 私教课置为有效
	SetCourseValid(PersonalCourseSrv_SetCourseValidServer) error
}

func RegisterPersonalCourseSrvServer(s *grpc.Server, srv PersonalCourseSrvServer) {
	s.RegisterService(&_PersonalCourseSrv_serviceDesc, srv)
}

func _PersonalCourseSrv_AddCourse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).AddCourse(&personalCourseSrvAddCourseServer{stream})
}

type PersonalCourseSrv_AddCourseServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*PersonalCourseRequest, error)
	grpc.ServerStream
}

type personalCourseSrvAddCourseServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvAddCourseServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvAddCourseServer) Recv() (*PersonalCourseRequest, error) {
	m := new(PersonalCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_EditCourse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).EditCourse(&personalCourseSrvEditCourseServer{stream})
}

type PersonalCourseSrv_EditCourseServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*PersonalCourseRequest, error)
	grpc.ServerStream
}

type personalCourseSrvEditCourseServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvEditCourseServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvEditCourseServer) Recv() (*PersonalCourseRequest, error) {
	m := new(PersonalCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_SetCourseShops_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).SetCourseShops(&personalCourseSrvSetCourseShopsServer{stream})
}

type PersonalCourseSrv_SetCourseShopsServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*SetCourseShopsRequest, error)
	grpc.ServerStream
}

type personalCourseSrvSetCourseShopsServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvSetCourseShopsServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseShopsServer) Recv() (*SetCourseShopsRequest, error) {
	m := new(SetCourseShopsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_SetSalePrice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).SetSalePrice(&personalCourseSrvSetSalePriceServer{stream})
}

type PersonalCourseSrv_SetSalePriceServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*SetSalePriceRequest, error)
	grpc.ServerStream
}

type personalCourseSrvSetSalePriceServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvSetSalePriceServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvSetSalePriceServer) Recv() (*SetSalePriceRequest, error) {
	m := new(SetSalePriceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_DelCourse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).DelCourse(&personalCourseSrvDelCourseServer{stream})
}

type PersonalCourseSrv_DelCourseServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*SetCourseStateRequest, error)
	grpc.ServerStream
}

type personalCourseSrvDelCourseServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvDelCourseServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvDelCourseServer) Recv() (*SetCourseStateRequest, error) {
	m := new(SetCourseStateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_SetCourseInvalid_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).SetCourseInvalid(&personalCourseSrvSetCourseInvalidServer{stream})
}

type PersonalCourseSrv_SetCourseInvalidServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*SetCourseStateRequest, error)
	grpc.ServerStream
}

type personalCourseSrvSetCourseInvalidServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvSetCourseInvalidServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseInvalidServer) Recv() (*SetCourseStateRequest, error) {
	m := new(SetCourseStateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonalCourseSrv_SetCourseValid_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonalCourseSrvServer).SetCourseValid(&personalCourseSrvSetCourseValidServer{stream})
}

type PersonalCourseSrv_SetCourseValidServer interface {
	Send(*PersonalCourseResponse) error
	Recv() (*SetCourseStateRequest, error)
	grpc.ServerStream
}

type personalCourseSrvSetCourseValidServer struct {
	grpc.ServerStream
}

func (x *personalCourseSrvSetCourseValidServer) Send(m *PersonalCourseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personalCourseSrvSetCourseValidServer) Recv() (*SetCourseStateRequest, error) {
	m := new(SetCourseStateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PersonalCourseSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.PersonalCourseSrv",
	HandlerType: (*PersonalCourseSrvServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddCourse",
			Handler:       _PersonalCourseSrv_AddCourse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "EditCourse",
			Handler:       _PersonalCourseSrv_EditCourse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SetCourseShops",
			Handler:       _PersonalCourseSrv_SetCourseShops_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SetSalePrice",
			Handler:       _PersonalCourseSrv_SetSalePrice_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DelCourse",
			Handler:       _PersonalCourseSrv_DelCourse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SetCourseInvalid",
			Handler:       _PersonalCourseSrv_SetCourseInvalid_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SetCourseValid",
			Handler:       _PersonalCourseSrv_SetCourseValid_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "personal_course_tpl.proto",
}

func init() {
	proto.RegisterFile("personal_course_tpl.proto", fileDescriptor_personal_course_tpl_a706bec0ed83f88f)
}

var fileDescriptor_personal_course_tpl_a706bec0ed83f88f = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0xd5, 0x3a, 0x89, 0x1d, 0x5f, 0x7f, 0xb4, 0x0c, 0x4d, 0xbb, 0x4d, 0x81, 0x1a, 0x03, 0xc2,
	0x0f, 0xe0, 0xa0, 0xf4, 0x0d, 0x5e, 0xda, 0x06, 0xa4, 0x5a, 0x82, 0x62, 0xd6, 0x04, 0x24, 0x84,
	0xb4, 0x9a, 0xec, 0xdc, 0xd8, 0x23, 0x76, 0x67, 0x96, 0x99, 0xd9, 0xa8, 0x79, 0xe6, 0xdf, 0xf1,
	0x0f, 0x78, 0xe1, 0x89, 0xff, 0xc0, 0x23, 0x68, 0x3e, 0x6c, 0xaf, 0xe3, 0x14, 0x45, 0x55, 0x78,
	0x88, 0xac, 0x7b, 0xce, 0x9d, 0x3b, 0xf7, 0xce, 0x3d, 0xc7, 0x31, 0x3c, 0x2c, 0x51, 0x69, 0x29,
	0x68, 0x9e, 0x66, 0xb2, 0x52, 0x1a, 0x53, 0x53, 0xe6, 0xe3, 0x52, 0x49, 0x23, 0x49, 0xab, 0x54,
	0x92, 0x55, 0x99, 0x39, 0xec, 0x66, 0xb2, 0x28, 0xa4, 0xf0, 0xf0, 0x61, 0x2f, 0xc0, 0x3e, 0x1c,
	0xfe, 0x16, 0xc1, 0xc1, 0x0c, 0xcd, 0x89, 0x3b, 0x3d, 0x33, 0xd4, 0x60, 0x82, 0xbf, 0x56, 0xa8,
	0x0d, 0xf9, 0x04, 0x9a, 0x0b, 0xa4, 0x0c, 0x55, 0x1c, 0x0d, 0xa2, 0x51, 0xe7, 0xf8, 0xde, 0x38,
	0xd4, 0x39, 0x71, 0x1f, 0x2f, 0x1c, 0x97, 0x84, 0x1c, 0xf2, 0x08, 0xda, 0xa1, 0x03, 0xce, 0xe2,
	0xc6, 0x20, 0x1a, 0xf5, 0x92, 0x7d, 0x0f, 0x4c, 0x18, 0x79, 0x0c, 0x1d, 0x59, 0xa2, 0xa2, 0x46,
	0x2a, 0x4b, 0xef, 0x38, 0x1a, 0x96, 0xd0, 0x84, 0x0d, 0xff, 0xd8, 0x81, 0x83, 0x69, 0x98, 0xc4,
	0xb7, 0xf2, 0x66, 0x5d, 0x3c, 0x84, 0xfd, 0x33, 0x45, 0x05, 0x5b, 0x37, 0xd1, 0x72, 0xf1, 0x84,
	0x91, 0x07, 0xd0, 0xd2, 0x0b, 0x59, 0xae, 0xef, 0x6f, 0xda, 0x70, 0xc2, 0x36, 0x3b, 0xdf, 0xdd,
	0xee, 0x3c, 0x90, 0x82, 0x16, 0x18, 0xef, 0x0d, 0xa2, 0x51, 0x3b, 0x01, 0x0f, 0xbd, 0xa4, 0x05,
	0x92, 0x8f, 0xe1, 0x4e, 0x48, 0xc8, 0xa8, 0xc1, 0xb9, 0x54, 0x97, 0x71, 0xd3, 0xd5, 0xe8, 0x7b,
	0xf8, 0x24, 0xa0, 0xf6, 0x1a, 0xa3, 0x28, 0x17, 0x29, 0xe5, 0x45, 0xdc, 0x1a, 0xec, 0xd8, 0x6b,
	0x1c, 0xf0, 0x8c, 0x17, 0xe4, 0x10, 0xf6, 0x59, 0xa5, 0xa8, 0xe1, 0x52, 0xc4, 0xfb, 0xbe, 0x85,
	0x65, 0x4c, 0xee, 0xc1, 0x5e, 0xa9, 0x78, 0x86, 0x71, 0xdb, 0x11, 0x3e, 0x20, 0xef, 0x02, 0x2c,
	0xbb, 0x2e, 0xe6, 0x31, 0xb8, 0xbe, 0xc2, 0x1c, 0x93, 0x62, 0x4e, 0x06, 0xd0, 0x61, 0xa8, 0x33,
	0xc5, 0x4b, 0x57, 0xb3, 0xe3, 0xf8, 0x3a, 0x44, 0x3e, 0x82, 0x3e, 0x9e, 0x9f, 0x63, 0x66, 0xf8,
	0x05, 0xa6, 0x95, 0xe0, 0x26, 0xee, 0xba, 0xfa, 0xbd, 0x15, 0x7a, 0x2a, 0xb8, 0x21, 0x4f, 0xe1,
	0x4e, 0x59, 0x9d, 0xe5, 0x5c, 0x2f, 0xd2, 0x6c, 0x41, 0x85, 0xc0, 0x3c, 0xee, 0x0d, 0xa2, 0x51,
	0xff, 0xf8, 0xc1, 0x78, 0x29, 0xa4, 0xa9, 0xe7, 0x4f, 0x3c, 0x9d, 0xf4, 0xcb, 0x8d, 0x78, 0x88,
	0x70, 0xff, 0xea, 0x6a, 0x75, 0x29, 0x85, 0x46, 0x32, 0x86, 0xa6, 0x36, 0xd4, 0x54, 0x3a, 0xec,
	0xf6, 0xfe, 0x72, 0xb7, 0xcf, 0x2b, 0xcd, 0x05, 0x6a, 0x3d, 0x73, 0x6c, 0x12, 0xb2, 0xfe, 0x53,
	0x63, 0xc3, 0x3f, 0x37, 0x84, 0xbc, 0x90, 0xa5, 0xfe, 0x1f, 0x84, 0xfc, 0x3e, 0x74, 0x9d, 0x88,
	0x34, 0x1a, 0xc3, 0xc5, 0x3c, 0x28, 0xa9, 0x63, 0xb1, 0x99, 0x87, 0xac, 0x04, 0x83, 0xce, 0x74,
	0xbc, 0xeb, 0xd6, 0xdc, 0xf2, 0x42, 0x0b, 0xfd, 0xd3, 0x6c, 0xe1, 0xb8, 0x3d, 0x2f, 0x01, 0x07,
	0x58, 0xf2, 0x8a, 0x47, 0x9a, 0x5b, 0x1e, 0xf9, 0x3b, 0x82, 0xb7, 0x67, 0x68, 0x66, 0x34, 0xc7,
	0xa9, 0x95, 0xc0, 0x9b, 0x8d, 0x57, 0xb3, 0x41, 0xe3, 0xf5, 0x36, 0xd8, 0xb9, 0x32, 0xf7, 0xe7,
	0xd0, 0x73, 0xb2, 0x5b, 0x0d, 0xbe, 0xeb, 0x34, 0x70, 0xb0, 0xd2, 0xc0, 0x0c, 0x8d, 0x6b, 0xea,
	0xfb, 0xcb, 0x12, 0x93, 0xae, 0xcb, 0x5d, 0x3e, 0xc8, 0x53, 0xe8, 0x85, 0xc2, 0x0e, 0xf6, 0x93,
	0x77, 0x8e, 0x1f, 0xad, 0xce, 0xfa, 0x95, 0x4d, 0x6b, 0x67, 0x92, 0x6e, 0xb6, 0xc6, 0xf4, 0xf0,
	0x9f, 0x08, 0xc8, 0x76, 0x12, 0xe9, 0x43, 0x83, 0x33, 0x37, 0x74, 0x2f, 0x69, 0xf0, 0xe0, 0x55,
	0xfb, 0xbc, 0x39, 0x5e, 0x60, 0x1e, 0xc6, 0x03, 0x07, 0x7d, 0x6d, 0x11, 0xb7, 0x1a, 0x9a, 0x63,
	0x5a, 0x70, 0x11, 0x26, 0x6c, 0xd9, 0xf8, 0x1b, 0x2e, 0xd6, 0x14, 0x7d, 0x15, 0xbe, 0x03, 0x3c,
	0x45, 0x5f, 0xad, 0xfd, 0xb7, 0x57, 0xf7, 0xdf, 0x07, 0xd0, 0x33, 0x8a, 0x0a, 0x7d, 0x8e, 0x2a,
	0x35, 0x97, 0x25, 0x86, 0x85, 0x75, 0x97, 0xa0, 0x7d, 0x08, 0x2b, 0x97, 0x55, 0x92, 0xa8, 0xac,
	0xed, 0x9d, 0x5c, 0x96, 0xd8, 0xcb, 0xaa, 0x20, 0x1f, 0x42, 0x9f, 0xeb, 0x54, 0x8a, 0x9c, 0x0b,
	0x4c, 0xed, 0x95, 0xc1, 0xff, 0x5d, 0xae, 0xbf, 0x75, 0xa0, 0xdd, 0xf7, 0xf1, 0x5f, 0xbb, 0xf0,
	0xd6, 0xa6, 0x89, 0x66, 0xea, 0x82, 0x4c, 0xa1, 0xfd, 0x8c, 0x31, 0x1f, 0x93, 0xf7, 0xd6, 0x7e,
	0xbc, 0xee, 0x8b, 0xf4, 0xf0, 0xf1, 0x6b, 0x79, 0xef, 0xc6, 0x51, 0xf4, 0x59, 0x44, 0xbe, 0x03,
	0xf8, 0x8a, 0x71, 0x73, 0x9b, 0x25, 0x4f, 0xa1, 0xbf, 0x69, 0xcb, 0x5a, 0xd9, 0x6b, 0xfd, 0x7a,
	0xd3, 0x4e, 0xbb, 0x75, 0x33, 0x90, 0x77, 0xea, 0x45, 0xaf, 0x7a, 0xe4, 0x66, 0x25, 0xa7, 0xd0,
	0xfe, 0x12, 0xf3, 0xad, 0xd9, 0xaf, 0xfd, 0xef, 0x78, 0xb3, 0x8a, 0x3f, 0xc2, 0xdd, 0xd5, 0xe9,
	0x89, 0xb8, 0xa0, 0x39, 0x67, 0xb7, 0x53, 0xb8, 0xfe, 0xa8, 0x3f, 0xdc, 0x5a, 0xd9, 0xe7, 0x2f,
	0x7e, 0x7a, 0x32, 0xe7, 0x26, 0xa7, 0x67, 0x63, 0x6d, 0x2e, 0xd9, 0x38, 0x13, 0x47, 0x95, 0x91,
	0xf6, 0xef, 0x53, 0xcd, 0x7e, 0x39, 0x9a, 0xcb, 0x23, 0xf7, 0x9b, 0xe1, 0x28, 0x54, 0xfa, 0x22,
	0x7c, 0xfe, 0xde, 0xb8, 0x7b, 0x6a, 0xe4, 0xd4, 0x52, 0x3f, 0x4f, 0x3d, 0x74, 0xd6, 0x74, 0x99,
	0x4f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x0b, 0x50, 0x4b, 0xa0, 0x08, 0x00, 0x00,
}