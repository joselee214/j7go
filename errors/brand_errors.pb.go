// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brand_errors.proto

package business_errors

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

type BrandError int32

const (
	// 初始值, 无意义
	BrandError_INIT BrandError = 0
	// 获取品牌列表失败
	BrandError_GET_SHOP_LIST_ERROR BrandError = 60610
)

var BrandError_name = map[int32]string{
	0:     "INIT",
	60610: "GET_SHOP_LIST_ERROR",
}
var BrandError_value = map[string]int32{
	"INIT":                0,
	"GET_SHOP_LIST_ERROR": 60610,
}

func (x BrandError) String() string {
	return proto.EnumName(BrandError_name, int32(x))
}
func (BrandError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_brand_errors_aa5e6cbb56466a16, []int{0}
}

func init() {
	proto.RegisterEnum("product.BrandError", BrandError_name, BrandError_value)
}

func init() { proto.RegisterFile("brand_errors.proto", fileDescriptor_brand_errors_aa5e6cbb56466a16) }

var fileDescriptor_brand_errors_aa5e6cbb56466a16 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2a, 0x4a, 0xcc,
	0x4b, 0x89, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xd1, 0x32, 0xe4, 0xe2, 0x72, 0x02, 0x49, 0xbb, 0x82,
	0x64, 0x85, 0x38, 0xb8, 0x58, 0x3c, 0xfd, 0x3c, 0x43, 0x04, 0x18, 0x84, 0x24, 0xb9, 0x84, 0xdd,
	0x5d, 0x43, 0xe2, 0x83, 0x3d, 0xfc, 0x03, 0xe2, 0x7d, 0x3c, 0x83, 0x43, 0xe2, 0x5d, 0x83, 0x82,
	0xfc, 0x83, 0x04, 0x0e, 0xdd, 0x64, 0x76, 0x52, 0x8f, 0xe2, 0x4f, 0x2a, 0x2d, 0xce, 0xcc, 0x4b,
	0x2d, 0x2e, 0x86, 0x1a, 0x7a, 0x8a, 0x49, 0x38, 0xb4, 0x24, 0xdf, 0xb5, 0x22, 0x39, 0xb5, 0xa0,
	0x24, 0x33, 0x3f, 0xaf, 0x38, 0x06, 0x6c, 0x64, 0x12, 0x1b, 0xd8, 0x2e, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0x8c, 0x49, 0x7a, 0x81, 0x00, 0x00, 0x00,
}
