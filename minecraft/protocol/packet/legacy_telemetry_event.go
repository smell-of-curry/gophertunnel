package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LegacyTelemetryEvent is sent by the server to send an event with additional data. It is typically sent to the client for
// telemetry reasons, much like the SimpleEvent packet.
type LegacyTelemetryEvent struct {
	// TargetActorID is the ID of the target actor.
	TargetActorID int64
	// EventType is the type of event being transmitted.
	EventType uint32
	// UsePlayerID indicates if the player ID should be used.
	UsePlayerID byte
	// Fields contains the dynamic fields based on the EventType.
	Fields map[string]interface{}
}

// ID returns the packet ID.
func (*LegacyTelemetryEvent) ID() uint32 {
	return IDLegacyTelemetryEvent
}

// Marshal writes the packet fields to the protocol IO.
func (pk *LegacyTelemetryEvent) Marshal(io protocol.IO) {
	io.Varint64(&pk.TargetActorID)
	io.Varuint32(&pk.EventType)
	io.Uint8(&pk.UsePlayerID)

	// Marshal the dynamic fields based on the EventType
	switch pk.EventType {
	case 0:
		io.Varuint32(pk.Fields["AchievementID"].(*uint32))
	case 1:
		io.Varint64(pk.Fields["InteractedEntityID"].(*int64))
		io.Varint32(pk.Fields["InteractionType"].(*int32))
		io.Varint32(pk.Fields["InteractionActorType"].(*int32))
		io.Varint32(pk.Fields["InteractionActorVariant"].(*int32))
		io.Int8(pk.Fields["InteractionActorColor"].(*int8))
	case 2:
		// Currently supported: (0 -> Overworld, 1 -> Nether, 2 -> The End, 3 -> Undefined)
		io.Varuint32(pk.Fields["DimensionID"].(*uint32))
	case 3:
		// Currently supported: (0 -> Overworld, 1 -> Nether, 2 -> The End, 3 -> Undefined)
		io.Varuint32(pk.Fields["SourceDimensionID"].(*uint32))
		// Currently supported: (0 -> Overworld, 1 -> Nether, 2 -> The End, 3 -> Undefined)
		io.Varuint32(pk.Fields["TargetDimensionID"].(*uint32))
	case 4:
		io.Varint64(pk.Fields["InstigatorActorID"].(*int64))
		io.Varint64(pk.Fields["TargetActorID"].(*int64))
		io.Varint32(pk.Fields["InstigatorChildActorType"].(*int32))
		io.Varint32(pk.Fields["DamageSource"].(*int32))
		// -1 if not a trading actor.
		io.Varint32(pk.Fields["TradeTier"].(*int32))
		// Empty if not a trading actor.
		io.String(pk.Fields["TraderName"].(*string))
	case 5:
		io.Varuint32(pk.Fields["ContentsColor"].(*uint32))
		io.Varint32(pk.Fields["ContentsType"].(*int32))
		io.Varint32(pk.Fields["FillLevel"].(*int32))
	case 6:
		io.Varint32(pk.Fields["InstigatorActorID"].(*int32))
		io.Varint32(pk.Fields["InstigatorMobVariant"].(*int32))
		io.Varint32(pk.Fields["DamageSource"].(*int32))
		io.Bool(pk.Fields["DiedInRaid"].(*bool))
	case 7:
		io.Varint64(pk.Fields["BossActorID"].(*int64))
		io.Varint32(pk.Fields["PartySize"].(*int32))
		io.Varint32(pk.Fields["BossType"].(*int32))
	case 8:
		io.Varint32(pk.Fields["Result"].(*int32))
		io.Varint32(pk.Fields["ResultNumber"].(*int32))
		io.String(pk.Fields["CommandName"].(*string))
		io.String(pk.Fields["ResultKey"].(*string))
		io.String(pk.Fields["ResultString"].(*string))
	case 9, 10:
		// No data
	case 11:
		io.Varint32(pk.Fields["SuccessCount"].(*int32))
		io.Varint32(pk.Fields["ErrorCount"].(*int32))
		io.String(pk.Fields["CommandName"].(*string))
		io.String(pk.Fields["ErrorList"].(*string))
	case 12:
		// No data
	case 13:
		io.Varint32(pk.Fields["BornBabyEntityType"].(*int32))
		io.Varint32(pk.Fields["BornBabyEntityVariant"].(*int32))
		io.Int8(pk.Fields["BornBabyEntityColor"].(*int8))
	case 14:
		// No data
	case 15, 16:
		io.Varint32(pk.Fields["BlockInteractionType"].(*int32))
		// Id of the relevant item used in the interaction.
		io.Varint32(pk.Fields["ItemId"].(*int32))
	case 17:
		// Id of the relevant item used in the interaction.
		io.Varint32(pk.Fields["ItemId"].(*int32))
	case 18:
		io.String(pk.Fields["EventName"].(*string))
	case 19:
		// Marked as varint, but should only be positive.
		io.Varuint32(pk.Fields["CurrentRaidWave"].(*uint32))
		// Marked as varint, but should only be positive.
		io.Varuint32(pk.Fields["TotalRaidWaves"].(*uint32))
		io.Bool(pk.Fields["WonRaid"].(*bool))
	case 20, 21, 22:
		// No data
	case 23:
		// Marked as varint, but should only be positive.
		io.Varuint32(pk.Fields["RedstoneLevel"].(*uint32))
	case 24:
		// Id of the relevant item used in the interaction.
		io.Varint32(pk.Fields["ItemId"].(*int32))
		io.Bool(pk.Fields["WasTargetingBarteringPlayer"].(*bool))
	case 25:
		io.Varint32(pk.Fields["PlayerWaxedOrUnwaxedCopperBlockID"].(*int32))
	case 26:
		io.String(pk.Fields["CodeBuilderRuntimeAction"].(*string))
	case 27:
		io.String(pk.Fields["ObjectiveName"].(*string))
		io.Varint32(pk.Fields["CodeBuilderScoreboardScore"].(*int32))
	case 28, 29, 30:
		// No data
	case 31:
		// Id of the relevant item used in the interaction.
		// Marked as `short`
		io.Int16(pk.Fields["ItemID"].(*int16))
		io.Int32(pk.Fields["ItemAux"].(*int32))
		io.Int32(pk.Fields["UseMethod"].(*int32))
		io.Int32(pk.Fields["UseCount"].(*int32))
	default:
		// Handle default or unknown EventType
	}
}
