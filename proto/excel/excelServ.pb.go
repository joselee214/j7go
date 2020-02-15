// Code generated by protoc-gen-go. DO NOT EDIT.
// source: excelServ.proto

package excel // import "j7go/proto/excel"

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

// Excel保存在哪里
type ExcelSaveType int32

const (
	// 阿里oss
	ExcelSaveType_OSS ExcelSaveType = 0
	// 服务本地
	ExcelSaveType_LOCAL ExcelSaveType = 1
)

var ExcelSaveType_name = map[int32]string{
	0: "OSS",
	1: "LOCAL",
}
var ExcelSaveType_value = map[string]int32{
	"OSS":   0,
	"LOCAL": 1,
}

func (x ExcelSaveType) String() string {
	return proto.EnumName(ExcelSaveType_name, int32(x))
}
func (ExcelSaveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{0}
}

type DataToExcelRequest struct {
	Header               *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FileName             string               `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ExcelBuildData       *ExcelBuildData      `protobuf:"bytes,3,opt,name=excel_build_data,json=excelBuildData,proto3" json:"excel_build_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataToExcelRequest) Reset()         { *m = DataToExcelRequest{} }
func (m *DataToExcelRequest) String() string { return proto.CompactTextString(m) }
func (*DataToExcelRequest) ProtoMessage()    {}
func (*DataToExcelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{0}
}
func (m *DataToExcelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataToExcelRequest.Unmarshal(m, b)
}
func (m *DataToExcelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataToExcelRequest.Marshal(b, m, deterministic)
}
func (dst *DataToExcelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataToExcelRequest.Merge(dst, src)
}
func (m *DataToExcelRequest) XXX_Size() int {
	return xxx_messageInfo_DataToExcelRequest.Size(m)
}
func (m *DataToExcelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataToExcelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataToExcelRequest proto.InternalMessageInfo

func (m *DataToExcelRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DataToExcelRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *DataToExcelRequest) GetExcelBuildData() *ExcelBuildData {
	if m != nil {
		return m.ExcelBuildData
	}
	return nil
}

type DataToExcelResponse struct {
	Status *common.BusinessStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 保存类型
	Type ExcelSaveType `protobuf:"varint,2,opt,name=type,proto3,enum=utoProto.excel.ExcelSaveType" json:"type,omitempty"`
	// 获取资源的url
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataToExcelResponse) Reset()         { *m = DataToExcelResponse{} }
func (m *DataToExcelResponse) String() string { return proto.CompactTextString(m) }
func (*DataToExcelResponse) ProtoMessage()    {}
func (*DataToExcelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{1}
}
func (m *DataToExcelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataToExcelResponse.Unmarshal(m, b)
}
func (m *DataToExcelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataToExcelResponse.Marshal(b, m, deterministic)
}
func (dst *DataToExcelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataToExcelResponse.Merge(dst, src)
}
func (m *DataToExcelResponse) XXX_Size() int {
	return xxx_messageInfo_DataToExcelResponse.Size(m)
}
func (m *DataToExcelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataToExcelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataToExcelResponse proto.InternalMessageInfo

func (m *DataToExcelResponse) GetStatus() *common.BusinessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DataToExcelResponse) GetType() ExcelSaveType {
	if m != nil {
		return m.Type
	}
	return ExcelSaveType_OSS
}

func (m *DataToExcelResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ExcelToDataRequest struct {
	Header *common.CommonHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 保存类型
	Type ExcelSaveType `protobuf:"varint,2,opt,name=type,proto3,enum=utoProto.excel.ExcelSaveType" json:"type,omitempty"`
	// 获取资源的url
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExcelToDataRequest) Reset()         { *m = ExcelToDataRequest{} }
func (m *ExcelToDataRequest) String() string { return proto.CompactTextString(m) }
func (*ExcelToDataRequest) ProtoMessage()    {}
func (*ExcelToDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{2}
}
func (m *ExcelToDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelToDataRequest.Unmarshal(m, b)
}
func (m *ExcelToDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelToDataRequest.Marshal(b, m, deterministic)
}
func (dst *ExcelToDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelToDataRequest.Merge(dst, src)
}
func (m *ExcelToDataRequest) XXX_Size() int {
	return xxx_messageInfo_ExcelToDataRequest.Size(m)
}
func (m *ExcelToDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelToDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelToDataRequest proto.InternalMessageInfo

func (m *ExcelToDataRequest) GetHeader() *common.CommonHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ExcelToDataRequest) GetType() ExcelSaveType {
	if m != nil {
		return m.Type
	}
	return ExcelSaveType_OSS
}

func (m *ExcelToDataRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ExcelToDataResponse struct {
	Status               *common.BusinessStatus       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Sheets               []*ExcelToDataResponse_Sheet `protobuf:"bytes,2,rep,name=sheets,proto3" json:"sheets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ExcelToDataResponse) Reset()         { *m = ExcelToDataResponse{} }
func (m *ExcelToDataResponse) String() string { return proto.CompactTextString(m) }
func (*ExcelToDataResponse) ProtoMessage()    {}
func (*ExcelToDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{3}
}
func (m *ExcelToDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelToDataResponse.Unmarshal(m, b)
}
func (m *ExcelToDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelToDataResponse.Marshal(b, m, deterministic)
}
func (dst *ExcelToDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelToDataResponse.Merge(dst, src)
}
func (m *ExcelToDataResponse) XXX_Size() int {
	return xxx_messageInfo_ExcelToDataResponse.Size(m)
}
func (m *ExcelToDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelToDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelToDataResponse proto.InternalMessageInfo

func (m *ExcelToDataResponse) GetStatus() *common.BusinessStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ExcelToDataResponse) GetSheets() []*ExcelToDataResponse_Sheet {
	if m != nil {
		return m.Sheets
	}
	return nil
}

type ExcelToDataResponse_Col struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Style                string   `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExcelToDataResponse_Col) Reset()         { *m = ExcelToDataResponse_Col{} }
func (m *ExcelToDataResponse_Col) String() string { return proto.CompactTextString(m) }
func (*ExcelToDataResponse_Col) ProtoMessage()    {}
func (*ExcelToDataResponse_Col) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{3, 0}
}
func (m *ExcelToDataResponse_Col) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelToDataResponse_Col.Unmarshal(m, b)
}
func (m *ExcelToDataResponse_Col) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelToDataResponse_Col.Marshal(b, m, deterministic)
}
func (dst *ExcelToDataResponse_Col) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelToDataResponse_Col.Merge(dst, src)
}
func (m *ExcelToDataResponse_Col) XXX_Size() int {
	return xxx_messageInfo_ExcelToDataResponse_Col.Size(m)
}
func (m *ExcelToDataResponse_Col) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelToDataResponse_Col.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelToDataResponse_Col proto.InternalMessageInfo

