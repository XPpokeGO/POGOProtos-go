// Code generated by protoc-gen-go.
// source: networking_requests.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestType int32

const (
	RequestType_METHOD_UNSET                   RequestType = 0
	RequestType_PLAYER_UPDATE                  RequestType = 1
	RequestType_GET_PLAYER                     RequestType = 2
	RequestType_GET_INVENTORY                  RequestType = 4
	RequestType_DOWNLOAD_SETTINGS              RequestType = 5
	RequestType_DOWNLOAD_ITEM_TEMPLATES        RequestType = 6
	RequestType_DOWNLOAD_REMOTE_CONFIG_VERSION RequestType = 7
	RequestType_FORT_SEARCH                    RequestType = 101
	RequestType_ENCOUNTER                      RequestType = 102
	RequestType_CATCH_POKEMON                  RequestType = 103
	RequestType_FORT_DETAILS                   RequestType = 104
	RequestType_ITEM_USE                       RequestType = 105
	RequestType_GET_MAP_OBJECTS                RequestType = 106
	RequestType_FORT_DEPLOY_POKEMON            RequestType = 110
	RequestType_FORT_RECALL_POKEMON            RequestType = 111
	RequestType_RELEASE_POKEMON                RequestType = 112
	RequestType_USE_ITEM_POTION                RequestType = 113
	RequestType_USE_ITEM_CAPTURE               RequestType = 114
	RequestType_USE_ITEM_FLEE                  RequestType = 115
	RequestType_USE_ITEM_REVIVE                RequestType = 116
	RequestType_TRADE_SEARCH                   RequestType = 117
	RequestType_TRADE_OFFER                    RequestType = 118
	RequestType_TRADE_RESPONSE                 RequestType = 119
	RequestType_TRADE_RESULT                   RequestType = 120
	RequestType_GET_PLAYER_PROFILE             RequestType = 121
	RequestType_GET_ITEM_PACK                  RequestType = 122
	RequestType_BUY_ITEM_PACK                  RequestType = 123
	RequestType_BUY_GEM_PACK                   RequestType = 124
	RequestType_EVOLVE_POKEMON                 RequestType = 125
	RequestType_GET_HATCHED_EGGS               RequestType = 126
	RequestType_ENCOUNTER_TUTORIAL_COMPLETE    RequestType = 127
	RequestType_LEVEL_UP_REWARDS               RequestType = 128
	RequestType_CHECK_AWARDED_BADGES           RequestType = 129
	RequestType_USE_ITEM_GYM                   RequestType = 133
	RequestType_GET_GYM_DETAILS                RequestType = 134
	RequestType_START_GYM_BATTLE               RequestType = 135
	RequestType_ATTACK_GYM                     RequestType = 136
	RequestType_RECYCLE_INVENTORY_ITEM         RequestType = 137
	RequestType_COLLECT_DAILY_BONUS            RequestType = 138
	RequestType_USE_ITEM_XP_BOOST              RequestType = 139
	RequestType_USE_ITEM_EGG_INCUBATOR         RequestType = 140
	RequestType_USE_INCENSE                    RequestType = 141
	RequestType_GET_INCENSE_POKEMON            RequestType = 142
	RequestType_INCENSE_ENCOUNTER              RequestType = 143
	RequestType_ADD_FORT_MODIFIER              RequestType = 144
	RequestType_DISK_ENCOUNTER                 RequestType = 145
	RequestType_COLLECT_DAILY_DEFENDER_BONUS   RequestType = 146
	RequestType_UPGRADE_POKEMON                RequestType = 147
	RequestType_SET_FAVORITE_POKEMON           RequestType = 148
	RequestType_NICKNAME_POKEMON               RequestType = 149
	RequestType_EQUIP_BADGE                    RequestType = 150
	RequestType_SET_CONTACT_SETTINGS           RequestType = 151
	RequestType_SET_BUDDY_POKEMON              RequestType = 152
	RequestType_GET_BUDDY_WALKED               RequestType = 153
	RequestType_GET_ASSET_DIGEST               RequestType = 300
	RequestType_GET_DOWNLOAD_URLS              RequestType = 301
	RequestType_GET_SUGGESTED_CODENAMES        RequestType = 401
	RequestType_CHECK_CODENAME_AVAILABLE       RequestType = 402
	RequestType_CLAIM_CODENAME                 RequestType = 403
	RequestType_SET_AVATAR                     RequestType = 404
	RequestType_SET_PLAYER_TEAM                RequestType = 405
	RequestType_MARK_TUTORIAL_COMPLETE         RequestType = 406
	RequestType_LOAD_SPAWN_POINTS              RequestType = 500
	RequestType_CHECK_CHALLENGE                RequestType = 600
	RequestType_VERIFY_CHALLENGE               RequestType = 601
	RequestType_ECHO                           RequestType = 666
	RequestType_DEBUG_UPDATE_INVENTORY         RequestType = 700
	RequestType_DEBUG_DELETE_PLAYER            RequestType = 701
	RequestType_SFIDA_REGISTRATION             RequestType = 800
	RequestType_SFIDA_ACTION_LOG               RequestType = 801
	RequestType_SFIDA_CERTIFICATION            RequestType = 802
	RequestType_SFIDA_UPDATE                   RequestType = 803
	RequestType_SFIDA_ACTION                   RequestType = 804
	RequestType_SFIDA_DOWSER                   RequestType = 805
	RequestType_SFIDA_CAPTURE                  RequestType = 806
)

