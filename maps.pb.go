// Code generated by protoc-gen-go.
// source: maps.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of MapPokemon from map_pokemon.proto

// Ignoring public import of NearbyPokemon from map_pokemon.proto

// Ignoring public import of WildPokemon from map_pokemon.proto

type MapObjectsStatus int32

const (
	MapObjectsStatus_UNSET_STATUS   MapObjectsStatus = 0
	MapObjectsStatus_SUCCESS        MapObjectsStatus = 1
	MapObjectsStatus_LOCATION_UNSET MapObjectsStatus = 2
)

var MapObjectsStatus_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "SUCCESS",
	2: "LOCATION_UNSET",
}
var MapObjectsStatus_value = map[string]int32{
	"UNSET_STATUS":   0,
	"SUCCESS":        1,
	"LOCATION_UNSET": 2,
}

func (x MapObjectsStatus) String() string {
	return proto.EnumName(MapObjectsStatus_name, int32(x))
}
func (MapObjectsStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type MapCell struct {
	// S2 geographic area that the cell covers (http://s2map.com/) (https://code.google.com/archive/p/s2-geometry-library/)
	S2CellId             uint64         `protobuf:"varint,1,opt,name=s2_cell_id,json=s2CellId" json:"s2_cell_id,omitempty"`
	CurrentTimestampMs   int64          `protobuf:"varint,2,opt,name=current_timestamp_ms,json=currentTimestampMs" json:"current_timestamp_ms,omitempty"`
	Forts                []*FortData    `protobuf:"bytes,3,rep,name=forts" json:"forts,omitempty"`
	SpawnPoints          []*SpawnPoint  `protobuf:"bytes,4,rep,name=spawn_points,json=spawnPoints" json:"spawn_points,omitempty"`
	DeletedObjects       []string       `protobuf:"bytes,6,rep,name=deleted_objects,json=deletedObjects" json:"deleted_objects,omitempty"`
	IsTruncatedList      bool           `protobuf:"varint,7,opt,name=is_truncated_list,json=isTruncatedList" json:"is_truncated_list,omitempty"`
	FortSummaries        []*FortSummary `protobuf:"bytes,8,rep,name=fort_summaries,json=fortSummaries" json:"fort_summaries,omitempty"`
	DecimatedSpawnPoints []*SpawnPoint  `protobuf:"bytes,9,rep,name=decimated_spawn_points,json=decimatedSpawnPoints" json:"decimated_spawn_points,omitempty"`
	// Pokemon within 2 steps or less.
	WildPokemons []*WildPokemon `protobuf:"bytes,5,rep,name=wild_pokemons,json=wildPokemons" json:"wild_pokemons,omitempty"`
	// Pokemon within 1 step or none.
	CatchablePokemons []*MapPokemon `protobuf:"bytes,10,rep,name=catchable_pokemons,json=catchablePokemons" json:"catchable_pokemons,omitempty"`
	// Pokemon farther away than 2 steps, but still in the area.
	NearbyPokemons []*NearbyPokemon `protobuf:"bytes,11,rep,name=nearby_pokemons,json=nearbyPokemons" json:"nearby_pokemons,omitempty"`
}

func (m *MapCell) Reset()                    { *m = MapCell{} }
func (m *MapCell) String() string            { return proto.CompactTextString(m) }
func (*MapCell) ProtoMessage()               {}
func (*MapCell) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *MapCell) GetForts() []*FortData {
	if m != nil {
		return m.Forts
	}
	return nil
}

func (m *MapCell) GetSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.SpawnPoints
	}
	return nil
}

func (m *MapCell) GetFortSummaries() []*FortSummary {
	if m != nil {
		return m.FortSummaries
	}
	return nil
}

func (m *MapCell) GetDecimatedSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.DecimatedSpawnPoints
	}
	return nil
}

func (m *MapCell) GetWildPokemons() []*WildPokemon {
	if m != nil {
		return m.WildPokemons
	}
	return nil
}

