// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	data.proto
	data_battle.proto
	data_capture.proto
	data_gym.proto
	data_logs.proto
	data_player.proto
	enums.proto
	inventory.proto
	inventory_item.proto
	map_fort.proto
	map_pokemon.proto
	maps.proto
	networking_envelopes.proto
	networking_requests.proto
	networking_requests_messages.proto
	networking_responses.proto
	settings.proto
	settings_master.proto
	settings_master_item.proto
	settings_master_pokemon.proto

It has these top-level messages:
	AssetDigestEntry
	DownloadUrlEntry
	PlayerBadge
	PlayerData
	PokedexEntry
	PokemonData
	BattleAction
	BattleLog
	BattleParticipant
	BattlePokemonInfo
	BattleResults
	CaptureAward
	CaptureProbability
	GymMembership
	GymState
	ActionLogEntry
	CatchPokemonLogEntry
	FortSearchLogEntry
	ContactSettings
	Currency
	DailyBonus
	EquippedBadge
	PlayerAvatar
	PlayerCamera
	PlayerCurrency
	PlayerPublicProfile
	PlayerStats
	AppliedItem
	AppliedItems
	EggIncubator
	EggIncubators
	InventoryDelta
	InventoryItem
	InventoryItemData
	InventoryUpgrade
	InventoryUpgrades
	PokemonFamily
	ItemAward
	ItemData
	FortData
	FortLureInfo
	FortModifier
	FortSummary
	MapPokemon
	NearbyPokemon
	WildPokemon
	MapCell
	SpawnPoint
	AuthTicket
	RequestEnvelope
	ResponseEnvelope
	Unknown6
	Unknown6Response
	Request
	AddFortModifierMessage
	AttackGymMessage
	CatchPokemonMessage
	CheckAwardedBadgesMessage
	CheckCodenameAvailableMessage
	ClaimCodenameMessage
	CollectDailyBonusMessage
	CollectDailyDefenderBonusMessage
	DiskEncounterMessage
	DownloadItemTemplatesMessage
	DownloadRemoteConfigVersionMessage
	DownloadSettingsMessage
	EchoMessage
	EncounterMessage
	EncounterTutorialCompleteMessage
	EquipBadgeMessage
	EvolvePokemonMessage
	FortDeployPokemonMessage
	FortDetailsMessage
	FortRecallPokemonMessage
	FortSearchMessage
	GetAssetDigestMessage
	GetDownloadUrlsMessage
	GetGymDetailsMessage
	GetHatchedEggsMessage
	GetIncensePokemonMessage
	GetInventoryMessage
	GetMapObjectsMessage
	GetPlayerMessage
	GetPlayerProfileMessage
	GetSuggestedCodenamesMessage
	IncenseEncounterMessage
	LevelUpRewardsMessage
	MarkTutorialCompleteMessage
	NicknamePokemonMessage
	PlayerUpdateMessage
	RecycleInventoryItemMessage
	ReleasePokemonMessage
	SetAvatarMessage
	SetContactSettingsMessage
	SetFavoritePokemonMessage
	SetPlayerTeamMessage
	SfidaActionLogMessage
	StartGymBattleMessage
	UpgradePokemonMessage
	UseIncenseMessage
	UseItemCaptureMessage
	UseItemEggIncubatorMessage
	UseItemGymMessage
	UseItemPotionMessage
	UseItemReviveMessage
	UseItemXpBoostMessage
	AddFortModifierResponse
	AttackGymResponse
	CatchPokemonResponse
	CheckAwardedBadgesResponse
	CheckCodenameAvailableResponse
	ClaimCodenameResponse
	CollectDailyBonusResponse
	CollectDailyDefenderBonusResponse
	DiskEncounterResponse
	DownloadItemTemplatesResponse
	DownloadRemoteConfigVersionResponse
	DownloadSettingsResponse
	EchoResponse
	EncounterResponse
	EncounterTutorialCompleteResponse
	EquipBadgeResponse
	EvolvePokemonResponse
	FortDeployPokemonResponse
	FortDetailsResponse
	FortRecallPokemonResponse
	FortSearchResponse
	GetAssetDigestResponse
	GetDownloadUrlsResponse
	GetGymDetailsResponse
	GetHatchedEggsResponse
	GetIncensePokemonResponse
	GetInventoryResponse
	GetMapObjectsResponse
	GetPlayerProfileResponse
	GetPlayerResponse
	GetSuggestedCodenamesResponse
	IncenseEncounterResponse
	LevelUpRewardsResponse
	MarkTutorialCompleteResponse
	NicknamePokemonResponse
	PlayerUpdateResponse
	RecycleInventoryItemResponse
	ReleasePokemonResponse
	SetAvatarResponse
	SetContactSettingsResponse
	SetFavoritePokemonResponse
	SetPlayerTeamResponse
	SfidaActionLogResponse
	StartGymBattleResponse
	UpgradePokemonResponse
	UseIncenseResponse
	UseItemCaptureResponse
	UseItemEggIncubatorResponse
	UseItemGymResponse
	UseItemPotionResponse
	UseItemReviveResponse
	UseItemXpBoostResponse
	DownloadSettingsAction
	FortSettings
	GlobalSettings
	InventorySettings
	LevelSettings
	MapSettings
	BadgeSettings
	CameraSettings
	EncounterSettings
	EquippedBadgeSettings
	GymBattleSettings
	GymLevelSettings
	IapItemDisplay
	IapSettings
	ItemSettings
	MoveSequenceSettings
	MoveSettings
	PlayerLevelSettings
	PokemonSettings
	PokemonUpgradeSettings
	TypeEffectiveSettings
	BattleAttributes
	EggIncubatorAttributes
	ExperienceBoostAttributes
	FoodAttributes
	FortModifierAttributes
	IncenseAttributes
	InventoryUpgradeAttributes
	PokeballAttributes
	PotionAttributes
	ReviveAttributes
	CameraAttributes
	EncounterAttributes
	StatsAttributes