func (m *ExcelToDataResponse_Col) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ExcelToDataResponse_Col) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

type ExcelToDataResponse_Row struct {
	Row                  []*ExcelToDataResponse_Col `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExcelToDataResponse_Row) Reset()         { *m = ExcelToDataResponse_Row{} }
func (m *ExcelToDataResponse_Row) String() string { return proto.CompactTextString(m) }
func (*ExcelToDataResponse_Row) ProtoMessage()    {}
func (*ExcelToDataResponse_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{3, 1}
}
func (m *ExcelToDataResponse_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelToDataResponse_Row.Unmarshal(m, b)
}
func (m *ExcelToDataResponse_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelToDataResponse_Row.Marshal(b, m, deterministic)
}
func (dst *ExcelToDataResponse_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelToDataResponse_Row.Merge(dst, src)
}
func (m *ExcelToDataResponse_Row) XXX_Size() int {
	return xxx_messageInfo_ExcelToDataResponse_Row.Size(m)
}
func (m *ExcelToDataResponse_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelToDataResponse_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelToDataResponse_Row proto.InternalMessageInfo

func (m *ExcelToDataResponse_Row) GetRow() []*ExcelToDataResponse_Col {
	if m != nil {
		return m.Row
	}
	return nil
}

type ExcelToDataResponse_Sheet struct {
	SheetName            string                     `protobuf:"bytes,1,opt,name=sheet_name,json=sheetName,proto3" json:"sheet_name,omitempty"`
	Rows                 []*ExcelToDataResponse_Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExcelToDataResponse_Sheet) Reset()         { *m = ExcelToDataResponse_Sheet{} }
func (m *ExcelToDataResponse_Sheet) String() string { return proto.CompactTextString(m) }
func (*ExcelToDataResponse_Sheet) ProtoMessage()    {}
func (*ExcelToDataResponse_Sheet) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{3, 2}
}
func (m *ExcelToDataResponse_Sheet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelToDataResponse_Sheet.Unmarshal(m, b)
}
func (m *ExcelToDataResponse_Sheet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelToDataResponse_Sheet.Marshal(b, m, deterministic)
}
func (dst *ExcelToDataResponse_Sheet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelToDataResponse_Sheet.Merge(dst, src)
}
func (m *ExcelToDataResponse_Sheet) XXX_Size() int {
	return xxx_messageInfo_ExcelToDataResponse_Sheet.Size(m)
}
func (m *ExcelToDataResponse_Sheet) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelToDataResponse_Sheet.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelToDataResponse_Sheet proto.InternalMessageInfo

func (m *ExcelToDataResponse_Sheet) GetSheetName() string {
	if m != nil {
		return m.SheetName
	}
	return ""
}

func (m *ExcelToDataResponse_Sheet) GetRows() []*ExcelToDataResponse_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ExcelBuildData struct {
	Sheets               []*ExcelBuildData_Sheet `protobuf:"bytes,1,rep,name=sheets,proto3" json:"sheets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ExcelBuildData) Reset()         { *m = ExcelBuildData{} }
func (m *ExcelBuildData) String() string { return proto.CompactTextString(m) }
func (*ExcelBuildData) ProtoMessage()    {}
func (*ExcelBuildData) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{4}
}
func (m *ExcelBuildData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelBuildData.Unmarshal(m, b)
}
func (m *ExcelBuildData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelBuildData.Marshal(b, m, deterministic)
}
func (dst *ExcelBuildData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelBuildData.Merge(dst, src)
}
func (m *ExcelBuildData) XXX_Size() int {
	return xxx_messageInfo_ExcelBuildData.Size(m)
}
func (m *ExcelBuildData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelBuildData.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelBuildData proto.InternalMessageInfo

func (m *ExcelBuildData) GetSheets() []*ExcelBuildData_Sheet {
	if m != nil {
		return m.Sheets
	}
	return nil
}

type ExcelBuildData_Col struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Style                string   `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExcelBuildData_Col) Reset()         { *m = ExcelBuildData_Col{} }
func (m *ExcelBuildData_Col) String() string { return proto.CompactTextString(m) }
func (*ExcelBuildData_Col) ProtoMessage()    {}
func (*ExcelBuildData_Col) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{4, 0}
}
func (m *ExcelBuildData_Col) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelBuildData_Col.Unmarshal(m, b)
}
func (m *ExcelBuildData_Col) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelBuildData_Col.Marshal(b, m, deterministic)
}
func (dst *ExcelBuildData_Col) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelBuildData_Col.Merge(dst, src)
}
func (m *ExcelBuildData_Col) XXX_Size() int {
	return xxx_messageInfo_ExcelBuildData_Col.Size(m)
}
func (m *ExcelBuildData_Col) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelBuildData_Col.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelBuildData_Col proto.InternalMessageInfo

func (m *ExcelBuildData_Col) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ExcelBuildData_Col) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

type ExcelBuildData_Row struct {
	Row                  []*ExcelBuildData_Col `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExcelBuildData_Row) Reset()         { *m = ExcelBuildData_Row{} }
