// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/bidding_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible bidding errors.
type BiddingErrorEnum_BiddingError int32

const (
	// Enum unspecified.
	BiddingErrorEnum_UNSPECIFIED BiddingErrorEnum_BiddingError = 0
	// The received error code is not known in this version.
	BiddingErrorEnum_UNKNOWN BiddingErrorEnum_BiddingError = 1
	// Cannot transition to new bidding strategy.
	BiddingErrorEnum_BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED BiddingErrorEnum_BiddingError = 2
	// Cannot attach bidding strategy to campaign.
	BiddingErrorEnum_CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN BiddingErrorEnum_BiddingError = 7
	// Bidding strategy is not supported or cannot be used as anonymous.
	BiddingErrorEnum_INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 10
	// The type does not match the named strategy's type.
	BiddingErrorEnum_INVALID_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 14
	// The bid is invalid.
	BiddingErrorEnum_INVALID_BID BiddingErrorEnum_BiddingError = 17
	// Bidding strategy is not available for the account type.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE BiddingErrorEnum_BiddingError = 18
	// Conversion tracking is not enabled for the campaign for VBB transition.
	BiddingErrorEnum_CONVERSION_TRACKING_NOT_ENABLED BiddingErrorEnum_BiddingError = 19
	// Not enough conversions tracked for VBB transitions.
	BiddingErrorEnum_NOT_ENOUGH_CONVERSIONS BiddingErrorEnum_BiddingError = 20
	// Campaign can not be created with given bidding strategy. It can be
	// transitioned to the strategy, once eligible.
	BiddingErrorEnum_CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 21
	// Cannot target content network only as campaign uses Page One Promoted
	// bidding strategy.
	BiddingErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 23
	// Budget Optimizer and Target Spend bidding strategies are not supported
	// for campaigns with AdSchedule targeting.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE BiddingErrorEnum_BiddingError = 24
	// Pay per conversion is not available to all the customer, only few
	// whitelisted customers can use this.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER BiddingErrorEnum_BiddingError = 25
	// Pay per conversion is not allowed with Target CPA.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA BiddingErrorEnum_BiddingError = 26
	// Cannot set bidding strategy to Manual CPM for search network only
	// campaigns.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS BiddingErrorEnum_BiddingError = 27
	// The bidding strategy is not supported for use in drafts or experiments.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS BiddingErrorEnum_BiddingError = 28
	// Bidding strategy type does not support product type ad group criterion.
	BiddingErrorEnum_BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION BiddingErrorEnum_BiddingError = 29
	// Bid amount is too small.
	BiddingErrorEnum_BID_TOO_SMALL BiddingErrorEnum_BiddingError = 30
	// Bid amount is too big.
	BiddingErrorEnum_BID_TOO_BIG BiddingErrorEnum_BiddingError = 31
	// Bid has too many fractional digit precision.
	BiddingErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS BiddingErrorEnum_BiddingError = 32
	// Invalid domain name specified.
	BiddingErrorEnum_INVALID_DOMAIN_NAME BiddingErrorEnum_BiddingError = 33
	// The field is not compatible with payment mode.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_PAYMENT_MODE BiddingErrorEnum_BiddingError = 34
)

