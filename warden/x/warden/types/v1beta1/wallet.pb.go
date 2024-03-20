// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/v1beta1/wallet.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// WalletType specifies the Layer 1 blockchain that this wallet will be used
// for.
type WalletType int32

const (
	// The wallet type is missing
	WalletType_WALLET_TYPE_UNSPECIFIED WalletType = 0
	// The wallet type for Ethereum
	WalletType_WALLET_TYPE_ETH WalletType = 1
	// The wallet type for Celestia
	WalletType_WALLET_TYPE_CELESTIA WalletType = 2
	// The wallet type for Sui
	WalletType_WALLET_TYPE_SUI WalletType = 3
)

var WalletType_name = map[int32]string{
	0: "WALLET_TYPE_UNSPECIFIED",
	1: "WALLET_TYPE_ETH",
	2: "WALLET_TYPE_CELESTIA",
	3: "WALLET_TYPE_SUI",
}

var WalletType_value = map[string]int32{
	"WALLET_TYPE_UNSPECIFIED": 0,
	"WALLET_TYPE_ETH":         1,
	"WALLET_TYPE_CELESTIA":    2,
	"WALLET_TYPE_SUI":         3,
}

func (x WalletType) String() string {
	return proto.EnumName(WalletType_name, int32(x))
}

func (WalletType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc128a4841264a1b, []int{0}
}

func init() {
	proto.RegisterEnum("warden.warden.v1beta1.WalletType", WalletType_name, WalletType_value)
}

func init() {
	proto.RegisterFile("warden/warden/v1beta1/wallet.proto", fileDescriptor_fc128a4841264a1b)
}

var fileDescriptor_fc128a4841264a1b = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4f, 0x2c, 0x4a,
	0x49, 0xcd, 0xd3, 0x87, 0x52, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa, 0xe5, 0x89, 0x39,
	0x39, 0xa9, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xa2, 0x10, 0x49, 0x3d, 0x28, 0x05,
	0x55, 0xa3, 0x95, 0xcb, 0xc5, 0x15, 0x0e, 0x56, 0x16, 0x52, 0x59, 0x90, 0x2a, 0x24, 0xcd, 0x25,
	0x1e, 0xee, 0xe8, 0xe3, 0xe3, 0x1a, 0x12, 0x1f, 0x12, 0x19, 0xe0, 0x1a, 0x1f, 0xea, 0x17, 0x1c,
	0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0x24, 0xcc, 0xc5, 0x8f, 0x2c, 0xe9,
	0x1a, 0xe2, 0x21, 0xc0, 0x28, 0x24, 0xc1, 0x25, 0x82, 0x2c, 0xe8, 0xec, 0xea, 0xe3, 0x1a, 0x1c,
	0xe2, 0xe9, 0x28, 0xc0, 0x84, 0xae, 0x3c, 0x38, 0xd4, 0x53, 0x80, 0xd9, 0x29, 0xf1, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xdc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xa1, 0xfe, 0xd0, 0x05, 0x3b, 0x3c, 0x39, 0x3f, 0x07, 0xca, 0x47, 0xe3, 0xea,
	0x57, 0xc0, 0x18, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x30, 0x5f, 0x27, 0xb1, 0x81, 0xd5, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x8a, 0x88, 0x9b, 0x15, 0x01, 0x00, 0x00,
}
