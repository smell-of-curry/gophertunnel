package protocol

const (
	HeightMapDataNone = iota
	HeightMapDataHasData
	HeightMapDataTooHigh
	HeightMapDataTooLow
	HeightMapDataAllCopied
)

const (
	SubChunkResultUndefined = iota
	SubChunkResultSuccess
	SubChunkResultChunkNotFound
	SubChunkResultInvalidDimension
	SubChunkResultPlayerNotFound
	SubChunkResultIndexOutOfBounds
	SubChunkResultSuccessAllAir
)

// SubChunkEntry contains the data of a sub-chunk entry relative to a center sub chunk position, used for the sub-chunk
// requesting system introduced in v1.18.10.
type SubChunkEntry struct {
	// Offset contains the offset between the sub-chunk position and the center position.
	Offset SubChunkOffset
	// Result is always one of the constants defined in the SubChunkResult constants.
	Result byte
	// RawPayload contains the serialized sub-chunk data, if present.
	RawPayload Optional[[]byte]
	// HeightMapType is always one of the constants defined in the HeightMapData constants.
	HeightMapType byte
	// HeightMapData is the data for the height map, if present.
	HeightMapData Optional[[]int8]
	// RenderHeightMapType is always one of the constants defined in the HeightMapData constants.
	RenderHeightMapType byte
	// RenderHeightMapData is the data for the render height map, if present.
	RenderHeightMapData Optional[[]int8]
	// BlobHash is the hash of the blob, if present.
	BlobHash Optional[uint64]
}

// Marshal encodes/decodes a SubChunkEntry assuming the blob cache is enabled.
//
// Wire format (unchanged through 1.26.40 / protocol 2168): payload presence is
// gated by Result (skipped for SuccessAllAir), height-map bytes are gated by
// HeightMapType / RenderHeightMapType, and BlobHash is always a raw uint64.
// Do NOT use OptionalFunc bool prefixes here — that desyncs the client and
// produces a Block disconnect.
func (x *SubChunkEntry) Marshal(r IO) {
	Single(r, &x.Offset)
	r.Uint8(&x.Result)
	if x.Result != SubChunkResultSuccessAllAir {
		payload, _ := x.RawPayload.Value()
		r.ByteSlice(&payload)
		x.RawPayload = Option(payload)
	} else {
		x.RawPayload = Optional[[]byte]{}
	}
	marshalHeightMaps(r, x)
	hash, _ := x.BlobHash.Value()
	r.Uint64(&hash)
	x.BlobHash = Option(hash)
}

// SubChunkEntryNoCache encodes/decodes a SubChunkEntry assuming the blob cache is not enabled.
func SubChunkEntryNoCache(r IO, x *SubChunkEntry) {
	Single(r, &x.Offset)
	r.Uint8(&x.Result)
	payload, _ := x.RawPayload.Value()
	r.ByteSlice(&payload)
	x.RawPayload = Option(payload)
	marshalHeightMaps(r, x)
	x.BlobHash = Optional[uint64]{}
}

func marshalHeightMaps(r IO, x *SubChunkEntry) {
	r.Uint8(&x.HeightMapType)
	if x.HeightMapType == HeightMapDataHasData {
		data, _ := x.HeightMapData.Value()
		if data == nil {
			data = make([]int8, 256)
		}
		FuncSliceOfLen(r, 256, &data, r.Int8)
		x.HeightMapData = Option(data)
	} else {
		x.HeightMapData = Optional[[]int8]{}
	}
	r.Uint8(&x.RenderHeightMapType)
	if x.RenderHeightMapType == HeightMapDataHasData {
		data, _ := x.RenderHeightMapData.Value()
		if data == nil {
			data = make([]int8, 256)
		}
		FuncSliceOfLen(r, 256, &data, r.Int8)
		x.RenderHeightMapData = Option(data)
	} else {
		x.RenderHeightMapData = Optional[[]int8]{}
	}
}

// SubChunkOffset represents an offset from the base position of another sub chunk.
type SubChunkOffset [3]int8

// Marshal encodes/decodes a SubChunkOffset.
func (x *SubChunkOffset) Marshal(r IO) {
	r.Int8(&x[0])
	r.Int8(&x[1])
	r.Int8(&x[2])
}
