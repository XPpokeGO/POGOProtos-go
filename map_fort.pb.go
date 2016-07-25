// Code generated by protoc-gen-go.
// source: map_fort.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type FortRenderingType int32

const (
	FortRenderingType_DEFAULT       FortRenderingType = 0
	FortRenderingType_INTERNAL_TEST FortRenderingType = 1
)

var FortRenderingType_name = map[int32]string{
	0: "DEFAULT",
	1: "INTERNAL_TEST",
}
var FortRenderingType_value = map[string]int32{
	"DEFAULT":       0,
	"INTERNAL_TEST": 1,
}

func (x FortRenderingType) String() string {
	return proto.EnumName(FortRenderingType_name, int32(x))
}
func (FortRenderingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type FortSponsor int32

const (
	FortSponsor_UNSET_SPONSOR FortSponsor = 0
	FortSponsor_MCDONALDS     FortSponsor = 1
	FortSponsor_POKEMON_STORE FortSponsor = 2
)

var FortSponsor_name = map[int32]string{
	0: "UNSET_SPONSOR",
	1: "MCDONALDS",
	2: "POKEMON_STORE",
}
var FortSponsor_value = map[string]int32{
	"UNSET_SPONSOR": 0,
	"MCDONALDS":     1,
	"POKEMON_STORE": 2,
}

func (x FortSponsor) String() string {
	return proto.EnumName(FortSponsor_name, int32(x))
}
func (FortSponsor) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

type FortType int32

const (
	FortType_GYM        FortType = 0
	FortType_CHECKPOINT FortType = 1
)

var FortType_name = map[int32]string{
	0: "GYM",
	1: "CHECKPOINT",
}
var FortType_value = map[string]int32{
	"GYM":        0,
	"CHECKPOINT": 1,
}

func (x FortType) String() string {
	return proto.EnumName(FortType_name, int32(x))
}
func (FortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

type FortData struct {
	Id                      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastModifiedTimestampMs int64    `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64  `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64  `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	Enabled                 bool     `protobuf:"varint,8,opt,name=enabled" json:"enabled,omitempty"`
	Type                    FortType `protobuf:"varint,9,opt,name=type,enum=POGOProtos.Map.Fort.FortType" json:"type,omitempty"`
	// Team that owns the gym
	OwnedByTeam TeamColor `protobuf:"varint,5,opt,name=owned_by_team,json=ownedByTeam,enum=POGOProtos.Enums.TeamColor" json:"owned_by_team,omitempty"`
	// Highest CP Pokemon at the gym
	GuardPokemonId PokemonId `protobuf:"varint,6,opt,name=guard_pokemon_id,json=guardPokemonId,enum=POGOProtos.Enums.PokemonId" json:"guard_pokemon_id,omitempty"`
	GuardPokemonCp int32     `protobuf:"varint,7,opt,name=guard_pokemon_cp,json=guardPokemonCp" json:"guard_pokemon_cp,omitempty"`
	// Prestigate / experience of the gym
	GymPoints int64 `protobuf:"varint,10,opt,name=gym_points,json=gymPoints" json:"gym_points,omitempty"`
	// Whether someone is battling at the gym currently
	IsInBattle bool `protobuf:"varint,11,opt,name=is_in_battle,json=isInBattle" json:"is_in_battle,omitempty"`
	// Timestamp when the pokestop can be activated again to get items / xp
	CooldownCompleteTimestampMs int64             `protobuf:"varint,14,opt,name=cooldown_complete_timestamp_ms,json=cooldownCompleteTimestampMs" json:"cooldown_complete_timestamp_ms,omitempty"`
	Sponsor                     FortSponsor       `protobuf:"varint,15,opt,name=sponsor,enum=POGOProtos.Map.Fort.FortSponsor" json:"sponsor,omitempty"`
	RenderingType               FortRenderingType `protobuf:"varint,16,opt,name=rendering_type,json=renderingType,enum=POGOProtos.Map.Fort.FortRenderingType" json:"rendering_type,omitempty"`
	// Might represent the type of item applied to the pokestop, right now only lures can be applied
	ActiveFortModifier []byte        `protobuf:"bytes,12,opt,name=active_fort_modifier,json=activeFortModifier,proto3" json:"active_fort_modifier,omitempty"`
	LureInfo           *FortLureInfo `protobuf:"bytes,13,opt,name=lure_info,json=lureInfo" json:"lure_info,omitempty"`
}

func (m *FortData) Reset()                    { *m = FortData{} }
func (m *FortData) String() string            { return proto.CompactTextString(m) }
func (*FortData) ProtoMessage()               {}
func (*FortData) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *FortData) GetLureInfo() *FortLureInfo {
	if m != nil {
		return m.LureInfo
	}
	return nil
}

type FortLureInfo struct {
	FortId                 string    `protobuf:"bytes,1,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	EncounterId            uint64    `protobuf:"fixed64,2,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	ActivePokemonId        PokemonId `protobuf:"varint,3,opt,name=active_pokemon_id,json=activePokemonId,enum=POGOProtos.Enums.PokemonId" json:"active_pokemon_id,omitempty"`
	LureExpiresTimestampMs int64     `protobuf:"varint,4,opt,name=lure_expires_timestamp_ms,json=lureExpiresTimestampMs" json:"lure_expires_timestamp_ms,omitempty"`
}

func (m *FortLureInfo) Reset()                    { *m = FortLureInfo{} }
func (m *FortLureInfo) String() string            { return proto.CompactTextString(m) }
func (*FortLureInfo) ProtoMessage()               {}
func (*FortLureInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

type FortModifier struct {
	ItemId                 ItemId `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ExpirationTimestampMs  int64  `protobuf:"varint,2,opt,name=expiration_timestamp_ms,json=expirationTimestampMs" json:"expiration_timestamp_ms,omitempty"`
	DeployerPlayerCodename string `protobuf:"bytes,3,opt,name=deployer_player_codename,json=deployerPlayerCodename" json:"deployer_player_codename,omitempty"`
}

func (m *FortModifier) Reset()                    { *m = FortModifier{} }
func (m *FortModifier) String() string            { return proto.CompactTextString(m) }
func (*FortModifier) ProtoMessage()               {}
func (*FortModifier) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

type FortSummary struct {
	FortSummaryId           string  `protobuf:"bytes,1,opt,name=fort_summary_id,json=fortSummaryId" json:"fort_summary_id,omitempty"`
	LastModifiedTimestampMs int64   `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *FortSummary) Reset()                    { *m = FortSummary{} }
func (m *FortSummary) String() string            { return proto.CompactTextString(m) }
func (*FortSummary) ProtoMessage()               {}
func (*FortSummary) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func init() {
	proto.RegisterType((*FortData)(nil), "POGOProtos.Map.Fort.FortData")
	proto.RegisterType((*FortLureInfo)(nil), "POGOProtos.Map.Fort.FortLureInfo")
	proto.RegisterType((*FortModifier)(nil), "POGOProtos.Map.Fort.FortModifier")
	proto.RegisterType((*FortSummary)(nil), "POGOProtos.Map.Fort.FortSummary")
	proto.RegisterEnum("POGOProtos.Map.Fort.FortRenderingType", FortRenderingType_name, FortRenderingType_value)
	proto.RegisterEnum("POGOProtos.Map.Fort.FortSponsor", FortSponsor_name, FortSponsor_value)
	proto.RegisterEnum("POGOProtos.Map.Fort.FortType", FortType_name, FortType_value)
}

func init() { proto.RegisterFile("map_fort.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x54, 0xdb, 0x4e, 0xf3, 0x46,
	0x10, 0xc6, 0x04, 0x72, 0x98, 0x24, 0x26, 0x6c, 0x29, 0xb8, 0xa1, 0x54, 0x90, 0x4a, 0x08, 0x71,
	0x11, 0xb5, 0x20, 0x55, 0x3d, 0x48, 0xad, 0x88, 0x63, 0x68, 0x04, 0x89, 0x2d, 0xc7, 0x5c, 0xb4,
	0x37, 0x96, 0x13, 0x6f, 0x22, 0xab, 0xb6, 0xd7, 0xb2, 0x37, 0xb4, 0x79, 0x8e, 0xbe, 0x46, 0x1f,
	0xa1, 0xcf, 0xd0, 0x67, 0xea, 0x78, 0x6d, 0xe7, 0xd0, 0x12, 0xfd, 0x97, 0xff, 0x8d, 0xbd, 0x3b,
	0xdf, 0x37, 0xb3, 0x3b, 0x33, 0xdf, 0x0e, 0xc8, 0x81, 0x13, 0xd9, 0x33, 0x16, 0xf3, 0x6e, 0x14,
	0x33, 0xce, 0xc8, 0x27, 0x86, 0xfe, 0xa4, 0x1b, 0xe9, 0x32, 0xe9, 0x0e, 0x9d, 0xa8, 0xfb, 0x88,
	0x50, 0xbb, 0x4e, 0xc3, 0x45, 0x90, 0x64, 0x8c, 0xf6, 0x89, 0x17, 0xbe, 0xd1, 0x90, 0xb3, 0x78,
	0x69, 0x7b, 0x9c, 0x06, 0x99, 0xb5, 0xf3, 0x67, 0x19, 0xaa, 0x29, 0xb7, 0xef, 0x70, 0x87, 0xc8,
	0xb0, 0xef, 0xb9, 0x8a, 0x74, 0x29, 0xdd, 0xd4, 0x4c, 0x5c, 0x91, 0x1f, 0xa0, 0xed, 0x3b, 0x09,
	0xb7, 0x03, 0xe6, 0x7a, 0x33, 0x8f, 0xba, 0x36, 0xf7, 0x02, 0x9a, 0x70, 0x27, 0x88, 0xec, 0x20,
	0x51, 0xf6, 0x91, 0x57, 0x32, 0xcf, 0x52, 0xc6, 0x30, 0x27, 0x58, 0x05, 0x3e, 0x4c, 0x48, 0x1b,
	0xaa, 0xbe, 0xc3, 0x3d, 0xbe, 0x70, 0xa9, 0x52, 0x42, 0xaa, 0x64, 0xae, 0xf6, 0xe4, 0x73, 0xa8,
	0xf9, 0x2c, 0x9c, 0x67, 0xe0, 0x81, 0x00, 0xd7, 0x06, 0xa2, 0x40, 0x85, 0x86, 0xce, 0xc4, 0xa7,
	0xae, 0x52, 0x45, 0xac, 0x6a, 0x16, 0x5b, 0xf2, 0x35, 0x1c, 0xf0, 0x65, 0x44, 0x95, 0x1a, 0x9a,
	0xe5, 0xbb, 0x8b, 0xee, 0x3b, 0x49, 0x8b, 0x8f, 0x85, 0x24, 0x53, 0x50, 0xc9, 0x4f, 0xd0, 0x64,
	0xbf, 0x87, 0x78, 0xf7, 0xc9, 0xd2, 0xe6, 0xd4, 0x09, 0x94, 0x43, 0xe1, 0x7b, 0xbe, 0xe9, 0xab,
	0x89, 0x32, 0x59, 0x88, 0xaa, 0xcc, 0x67, 0xb1, 0x59, 0x17, 0x1e, 0xbd, 0x65, 0x6a, 0x21, 0x1a,
	0xb4, 0xe6, 0x0b, 0x27, 0x76, 0xed, 0x88, 0xfd, 0x46, 0x03, 0x16, 0xda, 0x58, 0xa2, 0xf2, 0xae,
	0x18, 0x46, 0xc6, 0x19, 0xb8, 0xa6, 0x2c, 0x9c, 0x56, 0x7b, 0x72, 0xf3, 0xdf, 0x30, 0xd3, 0x48,
	0xa9, 0x60, 0x98, 0xc3, 0x6d, 0xa6, 0x1a, 0x91, 0x0b, 0x80, 0xf9, 0x32, 0x40, 0x9e, 0x17, 0xf2,
	0x44, 0x01, 0x51, 0xe5, 0x1a, 0x5a, 0x0c, 0x61, 0x20, 0x97, 0xd0, 0xf0, 0x12, 0xdb, 0x0b, 0xed,
	0x89, 0xc3, 0xb9, 0x4f, 0x95, 0xba, 0x28, 0x11, 0x78, 0xc9, 0x20, 0xec, 0x09, 0x0b, 0x51, 0xe1,
	0x8b, 0x29, 0x63, 0xbe, 0x8b, 0x59, 0xd8, 0x53, 0x16, 0x44, 0x3e, 0xe5, 0x74, 0xbb, 0x75, 0xb2,
	0x08, 0x7a, 0x5e, 0xb0, 0xd4, 0x9c, 0xb4, 0xd9, 0xbe, 0xef, 0xa1, 0x92, 0x44, 0x2c, 0x4c, 0x58,
	0xac, 0x1c, 0x89, 0x6c, 0x2f, 0x77, 0x56, 0x7b, 0x9c, 0xf1, 0xcc, 0xc2, 0x81, 0x0c, 0x41, 0x8e,
	0x69, 0xe8, 0xd2, 0xd8, 0x0b, 0xe7, 0xb6, 0x68, 0x58, 0x4b, 0x84, 0xb8, 0xde, 0x19, 0xc2, 0x2c,
	0xe8, 0xa2, 0x73, 0xcd, 0x78, 0x73, 0x4b, 0xbe, 0x82, 0x13, 0x67, 0xca, 0xbd, 0x37, 0x2a, 0x04,
	0x5f, 0xa8, 0x31, 0x56, 0x1a, 0x18, 0xb4, 0x61, 0x92, 0x0c, 0x4b, 0xa3, 0xe4, 0x32, 0x8c, 0xc9,
	0x8f, 0xa8, 0xaf, 0x45, 0x4c, 0xb1, 0x4a, 0x33, 0xa6, 0x34, 0x91, 0x56, 0xbf, 0xbb, 0xda, 0x79,
	0xf6, 0x0b, 0x32, 0x07, 0x48, 0x44, 0x7d, 0xe6, 0xab, 0xce, 0x3f, 0x12, 0x34, 0x36, 0x21, 0x72,
	0x06, 0x15, 0x71, 0xf6, 0xea, 0x79, 0x94, 0xd3, 0x2d, 0xb6, 0xf5, 0x0a, 0x1a, 0x34, 0x9c, 0xb2,
	0x45, 0xc8, 0x69, 0x9c, 0xa2, 0xe9, 0xa3, 0x28, 0x9b, 0xf5, 0x95, 0x0d, 0x29, 0x4f, 0x70, 0x9c,
	0x5f, 0x7f, 0x43, 0x41, 0xa5, 0x0f, 0x2b, 0xe8, 0x28, 0xf3, 0x5a, 0x4b, 0xe8, 0x3b, 0xf8, 0x4c,
	0x64, 0x45, 0xff, 0x88, 0xbc, 0x98, 0x26, 0xdb, 0x2d, 0x3d, 0x10, 0x2d, 0x3d, 0x4d, 0x09, 0x5a,
	0x86, 0x6f, 0x74, 0xb3, 0xf3, 0x77, 0x9e, 0xd0, 0xaa, 0x42, 0xd8, 0xde, 0x74, 0x0a, 0x14, 0x09,
	0xc9, 0xdb, 0xf5, 0x19, 0x14, 0xa3, 0xa2, 0x3b, 0x48, 0x47, 0x45, 0xfa, 0xc1, 0x0b, 0x95, 0x3d,
	0xf1, 0x27, 0xdf, 0xc0, 0x99, 0xb8, 0x02, 0xbe, 0x66, 0x4c, 0xe6, 0x9d, 0x99, 0xf0, 0xe9, 0x1a,
	0xde, 0x94, 0xd4, 0xb7, 0xa0, 0xb8, 0x34, 0xf2, 0xd9, 0x12, 0x4b, 0x15, 0xf9, 0x4e, 0xfa, 0x9b,
	0x32, 0x17, 0xdf, 0x76, 0x90, 0x4d, 0x88, 0x9a, 0x79, 0x5a, 0xe0, 0x86, 0x80, 0xd5, 0x1c, 0xed,
	0xfc, 0x25, 0x41, 0x5d, 0x28, 0x6d, 0x11, 0x04, 0x4e, 0xbc, 0x24, 0xd7, 0x70, 0x24, 0xda, 0x91,
	0x64, 0xfb, 0x75, 0x5b, 0x9a, 0xb3, 0x35, 0x6b, 0xf0, 0xb1, 0x06, 0xd8, 0xed, 0x3d, 0x1c, 0xff,
	0x4f, 0xd4, 0xa4, 0x0e, 0x95, 0xbe, 0xf6, 0xf8, 0xf0, 0xfa, 0x62, 0xb5, 0xf6, 0xc8, 0x31, 0x34,
	0x07, 0x23, 0x4b, 0x33, 0x47, 0x0f, 0x2f, 0xb6, 0xa5, 0x8d, 0xad, 0x96, 0x74, 0xdb, 0xcb, 0x53,
	0xcc, 0xdf, 0x10, 0x32, 0x5e, 0x47, 0x63, 0xcd, 0xb2, 0xc7, 0x86, 0x3e, 0x1a, 0xeb, 0x26, 0x3a,
	0x35, 0xa1, 0x36, 0x54, 0xfb, 0x3a, 0xfa, 0xf4, 0xc7, 0x2d, 0x29, 0x65, 0x18, 0xfa, 0xb3, 0x36,
	0xd4, 0x47, 0xf6, 0xd8, 0xd2, 0x4d, 0xad, 0xb5, 0x7f, 0xfb, 0x65, 0x36, 0xcc, 0xc5, 0x79, 0x15,
	0x28, 0x3d, 0xfd, 0x32, 0x44, 0x37, 0x19, 0x40, 0xfd, 0x59, 0x53, 0x9f, 0x0d, 0x1d, 0x8f, 0x6c,
	0x49, 0xbd, 0xea, 0xaf, 0x65, 0x31, 0xfb, 0x13, 0x63, 0xcf, 0x90, 0x26, 0xd9, 0xfa, 0xfe, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x9b, 0xea, 0x98, 0x51, 0x06, 0x00, 0x00,
}