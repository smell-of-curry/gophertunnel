package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LegacyTelemetryEvent is sent by the server to send an event with additional data. It is typically sent to the client for
// telemetry reasons, much like the SimpleEvent packet.
type LegacyTelemetryEvent struct {
	// TargetActorID is the ID of the target actor.
	TargetActorID int64
	// UsePlayerID indicates if the player ID should be used.
	UsePlayerID byte
	// Event is the event that is transmitted.
	Event protocol.Event
}

// ID returns the packet ID.
func (*LegacyTelemetryEvent) ID() uint32 {
	return IDLegacyTelemetryEvent
}

// Marshal writes the packet fields to the protocol IO.
func (pk *LegacyTelemetryEvent) Marshal(io protocol.IO) {
	io.Varint64(&pk.TargetActorID)
	io.EventType(&pk.Event)
	io.Uint8(&pk.UsePlayerID)
	pk.Event.Marshal(io)
}
