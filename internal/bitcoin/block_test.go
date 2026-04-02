package bitcoin

import (
	"encoding/hex"
	"testing"
)

// Genesis block header (block 0).
var genesisHeaderHex = "0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c"

func TestParseBlockHeader(t *testing.T) {
	data, err := hex.DecodeString(genesisHeaderHex)
	if err != nil {
		t.Fatal(err)
	}
	h, err := ParseBlockHeader(data)
	if err != nil {
		t.Fatal(err)
	}
	if h.Version != 1 {
		t.Errorf("version = %d", h.Version)
	}
	if h.Nonce == 0 {
		t.Error("nonce should be nonzero")
	}
	if h.Bits != 0x1d00ffff {
		t.Errorf("bits = 0x%x, want 0x1d00ffff", h.Bits)
	}
}

func TestBlockHash_Genesis(t *testing.T) {
	data, _ := hex.DecodeString(genesisHeaderHex)
	h, _ := ParseBlockHeader(data)
	hash := h.BlockHash()
	want := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	if hash != want {
		t.Errorf("block hash = %s, want %s", hash, want)
	}
}

func TestBlockHeaderRoundTrip(t *testing.T) {
	data, _ := hex.DecodeString(genesisHeaderHex)
	h, _ := ParseBlockHeader(data)
	reserialized := hex.EncodeToString(h.Serialize())
	if reserialized != genesisHeaderHex {
		t.Error("round-trip mismatch")
	}
}

func TestParseBlockHeader_WrongSize(t *testing.T) {
	_, err := ParseBlockHeader([]byte{1, 2, 3})
	if err == nil {
		t.Error("expected error for wrong size")
	}
}
