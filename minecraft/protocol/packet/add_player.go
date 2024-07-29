package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddPlayer is sent by the server to the clients when a new player joins the game. The server
// sends this packet to the *other* players.
type AddPlayer struct {
	// UUID is the UUID of the player. It is the same UUID that the client sent in the Login packet at the
	// start of the session. A player with this UUID must exist in the player list (built up using the
	// PlayerList packet), for it to show up in-game.
	UUID uuid.UUID
	// PlayerName is the name of the player. This username is the username that will be set as the initial
	// name tag of the player.
	PlayerName string
	// TargetRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	TargetRuntimeID uint64
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	PlatformChatID string
	// Position is the position to spawn the player on. If the player is on a distance that the viewer cannot
	// see it, the player will still show up if the viewer moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the player spawns with. This velocity will initiate client side
	// movement of the player.
	Velocity mgl32.Vec3
	// Rotation is the rotation of the entity. It is a Vec2, with the X value being the yaw and the Y value being pitch.
	Rotation mgl32.Vec2
	// YHeadRotation is the vertical rotation of the entity. YHeadRotation is measured in degrees.
	// A different value for YHeadRotation than Yaw means that the entity will have its head turned.
	YHeadRotation float32
	// CarriedItem is the item that the player is holding. The item is shown to the viewer as soon as the player
	// itself shows up. Needless to say that this field is rather pointless, as additional packets still must
	// be sent for armour to show up.
	CarriedItem protocol.ItemInstance
	// PlayerGameType is the game type of the player. If set to GameTypeSpectator, the player will not be shown to viewers.
	PlayerGameType int32
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the player looks. Flags include ones such as 'on fire' and 'sprinting'.
	// The metadata values are indexed by their property key.
	// TODO: This is actually different from the AddActor packet, check https://mojang.github.io/bedrock-protocol-docs/html/AddPlayerPacket.html
	EntityMetadata map[uint32]any
	// SyncedProperties is a list of properties that the entity inhibits. These properties define and alter specific
	// attributes of the entity.
	SyncedProperties protocol.PropertySyncData
	// AbilityData represents various data about the abilities of a player, such as ability layers or permissions.
	AbilitiesData protocol.AbilityData
	// ActorLinks is a list of entity links that are currently active on the player. These links alter the
	// way the player shows up when first spawned in terms of it shown as riding an entity. Setting these
	// links is important for new viewers to see the player is riding another entity.
	ActorLinks []protocol.ActorLink
	// DeviceID is the device ID set in one of the files found in the storage of the device of the player. It
	// may be changed freely, so it should not be relied on for anything.
	DeviceID string
	// BuildPlatform is the build platform/device OS of the player that is about to be added, as it sent in
	// the Login packet when joining.
	BuildPlatform int32
}

// ID ...
func (*AddPlayer) ID() uint32 {
	return IDAddPlayer
}

func (pk *AddPlayer) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	io.String(&pk.PlayerName)
	io.Varuint64(&pk.TargetRuntimeID)
	io.String(&pk.PlatformChatID)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Vec2(&pk.Rotation)
	io.Float32(&pk.YHeadRotation)
	io.ItemInstance(&pk.CarriedItem)
	io.Varint32(&pk.PlayerGameType)
	io.EntityMetadata(&pk.EntityMetadata)
	protocol.Single(io, &pk.SyncedProperties)
	protocol.Single(io, &pk.AbilitiesData)
	protocol.Slice(io, &pk.ActorLinks)
	io.String(&pk.DeviceID)
	io.Int32(&pk.BuildPlatform)
}
