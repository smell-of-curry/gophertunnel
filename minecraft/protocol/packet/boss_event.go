package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BossEventAdd = iota
	BossEventPlayerAdded
	BossEventRemove
	BossEventPlayerRemoved
	BossEventUpdatePercent
	BossEventUpdateName
	BossEventUpdateProperties
	BossEventUpdateStyle
	BossEventQuery
)

const (
	BossEventColorGrey = iota
	BossEventColorBlue
	BossEventColorRed
	BossEventColorGreen
	BossEventColorYellow
	BossEventColorPurple
	BossEventColorWhite
)

// BossEvent is sent by the server to make a specific 'boss event' occur in the
// world. It includes features such as showing a boss bar to the player and
// turning the sky dark.
type BossEvent struct {
	// BossEntityUniqueID is the unique ID of the boss entity that the boss
	// event sent involves. By default, the health percentage and title of the
	// boss bar depend on the health and name tag of this entity. If
	// BossEntityUniqueID is the same as the client's entity unique ID, its
	// HealthPercentage and BossBarTitle can be freely altered.
	BossEntityUniqueID int64
	// EventType is the type of the event. The fields written depend on the
	// event type set, and some event types are sent by the client, whereas
	// others are sent by the server. The event type is one of the constants
	// above.
	EventType uint32
	// PlayerUniqueID is the unique ID of the player that is registered to or
	// unregistered from the boss fight. It is set if EventType is either
	// BossEventRegisterPlayer or BossEventUnregisterPlayer.
	PlayerUniqueID int64
	// BossBarTitle is the title shown above the boss bar. It may be set to set
	// a different title if the BossEntityUniqueID matches the client's entity
	// unique ID.
	BossBarTitle string
	// FilteredBossBarTitle is a filtered version of BossBarTitle with all the
	// profanity removed. The client will use this over BossBarTitle if this
	// field is not empty and they have the "Filter Profanity" setting enabled.
	FilteredBossBarTitle string
	// HealthPercentage is the percentage of health that is shown in the boss
	// bar (0.0-1.0). The HealthPercentage may be set to a specific value if the
	// BossEntityUniqueID matches the client's entity unique ID.
	HealthPercentage float32
	// ScreenDarkening currently seems not to do anything.
	ScreenDarkening uint16
	// Color is the color of the boss bar that is shown when a player is
	// subscribed. It is only set if the EventType is BossEventShow,
	// BossEventAppearanceProperties or BossEventTexture. This is functional as
	// of 1.18 and can be any of the BossEventColor constants listed above.
	Color uint32
	// Overlay is the overlay of the boss bar that is shown on top of the boss
	// bar when a player is subscribed. It currently does not function. It is
	// only set if the EventType is BossEventShow, BossEventAppearanceProperties
	// or BossEventTexture.
	Overlay uint32
}

// ID ...
func (*BossEvent) ID() uint32 {
	return IDBossEvent
}

func (pk *BossEvent) Marshal(io protocol.IO) {
	io.Varint64(&pk.BossEntityUniqueID)
	io.Varuint32(&pk.EventType)
	switch pk.EventType {
	case BossEventAdd:
		io.String(&pk.BossBarTitle)
		io.String(&pk.FilteredBossBarTitle)
		io.Float32(&pk.HealthPercentage)
		io.Uint16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Color)
		io.Varuint32(&pk.Overlay)
	case BossEventPlayerAdded, BossEventPlayerRemoved, BossEventQuery:
		io.Varint64(&pk.PlayerUniqueID)
	case BossEventRemove:
		// No extra payload for this boss event type.
	case BossEventUpdatePercent:
		io.Float32(&pk.HealthPercentage)
	case BossEventUpdateName:
		io.String(&pk.BossBarTitle)
		io.String(&pk.FilteredBossBarTitle)
	case BossEventUpdateProperties:
		io.Uint16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Color)
		io.Varuint32(&pk.Overlay)
	case BossEventUpdateStyle:
		io.Varuint32(&pk.Color)
		io.Varuint32(&pk.Overlay)
	default:
		io.UnknownEnumOption(pk.EventType, "boss event type")
	}
}