func (m *ExcelBuildData_Row) String() string { return proto.CompactTextString(m) }
func (*ExcelBuildData_Row) ProtoMessage()    {}
func (*ExcelBuildData_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{4, 1}
}
func (m *ExcelBuildData_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelBuildData_Row.Unmarshal(m, b)
}
func (m *ExcelBuildData_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelBuildData_Row.Marshal(b, m, deterministic)
}
func (dst *ExcelBuildData_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelBuildData_Row.Merge(dst, src)
}
func (m *ExcelBuildData_Row) XXX_Size() int {
	return xxx_messageInfo_ExcelBuildData_Row.Size(m)
}
func (m *ExcelBuildData_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelBuildData_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelBuildData_Row proto.InternalMessageInfo

func (m *ExcelBuildData_Row) GetRow() []*ExcelBuildData_Col {
	if m != nil {
		return m.Row
	}
	return nil
}

type ExcelBuildData_Sheet struct {
	SheetName            string                `protobuf:"bytes,1,opt,name=sheet_name,json=sheetName,proto3" json:"sheet_name,omitempty"`
	Rows                 []*ExcelBuildData_Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExcelBuildData_Sheet) Reset()         { *m = ExcelBuildData_Sheet{} }
func (m *ExcelBuildData_Sheet) String() string { return proto.CompactTextString(m) }
func (*ExcelBuildData_Sheet) ProtoMessage()    {}
func (*ExcelBuildData_Sheet) Descriptor() ([]byte, []int) {
	return fileDescriptor_excelServ_b493313b28f37312, []int{4, 2}
}
func (m *ExcelBuildData_Sheet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcelBuildData_Sheet.Unmarshal(m, b)
}
func (m *ExcelBuildData_Sheet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcelBuildData_Sheet.Marshal(b, m, deterministic)
}
func (dst *ExcelBuildData_Sheet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcelBuildData_Sheet.Merge(dst, src)
}
func (m *ExcelBuildData_Sheet) XXX_Size() int {
	return xxx_messageInfo_ExcelBuildData_Sheet.Size(m)
}
func (m *ExcelBuildData_Sheet) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcelBuildData_Sheet.DiscardUnknown(m)
}