*/
package protos

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

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type AssetDigestEntry struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	BundleName string `protobuf:"bytes,2,opt,name=bundle_name,json=bundleName" json:"bundle_name,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Checksum   uint32 `protobuf:"varint,4,opt,name=checksum" json:"checksum,omitempty"`
	Size       int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Key        []byte `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AssetDigestEntry) Reset()                    { *m = AssetDigestEntry{} }
func (m *AssetDigestEntry) String() string            { return proto.CompactTextString(m) }
func (*AssetDigestEntry) ProtoMessage()               {}
func (*AssetDigestEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DownloadUrlEntry struct {
	AssetId  string `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Checksum uint32 `protobuf:"varint,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *DownloadUrlEntry) Reset()                    { *m = DownloadUrlEntry{} }
func (m *DownloadUrlEntry) String() string            { return proto.CompactTextString(m) }
func (*DownloadUrlEntry) ProtoMessage()               {}
func (*DownloadUrlEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PlayerBadge struct {
	BadgeType    BadgeType `protobuf:"varint,1,opt,name=badge_type,json=badgeType,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	Rank         int32     `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"`
	StartValue   int32     `protobuf:"varint,3,opt,name=start_value,json=startValue" json:"start_value,omitempty"`
	EndValue     int32     `protobuf:"varint,4,opt,name=end_value,json=endValue" json:"end_value,omitempty"`
	CurrentValue float64   `protobuf:"fixed64,5,opt,name=current_value,json=currentValue" json:"current_value,omitempty"`
}

func (m *PlayerBadge) Reset()                    { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string            { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()               {}
func (*PlayerBadge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PlayerData struct {
	CreationTimestampMs int64            `protobuf:"varint,1,opt,name=creation_timestamp_ms,json=creationTimestampMs" json:"creation_timestamp_ms,omitempty"`
	Username            string           `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team                TeamColor        `protobuf:"varint,5,opt,name=team,enum=POGOProtos.Enums.TeamColor" json:"team,omitempty"`
	TutorialState       []TutorialState  `protobuf:"varint,7,rep,packed,name=tutorial_state,json=tutorialState,enum=POGOProtos.Enums.TutorialState" json:"tutorial_state,omitempty"`
	Avatar              *PlayerAvatar    `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	MaxPokemonStorage   int32            `protobuf:"varint,9,opt,name=max_pokemon_storage,json=maxPokemonStorage" json:"max_pokemon_storage,omitempty"`
	MaxItemStorage      int32            `protobuf:"varint,10,opt,name=max_item_storage,json=maxItemStorage" json:"max_item_storage,omitempty"`
	DailyBonus          *DailyBonus      `protobuf:"bytes,11,opt,name=daily_bonus,json=dailyBonus" json:"daily_bonus,omitempty"`
	EquippedBadge       *EquippedBadge   `protobuf:"bytes,12,opt,name=equipped_badge,json=equippedBadge" json:"equipped_badge,omitempty"`
	ContactSettings     *ContactSettings `protobuf:"bytes,13,opt,name=contact_settings,json=contactSettings" json:"contact_settings,omitempty"`
	Currencies          []*Currency      `protobuf:"bytes,14,rep,name=currencies" json:"currencies,omitempty"`
}

func (m *PlayerData) Reset()                    { *m = PlayerData{} }
func (m *PlayerData) String() string            { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()               {}
func (*PlayerData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlayerData) GetAvatar() *PlayerAvatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *PlayerData) GetDailyBonus() *DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *PlayerData) GetEquippedBadge() *EquippedBadge {
	if m != nil {
		return m.EquippedBadge
	}
	return nil
}

func (m *PlayerData) GetContactSettings() *ContactSettings {
	if m != nil {
		return m.ContactSettings
	}
	return nil
}

func (m *PlayerData) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

type PokedexEntry struct {
	PokemonId            PokemonId `protobuf:"varint,1,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	TimesEncountered     int32     `protobuf:"varint,2,opt,name=times_encountered,json=timesEncountered" json:"times_encountered,omitempty"`
	TimesCaptured        int32     `protobuf:"varint,3,opt,name=times_captured,json=timesCaptured" json:"times_captured,omitempty"`
	EvolutionStonePieces int32     `protobuf:"varint,4,opt,name=evolution_stone_pieces,json=evolutionStonePieces" json:"evolution_stone_pieces,omitempty"`
	EvolutionStones      int32     `protobuf:"varint,5,opt,name=evolution_stones,json=evolutionStones" json:"evolution_stones,omitempty"`
}

func (m *PokedexEntry) Reset()                    { *m = PokedexEntry{} }
func (m *PokedexEntry) String() string            { return proto.CompactTextString(m) }
func (*PokedexEntry) ProtoMessage()               {}
func (*PokedexEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PokemonData struct {
	Id                     uint64      `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	PokemonId              PokemonId   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Cp                     int32       `protobuf:"varint,3,opt,name=cp" json:"cp,omitempty"`
	Stamina                int32       `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	StaminaMax             int32       `protobuf:"varint,5,opt,name=stamina_max,json=staminaMax" json:"stamina_max,omitempty"`
	Move_1                 PokemonMove `protobuf:"varint,6,opt,name=move_1,json=move1,enum=POGOProtos.Enums.PokemonMove" json:"move_1,omitempty"`
	Move_2                 PokemonMove `protobuf:"varint,7,opt,name=move_2,json=move2,enum=POGOProtos.Enums.PokemonMove" json:"move_2,omitempty"`
	DeployedFortId         string      `protobuf:"bytes,8,opt,name=deployed_fort_id,json=deployedFortId" json:"deployed_fort_id,omitempty"`
	OwnerName              string      `protobuf:"bytes,9,opt,name=owner_name,json=ownerName" json:"owner_name,omitempty"`
	IsEgg                  bool        `protobuf:"varint,10,opt,name=is_egg,json=isEgg" json:"is_egg,omitempty"`
	EggKmWalkedTarget      float64     `protobuf:"fixed64,11,opt,name=egg_km_walked_target,json=eggKmWalkedTarget" json:"egg_km_walked_target,omitempty"`
	EggKmWalkedStart       float64     `protobuf:"fixed64,12,opt,name=egg_km_walked_start,json=eggKmWalkedStart" json:"egg_km_walked_start,omitempty"`
	Origin                 int32       `protobuf:"varint,14,opt,name=origin" json:"origin,omitempty"`
	HeightM                float32     `protobuf:"fixed32,15,opt,name=height_m,json=heightM" json:"height_m,omitempty"`
	WeightKg               float32     `protobuf:"fixed32,16,opt,name=weight_kg,json=weightKg" json:"weight_kg,omitempty"`
	IndividualAttack       int32       `protobuf:"varint,17,opt,name=individual_attack,json=individualAttack" json:"individual_attack,omitempty"`
	IndividualDefense      int32       `protobuf:"varint,18,opt,name=individual_defense,json=individualDefense" json:"individual_defense,omitempty"`
	IndividualStamina      int32       `protobuf:"varint,19,opt,name=individual_stamina,json=individualStamina" json:"individual_stamina,omitempty"`
	CpMultiplier           float32     `protobuf:"fixed32,20,opt,name=cp_multiplier,json=cpMultiplier" json:"cp_multiplier,omitempty"`
	Pokeball               ItemId      `protobuf:"varint,21,opt,name=pokeball,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball,omitempty"`
	CapturedCellId         uint64      `protobuf:"varint,22,opt,name=captured_cell_id,json=capturedCellId" json:"captured_cell_id,omitempty"`
	BattlesAttacked        int32       `protobuf:"varint,23,opt,name=battles_attacked,json=battlesAttacked" json:"battles_attacked,omitempty"`
	BattlesDefended        int32       `protobuf:"varint,24,opt,name=battles_defended,json=battlesDefended" json:"battles_defended,omitempty"`
	EggIncubatorId         string      `protobuf:"bytes,25,opt,name=egg_incubator_id,json=eggIncubatorId" json:"egg_incubator_id,omitempty"`
	CreationTimeMs         uint64      `protobuf:"varint,26,opt,name=creation_time_ms,json=creationTimeMs" json:"creation_time_ms,omitempty"`
	NumUpgrades            int32       `protobuf:"varint,27,opt,name=num_upgrades,json=numUpgrades" json:"num_upgrades,omitempty"`
	AdditionalCpMultiplier float32     `protobuf:"fixed32,28,opt,name=additional_cp_multiplier,json=additionalCpMultiplier" json:"additional_cp_multiplier,omitempty"`
	Favorite               int32       `protobuf:"varint,29,opt,name=favorite" json:"favorite,omitempty"`
	Nickname               string      `protobuf:"bytes,30,opt,name=nickname" json:"nickname,omitempty"`
	FromFort               int32       `protobuf:"varint,31,opt,name=from_fort,json=fromFort" json:"from_fort,omitempty"`
}

func (m *PokemonData) Reset()                    { *m = PokemonData{} }
func (m *PokemonData) String() string            { return proto.CompactTextString(m) }
func (*PokemonData) ProtoMessage()               {}
func (*PokemonData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*AssetDigestEntry)(nil), "POGOProtos.Data.AssetDigestEntry")
	proto.RegisterType((*DownloadUrlEntry)(nil), "POGOProtos.Data.DownloadUrlEntry")
	proto.RegisterType((*PlayerBadge)(nil), "POGOProtos.Data.PlayerBadge")
	proto.RegisterType((*PlayerData)(nil), "POGOProtos.Data.PlayerData")
	proto.RegisterType((*PokedexEntry)(nil), "POGOProtos.Data.PokedexEntry")
	proto.RegisterType((*PokemonData)(nil), "POGOProtos.Data.PokemonData")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0xb9, 0x38, 0xf6, 0xf8, 0x12, 0x67, 0x92, 0x86, 0x6d, 0x4a, 0x69, 0x6b, 0xa8, 0x28,
	0x42, 0x18, 0xd5, 0xf0, 0x03, 0x21, 0x90, 0x48, 0xec, 0x80, 0x4c, 0x09, 0xb5, 0xc6, 0x29, 0x48,
	0xfc, 0x59, 0x8d, 0x77, 0x27, 0xdb, 0x91, 0xf7, 0xc6, 0xee, 0xac, 0x5b, 0xf3, 0x1e, 0xfc, 0xe1,
	0x11, 0x78, 0x0a, 0xde, 0x0c, 0xce, 0x9c, 0x99, 0x75, 0xd6, 0x51, 0xc3, 0xe5, 0x4f, 0x32, 0xf3,
	0x7d, 0xdf, 0x5c, 0xf6, 0x9c, 0xf3, 0x9d, 0x31, 0x21, 0x3e, 0x57, 0x7c, 0x90, 0x66, 0x89, 0x4a,
	0xe8, 0xfe, 0xf4, 0xf9, 0xb7, 0xcf, 0xa7, 0x7a, 0x98, 0x0f, 0xc6, 0x00, 0x9f, 0xb4, 0x44, 0x5c,
	0x44, 0xb9, 0x61, 0x4f, 0x0e, 0xb4, 0xd2, 0x4d, 0x43, 0xbe, 0x12, 0x99, 0x85, 0x8e, 0x64, 0xbc,
	0x14, 0xb1, 0x4a, 0xb2, 0x95, 0x2b, 0x95, 0x88, 0x0c, 0xda, 0xff, 0xa3, 0x46, 0x7a, 0xa7, 0x79,
	0x2e, 0xd4, 0x58, 0x06, 0x22, 0x57, 0xe7, 0xb1, 0xca, 0x56, 0xf4, 0x2e, 0x69, 0x70, 0x8d, 0xb9,
	0xd2, 0x77, 0x6a, 0x0f, 0x6b, 0x4f, 0x9a, 0x6c, 0x0f, 0xe7, 0x13, 0x9f, 0x3e, 0x20, 0xad, 0x79,
	0x11, 0xfb, 0xa1, 0x70, 0x63, 0x1e, 0x09, 0x67, 0x0b, 0x59, 0x62, 0xa0, 0x1f, 0x00, 0xa1, 0x0e,
	0xd9, 0x5b, 0x8a, 0x2c, 0x97, 0x49, 0xec, 0x6c, 0x03, 0xb9, 0xcd, 0xca, 0x29, 0x3d, 0x21, 0x0d,
	0xef, 0xa5, 0xf0, 0x16, 0x79, 0x11, 0x39, 0x3b, 0x40, 0x75, 0xd8, 0x7a, 0x4e, 0x29, 0xd9, 0xc9,
	0xe5, 0xaf, 0xc2, 0xd9, 0x05, 0x7c, 0x97, 0xe1, 0x98, 0xf6, 0xc8, 0xf6, 0x42, 0xac, 0x9c, 0x3a,
	0x40, 0x6d, 0xa6, 0x87, 0xfd, 0x84, 0xf4, 0xc6, 0xc9, 0xab, 0x38, 0x4c, 0xb8, 0xff, 0x22, 0x0b,
	0xff, 0xf5, 0xae, 0xb0, 0x41, 0x91, 0x85, 0xf6, 0x8e, 0x7a, 0xb8, 0x3e, 0x66, 0xbb, 0x72, 0xcc,
	0x3f, 0x5c, 0xab, 0xff, 0x67, 0x8d, 0xb4, 0xa6, 0x18, 0xc4, 0x33, 0xee, 0x07, 0x82, 0x7e, 0x41,
	0xc8, 0x5c, 0x0f, 0x5c, 0xb5, 0x4a, 0x05, 0x1e, 0xd7, 0x1d, 0xde, 0x1b, 0x54, 0x32, 0x71, 0x8e,
	0x39, 0x40, 0xf1, 0x25, 0x48, 0x58, 0x73, 0x5e, 0x0e, 0xf5, 0xd9, 0x19, 0x8f, 0x17, 0x78, 0x1d,
	0x38, 0x5b, 0x8f, 0x75, 0x34, 0x73, 0xc5, 0x33, 0xe5, 0x2e, 0x79, 0x58, 0x94, 0xd7, 0x22, 0x08,
	0xfd, 0xa8, 0x11, 0x7a, 0x8f, 0x34, 0x45, 0xec, 0x5b, 0x7a, 0x07, 0xe9, 0x06, 0x00, 0x86, 0x7c,
	0x8f, 0x74, 0xbc, 0x22, 0xcb, 0x20, 0xa9, 0x56, 0xa0, 0xa3, 0x57, 0x63, 0x6d, 0x0b, 0xa2, 0xa8,
	0xff, 0xfb, 0x2e, 0x21, 0xe6, 0x13, 0x74, 0x95, 0xd0, 0x21, 0xb9, 0xe3, 0x65, 0x82, 0x2b, 0x48,
	0x88, 0xab, 0x64, 0x04, 0x29, 0xe7, 0x51, 0xea, 0x46, 0x39, 0x7e, 0xcc, 0x36, 0x3b, 0x2c, 0xc9,
	0xcb, 0x92, 0xbb, 0xc8, 0x75, 0x84, 0x8a, 0x5c, 0x64, 0x95, 0x84, 0xaf, 0xe7, 0xf4, 0x13, 0xb2,
	0xa3, 0x04, 0x8f, 0xf0, 0xe8, 0x37, 0xc6, 0xe2, 0x12, 0xd8, 0x51, 0x12, 0x26, 0x19, 0x43, 0x21,
	0xfd, 0x8e, 0x74, 0x55, 0x01, 0x55, 0x28, 0x79, 0xe8, 0xc2, 0x01, 0x4a, 0x38, 0x7b, 0x0f, 0xb7,
	0x61, 0xe9, 0x83, 0x37, 0x2c, 0xb5, 0xba, 0x99, 0x96, 0x9d, 0x6d, 0xf5, 0x6a, 0xac, 0xa3, 0xaa,
	0x10, 0xfd, 0x92, 0xd4, 0xf9, 0x12, 0xbe, 0x2a, 0x73, 0x1a, 0x70, 0x7c, 0x6b, 0xf8, 0xfe, 0xe0,
	0x86, 0x29, 0x06, 0xe6, 0xcb, 0xed, 0xbf, 0x53, 0xd4, 0x32, 0xbb, 0x86, 0x0e, 0xc8, 0x61, 0xc4,
	0x5f, 0xbb, 0x69, 0xb2, 0x10, 0x11, 0x44, 0x23, 0x87, 0x9d, 0x79, 0x20, 0x9c, 0x26, 0x46, 0xf9,
	0x00, 0xa8, 0xa9, 0x61, 0x66, 0x86, 0xa0, 0x4f, 0x48, 0x4f, 0xeb, 0xb5, 0x79, 0xd6, 0x62, 0x82,
	0xe2, 0x2e, 0xe0, 0x13, 0x80, 0x4b, 0xe5, 0x88, 0xb4, 0x7c, 0x2e, 0xc3, 0x95, 0x3b, 0x4f, 0xe2,
	0x22, 0x77, 0x5a, 0x78, 0xb9, 0xfe, 0x6d, 0x97, 0x1b, 0x6b, 0xe9, 0x99, 0x56, 0x32, 0xe2, 0xaf,
	0xc7, 0xf4, 0x7b, 0xd2, 0x15, 0xbf, 0x14, 0x32, 0x4d, 0x85, 0xef, 0x62, 0x15, 0x39, 0x6d, 0xdc,
	0xe7, 0xf1, 0x6d, 0xfb, 0x9c, 0x5b, 0x35, 0x56, 0x1f, 0xeb, 0x88, 0xea, 0x94, 0x32, 0xd2, 0xf3,
	0x92, 0x58, 0x71, 0x4f, 0xb9, 0x60, 0x0e, 0x25, 0xe3, 0x20, 0x77, 0x3a, 0xb8, 0xdf, 0x07, 0xb7,
	0xed, 0x37, 0x32, 0xfa, 0x99, 0x95, 0xb3, 0x7d, 0x6f, 0x13, 0xa0, 0x5f, 0x13, 0x62, 0x4a, 0xcd,
	0x93, 0x22, 0x77, 0xba, 0x90, 0xc6, 0xd6, 0xf0, 0xe1, 0xad, 0xbb, 0x19, 0xe5, 0x8a, 0x55, 0xd6,
	0xf4, 0xff, 0xaa, 0x91, 0xb6, 0x8e, 0xb2, 0x2f, 0x5e, 0x1b, 0x37, 0x83, 0xc1, 0xca, 0x7c, 0x58,
	0x3f, 0xbf, 0xb1, 0xa8, 0x6c, 0x66, 0x26, 0x3e, 0x6b, 0xa6, 0xe5, 0x90, 0x7e, 0x44, 0x0e, 0xb0,
	0xa2, 0x5d, 0xd8, 0x3c, 0x29, 0x62, 0x25, 0x32, 0xe1, 0x5b, 0xb7, 0xf5, 0x90, 0x38, 0xbf, 0xc6,
	0xe9, 0x63, 0x28, 0x43, 0x14, 0x7b, 0x3c, 0x55, 0x85, 0x56, 0x1a, 0xf3, 0x75, 0x10, 0x1d, 0x59,
	0x90, 0x7e, 0x46, 0x8e, 0xc5, 0x32, 0x09, 0x0b, 0xf4, 0x0b, 0x24, 0x3d, 0x16, 0x6e, 0x2a, 0x85,
	0x07, 0x9f, 0x6b, 0xcc, 0x78, 0xb4, 0x66, 0x67, 0x9a, 0x9c, 0x22, 0x47, 0x3f, 0x24, 0xbd, 0x1b,
	0xab, 0x72, 0xdb, 0xd9, 0xf6, 0x37, 0xf5, 0x79, 0xff, 0xb7, 0x26, 0x74, 0x18, 0xf3, 0x09, 0xe8,
	0xcf, 0x2e, 0xd9, 0xb2, 0x1f, 0x5e, 0x67, 0x30, 0xba, 0x11, 0x90, 0xad, 0xff, 0x15, 0x10, 0xd8,
	0xcb, 0x4b, 0xed, 0x77, 0xc1, 0x48, 0xb7, 0x66, 0x6d, 0x69, 0x19, 0x73, 0x7b, 0xfb, 0x72, 0x6a,
	0xfb, 0x90, 0x1e, 0xba, 0x50, 0xca, 0xf6, 0xae, 0xc4, 0x42, 0x17, 0xfc, 0x35, 0xc4, 0xa1, 0x1e,
	0x25, 0x4b, 0xe1, 0x3e, 0xc5, 0x76, 0xdc, 0x1d, 0xde, 0xbf, 0xf5, 0x0a, 0x17, 0x20, 0x63, 0xbb,
	0x5a, 0xfc, 0x74, 0xbd, 0x6a, 0x08, 0x1e, 0xff, 0xaf, 0xab, 0x86, 0xda, 0x67, 0xbe, 0x48, 0xc3,
	0x64, 0x05, 0x85, 0x7f, 0x95, 0x64, 0xd8, 0xd9, 0x1b, 0xd8, 0x76, 0xba, 0x25, 0xfe, 0x0d, 0xc0,
	0xf0, 0x81, 0xf7, 0x09, 0x81, 0xe7, 0x40, 0x64, 0xe6, 0x2d, 0x6a, 0xa2, 0xa6, 0x89, 0x08, 0x3e,
	0x45, 0x77, 0x48, 0x5d, 0x42, 0x35, 0x04, 0x01, 0xda, 0xb4, 0xc1, 0x76, 0x65, 0x7e, 0x1e, 0x04,
	0xd0, 0xb2, 0x8e, 0x00, 0x73, 0x17, 0x91, 0xfb, 0x8a, 0x87, 0x50, 0x7b, 0x2e, 0x34, 0x83, 0x40,
	0x28, 0xb4, 0x69, 0x8d, 0x1d, 0x00, 0xf7, 0x2c, 0xfa, 0x09, 0x99, 0x4b, 0x24, 0xe8, 0xc7, 0xe4,
	0x70, 0x73, 0x01, 0x36, 0x68, 0xb4, 0x63, 0x8d, 0xf5, 0x2a, 0xfa, 0x99, 0xc6, 0xe9, 0x31, 0xa9,
	0x43, 0x8f, 0x0a, 0x64, 0x0c, 0x96, 0xd0, 0x71, 0xb4, 0x33, 0xfd, 0x52, 0xbd, 0x14, 0x32, 0x78,
	0xa9, 0xdc, 0xc8, 0xd9, 0x07, 0x66, 0x8b, 0xed, 0x99, 0xf9, 0x85, 0x6e, 0xf3, 0xaf, 0x0c, 0xb5,
	0x08, 0x9c, 0x1e, 0x72, 0x0d, 0x03, 0x3c, 0x0b, 0x74, 0x5d, 0xcb, 0xd8, 0x97, 0x4b, 0xe9, 0x17,
	0xd0, 0x33, 0xb9, 0x02, 0x0f, 0x2e, 0x9c, 0x03, 0x53, 0xd7, 0xd7, 0xc4, 0x29, 0xe2, 0x70, 0x57,
	0x5a, 0x11, 0xfb, 0xe2, 0x4a, 0xc4, 0xb9, 0x70, 0xa8, 0xe9, 0x69, 0xd7, 0xcc, 0xd8, 0x10, 0x37,
	0xe4, 0x65, 0x75, 0x1c, 0xde, 0x94, 0xcf, 0x6c, 0x9d, 0xe8, 0x17, 0x07, 0x5e, 0x8b, 0x22, 0x54,
	0x32, 0x0d, 0xa5, 0xc8, 0x9c, 0x23, 0xbc, 0x6b, 0xdb, 0x4b, 0x2f, 0xd6, 0x18, 0xfd, 0x8a, 0x34,
	0x74, 0x0d, 0xce, 0x79, 0x18, 0x3a, 0x77, 0x30, 0xef, 0x8f, 0xaa, 0x79, 0x9f, 0x94, 0x3f, 0x43,
	0x06, 0xba, 0x65, 0xe2, 0x1f, 0x28, 0xdb, 0xf5, 0x12, 0x9d, 0xfe, 0xd2, 0x93, 0xae, 0x27, 0xc2,
	0x50, 0xa7, 0xff, 0x18, 0xb6, 0xd9, 0x61, 0xdd, 0x12, 0x1f, 0x01, 0x0c, 0xe9, 0x07, 0x9b, 0xcd,
	0x21, 0x1c, 0x21, 0xb8, 0xd8, 0x44, 0x05, 0x5c, 0xfc, 0xb6, 0xb1, 0x99, 0xc5, 0x4f, 0x2d, 0x5c,
	0x95, 0x62, 0x4c, 0x7c, 0x90, 0x3a, 0x1b, 0xd2, 0xb1, 0x85, 0xf5, 0xf9, 0x3a, 0xdb, 0x32, 0xf6,
	0x0a, 0xa0, 0x92, 0x4c, 0x9f, 0x7f, 0xd7, 0x94, 0x1f, 0xe0, 0x93, 0x12, 0x9e, 0xa0, 0x72, 0xe3,
	0x2d, 0xd5, 0xcf, 0xe8, 0x89, 0xbd, 0x69, 0xe5, 0x19, 0x85, 0x17, 0xf4, 0x11, 0x69, 0x43, 0xb5,
	0xbb, 0x45, 0x1a, 0x64, 0xdc, 0x87, 0x66, 0x70, 0x0f, 0x8f, 0x6e, 0x01, 0xf6, 0xc2, 0x42, 0xf4,
	0x73, 0xe2, 0x70, 0xdf, 0x97, 0x7a, 0x11, 0x64, 0x62, 0x33, 0xca, 0xef, 0x60, 0x94, 0x8f, 0xaf,
	0xf9, 0x51, 0x35, 0xde, 0xf0, 0x3c, 0x5f, 0xf1, 0x25, 0x14, 0x19, 0xbc, 0xa5, 0xf7, 0xcd, 0x4f,
	0x84, 0x72, 0xae, 0xb9, 0x58, 0x7a, 0x0b, 0xf4, 0xc7, 0xbb, 0xe6, 0xe9, 0x2e, 0xe7, 0xba, 0xe8,
	0xae, 0xb2, 0x24, 0x42, 0x8f, 0x39, 0x0f, 0xec, 0x42, 0x00, 0xb4, 0xb9, 0xce, 0x1a, 0x3f, 0xd7,
	0xf1, 0x07, 0x62, 0x3e, 0x7d, 0x6b, 0x5a, 0x9b, 0x6e, 0xcd, 0xcd, 0xec, 0xd3, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x63, 0x51, 0x2d, 0xe7, 0x83, 0x0a, 0x00, 0x00,
}
