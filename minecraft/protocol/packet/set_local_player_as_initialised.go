package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetLocalPlayerAsInit is sent by the client in response to a PlayStatus packet with the status set
// to 3. The packet marks the moment at which the client is fully initialized and can receive any packet
// without discarding it.
type SetLocalPlayerAsInit struct {
	// EntityRuntimeID is the entity runtime ID the player was assigned earlier in the login sequence in the
	// StartGame packet.
	EntityRuntimeID uint64
}

// ID ...
func (*SetLocalPlayerAsInit) ID() uint32 {
	return IDSetLocalPlayerAsInit
}

func (pk *SetLocalPlayerAsInit) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
}