var xxx_messageInfo_ExcelBuildData_Sheet proto.InternalMessageInfo

func (m *ExcelBuildData_Sheet) GetSheetName() string {
	if m != nil {
		return m.SheetName
	}
	return ""
}

func (m *ExcelBuildData_Sheet) GetRows() []*ExcelBuildData_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*DataToExcelRequest)(nil), "utoProto.excel.DataToExcelRequest")
	proto.RegisterType((*DataToExcelResponse)(nil), "utoProto.excel.DataToExcelResponse")
	proto.RegisterType((*ExcelToDataRequest)(nil), "utoProto.excel.ExcelToDataRequest")
	proto.RegisterType((*ExcelToDataResponse)(nil), "utoProto.excel.ExcelToDataResponse")
	proto.RegisterType((*ExcelToDataResponse_Col)(nil), "utoProto.excel.ExcelToDataResponse.Col")
	proto.RegisterType((*ExcelToDataResponse_Row)(nil), "utoProto.excel.ExcelToDataResponse.Row")
	proto.RegisterType((*ExcelToDataResponse_Sheet)(nil), "utoProto.excel.ExcelToDataResponse.Sheet")
	proto.RegisterType((*ExcelBuildData)(nil), "utoProto.excel.ExcelBuildData")
	proto.RegisterType((*ExcelBuildData_Col)(nil), "utoProto.excel.ExcelBuildData.Col")
	proto.RegisterType((*ExcelBuildData_Row)(nil), "utoProto.excel.ExcelBuildData.Row")
	proto.RegisterType((*ExcelBuildData_Sheet)(nil), "utoProto.excel.ExcelBuildData.Sheet")
	proto.RegisterEnum("utoProto.excel.ExcelSaveType", ExcelSaveType_name, ExcelSaveType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExcelServiceClient is the client API for ExcelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExcelServiceClient interface {
	// 通过data返回Excel保存路径
	DataToExcel(ctx context.Context, opts ...grpc.CallOption) (ExcelService_DataToExcelClient, error)
	// 通过Excel返回excel里边的data
	ExcelToData(ctx context.Context, opts ...grpc.CallOption) (ExcelService_ExcelToDataClient, error)
}

type excelServiceClient struct {
	cc *grpc.ClientConn
}

func NewExcelServiceClient(cc *grpc.ClientConn) ExcelServiceClient {
	return &excelServiceClient{cc}
}

func (c *excelServiceClient) DataToExcel(ctx context.Context, opts ...grpc.CallOption) (ExcelService_DataToExcelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExcelService_serviceDesc.Streams[0], "/utoProto.excel.ExcelService/DataToExcel", opts...)
	if err != nil {
		return nil, err
	}
	x := &excelServiceDataToExcelClient{stream}
	return x, nil
}

