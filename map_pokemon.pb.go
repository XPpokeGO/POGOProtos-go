// Code generated by protoc-gen-go.
// source: map_pokemon.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PokemonData from data.proto

type MapPokemon struct {
	SpawnPointId string    `protobuf:"bytes,1,opt,name=spawn_point_id,json=spawnPointId" json:"spawn_point_id,omitempty"`
	EncounterId  uint64    `protobuf:"fixed64,2,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	PokemonId    PokemonId `protobuf:"varint,3,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	// After this timestamp, the pokemon will be gone.
	ExpirationTimestampMs int64   `protobuf:"varint,4,opt,name=expiration_timestamp_ms,json=expirationTimestampMs" json:"expiration_timestamp_ms,omitempty"`
	Latitude              float64 `protobuf:"fixed64,5,opt,name=latitude" json:"latitude,omitempty"`
	Longitude             float64 `protobuf:"fixed64,6,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *MapPokemon) Reset()                    { *m = MapPokemon{} }
func (m *MapPokemon) String() string            { return proto.CompactTextString(m) }
func (*MapPokemon) ProtoMessage()               {}
func (*MapPokemon) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

type NearbyPokemon struct {
	PokemonId        PokemonId `protobuf:"varint,1,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	DistanceInMeters float32   `protobuf:"fixed32,2,opt,name=distance_in_meters,json=distanceInMeters" json:"distance_in_meters,omitempty"`
	EncounterId      uint64    `protobuf:"fixed64,3,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
}

func (m *NearbyPokemon) Reset()                    { *m = NearbyPokemon{} }
func (m *NearbyPokemon) String() string            { return proto.CompactTextString(m) }
func (*NearbyPokemon) ProtoMessage()               {}
func (*NearbyPokemon) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

type WildPokemon struct {
	EncounterId             uint64       `protobuf:"fixed64,1,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	LastModifiedTimestampMs int64        `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64      `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64      `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	SpawnPointId            string       `protobuf:"bytes,5,opt,name=spawn_point_id,json=spawnPointId" json:"spawn_point_id,omitempty"`
	PokemonData             *PokemonData `protobuf:"bytes,7,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	TimeTillHiddenMs        int32        `protobuf:"varint,11,opt,name=time_till_hidden_ms,json=timeTillHiddenMs" json:"time_till_hidden_ms,omitempty"`
}

func (m *WildPokemon) Reset()                    { *m = WildPokemon{} }
func (m *WildPokemon) String() string            { return proto.CompactTextString(m) }
func (*WildPokemon) ProtoMessage()               {}
func (*WildPokemon) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

func (m *WildPokemon) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func init() {
	proto.RegisterType((*MapPokemon)(nil), "POGOProtos.Map.Pokemon.MapPokemon")
	proto.RegisterType((*NearbyPokemon)(nil), "POGOProtos.Map.Pokemon.NearbyPokemon")
	proto.RegisterType((*WildPokemon)(nil), "POGOProtos.Map.Pokemon.WildPokemon")
}

func init() { proto.RegisterFile("map_pokemon.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xd9, 0x4d, 0x13, 0xc8, 0x6c, 0xa8, 0x8a, 0x11, 0x34, 0x0a, 0x11, 0x82, 0x88, 0x03,
	0x07, 0xba, 0x95, 0x8a, 0xc4, 0x01, 0x0e, 0x48, 0x08, 0x04, 0x39, 0x2c, 0x5d, 0x59, 0x95, 0x90,
	0xb8, 0xac, 0x9c, 0xd8, 0xb4, 0x56, 0xd7, 0x7f, 0x14, 0x3b, 0x02, 0xce, 0xbc, 0x06, 0xaf, 0xc2,
	0xbb, 0x31, 0x76, 0x76, 0x9b, 0xb6, 0x09, 0x48, 0xbd, 0x8d, 0xe7, 0xfb, 0xec, 0xfd, 0xe6, 0xa7,
	0x59, 0xb8, 0xa7, 0x98, 0xad, 0xac, 0x39, 0x17, 0xca, 0xe8, 0xdc, 0x2e, 0x8c, 0x37, 0xe4, 0x61,
	0x79, 0xfc, 0xf1, 0xb8, 0x0c, 0xa5, 0xcb, 0x0b, 0x66, 0xf3, 0x72, 0xa5, 0x8e, 0x32, 0xa1, 0x97,
	0xca, 0xad, 0x4c, 0x23, 0xe0, 0xcc, 0xb3, 0x55, 0x3d, 0xf9, 0x95, 0x02, 0xa0, 0xb1, 0xf1, 0x91,
	0x67, 0xb0, 0xeb, 0x2c, 0xfb, 0xae, 0xf1, 0x59, 0xa9, 0x7d, 0x25, 0xf9, 0x30, 0x79, 0x92, 0x3c,
	0xef, 0xd3, 0x41, 0xec, 0x96, 0xa1, 0x39, 0xe5, 0xe4, 0x29, 0x0c, 0x84, 0x9e, 0x9b, 0xa5, 0xf6,
	0x62, 0x11, 0x3c, 0x29, 0x7a, 0x7a, 0x34, 0xbb, 0xe8, 0xa1, 0xe5, 0x35, 0x40, 0x93, 0x2c, 0x18,
	0x3a, 0x68, 0xd8, 0x3d, 0x7a, 0x94, 0x5f, 0x4a, 0xf7, 0x21, 0x06, 0x6a, 0xbe, 0x3b, 0xe5, 0xb4,
	0x6f, 0xdb, 0x92, 0xbc, 0x82, 0x7d, 0xf1, 0xc3, 0xca, 0x05, 0xf3, 0x12, 0xaf, 0x7b, 0xa9, 0x84,
	0xf3, 0x4c, 0xd9, 0x4a, 0xb9, 0xe1, 0x0e, 0x3e, 0xd4, 0xa1, 0x0f, 0xd6, 0xf2, 0x49, 0xab, 0x16,
	0x8e, 0x8c, 0xe0, 0x4e, 0x8d, 0x4d, 0xbf, 0xe4, 0x62, 0xd8, 0x45, 0x63, 0x42, 0x2f, 0xce, 0x64,
	0x0c, 0xfd, 0xda, 0xe8, 0xd3, 0x95, 0xd8, 0x8b, 0xe2, 0xba, 0x31, 0xf9, 0x9d, 0xc0, 0xdd, 0xcf,
	0x82, 0x2d, 0x66, 0x3f, 0x5b, 0x10, 0x57, 0xf3, 0x27, 0x37, 0xca, 0xff, 0x02, 0x08, 0x97, 0x18,
	0x4a, 0xcf, 0x45, 0x25, 0x75, 0xa5, 0x04, 0x22, 0x71, 0x11, 0x52, 0x4a, 0xf7, 0x5a, 0x65, 0xaa,
	0x8b, 0xd8, 0xdf, 0x80, 0xd9, 0xd9, 0x80, 0x39, 0xf9, 0x93, 0x42, 0xf6, 0x45, 0xd6, 0xbc, 0x0d,
	0x77, 0xfd, 0x4a, 0xb2, 0xc9, 0xff, 0x0d, 0x8c, 0x6a, 0xe6, 0x7c, 0xa5, 0x0c, 0x97, 0xdf, 0xa4,
	0xe0, 0x57, 0x31, 0xa6, 0x11, 0xe3, 0x7e, 0x70, 0x14, 0x8d, 0xe1, 0x5f, 0x20, 0x3b, 0xff, 0x03,
	0xb9, 0x73, 0x0d, 0xe4, 0x96, 0xfd, 0xe9, 0x6e, 0xd9, 0x9f, 0xb7, 0x30, 0x68, 0xe1, 0x86, 0x55,
	0x1c, 0xde, 0x46, 0x4f, 0x76, 0x34, 0xbe, 0x8c, 0xf7, 0x7d, 0x58, 0xd1, 0x66, 0xde, 0x50, 0xd3,
	0xcc, 0xae, 0x0f, 0xe4, 0x00, 0xee, 0x87, 0x79, 0x70, 0xa8, 0xba, 0xae, 0xce, 0x24, 0xe7, 0x42,
	0x87, 0xb1, 0x32, 0x7c, 0xa7, 0x4b, 0xf7, 0x82, 0x74, 0x82, 0xca, 0xa7, 0x28, 0x14, 0xee, 0xdd,
	0xe3, 0xaf, 0x63, 0xcc, 0x77, 0xb6, 0x9c, 0xe5, 0x73, 0xa3, 0x0e, 0xed, 0xb9, 0xd2, 0xa7, 0xe6,
	0x00, 0xe7, 0x3e, 0x8c, 0xff, 0x80, 0x2b, 0x6f, 0x95, 0xc9, 0xac, 0x17, 0xeb, 0x97, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xee, 0x9e, 0xf6, 0x69, 0x55, 0x03, 0x00, 0x00,
}
