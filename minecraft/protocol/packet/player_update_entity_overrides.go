package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PlayerUpdateEntityOverridesTypeClearAll = iota
	PlayerUpdateEntityOverridesTypeRemove
	PlayerUpdateEntityOverridesTypeInt
	PlayerUpdateEntityOverridesTypeFloat
)

// Cereal type-name consts for each Update oneOf alternative (no Enum-as-Value).
// See Mojang bedrock-protocol-docs IntOverride.json / ClearOverride.json / …
const (
	playerUpdateEntityOverridesNameClear = "clearoverrides"
	playerUpdateEntityOverridesNameRemove = "removeoverride"
	playerUpdateEntityOverridesNameInt    = "setintoverride"
	playerUpdateEntityOverridesNameFloat  = "setfloatoverride"
)

// PlayerUpdateEntityOverrides is sent by the server to modify an entity's properties individually.
type PlayerUpdateEntityOverrides struct {
	// EntityUniqueID is the unique ID of the entity. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the entity out for
	// this field.
	EntityUniqueID int64
	// PropertyIndex is the index of the property to modify. The index is unique for each property of an entity.
	PropertyIndex uint32
	// Type is the type of action to perform with the property. It is one of the constants above.
	Type byte
	// IntValue is the new integer value of the property. It is only used when Type is set to
	// PlayerUpdateEntityOverridesTypeInt.
	IntValue int32
	// FloatValue is the new float value of the property. It is only used when Type is set to
	// PlayerUpdateEntityOverridesTypeFloat.
	FloatValue float32
}

// ID ...
func (*PlayerUpdateEntityOverrides) ID() uint32 {
	return IDPlayerUpdateEntityOverrides
}

func playerUpdateEntityOverridesTypeName(t byte) string {
	switch t {
	case PlayerUpdateEntityOverridesTypeClearAll:
		return playerUpdateEntityOverridesNameClear
	case PlayerUpdateEntityOverridesTypeRemove:
		return playerUpdateEntityOverridesNameRemove
	case PlayerUpdateEntityOverridesTypeInt:
		return playerUpdateEntityOverridesNameInt
	case PlayerUpdateEntityOverridesTypeFloat:
		return playerUpdateEntityOverridesNameFloat
	}
	return ""
}

func (pk *PlayerUpdateEntityOverrides) Marshal(io protocol.IO) {
	// 2168+ cereal: ActorUniqueID + propertyIndex + oneOf control (varuint32) +
	// Type string (no Enum-as-Value) + optional value. NOT a duplicate Uint8 type.
	io.ActorUniqueID(&pk.EntityUniqueID)
	io.Varuint32(&pk.PropertyIndex)

	variant := uint32(pk.Type)
	io.Varuint32(&variant)
	pk.Type = byte(variant)

	typeName := playerUpdateEntityOverridesTypeName(pk.Type)
	io.String(&typeName)

	switch pk.Type {
	case PlayerUpdateEntityOverridesTypeClearAll, PlayerUpdateEntityOverridesTypeRemove:
	case PlayerUpdateEntityOverridesTypeInt:
		io.Int32(&pk.IntValue)
	case PlayerUpdateEntityOverridesTypeFloat:
		io.Float32(&pk.FloatValue)
	default:
		io.UnknownEnumOption(pk.Type, "entity override type")
	}
}
