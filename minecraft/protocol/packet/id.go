package packet

const (
	IDKeepAlive = iota
	IDLogin
	IDPlayStatus
	IDServerToClientHandshake
	IDClientToServerHandshake
	IDDisconnect
	IDResourcePacksInfo
	IDResourcePackStack
	IDResourcePackClientResponse
	IDText
	IDSetTime
	IDStartGame
	IDAddPlayer
	IDAddActor
	IDRemoveActor
	IDAddItemActor
	IDServerPlayerPostMovePosition
	IDTakeItemActor
	IDMoveAbsoluteActor
	IDMovePlayer
	IDPassengerJump
	IDUpdateBlock
	IDAddPainting
	// Deprecated: IDTickSync is deprecated.
	IDTickSync
	IDLevelSoundEventV1
	IDLevelEvent
	IDTileEvent
	IDActorEvent
	IDMobEffect
	IDUpdateAttributes
	IDInventoryTransaction
	IDPlayerEquipment
	IDMobArmorEquipment
	IDInteract
	IDBlockPickRequest
	IDActorPickRequest
	IDPlayerAction
	// Deprecated: IDActorFall is deprecated.
	IDActorFall
	IDHurtArmor
	IDSetActorData
	IDSetActorMotion
	IDSetActorLink
	IDSetHealth
	IDSetSpawnPosition
	IDAnimate
	IDRespawn
	IDContainerOpen
	IDContainerClose
	IDPlayerHotbar
	IDInventoryContent
	IDInventorySlot
	IDContainerSetData
	IDCraftingData
	// Deprecated: IDCraftingEvent is deprecated.
	IDCraftingEvent
	IDGuiDataPickItem
	// Deprecated: IDAdventureSettings is deprecated.
	IDAdventureSettings
	IDBlockActorData
	IDPlayerInput
	IDFullChunkData
	IDSetCommandsEnabled
	IDSetDifficulty
	IDChangeDimension
	IDSetPlayerGameType
	IDPlayerList
	IDSimpleEvent
	IDLegacyTelemetryEvent
	IDSpawnExperienceOrb
	IDMapData
	IDMapInfoRequest
	IDRequestChunkRadius
	IDChunkRadiusUpdated
	// Deprecated: IDItemFrameDropItem is deprecated.
	IDItemFrameDropItem
	IDGameRulesChanged
	IDCamera
	IDBossEvent
	IDShowCredits
	IDAvailableCommands
	IDCommandRequest
	IDCommandBlockUpdate
	IDCommandOutput
	IDUpdateTrade
	IDUpdateEquip
	IDResourcePackDataInfo
	IDResourcePackChunkData
	IDResourcePackChunkRequest
	IDTransfer
	IDPlaySound
	IDStopSound
	IDSetTitle
	IDAddBehaviorTree
	IDStructureBlockUpdate
	IDShowStoreOffer
	IDPurchaseReceipt
	IDPlayerSkin
	IDSubclientLogin
	IDAutomationClientConnect
	IDSetLastHurtBy
	IDBookEdit
	IDNPCRequest
	IDPhotoTransfer
	IDShowModalForm
	IDModalFormResponse
	IDServerSettingsRequest
	IDServerSettingsResponse
	IDShowProfile
	IDSetDefaultGameType
	IDRemoveObjective
	IDSetDisplayObjective
	IDSetScore
	IDLabTable
	IDUpdateBlockSynced
	IDMoveDeltaActor
	IDSetScoreboardIdentity
	IDSetLocalPlayerAsInit
	IDUpdateSoftEnum
	IDPing
	IDBlockPalette
	IDScriptCustomEvent
	IDSpawnParticleEffect
	IDAvailableActorIDList
	IDLevelSoundEventV2
	IDNetworkChunkPublisherUpdate
	IDBiomeDefinitionList
	IDLevelSoundEvent
	IDLevelEventGeneric
	IDLecternUpdate
	// Deprecated: IDVideoStreamConnect is deprecated.
	IDVideoStreamConnect
	// Deprecated: IDAddEntity is deprecated.
	IDAddEntity
	// Deprecated: IDRemoveEntity is deprecated.
	IDRemoveEntity
	IDClientCacheStatus
	IDOnScreenTextureAnimation
	IDMapCreateLockedCopy
	IDStructureTemplateDataExportRequest
	IDStructureTemplateDataExportResponse
	_ // Marked as Unused
	IDClientCacheBlobStatus
	IDClientCacheMissResponse
	IDEducationSettings
	IDEmote
	IDMultiplayerSettings
	IDSettingsCommand
	IDAnvilDamage
	IDCompletedUsingItem
	IDNetworkSettings
	IDPlayerAuthInput
	IDCreativeContent
	IDPlayerEnchantOptions
	IDItemStackRequest
	IDItemStackResponse
	IDPlayerArmorDamage
	IDCodeBuilder
	IDUpdatePlayerGameType
	IDEmoteList
	IDPositionTrackingDBServerBroadcast
	IDPositionTrackingDBClientRequest
	IDDebugInfo
	IDPacketViolationWarning
	IDMotionPredictionHints
	IDTriggerAnimation
	IDCameraShake
	IDPlayerFogSetting
	IDCorrectPlayerMovePrediction
	IDItemRegistry
	// Deprecated: IDFilterText is deprecated.
	IDFilterText
	IDClientBoundDebugRenderer
	IDSyncActorProperty
	IDAddVolumeEntity
	IDRemoveVolumeEntity
	IDSimulationType
	IDNpcDialogue
	IDEduUriResource
	IDCreatePhoto
	IDUpdateSubChunkBlocks
	// Deprecated: IDPhotoInfoRequest is deprecated.
	IDPhotoInfoRequest
	IDSubChunk
	IDSubChunkRequest
	IDPlayerStartItemCooldown
	IDScriptMessage
	IDCodeBuilderSource
	IDTickingAreasLoadStatus
	IDDimensionData
	IDAgentAction
	IDChangeMobProperty
	IDLessonProgress
	IDRequestAbility
	IDRequestPermissions
	IDToastRequest
	IDUpdateAbilities
	IDUpdateAdventureSettings
	IDDeathInfo
	IDEditorNetwork
	IDFeatureRegistry
	IDServerStats
	IDRequestNetworkSettings
	IDGameTestRequest
	IDGameTestResults
	IDPlayerClientInputPermissions
	// Deprecated: IDClientCheatAbility is deprecated.
	IDClientCheatAbility
	IDCameraPresets
	IDUnlockedRecipes
)

const (
	IDCameraInstruction = iota + 300
	IDCompressedBiomeDefinitionList
	IDTrimData
	IDOpenSign
	IDAgentAnimation
	IDRefreshEntitlements
	IDPlayerToggleCrafterSlotRequest
	IDSetPlayerInventoryOptions
	IDSetHud
	IDAwardAchievement
	IDClientBoundCloseForm
	_
	IDServerBoundLoadingScreen
	IDJigsawStructureData
	IDCurrentStructureFeature
	IDServerBoundDiagnostics
	IDCameraAimAssist
	IDContainerRegistryCleanup
	IDMovementEffect
	IDSetMovementAuthority
	IDCameraAimAssistPresets
	IDClientCameraAimAssist
	IDClientMovementPredictionSync
)
