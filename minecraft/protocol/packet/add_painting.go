package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddPainting is sent by the server to the client to make a painting entity show up. It is one of the few
// entities that cannot be sent using the AddActor packet.
type AddPainting struct {
	// TargetActorID is the unique ID of the entity. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the entity out for
	// this field.
	TargetActorID int64
	// TargetRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	TargetRuntimeID uint64
	// Position is the position to spawn the entity on. If the entity is on a distance that the player cannot
	// see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Direction is the facing direction of the painting.
	Direction int32
	// Motif is the title of the painting. The title of the painting must be valid.
	Motif string
}

// ID ...
func (*AddPainting) ID() uint32 {
	return IDAddPainting
}

func (pk *AddPainting) Marshal(io protocol.IO) {
	io.Varint64(&pk.TargetActorID)
	io.Varuint64(&pk.TargetRuntimeID)
	io.Vec3(&pk.Position)
	io.Varint32(&pk.Direction)
	io.String(&pk.Motif)
}
