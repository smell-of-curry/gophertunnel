package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddActor is sent by the server to the client to spawn an entity to the player. It is used for every entity
// except other players, for which the AddPlayer packet is used.
type AddActor struct {
	// TargetActorID is the unique ID of the entity. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the entity out for
	// this field.
	TargetActorID int64
	// TargetRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	TargetRuntimeID uint64
	// ActorType is the string entity type of the entity, for example 'minecraft:skeleton'. A list of these
	// entities may be found online.
	ActorType string
	// Position is the position to spawn the entity on. If the entity is on a distance that the player cannot
	// see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the entity spawns with. This velocity will initiate client side
	// movement of the entity.
	Velocity mgl32.Vec3
	// Rotation is the rotation of the entity. It is a Vec2, with the X value being the yaw and the Y value being pitch.
	Rotation mgl32.Vec2
	// YHeadRotation is the vertical rotation of the entity. YHeadRotation is measured in degrees.
	// A different value for YHeadRotation than Yaw means that the entity will have its head turned.
	YHeadRotation float32
	// YBodyRotation is the horizontal rotation of the entity. YBodyRotation is measured in degrees.
	// A different value for YBodyRotation than HeadYaw means that the entity will have its body turned
	YBodyRotation float32
	// Attributes is a slice of attributes that the entity has. It includes attributes such as its health,
	// movement speed, etc.
	Attributes []protocol.AttributeValue
	// ActorData is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the entity looks. Flags include ones such as 'on fire' and 'sprinting'.
	// The metadata values are indexed by their property key.
	// TODO: Map to https://mojang.github.io/bedrock-protocol-docs/html/AddActorPacket.html
	ActorData map[uint32]any
	// SynchedProperties is a list of properties that the entity inhibits. These properties define and alter specific
	// attributes of the entity.
	SynchedProperties protocol.PropertySyncData
	// EntityLinks is a list of entity links that are currently active on the entity. These links alter the
	// way the entity shows up when first spawned in terms of it shown as riding an entity. Setting these
	// links is important for new viewers to see the entity is riding another entity.
	ActorLinks []protocol.ActorLink
}

// ID ...
func (*AddActor) ID() uint32 {
	return IDAddActor
}

func (pk *AddActor) Marshal(io protocol.IO) {
	io.Varint64(&pk.TargetActorID)
	io.Varuint64(&pk.TargetRuntimeID)
	io.String(&pk.ActorType)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Vec2(&pk.Rotation)
	io.Float32(&pk.YHeadRotation)
	io.Float32(&pk.YBodyRotation)
	protocol.Slice(io, &pk.Attributes)
	io.EntityMetadata(&pk.ActorData)
	protocol.Single(io, &pk.SynchedProperties)
	protocol.Slice(io, &pk.ActorLinks)
}