func (m *MapCell) GetCatchablePokemons() []*MapPokemon {
	if m != nil {
		return m.CatchablePokemons
	}
	return nil
}

func (m *MapCell) GetNearbyPokemons() []*NearbyPokemon {
	if m != nil {
		return m.NearbyPokemons
	}
	return nil
}

type SpawnPoint struct {
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *SpawnPoint) Reset()                    { *m = SpawnPoint{} }
func (m *SpawnPoint) String() string            { return proto.CompactTextString(m) }
func (*SpawnPoint) ProtoMessage()               {}
func (*SpawnPoint) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func init() {
	proto.RegisterType((*MapCell)(nil), "POGOProtos.Map.MapCell")
	proto.RegisterType((*SpawnPoint)(nil), "POGOProtos.Map.SpawnPoint")
	proto.RegisterEnum("POGOProtos.Map.MapObjectsStatus", MapObjectsStatus_name, MapObjectsStatus_value)
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0x5d, 0x6b, 0x1a, 0x4d,
	0x14, 0xc7, 0x63, 0x8c, 0x51, 0x8f, 0x66, 0xd5, 0x21, 0x3c, 0x2c, 0xf2, 0x14, 0xc4, 0x52, 0x1a,
	0x72, 0x21, 0xc5, 0x5c, 0xf7, 0x22, 0xb5, 0x26, 0x0d, 0x24, 0xba, 0xdd, 0x5d, 0x29, 0xf4, 0x66,
	0x18, 0x77, 0x27, 0xed, 0xb4, 0xfb, 0xc6, 0x9e, 0x11, 0xc9, 0x57, 0xea, 0xa7, 0xec, 0xcc, 0xac,
	0xee, 0x26, 0x01, 0xe9, 0xcd, 0x72, 0x5e, 0xfe, 0xe7, 0x37, 0xe7, 0xcf, 0xcc, 0x02, 0xc4, 0x2c,
	0xc3, 0x49, 0x96, 0xa7, 0x32, 0x25, 0x96, 0xb3, 0xbc, 0x5d, 0x3a, 0x3a, 0xc4, 0xc9, 0x03, 0xcb,
	0x86, 0x96, 0xea, 0xd1, 0xc7, 0x34, 0x97, 0x45, 0x7f, 0x38, 0xd0, 0x79, 0x96, 0xfe, 0xe6, 0x71,
	0x9a, 0x14, 0xa5, 0xf1, 0x9f, 0x06, 0x34, 0x95, 0x74, 0xc6, 0xa3, 0x88, 0xfc, 0x0f, 0x80, 0x53,
	0x1a, 0xa8, 0x90, 0x8a, 0xd0, 0xae, 0x8d, 0x6a, 0x17, 0x27, 0x6e, 0x0b, 0xa7, 0xba, 0x77, 0x17,
	0x92, 0x0f, 0x70, 0x1e, 0x6c, 0xf2, 0x9c, 0x27, 0x92, 0x4a, 0x11, 0x73, 0x94, 0x2c, 0xce, 0x68,
	0x8c, 0xf6, 0xb1, 0xd2, 0xd5, 0x5d, 0xb2, 0xeb, 0xf9, 0xfb, 0xd6, 0x03, 0x92, 0x2b, 0x68, 0xe8,
	0xc3, 0xd1, 0xae, 0x8f, 0xea, 0x17, 0x9d, 0xe9, 0x9b, 0xc9, 0xcb, 0xf5, 0x26, 0x37, 0x7a, 0x33,
	0xfd, 0xf9, 0xcc, 0x24, 0x73, 0x0b, 0x2d, 0xf9, 0x08, 0x5d, 0xcc, 0xd8, 0x36, 0x51, 0x7b, 0x8a,
	0x44, 0xcd, 0x9e, 0x98, 0xd9, 0xe1, 0xeb, 0x59, 0x4f, 0x6b, 0x1c, 0x2d, 0x71, 0x3b, 0x58, 0xc6,
	0x48, 0xde, 0x43, 0x2f, 0xe4, 0x11, 0x97, 0x3c, 0xa4, 0xe9, 0xfa, 0x17, 0x0f, 0x14, 0xe1, 0x54,
	0x11, 0xda, 0xae, 0xb5, 0x2b, 0x2f, 0x8b, 0x2a, 0xb9, 0x84, 0x81, 0x40, 0x2a, 0xf3, 0x4d, 0x12,
	0x30, 0xad, 0x8e, 0x04, 0x4a, 0xbb, 0xa9, 0xbc, 0xb4, 0xdc, 0x9e, 0x40, 0x7f, 0x5f, 0xbf, 0x57,
	0x65, 0x72, 0x0b, 0x96, 0x5e, 0x8e, 0xe2, 0x26, 0x8e, 0x59, 0x2e, 0x38, 0xda, 0x2d, 0xb3, 0xd5,
	0xe8, 0xa0, 0x23, 0xcf, 0x28, 0x9f, 0xdc, 0xb3, 0xc7, 0x32, 0x51, 0x63, 0xc4, 0x81, 0xff, 0x42,
	0x1e, 0x88, 0xd8, 0x9c, 0xf8, 0xc2, 0x66, 0xfb, 0x9f, 0x36, 0xcf, 0xcb, 0x49, 0xef, 0x99, 0xdf,
	0x2f, 0x70, 0xb6, 0x15, 0x51, 0xb8, 0xbf, 0x55, 0xb4, 0x1b, 0x06, 0xf4, 0xf6, 0x35, 0xc8, 0xd9,
	0xdd, 0xfa, 0x37, 0x25, 0xde, 0xc5, 0x6e, 0x77, 0x5b, 0x25, 0x48, 0xbe, 0x02, 0x51, 0x8e, 0x83,
	0x9f, 0x6c, 0x1d, 0xf1, 0x0a, 0x07, 0x06, 0x37, 0x3e, 0x84, 0x53, 0xf1, 0x9e, 0x36, 0x28, 0xa7,
	0x4b, 0xe4, 0x02, 0x7a, 0x09, 0x67, 0xf9, 0xfa, 0xa9, 0xe2, 0x75, 0x0c, 0xef, 0xdd, 0x21, 0xde,
	0xc2, 0xc8, 0xf7, 0x48, 0x2b, 0x79, 0x9e, 0xe2, 0xf8, 0x06, 0xa0, 0xf2, 0x4e, 0x86, 0xd0, 0x8a,
	0x98, 0x14, 0x72, 0x13, 0x72, 0xf3, 0x08, 0x6b, 0x6e, 0x99, 0xab, 0xa7, 0xdc, 0x8e, 0xd2, 0xe4,
	0x47, 0xd1, 0xac, 0x9b, 0x66, 0x55, 0xb8, 0x9c, 0x43, 0x5f, 0x1d, 0xba, 0x7b, 0x09, 0x9e, 0x64,
	0x72, 0x83, 0xa4, 0x0f, 0xdd, 0xd5, 0xc2, 0x9b, 0xfb, 0xd4, 0xf3, 0xaf, 0xfd, 0x95, 0xd7, 0x3f,
	0x22, 0x1d, 0x68, 0x7a, 0xab, 0xd9, 0x6c, 0xee, 0x79, 0xfd, 0x1a, 0x21, 0x60, 0xdd, 0x2f, 0x67,
	0xd7, 0xfe, 0xdd, 0x72, 0x41, 0x8d, 0xae, 0x7f, 0xfc, 0xa9, 0xf5, 0xfd, 0xd4, 0xfc, 0x44, 0xe8,
	0x1c, 0x39, 0xb5, 0x75, 0x11, 0x5f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x27, 0x2c, 0xed,
	0x91, 0x03, 0x00, 0x00,
}