type ExcelService_DataToExcelClient interface {
	Send(*DataToExcelRequest) error
	Recv() (*DataToExcelResponse, error)
	grpc.ClientStream
}

type excelServiceDataToExcelClient struct {
	grpc.ClientStream
}

func (x *excelServiceDataToExcelClient) Send(m *DataToExcelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *excelServiceDataToExcelClient) Recv() (*DataToExcelResponse, error) {
	m := new(DataToExcelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *excelServiceClient) ExcelToData(ctx context.Context, opts ...grpc.CallOption) (ExcelService_ExcelToDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExcelService_serviceDesc.Streams[1], "/utoProto.excel.ExcelService/ExcelToData", opts...)
	if err != nil {
		return nil, err
	}
	x := &excelServiceExcelToDataClient{stream}
	return x, nil
}

type ExcelService_ExcelToDataClient interface {
	Send(*ExcelToDataRequest) error
	Recv() (*ExcelToDataResponse, error)
	grpc.ClientStream
}

type excelServiceExcelToDataClient struct {
	grpc.ClientStream
}

func (x *excelServiceExcelToDataClient) Send(m *ExcelToDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *excelServiceExcelToDataClient) Recv() (*ExcelToDataResponse, error) {
	m := new(ExcelToDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExcelServiceServer is the server API for ExcelService service.
type ExcelServiceServer interface {
	// 通过data返回Excel保存路径
	DataToExcel(ExcelService_DataToExcelServer) error
	// 通过Excel返回excel里边的data
	ExcelToData(ExcelService_ExcelToDataServer) error
}

func RegisterExcelServiceServer(s *grpc.Server, srv ExcelServiceServer) {
	s.RegisterService(&_ExcelService_serviceDesc, srv)
}

func _ExcelService_DataToExcel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExcelServiceServer).DataToExcel(&excelServiceDataToExcelServer{stream})
}

type ExcelService_DataToExcelServer interface {
	Send(*DataToExcelResponse) error
	Recv() (*DataToExcelRequest, error)
	grpc.ServerStream
}

type excelServiceDataToExcelServer struct {
	grpc.ServerStream
}

func (x *excelServiceDataToExcelServer) Send(m *DataToExcelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *excelServiceDataToExcelServer) Recv() (*DataToExcelRequest, error) {
	m := new(DataToExcelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExcelService_ExcelToData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExcelServiceServer).ExcelToData(&excelServiceExcelToDataServer{stream})
}

type ExcelService_ExcelToDataServer interface {
	Send(*ExcelToDataResponse) error
	Recv() (*ExcelToDataRequest, error)
	grpc.ServerStream
}

type excelServiceExcelToDataServer struct {
	grpc.ServerStream
}

