package protocol

const (
	// EntityLinkNone is set to remove the link between two entities.
	EntityLinkNone = iota
	// EntityLinkRiding is set for entities that have control over the entity they're riding, such as in a
	// minecart.
	EntityLinkRiding
	// EntityLinkPassenger is set for entities being a passenger of a vehicle they enter, such as the back
	// sit of a boat.
	EntityLinkPassenger
)

// ActorLink is a link between two entities, typically being one entity riding another.
type ActorLink struct {
	// ActorUniqueID_A is a unique ID of a entity. For a player sitting
	// in a boat, this is the unique ID of the boat.
	ActorUniqueID_A int64
	// ActorUniqueID_B is a unique ID of a entity. For a player sitting in a
	// boat, this is the unique ID of the player.
	ActorUniqueID_B int64
	// LinkType is one of the types above. It specifies the way the entity is linked to another entity.
	LinkType byte
	// Immediate is set to immediately dismount an entity from another. This should be set when the mount of
	// an entity is killed.
	Immediate bool
	// PassengerInitiated specifies if the link was changed by the passenger, for example the player starting to ride
	// a horse by itself. This is generally true in vanilla environment for players.
	PassengerInitiated bool
}

// Marshal encodes/decodes a single entity link.
func (x *ActorLink) Marshal(r IO) {
	r.Varint64(&x.ActorUniqueID_A)
	r.Varint64(&x.ActorUniqueID_B)
	r.Uint8(&x.LinkType)
	r.Bool(&x.Immediate)
	r.Bool(&x.PassengerInitiated)
}
