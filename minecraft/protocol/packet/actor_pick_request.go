package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ActorPickRequest is sent by the client when it tries to pick an entity, so that it gets a spawn egg which
// can spawn that entity.
type ActorPickRequest struct {
	// ActorID is the target unique ID of the entity that was attempted to be picked. The server must find the
	// type of that entity and provide the correct spawn egg to the player.
	ActorID int64
	// MaxSlots is the number of empty hotbar slots the player has
	// its used to decide whether to overwrite a slot or add it to an empty one
	MaxSlots byte
	// WithData is true if the pick request requests the entity metadata.
	WithData bool
}

// ID ...
func (*ActorPickRequest) ID() uint32 {
	return IDActorPickRequest
}

func (pk *ActorPickRequest) Marshal(io protocol.IO) {
	io.Int64(&pk.ActorID)
	io.Uint8(&pk.MaxSlots)
	io.Bool(&pk.WithData)
}
