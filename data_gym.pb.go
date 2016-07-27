// Code generated by protoc-gen-go.
// source: data_gym.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

type GymMembership struct {
	PokemonData          *PokemonData         `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	TrainerPublicProfile *PlayerPublicProfile `protobuf:"bytes,2,opt,name=trainer_public_profile,json=trainerPublicProfile" json:"trainer_public_profile,omitempty"`
}

func (m *GymMembership) Reset()                    { *m = GymMembership{} }
func (m *GymMembership) String() string            { return proto.CompactTextString(m) }
func (*GymMembership) ProtoMessage()               {}
func (*GymMembership) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *GymMembership) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *GymMembership) GetTrainerPublicProfile() *PlayerPublicProfile {
	if m != nil {
		return m.TrainerPublicProfile
	}
	return nil
}

type GymState struct {
	FortData    *FortData        `protobuf:"bytes,1,opt,name=fort_data,json=fortData" json:"fort_data,omitempty"`
	Memberships []*GymMembership `protobuf:"bytes,2,rep,name=memberships" json:"memberships,omitempty"`
}

func (m *GymState) Reset()                    { *m = GymState{} }
func (m *GymState) String() string            { return proto.CompactTextString(m) }
func (*GymState) ProtoMessage()               {}
func (*GymState) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *GymState) GetFortData() *FortData {
	if m != nil {
		return m.FortData
	}
	return nil
}

func (m *GymState) GetMemberships() []*GymMembership {
	if m != nil {
		return m.Memberships
	}
	return nil
}

func init() {
	proto.RegisterType((*GymMembership)(nil), "POGOProtos.Data.Gym.GymMembership")
	proto.RegisterType((*GymState)(nil), "POGOProtos.Data.Gym.GymState")
}

func init() { proto.RegisterFile("data_gym.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x4d, 0x05, 0xa9, 0xa9, 0x16, 0x8c, 0x22, 0xa5, 0x54, 0x91, 0x3d, 0x09, 0x62, 0x0a,
	0x7a, 0xf3, 0x22, 0x48, 0xb1, 0xa7, 0xd2, 0xb0, 0xde, 0xbc, 0x84, 0x6c, 0x9b, 0xdd, 0x86, 0x6e,
	0x36, 0x21, 0x9b, 0x1e, 0xf6, 0x1d, 0x7c, 0x1b, 0x5f, 0xd0, 0xfc, 0x59, 0x74, 0xab, 0x5e, 0x76,
	0x67, 0x26, 0xf3, 0x9b, 0xf9, 0xe6, 0x83, 0xc3, 0x35, 0xb3, 0x8c, 0x16, 0x8d, 0xc4, 0xda, 0x28,
	0xab, 0xd0, 0x39, 0x59, 0xce, 0x97, 0xc4, 0x87, 0x35, 0x9e, 0xb9, 0x27, 0x3c, 0x6f, 0xe4, 0x18,
	0xfa, 0xa6, 0xd8, 0x30, 0x3e, 0x0b, 0x80, 0x2e, 0x59, 0xc3, 0x4d, 0x5b, 0x1a, 0x4a, 0xa6, 0x69,
	0xae, 0x8c, 0x8d, 0x79, 0xf2, 0x09, 0xe0, 0xa9, 0xc3, 0x16, 0x5c, 0x66, 0xdc, 0xd4, 0x1b, 0xa1,
	0xd1, 0x33, 0x3c, 0xd1, 0x6a, 0xcb, 0xa5, 0xaa, 0xa8, 0xc7, 0x47, 0xe0, 0x06, 0xdc, 0x0e, 0x1e,
	0x26, 0xf8, 0xf7, 0x32, 0x12, 0x9b, 0x7c, 0x9c, 0x0e, 0xf4, 0x4f, 0x82, 0x18, 0xbc, 0xb4, 0x86,
	0x89, 0x8a, 0x1b, 0xaa, 0x77, 0x59, 0x29, 0x56, 0xd4, 0xad, 0xca, 0x45, 0xc9, 0x47, 0xbd, 0x30,
	0xea, 0xee, 0xef, 0xa8, 0xa8, 0x30, 0xfe, 0x48, 0x60, 0x48, 0x44, 0xd2, 0x8b, 0x76, 0xd4, 0x5e,
	0x35, 0xf9, 0x00, 0xb0, 0xef, 0x54, 0xbf, 0x59, 0x66, 0x39, 0x7a, 0x82, 0xc7, 0xfe, 0xa0, 0xae,
	0xda, 0xab, 0xee, 0x8a, 0x05, 0xd3, 0xf8, 0xd5, 0x5f, 0xec, 0x3f, 0x41, 0x6e, 0x3f, 0x6f, 0x23,
	0x34, 0x83, 0x03, 0xf9, 0x7d, 0x7a, 0xed, 0x04, 0x1e, 0x3a, 0x3a, 0xc1, 0xff, 0x18, 0x8b, 0xf7,
	0x5c, 0x4a, 0xbb, 0xd8, 0xcb, 0xf5, 0xfb, 0xa4, 0x10, 0x76, 0xb3, 0xcb, 0xf0, 0x4a, 0xc9, 0xa9,
	0xde, 0xca, 0xaa, 0x50, 0xf7, 0x6a, 0x2d, 0xa6, 0xc1, 0xe3, 0x9a, 0x1c, 0x10, 0x40, 0x7a, 0xd9,
	0x51, 0xc8, 0x1e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0x06, 0x5d, 0x06, 0xc7, 0x01, 0x00,
	0x00,
}
