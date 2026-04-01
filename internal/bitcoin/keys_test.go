package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	b := key.Bytes()
	if len(b) != 32 {
		t.Errorf("key length = %d, want 32", len(b))
	}

	// Two generated keys should differ.
	key2, _ := GenerateKey()
	if hex.EncodeToString(key.Bytes()) == hex.EncodeToString(key2.Bytes()) {
		t.Error("two generated keys should not be identical")
	}
}

func TestPrivateKeyFromBytes(t *testing.T) {
	// Known private key (Bitcoin wiki example).
	b, _ := hex.DecodeString("E8F32E723DECF4051AEFAC8E2C93C9C5B214313817CDB01A1494B917C8436B35")
	key, err := PrivateKeyFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	if hex.EncodeToString(key.Bytes()) != "e8f32e723decf4051aefac8e2c93c9c5b214313817cdb01a1494b917c8436b35" {
		t.Errorf("key bytes mismatch")
	}
}

func TestPrivateKeyFromBytes_InvalidLength(t *testing.T) {
	_, err := PrivateKeyFromBytes([]byte{1, 2, 3})
	if err == nil {
		t.Error("expected error for wrong length")
	}
}

func TestPrivateKeyFromBytes_Zero(t *testing.T) {
	_, err := PrivateKeyFromBytes(make([]byte, 32))
	if err == nil {
		t.Error("expected error for zero key")
	}
}

func TestPublicKeyCompressed(t *testing.T) {
	// Known keypair from Bitcoin wiki.
	b, _ := hex.DecodeString("E8F32E723DECF4051AEFAC8E2C93C9C5B214313817CDB01A1494B917C8436B35")
	key, _ := PrivateKeyFromBytes(b)
	pub := key.PublicKey()

	compressed := pub.SerializeCompressed()
	if len(compressed) != 33 {
		t.Errorf("compressed length = %d, want 33", len(compressed))
	}
	// First byte should be 02 or 03 (parity).
	if compressed[0] != 0x02 && compressed[0] != 0x03 {
		t.Errorf("first byte = %02x, want 02 or 03", compressed[0])
	}
}

func TestPublicKeyUncompressed(t *testing.T) {
	key, _ := GenerateKey()
	pub := key.PublicKey()
	uncompressed := pub.SerializeUncompressed()
	if len(uncompressed) != 65 {
		t.Errorf("uncompressed length = %d, want 65", len(uncompressed))
	}
	if uncompressed[0] != 0x04 {
		t.Errorf("first byte = %02x, want 04", uncompressed[0])
	}
}

func TestWIF_Mainnet_Compressed(t *testing.T) {
	// Known test vector: private key 1 (smallest valid key).
	b := make([]byte, 32)
	b[31] = 1
	key, _ := PrivateKeyFromBytes(b)
	wif := key.WIF(true, true)
	// WIF for private key 1, mainnet, compressed.
	want := "KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn"
	if wif != want {
		t.Errorf("WIF = %q, want %q", wif, want)
	}
}

func TestWIF_Testnet_Compressed(t *testing.T) {
	b := make([]byte, 32)
	b[31] = 1
	key, _ := PrivateKeyFromBytes(b)
	wif := key.WIF(false, true)
	want := "cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87JcbXMTcA"
	if wif != want {
		t.Errorf("WIF = %q, want %q", wif, want)
	}
}

func TestPublicKeyHash160(t *testing.T) {
	key, _ := GenerateKey()
	pub := key.PublicKey()
	hash := pub.Hash160()
	if len(hash) != 20 {
		t.Errorf("hash160 length = %d, want 20", len(hash))
	}
	// Should be deterministic.
	hash2 := pub.Hash160()
	if hash != hash2 {
		t.Error("hash160 should be deterministic")
	}
}

func TestGenerateEntropy(t *testing.T) {
	data, err := GenerateEntropy(32)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 32 {
		t.Errorf("length = %d", len(data))
	}
	// Should not be all zeros (astronomically unlikely).
	allZero := true
	for _, b := range data {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Error("entropy should not be all zeros")
	}
}
