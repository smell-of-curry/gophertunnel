package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// BiomeDefinitionList is sent by the server to let the client know all biomes that are available and
// implemented on the server side. It is much like the AvailableActorIdentifiers packet, but instead
// functions for biomes.
type BiomeDefinitionList struct {
	// SerializedBiomeDefinitions is a network NBT serialized compound of all definitions of biomes that are
	// available on the server.
	SerializedBiomeDefinitions []byte
}

// ID ...
func (*BiomeDefinitionList) ID() uint32 {
	return IDBiomeDefinitionList
}

func (pk *BiomeDefinitionList) Marshal(io protocol.IO) {
	io.Bytes(&pk.SerializedBiomeDefinitions)
}