var BiddingErrorEnum_BiddingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED",
	7:  "CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN",
	10: "INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE",
	14: "INVALID_BIDDING_STRATEGY_TYPE",
	17: "INVALID_BID",
	18: "BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
	19: "CONVERSION_TRACKING_NOT_ENABLED",
	20: "NOT_ENOUGH_CONVERSIONS",
	21: "CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY",
	23: "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY",
	24: "BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE",
	25: "PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER",
	26: "PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA",
	27: "BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS",
	28: "BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS",
	29: "BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION",
	30: "BID_TOO_SMALL",
	31: "BID_TOO_BIG",
	32: "BID_TOO_MANY_FRACTIONAL_DIGITS",
	33: "INVALID_DOMAIN_NAME",
	34: "NOT_COMPATIBLE_WITH_PAYMENT_MODE",
}
var BiddingErrorEnum_BiddingError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED":                                     2,
	"CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN":                                  7,
	"INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE":                                     10,
	"INVALID_BIDDING_STRATEGY_TYPE":                                               14,
	"INVALID_BID":                                                                 17,
	"BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                             18,
	"CONVERSION_TRACKING_NOT_ENABLED":                                             19,
	"NOT_ENOUGH_CONVERSIONS":                                                      20,
	"CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY":                                21,
	"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY": 23,
	"BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE":                             24,
	"PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER":                               25,
	"PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA":                              26,
	"BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS":                      27,
	"BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS":                     28,
	"BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION":       29,
	"BID_TOO_SMALL":                    30,
	"BID_TOO_BIG":                      31,
	"BID_TOO_MANY_FRACTIONAL_DIGITS":   32,
	"INVALID_DOMAIN_NAME":              33,
	"NOT_COMPATIBLE_WITH_PAYMENT_MODE": 34,
}

func (x BiddingErrorEnum_BiddingError) String() string {
	return proto.EnumName(BiddingErrorEnum_BiddingError_name, int32(x))
}
func (BiddingErrorEnum_BiddingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bidding_error_bb805a21cd606d55, []int{0, 0}
}

// Container for enum describing possible bidding errors.
type BiddingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingErrorEnum) Reset()         { *m = BiddingErrorEnum{} }
func (m *BiddingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingErrorEnum) ProtoMessage()    {}
func (*BiddingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_error_bb805a21cd606d55, []int{0}
}
func (m *BiddingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingErrorEnum.Unmarshal(m, b)
}
func (m *BiddingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingErrorEnum.Marshal(b, m, deterministic)
}
func (dst *BiddingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingErrorEnum.Merge(dst, src)
}
func (m *BiddingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingErrorEnum.Size(m)
}
func (m *BiddingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BiddingErrorEnum)(nil), "google.ads.googleads.v1.errors.BiddingErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.BiddingErrorEnum_BiddingError", BiddingErrorEnum_BiddingError_name, BiddingErrorEnum_BiddingError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/bidding_error.proto", fileDescriptor_bidding_error_bb805a21cd606d55)
}