var RequestType_name = map[int32]string{
	0:   "METHOD_UNSET",
	1:   "PLAYER_UPDATE",
	2:   "GET_PLAYER",
	4:   "GET_INVENTORY",
	5:   "DOWNLOAD_SETTINGS",
	6:   "DOWNLOAD_ITEM_TEMPLATES",
	7:   "DOWNLOAD_REMOTE_CONFIG_VERSION",
	101: "FORT_SEARCH",
	102: "ENCOUNTER",
	103: "CATCH_POKEMON",
	104: "FORT_DETAILS",
	105: "ITEM_USE",
	106: "GET_MAP_OBJECTS",
	110: "FORT_DEPLOY_POKEMON",
	111: "FORT_RECALL_POKEMON",
	112: "RELEASE_POKEMON",
	113: "USE_ITEM_POTION",
	114: "USE_ITEM_CAPTURE",
	115: "USE_ITEM_FLEE",
	116: "USE_ITEM_REVIVE",
	117: "TRADE_SEARCH",
	118: "TRADE_OFFER",
	119: "TRADE_RESPONSE",
	120: "TRADE_RESULT",
	121: "GET_PLAYER_PROFILE",
	122: "GET_ITEM_PACK",
	123: "BUY_ITEM_PACK",
	124: "BUY_GEM_PACK",
	125: "EVOLVE_POKEMON",
	126: "GET_HATCHED_EGGS",
	127: "ENCOUNTER_TUTORIAL_COMPLETE",
	128: "LEVEL_UP_REWARDS",
	129: "CHECK_AWARDED_BADGES",
	133: "USE_ITEM_GYM",
	134: "GET_GYM_DETAILS",
	135: "START_GYM_BATTLE",
	136: "ATTACK_GYM",
	137: "RECYCLE_INVENTORY_ITEM",
	138: "COLLECT_DAILY_BONUS",
	139: "USE_ITEM_XP_BOOST",
	140: "USE_ITEM_EGG_INCUBATOR",
	141: "USE_INCENSE",
	142: "GET_INCENSE_POKEMON",
	143: "INCENSE_ENCOUNTER",
	144: "ADD_FORT_MODIFIER",
	145: "DISK_ENCOUNTER",
	146: "COLLECT_DAILY_DEFENDER_BONUS",
	147: "UPGRADE_POKEMON",
	148: "SET_FAVORITE_POKEMON",
	149: "NICKNAME_POKEMON",
	150: "EQUIP_BADGE",
	151: "SET_CONTACT_SETTINGS",
	152: "SET_BUDDY_POKEMON",
	153: "GET_BUDDY_WALKED",
	300: "GET_ASSET_DIGEST",
	301: "GET_DOWNLOAD_URLS",
	401: "GET_SUGGESTED_CODENAMES",
	402: "CHECK_CODENAME_AVAILABLE",
	403: "CLAIM_CODENAME",
	404: "SET_AVATAR",
	405: "SET_PLAYER_TEAM",
	406: "MARK_TUTORIAL_COMPLETE",
	500: "LOAD_SPAWN_POINTS",
	600: "CHECK_CHALLENGE",
	601: "VERIFY_CHALLENGE",
	666: "ECHO",
	700: "DEBUG_UPDATE_INVENTORY",
	701: "DEBUG_DELETE_PLAYER",
	800: "SFIDA_REGISTRATION",
	801: "SFIDA_ACTION_LOG",
	802: "SFIDA_CERTIFICATION",
	803: "SFIDA_UPDATE",
	804: "SFIDA_ACTION",
	805: "SFIDA_DOWSER",
	806: "SFIDA_CAPTURE",
}
var RequestType_value = map[string]int32{
	"METHOD_UNSET":                   0,
	"PLAYER_UPDATE":                  1,
	"GET_PLAYER":                     2,
	"GET_INVENTORY":                  4,
	"DOWNLOAD_SETTINGS":              5,
	"DOWNLOAD_ITEM_TEMPLATES":        6,
	"DOWNLOAD_REMOTE_CONFIG_VERSION": 7,
	"FORT_SEARCH":                    101,
	"ENCOUNTER":                      102,
	"CATCH_POKEMON":                  103,
	"FORT_DETAILS":                   104,
	"ITEM_USE":                       105,
	"GET_MAP_OBJECTS":                106,
	"FORT_DEPLOY_POKEMON":            110,
	"FORT_RECALL_POKEMON":            111,
	"RELEASE_POKEMON":                112,
	"USE_ITEM_POTION":                113,
	"USE_ITEM_CAPTURE":               114,
	"USE_ITEM_FLEE":                  115,
	"USE_ITEM_REVIVE":                116,
	"TRADE_SEARCH":                   117,
	"TRADE_OFFER":                    118,
	"TRADE_RESPONSE":                 119,
	"TRADE_RESULT":                   120,
	"GET_PLAYER_PROFILE":             121,
	"GET_ITEM_PACK":                  122,
	"BUY_ITEM_PACK":                  123,
	"BUY_GEM_PACK":                   124,
	"EVOLVE_POKEMON":                 125,
	"GET_HATCHED_EGGS":               126,
	"ENCOUNTER_TUTORIAL_COMPLETE":    127,
	"LEVEL_UP_REWARDS":               128,
	"CHECK_AWARDED_BADGES":           129,
	"USE_ITEM_GYM":                   133,
	"GET_GYM_DETAILS":                134,
	"START_GYM_BATTLE":               135,
	"ATTACK_GYM":                     136,
	"RECYCLE_INVENTORY_ITEM":         137,
	"COLLECT_DAILY_BONUS":            138,
	"USE_ITEM_XP_BOOST":              139,
	"USE_ITEM_EGG_INCUBATOR":         140,
	"USE_INCENSE":                    141,
	"GET_INCENSE_POKEMON":            142,
	"INCENSE_ENCOUNTER":              143,
	"ADD_FORT_MODIFIER":              144,
	"DISK_ENCOUNTER":                 145,
	"COLLECT_DAILY_DEFENDER_BONUS":   146,
	"UPGRADE_POKEMON":                147,
	"SET_FAVORITE_POKEMON":           148,
	"NICKNAME_POKEMON":               149,
	"EQUIP_BADGE":                    150,
	"SET_CONTACT_SETTINGS":           151,
	"SET_BUDDY_POKEMON":              152,
	"GET_BUDDY_WALKED":               153,
	"GET_ASSET_DIGEST":               300,
	"GET_DOWNLOAD_URLS":              301,
	"GET_SUGGESTED_CODENAMES":        401,
	"CHECK_CODENAME_AVAILABLE":       402,
	"CLAIM_CODENAME":                 403,
	"SET_AVATAR":                     404,
	"SET_PLAYER_TEAM":                405,
	"MARK_TUTORIAL_COMPLETE":         406,
	"LOAD_SPAWN_POINTS":              500,
	"CHECK_CHALLENGE":                600,
	"VERIFY_CHALLENGE":               601,
	"ECHO":                           666,
	"DEBUG_UPDATE_INVENTORY": 700,
	"DEBUG_DELETE_PLAYER":    701,
	"SFIDA_REGISTRATION":     800,
	"SFIDA_ACTION_LOG":       801,
	"SFIDA_CERTIFICATION":    802,
	"SFIDA_UPDATE":           803,
	"SFIDA_ACTION":           804,
	"SFIDA_DOWSER":           805,
	"SFIDA_CAPTURE":          806,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type Request struct {
	RequestType    RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,enum=POGOProtos.Networking.Requests.RequestType" json:"request_type,omitempty"`
	RequestMessage []byte      `protobuf:"bytes,2,opt,name=request_message,json=requestMessage,proto3" json:"request_message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func init() {
	proto.RegisterType((*Request)(nil), "POGOProtos.Networking.Requests.Request")
	proto.RegisterEnum("POGOProtos.Networking.Requests.RequestType", RequestType_name, RequestType_value)
}

func init() { proto.RegisterFile("networking_requests.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 1067 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x55, 0xeb, 0x6e, 0x1c, 0x35,
	0x14, 0x66, 0xb2, 0xcd, 0xb6, 0x75, 0x6e, 0x8e, 0x53, 0x92, 0x54, 0x2d, 0x05, 0xfa, 0x07, 0x04,
	0x52, 0x7e, 0xc0, 0x13, 0x78, 0xc7, 0x67, 0x66, 0x87, 0x9d, 0x19, 0x0f, 0xb6, 0x67, 0xc3, 0xf2,
	0xc7, 0x02, 0x69, 0x09, 0x05, 0x91, 0xa4, 0xd9, 0x94, 0x52, 0xee, 0x20, 0x71, 0xbf, 0xf6, 0xca,
	0x45, 0x42, 0xe2, 0xfe, 0x0b, 0xde, 0x00, 0xde, 0x03, 0xde, 0x83, 0x07, 0xe0, 0xd8, 0xb3, 0xeb,
	0xd9, 0x4a, 0x88, 0x1f, 0x51, 0xc6, 0xdf, 0xf1, 0xf9, 0xce, 0xed, 0xf3, 0x59, 0x72, 0x76, 0x7f,
	0x7c, 0x7c, 0xf5, 0xe0, 0xe8, 0xc5, 0x4b, 0xfb, 0x7b, 0xf6, 0x68, 0x7c, 0xf9, 0xca, 0x78, 0x72,
	0x3c, 0xd9, 0x39, 0x3c, 0x3a, 0x38, 0x3e, 0x60, 0x17, 0x2a, 0x99, 0xca, 0xca, 0x7d, 0x4e, 0x76,
	0xca, 0x70, 0x6b, 0x47, 0x4d, 0x6f, 0x5d, 0x7c, 0x37, 0x22, 0x27, 0xa7, 0x07, 0x56, 0x92, 0xe5,
	0xa9, 0xb7, 0x3d, 0xbe, 0x76, 0x38, 0xde, 0x8e, 0x1e, 0x88, 0x1e, 0x5e, 0x7d, 0xec, 0xd1, 0x9d,
	0xff, 0xa7, 0x98, 0x7d, 0x18, 0x74, 0x51, 0x4b, 0x47, 0xed, 0x81, 0x3d, 0x44, 0xd6, 0x66, 0x7c,
	0x2f, 0x8d, 0x27, 0x93, 0x67, 0xf6, 0xc6, 0xdb, 0x0b, 0x48, 0xb9, 0xac, 0x56, 0xa7, 0x70, 0xd1,
	0xa0, 0x8f, 0x7c, 0xb7, 0x42, 0x96, 0xe6, 0x58, 0x18, 0x25, 0xcb, 0x05, 0x98, 0xbe, 0x14, 0xb6,
	0x2e, 0x35, 0x18, 0x7a, 0x0f, 0x5b, 0x27, 0x2b, 0x55, 0xce, 0x47, 0xa0, 0x6c, 0x5d, 0x09, 0x6e,
	0x80, 0x46, 0x6c, 0x95, 0x90, 0x14, 0x8c, 0x6d, 0x60, 0xba, 0xe0, 0xae, 0xb8, 0x73, 0x56, 0x0e,
	0xa1, 0x34, 0x52, 0x8d, 0xe8, 0x09, 0x76, 0x2f, 0x59, 0x17, 0x72, 0xb7, 0xcc, 0x25, 0x17, 0x16,
	0x79, 0x4c, 0x56, 0xa6, 0x9a, 0x2e, 0xb2, 0x73, 0x64, 0x2b, 0xc0, 0x99, 0x81, 0xc2, 0xe2, 0x1f,
	0xd2, 0x18, 0xd0, 0xb4, 0xcb, 0x2e, 0x92, 0x0b, 0xc1, 0xa8, 0xa0, 0x90, 0x06, 0x6c, 0x2c, 0xcb,
	0x24, 0x4b, 0xed, 0x10, 0x94, 0xce, 0x64, 0x49, 0x4f, 0xb2, 0x35, 0xb2, 0x94, 0x48, 0x65, 0x90,
	0x93, 0xab, 0xb8, 0x4f, 0xc7, 0x6c, 0x85, 0x9c, 0x86, 0x32, 0x96, 0x75, 0x69, 0x30, 0x95, 0xe7,
	0x5c, 0x2a, 0x31, 0x37, 0x71, 0xdf, 0x56, 0x72, 0x80, 0x14, 0x25, 0xdd, 0x73, 0x25, 0x79, 0x17,
	0x01, 0x86, 0x67, 0xb9, 0xa6, 0xcf, 0xb3, 0x65, 0x72, 0xca, 0x07, 0xaf, 0x35, 0xd0, 0x4b, 0x6c,
	0x83, 0xac, 0xb9, 0xec, 0x0b, 0x5e, 0x59, 0xd9, 0x7b, 0x02, 0x62, 0xa3, 0xe9, 0x0b, 0x6c, 0x8b,
	0x6c, 0x4c, 0x9d, 0xaa, 0x5c, 0x8e, 0x02, 0xdb, 0x7e, 0x30, 0x28, 0x88, 0x79, 0x9e, 0x07, 0xc3,
	0x81, 0xa3, 0x51, 0x90, 0x03, 0xd7, 0x10, 0xc0, 0x43, 0x07, 0x62, 0x90, 0xa6, 0xd4, 0x4a, 0x1a,
	0x57, 0xc3, 0x65, 0x76, 0x86, 0xd0, 0x00, 0xc6, 0xbc, 0x32, 0xb5, 0x02, 0x7a, 0xe4, 0x32, 0x0f,
	0x68, 0x92, 0x03, 0xd0, 0xc9, 0x5d, 0xde, 0x0a, 0x86, 0xd9, 0x10, 0xe8, 0xb1, 0x2b, 0xc7, 0x28,
	0x2e, 0x60, 0xd6, 0x82, 0x2b, 0xae, 0x27, 0x0d, 0x22, 0x93, 0x04, 0x9b, 0xf0, 0x32, 0x63, 0x64,
	0xb5, 0x01, 0x14, 0xe8, 0x4a, 0xe2, 0x1c, 0xe9, 0xd5, 0xd6, 0x0d, 0xb1, 0x3a, 0x37, 0xf4, 0x15,
	0xb6, 0x49, 0x58, 0x3b, 0x45, 0x5b, 0x29, 0x99, 0x64, 0x39, 0xd0, 0x6b, 0x61, 0x9a, 0x3e, 0x67,
	0x1e, 0x0f, 0xe8, 0xab, 0x0e, 0xea, 0xd5, 0xa3, 0x39, 0xe8, 0x35, 0xc7, 0xe7, 0xa0, 0x74, 0x86,
	0xbc, 0xee, 0xa2, 0xc2, 0x50, 0xe6, 0xc3, 0xb6, 0xfe, 0x37, 0x5c, 0xa9, 0x8e, 0xab, 0xef, 0x46,
	0x02, 0xc2, 0x42, 0x8a, 0x2a, 0x78, 0x93, 0xdd, 0x4f, 0xce, 0x85, 0x99, 0x59, 0x53, 0xa3, 0x64,
	0x32, 0x9e, 0xe3, 0xac, 0x51, 0x0a, 0x80, 0x02, 0x7b, 0x0b, 0xd5, 0x43, 0x73, 0x18, 0x42, 0x8e,
	0x92, 0xc3, 0x7c, 0x77, 0xb9, 0x12, 0x9a, 0xbe, 0x1d, 0xb1, 0xb3, 0xe4, 0x0c, 0xd2, 0xc4, 0x03,
	0xcb, 0x1d, 0x84, 0x7c, 0x3d, 0x2e, 0x52, 0x94, 0xce, 0x3b, 0x11, 0x66, 0xb8, 0x1c, 0x5a, 0x95,
	0x8e, 0x0a, 0xfa, 0x5e, 0x84, 0xb1, 0xfd, 0x5c, 0xf1, 0x14, 0x46, 0xff, 0x7e, 0xe4, 0xa8, 0xb5,
	0xe1, 0xaa, 0xc1, 0x7b, 0xdc, 0x18, 0xac, 0xf9, 0x83, 0x08, 0x7b, 0x48, 0xf0, 0x1b, 0x0b, 0xf1,
	0xde, 0x1f, 0x46, 0xa8, 0xd4, 0x4d, 0x1c, 0xf1, 0x28, 0xce, 0xa1, 0xd5, 0xb5, 0xa7, 0xa7, 0x1f,
	0x45, 0x6c, 0x9b, 0x6c, 0xc4, 0x32, 0xcf, 0x51, 0x2b, 0x56, 0x20, 0xf1, 0xc8, 0xf6, 0x64, 0x59,
	0x6b, 0xfa, 0x71, 0x84, 0x4d, 0x5d, 0x0f, 0x79, 0x3c, 0x55, 0x21, 0x2e, 0xb5, 0xa1, 0x9f, 0x78,
	0xba, 0x80, 0x63, 0x17, 0x90, 0x33, 0xae, 0x31, 0xba, 0x54, 0xf4, 0xd3, 0x08, 0x7b, 0xb9, 0xe4,
	0x8d, 0x65, 0x0c, 0x6e, 0x58, 0x9f, 0xf9, 0x00, 0xcd, 0x8b, 0xf2, 0x48, 0x68, 0xe8, 0xe7, 0x3e,
	0xc0, 0x0c, 0x6d, 0x75, 0xff, 0x85, 0xc7, 0xb9, 0x10, 0xd6, 0x6b, 0xb3, 0x90, 0x22, 0x4b, 0x32,
	0xc4, 0xbf, 0x8c, 0x50, 0x43, 0xab, 0x22, 0xd3, 0x83, 0xb9, 0xcb, 0xd7, 0x23, 0xf6, 0x20, 0x39,
	0x7f, 0x77, 0xfe, 0x02, 0x12, 0x28, 0x05, 0x4e, 0xa3, 0x29, 0xe4, 0x86, 0xef, 0x5e, 0x5d, 0xa5,
	0x5e, 0x31, 0xb3, 0xe8, 0x37, 0xfd, 0x04, 0xf0, 0x35, 0xdb, 0x84, 0x0f, 0x71, 0x66, 0xa6, 0x35,
	0xdd, 0xf2, 0x8d, 0x2d, 0xb3, 0x78, 0x50, 0xf2, 0xa2, 0x85, 0x6f, 0xfb, 0xda, 0xe0, 0xc9, 0x3a,
	0xab, 0x9a, 0x59, 0xd1, 0x3b, 0x81, 0x03, 0x9f, 0x36, 0xf6, 0xdb, 0xb4, 0xdb, 0xe1, 0x2b, 0x5f,
	0x84, 0x33, 0xf5, 0x6a, 0x21, 0xda, 0x37, 0xf7, 0xb5, 0xe7, 0x4e, 0x03, 0xbe, 0xcb, 0xf3, 0x01,
	0x08, 0xfa, 0x4d, 0x80, 0xb9, 0x76, 0x4e, 0x22, 0x43, 0x29, 0x18, 0xfa, 0xdb, 0x82, 0x63, 0x71,
	0x70, 0xd8, 0x25, 0xb5, 0xc2, 0xd1, 0xff, 0xbe, 0xc0, 0xce, 0x93, 0x2d, 0x87, 0xeb, 0x3a, 0x75,
	0x37, 0x51, 0x3e, 0xb1, 0x14, 0xe0, 0xf2, 0xd5, 0xf4, 0x7a, 0x87, 0xdd, 0x47, 0xb6, 0x1b, 0x71,
	0xcd, 0x50, 0xcb, 0x87, 0xd8, 0x1b, 0xde, 0x43, 0x81, 0xdc, 0xe8, 0xb8, 0x3e, 0xc6, 0x39, 0xcf,
	0x8a, 0x60, 0xa6, 0x37, 0x3b, 0x4e, 0x35, 0x2e, 0x34, 0x5e, 0x44, 0x49, 0xd1, 0x5b, 0x1d, 0xd7,
	0x35, 0xdd, 0xbe, 0x29, 0x03, 0xbc, 0xa0, 0xb7, 0x3b, 0x6e, 0xf8, 0x05, 0x57, 0x83, 0xff, 0x90,
	0xfa, 0x9d, 0x8e, 0xcb, 0xb6, 0xd9, 0x92, 0x15, 0xdf, 0x2d, 0xb1, 0xe8, 0xac, 0xc4, 0x05, 0xf4,
	0x8f, 0xa7, 0x9a, 0xe6, 0xd3, 0xc7, 0x45, 0x03, 0x25, 0x36, 0xef, 0x2f, 0xb7, 0x57, 0x29, 0x2e,
	0xc3, 0x2c, 0x19, 0xcd, 0xc1, 0x7f, 0x9f, 0x60, 0xa7, 0xc9, 0x09, 0x88, 0xfb, 0x92, 0x7e, 0xeb,
	0x56, 0xec, 0xa6, 0x80, 0x5e, 0x9d, 0x4e, 0xd7, 0xf5, 0xdc, 0x56, 0xfe, 0x63, 0xd1, 0xe9, 0xaa,
	0x31, 0x0a, 0x70, 0xf1, 0x67, 0x2b, 0xfc, 0xcf, 0x45, 0xdc, 0x6b, 0x4c, 0x27, 0x99, 0xe0, 0xf8,
	0xde, 0xd2, 0x4c, 0xe3, 0xaa, 0xf0, 0xcb, 0xea, 0xfb, 0xae, 0x7f, 0x30, 0xde, 0x80, 0xc3, 0x42,
	0xc8, 0xe6, 0x32, 0xa5, 0x3f, 0x74, 0x1d, 0x53, 0x03, 0xc7, 0xa0, 0x0c, 0xaa, 0x2d, 0x6e, 0x1c,
	0x7e, 0xec, 0xba, 0xa7, 0xd8, 0x58, 0xa6, 0xbf, 0x17, 0x3f, 0xcd, 0x41, 0x0d, 0x07, 0xfd, 0x79,
	0x0e, 0xc2, 0x31, 0x69, 0x4c, 0xe1, 0x97, 0x2e, 0x2e, 0x90, 0x95, 0x29, 0xe5, 0x74, 0x29, 0xfe,
	0xda, 0xed, 0x9d, 0x7a, 0xba, 0xeb, 0x7f, 0x4d, 0x27, 0xcf, 0x36, 0xff, 0x1f, 0xff, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x04, 0xea, 0xd5, 0x72, 0x07, 0x00, 0x00,
}
