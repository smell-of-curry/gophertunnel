package packet

import (
	"bytes"
	"testing"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func TestPlayerUpdateEntityOverridesIntRoundTrip(t *testing.T) {
	in := &PlayerUpdateEntityOverrides{
		EntityUniqueID: 7,
		PropertyIndex:  3,
		Type:           PlayerUpdateEntityOverridesTypeInt,
		IntValue:       1,
	}
	var buf bytes.Buffer
	in.Marshal(protocol.NewWriter(&buf, 0))

	out := &PlayerUpdateEntityOverrides{}
	out.Marshal(protocol.NewReader(&buf, 0, false))
	if buf.Len() != 0 {
		t.Fatalf("unread bytes: %d %x", buf.Len(), buf.Bytes())
	}
	if *out != *in {
		t.Fatalf("got %+v want %+v", out, in)
	}
}

func TestPlayerUpdateEntityOverridesIntWireShape(t *testing.T) {
	// Control=2, string "setintoverride", int32=1 after uniqueID=1 index=0.
	pk := &PlayerUpdateEntityOverrides{
		EntityUniqueID: 1,
		PropertyIndex:  0,
		Type:           PlayerUpdateEntityOverridesTypeInt,
		IntValue:       1,
	}
	var buf bytes.Buffer
	pk.Marshal(protocol.NewWriter(&buf, 0))
	raw := buf.Bytes()
	if !bytes.Contains(raw, []byte("setintoverride")) {
		t.Fatalf("missing type name in wire: %x", raw)
	}
}
