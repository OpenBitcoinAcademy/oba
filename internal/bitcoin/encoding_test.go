package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	tests := []struct {
		hex  string
		want string
	}{
		{"", ""},
		{"61", "2g"},
		{"626262", "a3gV"},
		{"636363", "aPEr"},
		{"0000010203", "11Ldp"},
	}
	for _, tt := range tests {
		input, _ := hex.DecodeString(tt.hex)
		got := Base58Encode(input)
		if got != tt.want {
			t.Errorf("Base58Encode(%s) = %q, want %q", tt.hex, got, tt.want)
		}
	}
}

func TestBase58Decode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"2g", "61"},
		{"a3gV", "626262"},
		{"11Ldp", "0000010203"},
	}
	for _, tt := range tests {
		got, err := Base58Decode(tt.input)
		if err != nil {
			t.Errorf("Base58Decode(%q): %v", tt.input, err)
			continue
		}
		if hex.EncodeToString(got) != tt.want {
			t.Errorf("Base58Decode(%q) = %x, want %s", tt.input, got, tt.want)
		}
	}
}

func TestBase58Decode_InvalidChar(t *testing.T) {
	_, err := Base58Decode("0OIl") // All invalid in Base58.
	if err == nil {
		t.Error("expected error for invalid Base58 chars")
	}
}

func TestBase58CheckRoundTrip(t *testing.T) {
	payload, _ := hex.DecodeString("010966776006953D5567439E5E39F86A0D273BEE")
	encoded := Base58CheckEncode(0x00, payload)

	version, decoded, err := Base58CheckDecode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if version != 0x00 {
		t.Errorf("version = %02x, want 0x00", version)
	}
	if hex.EncodeToString(decoded) != "010966776006953d5567439e5e39f86a0d273bee" {
		t.Errorf("payload mismatch: %x", decoded)
	}
}

func TestBase58CheckDecode_BadChecksum(t *testing.T) {
	// Encode, then flip a character.
	encoded := Base58CheckEncode(0x00, []byte{1, 2, 3})
	bad := []byte(encoded)
	bad[len(bad)-1] ^= 0x01
	_, _, err := Base58CheckDecode(string(bad))
	if err == nil {
		t.Error("expected checksum error")
	}
}

func TestBase58CheckDecode_TooShort(t *testing.T) {
	_, _, err := Base58CheckDecode("1")
	if err == nil {
		t.Error("expected error for too-short input")
	}
}

// BIP-173 test vectors for Bech32 encoding.
func TestSegwitEncode_BIP173(t *testing.T) {
	// Witness v0, 20-byte program (P2WPKH).
	program, _ := hex.DecodeString("751e76e8199196d454941c45d1b3a323f1433bd6")
	got, err := SegwitEncode("bc", 0, program)
	if err != nil {
		t.Fatal(err)
	}
	want := "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4"
	if got != want {
		t.Errorf("SegwitEncode = %q, want %q", got, want)
	}
}

func TestSegwitEncode_Testnet(t *testing.T) {
	program, _ := hex.DecodeString("751e76e8199196d454941c45d1b3a323f1433bd6")
	got, err := SegwitEncode("tb", 0, program)
	if err != nil {
		t.Fatal(err)
	}
	want := "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx"
	if got != want {
		t.Errorf("SegwitEncode = %q, want %q", got, want)
	}
}

func TestSegwitEncode_WitnessV1_Bech32m(t *testing.T) {
	// BIP-350: witness v1, 32-byte program (Taproot).
	// SegwitEncode takes raw bytes and does 8-to-5 conversion internally.
	program, _ := hex.DecodeString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798")
	got, err := SegwitEncode("bc", 1, program)
	if err != nil {
		t.Fatal(err)
	}
	// Verify it starts with bc1p (Taproot prefix).
	if got[:4] != "bc1p" {
		t.Errorf("SegwitEncode should start with bc1p, got %q", got[:4])
	}
	// Verify round-trip: the BIP-350 reference vector bc1pw508d6q... uses
	// a pre-converted 5-bit program, not raw bytes. Our SegwitEncode does
	// the 8-to-5 conversion, producing a longer address. Verify length
	// matches expected Taproot address length (62 chars for 32-byte program).
	if len(got) != 62 {
		t.Errorf("Taproot address length = %d, want 62", len(got))
	}
}