func (x *excelServiceExcelToDataServer) Send(m *ExcelToDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *excelServiceExcelToDataServer) Recv() (*ExcelToDataRequest, error) {
	m := new(ExcelToDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExcelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "utoProto.excel.ExcelService",
	HandlerType: (*ExcelServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataToExcel",
			Handler:       _ExcelService_DataToExcel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ExcelToData",
			Handler:       _ExcelService_ExcelToData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "excelServ.proto",
}

func init() { proto.RegisterFile("excelServ.proto", fileDescriptor_excelServ_b493313b28f37312) }

var fileDescriptor_excelServ_b493313b28f37312 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6e, 0x12, 0x41,
	0x14, 0xc6, 0x1d, 0xb6, 0x45, 0xf7, 0x50, 0x91, 0x4c, 0x1b, 0x43, 0x30, 0x35, 0x64, 0x6b, 0x22,
	0x1a, 0x5d, 0x2c, 0x1a, 0x13, 0x83, 0x17, 0x16, 0x6c, 0xd2, 0x8b, 0xc6, 0x9a, 0x59, 0xbc, 0x69,
	0x8c, 0x64, 0x80, 0xb1, 0xdd, 0xb8, 0xec, 0xe0, 0xce, 0x2c, 0xc8, 0x13, 0x98, 0xf8, 0x10, 0xde,
	0xfb, 0x22, 0x5e, 0xf8, 0x18, 0x3e, 0x89, 0x99, 0xb3, 0x0b, 0x61, 0x53, 0xa4, 0xa4, 0x7a, 0x01,
	0xec, 0xcc, 0xf9, 0xf3, 0xfd, 0xf8, 0xe6, 0xec, 0xc0, 0x2d, 0xf1, 0xa5, 0x2f, 0x02, 0x4f, 0x44,
	0x63, 0x77, 0x14, 0x49, 0x2d, 0x69, 0x31, 0xd6, 0xf2, 0xad, 0x79, 0x72, 0x31, 0x52, 0xd9, 0xea,
	0xcb, 0xe1, 0x50, 0x86, 0x49, 0xd4, 0xf9, 0x41, 0x80, 0xbe, 0xe6, 0x9a, 0x77, 0xe4, 0xa1, 0x89,
	0x32, 0xf1, 0x39, 0x16, 0x4a, 0xd3, 0x47, 0x90, 0x3f, 0x17, 0x7c, 0x20, 0xa2, 0x32, 0xa9, 0x92,
	0x5a, 0xa1, 0xb1, 0xe3, 0xa6, 0x55, 0x6d, 0xfc, 0x39, 0xc2, 0x18, 0x4b, 0x73, 0xe8, 0x1d, 0xb0,
	0x3f, 0xfa, 0x81, 0xe8, 0x86, 0x7c, 0x28, 0xca, 0xb9, 0x2a, 0xa9, 0xd9, 0xec, 0x86, 0xd9, 0x78,
	0xc3, 0x87, 0x82, 0x1e, 0x41, 0x09, 0x85, 0xbb, 0xbd, 0xd8, 0x0f, 0x06, 0xdd, 0x01, 0xd7, 0xbc,
	0x6c, 0x61, 0xd3, 0xbb, 0x6e, 0x16, 0xcd, 0x45, 0x84, 0x96, 0x49, 0x33, 0x48, 0xac, 0x28, 0x32,
	0x6b, 0xe7, 0x1b, 0x81, 0xed, 0x0c, 0xab, 0x1a, 0xc9, 0x50, 0x09, 0xea, 0x42, 0x5e, 0x69, 0xae,
	0x63, 0x95, 0xc2, 0xde, 0x9e, 0xc1, 0xb6, 0x62, 0xe5, 0x87, 0x42, 0x29, 0x0f, 0xa3, 0x2c, 0xcd,
	0xa2, 0xfb, 0xb0, 0xa1, 0xa7, 0xa3, 0x84, 0xb4, 0xd8, 0xd8, 0x5d, 0x4a, 0xe1, 0xf1, 0xb1, 0xe8,
	0x4c, 0x47, 0x82, 0x61, 0x2a, 0x2d, 0x81, 0x15, 0x47, 0x01, 0x72, 0xdb, 0xcc, 0x3c, 0x3a, 0x5f,
	0x09, 0x50, 0xcc, 0xec, 0x48, 0x84, 0xbd, 0x92, 0x71, 0xff, 0x85, 0xe4, 0x77, 0x0e, 0xb6, 0x33,
	0x24, 0x57, 0xb4, 0xe5, 0x00, 0xf2, 0xea, 0x5c, 0x08, 0xad, 0xca, 0xb9, 0xaa, 0x55, 0x2b, 0x34,
	0x1e, 0x2c, 0xc5, 0xc9, 0x8a, 0xb8, 0x9e, 0xa9, 0x60, 0x69, 0x61, 0x65, 0x1f, 0xac, 0xb6, 0x0c,
	0xe8, 0x0e, 0x6c, 0x8e, 0x79, 0x10, 0x0b, 0x14, 0xb6, 0x59, 0xb2, 0x30, 0xbb, 0x4a, 0x4f, 0x83,
	0xd9, 0x84, 0x24, 0x8b, 0xca, 0x2b, 0xb0, 0x98, 0x9c, 0xd0, 0x17, 0x60, 0x45, 0x72, 0x52, 0x26,
	0xa8, 0x7c, 0x7f, 0x1d, 0xe5, 0xb6, 0x0c, 0x98, 0xa9, 0xa9, 0xf4, 0x61, 0x13, 0x29, 0xe8, 0x2e,
	0x00, 0x72, 0x24, 0x73, 0x98, 0x68, 0xdb, 0xb8, 0x83, 0x83, 0xd8, 0x84, 0x8d, 0x48, 0x4e, 0x66,
	0xff, 0x6e, 0x2d, 0x0d, 0x26, 0x27, 0x0c, 0x8b, 0x9c, 0xef, 0x39, 0x28, 0x66, 0xc7, 0x93, 0xbe,
	0x9c, 0xfb, 0x95, 0x50, 0xdf, 0x5b, 0x3d, 0xce, 0xff, 0x6e, 0x55, 0x33, 0xb1, 0xea, 0xd9, 0xa2,
	0x55, 0xce, 0x25, 0xa2, 0x73, 0x97, 0x3e, 0xac, 0xe9, 0xd2, 0xf3, 0x8c, 0x4b, 0x97, 0xb5, 0x9f,
	0x1b, 0xf4, 0x70, 0x0f, 0x6e, 0x66, 0xc6, 0x95, 0x5e, 0x07, 0xeb, 0xc4, 0xf3, 0x4a, 0xd7, 0xa8,
	0x0d, 0x9b, 0xc7, 0x27, 0xed, 0x83, 0xe3, 0x12, 0x69, 0xfc, 0x24, 0xb0, 0x75, 0x38, 0xbb, 0x9f,
	0xfc, 0xbe, 0xa0, 0xa7, 0x50, 0x58, 0x78, 0xa3, 0xe9, 0x05, 0xb9, 0x8b, 0x57, 0x53, 0x65, 0x6f,
	0x65, 0x4e, 0x72, 0x70, 0x35, 0xf2, 0x84, 0x98, 0xde, 0x0b, 0x67, 0x4a, 0x9d, 0x95, 0x07, 0xfe,
	0x97, 0xde, 0x4b, 0x86, 0xc2, 0xf4, 0x6e, 0xb5, 0x4e, 0xeb, 0x67, 0xbe, 0x0e, 0x78, 0xcf, 0x55,
	0x7a, 0x3a, 0x70, 0xfb, 0x61, 0x3d, 0xd6, 0xd2, 0x7c, 0x1e, 0xab, 0xc1, 0xa7, 0xfa, 0x99, 0xac,
	0xe3, 0xed, 0x5a, 0xc7, 0x26, 0x4d, 0xfc, 0xfe, 0x95, 0x2b, 0xbe, 0x4b, 0x7b, 0xbf, 0xc7, 0xae,
	0xbd, 0x3c, 0xe6, 0x3c, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x95, 0x28, 0xbc, 0xb2, 0x05,
	0x00, 0x00,
}
