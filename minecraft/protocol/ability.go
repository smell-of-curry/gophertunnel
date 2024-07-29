package protocol

const (
	Invalid      = -1
	AbilityBuild = 1 << iota
	AbilityMine
	AbilityDoorsAndSwitches
	AbilityOpenContainers
	AbilityAttackPlayers
	AbilityAttackMobs
	AbilityOperatorCommands
	AbilityTeleport
	AbilityInvulnerable
	AbilityFlying
	AbilityMayFly
	AbilityInstantBuild
	AbilityLightning
	AbilityFlySpeed
	AbilityWalkSpeed
	AbilityMuted
	AbilityWorldBuilder
	AbilityNoClip
	AbilityPrivilegedBuilder
	AbilityCount
)

const (
	AbilityLayerTypeCustomCache = iota
	AbilityLayerTypeBase
	AbilityLayerTypeSpectator
	AbilityLayerTypeCommands
	AbilityLayerTypeEditor
)

const (
	AbilityBaseFlySpeed  = 0.05
	AbilityBaseWalkSpeed = 0.1
)

// AbilityData represents various data about the abilities of a player, such as ability layers or permissions.
type AbilityData struct {
	// TargetPlayerRawID is a unique identifier of the player. It appears it is not required to fill this field
	// out with a correct value. Simply writing 0 seems to work.
	TargetPlayerRawID int64
	// PlayerPermissions is the permission level of the player as it shows up in the player list built up using
	// the PlayerList packet.
	PlayerPermissions byte
	// CommandPermissions is a set of permissions that specify what commands a player is allowed to execute.
	CommandPermissions byte
	// Layers contains all ability layers and their potential values. This should at least have one entry, being the
	// base layer.
	Layers []SerializedLayer
}

// Marshal encodes/decodes an AbilityData.
func (x *AbilityData) Marshal(r IO) {
	r.Int64(&x.TargetPlayerRawID)
	r.Uint8(&x.PlayerPermissions)
	r.Uint8(&x.CommandPermissions)
	SliceUint8Length(r, &x.Layers)
}

// SerializedLayer represents the abilities of a specific layer, such as the base layer or the spectator layer.
type SerializedLayer struct {
	// Type represents the type of the layer. This is one of the AbilityLayerType constants defined above.
	SerializedLayer uint16
	// AbilitiesSet is a set of abilities that are enabled for the layer. This is one of the Ability constants defined
	// above.
	AbilitiesSet uint32
	// AbilityValues is a set of values that are associated with the enabled abilities, representing the values of the
	// abilities.
	AbilityValues uint32
	// FlySpeed is the default fly speed of the layer.
	FlySpeed float32
	// WalkSpeed is the default walk speed of the layer.
	WalkSpeed float32
}

// Marshal encodes/decodes an AbilityLayer.
func (x *SerializedLayer) Marshal(r IO) {
	r.Uint16(&x.SerializedLayer)
	r.Uint32(&x.AbilitiesSet)
	r.Uint32(&x.AbilityValues)
	r.Float32(&x.FlySpeed)
	r.Float32(&x.WalkSpeed)
}
