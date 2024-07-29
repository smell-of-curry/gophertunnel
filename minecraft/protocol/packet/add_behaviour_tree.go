package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddBehaviorTree is sent by the server to the client. The packet is currently unused by both client and
// server.
type AddBehaviorTree struct {
	// BehaviorTree is a stringified JSON behavior file for a behavior tree.
	BehaviorTreeStructure string
}

// ID ...
func (*AddBehaviorTree) ID() uint32 {
	return IDAddBehaviorTree
}

func (pk *AddBehaviorTree) Marshal(io protocol.IO) {
	io.String(&pk.BehaviorTreeStructure)
}
