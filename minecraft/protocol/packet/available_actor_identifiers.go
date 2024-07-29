package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AvailableActorIdentifiers is sent by the server at the start of the game to let the client know all
// entities that are available on the server.
type AvailableActorIdentifiers struct {
	// SerializedEntityIdentifiers is a network NBT serialized compound of all entity identifiers that are
	// available in the server.
	SerializedEntityIdentifiers []byte
}

// ID ...
func (*AvailableActorIdentifiers) ID() uint32 {
	return IDAvailableActorIDList
}

func (pk *AvailableActorIdentifiers) Marshal(io protocol.IO) {
	io.Bytes(&pk.SerializedEntityIdentifiers)
}