var fileDescriptor_bidding_error_bb805a21cd606d55 = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xed, 0x8e, 0x1b, 0x35,
	0x14, 0x65, 0x83, 0x44, 0x24, 0x97, 0x0f, 0xaf, 0x0b, 0x14, 0x96, 0x36, 0x6d, 0x03, 0x12, 0x12,
	0x1f, 0x13, 0xd2, 0x4a, 0x20, 0x4d, 0x7f, 0xdd, 0x8c, 0x9d, 0x89, 0x95, 0x19, 0xdb, 0xb2, 0x3d,
	0x09, 0x41, 0x91, 0xac, 0x94, 0xac, 0xa2, 0x48, 0xbb, 0x99, 0x55, 0x66, 0xd9, 0x07, 0xe2, 0x27,
	0x8f, 0xb2, 0x2f, 0xc1, 0x7f, 0xc4, 0x43, 0x20, 0x8f, 0x33, 0x21, 0x52, 0xb2, 0xf0, 0x2b, 0x37,
	0x57, 0xe7, 0x1c, 0xfb, 0xdc, 0x39, 0xbe, 0xe8, 0xd5, 0xaa, 0x2c, 0x57, 0x57, 0x97, 0xbd, 0xc5,
	0xb2, 0xea, 0x85, 0xd2, 0x57, 0x77, 0xfd, 0xde, 0xe5, 0x76, 0x5b, 0x6e, 0xab, 0xde, 0xdb, 0xf5,
	0x72, 0xb9, 0xde, 0xac, 0x5c, 0xfd, 0x37, 0xba, 0xd9, 0x96, 0xb7, 0x25, 0xe9, 0x04, 0x60, 0xb4,
	0x58, 0x56, 0xd1, 0x9e, 0x13, 0xdd, 0xf5, 0xa3, 0xc0, 0xb9, 0x78, 0xda, 0x68, 0xde, 0xac, 0x7b,
	0x8b, 0xcd, 0xa6, 0xbc, 0x5d, 0xdc, 0xae, 0xcb, 0x4d, 0x15, 0xd8, 0xdd, 0x3f, 0xdb, 0x08, 0x0f,
	0x82, 0x2a, 0xf3, 0x78, 0xb6, 0xf9, 0xed, 0xba, 0x7b, 0xdf, 0x46, 0xef, 0x1f, 0x36, 0xc9, 0x47,
	0xe8, 0x51, 0x21, 0x8c, 0x62, 0x09, 0x1f, 0x72, 0x46, 0xf1, 0x3b, 0xe4, 0x11, 0x6a, 0x17, 0x62,
	0x2c, 0xe4, 0x54, 0xe0, 0x33, 0xf2, 0x2d, 0xfa, 0x7a, 0xc0, 0x29, 0xe5, 0x22, 0x75, 0xc6, 0x6a,
	0xb0, 0x2c, 0x9d, 0x39, 0xab, 0x41, 0x18, 0x6e, 0xb9, 0x14, 0x4e, 0x48, 0xeb, 0x20, 0xcb, 0xe4,
	0x94, 0x51, 0xdc, 0x22, 0x11, 0xfa, 0x26, 0x01, 0x51, 0xf7, 0xac, 0x85, 0x64, 0xe4, 0x8e, 0xa9,
	0xd2, 0x25, 0x90, 0x2b, 0xe0, 0xa9, 0xc0, 0x6d, 0x2f, 0xce, 0xc5, 0x04, 0x32, 0x4e, 0x1d, 0x08,
	0x29, 0x66, 0xb9, 0x2c, 0xcc, 0x09, 0xce, 0x4c, 0x31, 0x8c, 0xc8, 0x4b, 0xf4, 0xac, 0x01, 0x9f,
	0x86, 0x7c, 0xe8, 0xad, 0x1c, 0x40, 0xf0, 0x39, 0x79, 0x8d, 0x7a, 0x47, 0xd8, 0xfa, 0x7a, 0x13,
	0xe0, 0x19, 0x0c, 0x32, 0xe6, 0x86, 0x52, 0x3b, 0x48, 0x12, 0x59, 0x08, 0x1b, 0x54, 0x08, 0xf9,
	0x12, 0x3d, 0x4f, 0xa4, 0x98, 0x30, 0x6d, 0xbc, 0x43, 0xab, 0x21, 0x19, 0x7b, 0x01, 0xcf, 0x63,
	0xc2, 0x93, 0x28, 0x7e, 0x4c, 0x2e, 0xd0, 0xa7, 0xa1, 0x21, 0x8b, 0x74, 0xe4, 0xfe, 0xc5, 0x1b,
	0xfc, 0x31, 0xf9, 0x01, 0x7d, 0xb7, 0x1b, 0x43, 0xa2, 0x19, 0x58, 0xb6, 0xb7, 0xec, 0xa6, 0xdc,
	0x1e, 0x0f, 0x05, 0x7f, 0x42, 0x24, 0x1a, 0xef, 0x18, 0x16, 0x74, 0xca, 0xac, 0x17, 0xb4, 0x4c,
	0x58, 0x27, 0x98, 0x9d, 0x4a, 0x3d, 0x76, 0x52, 0x64, 0xb3, 0xc0, 0xde, 0x6b, 0x65, 0x6c, 0xc2,
	0x32, 0xa7, 0xa4, 0x3a, 0x16, 0x7c, 0xf2, 0xa0, 0x71, 0x53, 0x28, 0x25, 0xb5, 0x65, 0x34, 0x88,
	0x01, 0x75, 0x26, 0x19, 0x31, 0x5a, 0x64, 0x0c, 0x7f, 0x46, 0xfa, 0xe8, 0x7b, 0x05, 0x33, 0xa7,
	0x98, 0x3e, 0x30, 0x74, 0x62, 0x5e, 0x49, 0x61, 0xac, 0xcc, 0x99, 0xc6, 0x9f, 0x93, 0x57, 0x28,
	0x7a, 0x88, 0x12, 0x52, 0x11, 0xce, 0x69, 0x5c, 0x29, 0xc0, 0x17, 0x24, 0x46, 0x3f, 0x9e, 0xfe,
	0x28, 0x3b, 0x86, 0x3f, 0xc2, 0x30, 0xd0, 0xc9, 0x28, 0xd8, 0x6e, 0x1c, 0x1b, 0xfc, 0x05, 0x79,
	0x83, 0x7e, 0xfa, 0x1f, 0x5f, 0x5c, 0x38, 0xaa, 0x61, 0x68, 0x8d, 0x93, 0xda, 0xb1, 0x9f, 0x15,
	0xd3, 0x3c, 0x67, 0xc2, 0x1a, 0xfc, 0x94, 0x70, 0xc4, 0x4e, 0x26, 0xc7, 0x51, 0xc9, 0xcc, 0xa1,
	0x8c, 0x53, 0x5a, 0xd2, 0x22, 0x09, 0x89, 0x70, 0x40, 0x53, 0x2d, 0x0b, 0xe5, 0x12, 0xcd, 0x2d,
	0xd3, 0x5c, 0x0a, 0xfc, 0x8c, 0x9c, 0xa3, 0x0f, 0x06, 0x9c, 0x3a, 0x2b, 0xa5, 0x33, 0x39, 0x64,
	0x19, 0xee, 0xf8, 0xf0, 0x35, 0xad, 0x01, 0x4f, 0xf1, 0x73, 0xd2, 0x45, 0x9d, 0xa6, 0x91, 0x83,
	0x98, 0xb9, 0xa1, 0x86, 0xc4, 0x3f, 0x1a, 0xc8, 0x1c, 0xe5, 0x29, 0xb7, 0x06, 0xbf, 0x20, 0x4f,
	0xd0, 0xe3, 0x26, 0xb1, 0x54, 0xe6, 0xc0, 0x85, 0x13, 0x90, 0x33, 0xfc, 0x92, 0x7c, 0x85, 0x5e,
	0xd4, 0x01, 0x92, 0xb9, 0x02, 0xcb, 0xfd, 0xe4, 0xeb, 0x41, 0x2a, 0x98, 0x79, 0x37, 0x2e, 0x97,
	0x94, 0xe1, 0xee, 0xe0, 0xef, 0x33, 0xd4, 0xfd, 0xb5, 0xbc, 0x8e, 0xfe, 0x7b, 0x4d, 0x0c, 0xce,
	0x0f, 0x1f, 0xbc, 0xf2, 0xbb, 0x41, 0x9d, 0xfd, 0x42, 0x77, 0xa4, 0x55, 0x79, 0xb5, 0xd8, 0xac,
	0xa2, 0x72, 0xbb, 0xea, 0xad, 0x2e, 0x37, 0xf5, 0xe6, 0x68, 0xf6, 0xd3, 0xcd, 0xba, 0x7a, 0x68,
	0x5d, 0xbd, 0x09, 0x3f, 0xbf, 0xb7, 0xde, 0x4d, 0x01, 0xfe, 0x68, 0x75, 0xd2, 0x20, 0x06, 0xcb,
	0x2a, 0x0a, 0xa5, 0xaf, 0x26, 0xfd, 0xa8, 0x3e, 0xb2, 0xba, 0x6f, 0x00, 0x73, 0x58, 0x56, 0xf3,
	0x3d, 0x60, 0x3e, 0xe9, 0xcf, 0x03, 0xe0, 0xaf, 0x56, 0x37, 0x74, 0xe3, 0x18, 0x96, 0x55, 0x1c,
	0xef, 0x21, 0x71, 0x3c, 0xe9, 0xc7, 0x71, 0x00, 0xbd, 0x7d, 0xaf, 0xbe, 0xdd, 0xeb, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x01, 0x72, 0xb3, 0x4b, 0x05, 0x00, 0x00,
}