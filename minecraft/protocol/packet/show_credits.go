package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ShowCreditsStatusStart = iota
	ShowCreditsStatusEnd
)

// ShowCredits is sent by the server to show the Minecraft credits screen to the client. It is typically sent
// when the player beats the ender dragon and leaves the End.
type ShowCredits struct {
	// PlayerRuntimeID is the entity runtime ID of the player to show the credits to. It's not clear why this
	// field is actually here in the first place.
	PlayerRuntimeID uint64
	// StatusType is the status type of the credits. It is one of the constants above, and either starts or
	// stops the credits. It can be nil.
	CreditsState *int32
}

// ID ...
func (*ShowCredits) ID() uint32 {
	return IDShowCredits
}

func (pk *ShowCredits) Marshal(io protocol.IO) {
	io.Varuint64(&pk.PlayerRuntimeID)

	if pk.CreditsState != nil {
		io.Varint32(pk.CreditsState)
	} else {
		// Handle the nil case, if necessary, or set a default value
		var defaultCreditsState int32
		io.Varint32(&defaultCreditsState)
	}
}
