package bitcoin

import (
	"encoding/hex"
	"strings"
	"testing"
)

func testKey(t *testing.T) *PrivateKey {
	t.Helper()
	// Private key = 1 (smallest valid key).
	b := make([]byte, 32)
	b[31] = 1
	key, err := PrivateKeyFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func TestP2PKH_Mainnet(t *testing.T) {
	key := testKey(t)
	addr := P2PKH(key.PublicKey(), true)
	if !strings.HasPrefix(addr, "1") {
		t.Errorf("mainnet P2PKH should start with 1, got %q", addr)
	}
	// Known address for private key 1, compressed.
	want := "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"
	if addr != want {
		t.Errorf("P2PKH = %q, want %q", addr, want)
	}
}

func TestP2PKH_Testnet(t *testing.T) {
	key := testKey(t)
	addr := P2PKH(key.PublicKey(), false)
	if !strings.HasPrefix(addr, "m") && !strings.HasPrefix(addr, "n") {
		t.Errorf("testnet P2PKH should start with m or n, got %q", addr)
	}
}

func TestP2SHP2WPKH_Mainnet(t *testing.T) {
	key := testKey(t)
	addr := P2SHP2WPKH(key.PublicKey(), true)
	if !strings.HasPrefix(addr, "3") {
		t.Errorf("mainnet P2SH should start with 3, got %q", addr)
	}
}

func TestP2WPKH_Mainnet(t *testing.T) {
	key := testKey(t)
	addr, err := P2WPKH(key.PublicKey(), true)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(addr, "bc1q") {
		t.Errorf("mainnet P2WPKH should start with bc1q, got %q", addr)
	}
}

func TestP2WPKH_Testnet(t *testing.T) {
	key := testKey(t)
	addr, err := P2WPKH(key.PublicKey(), false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(addr, "tb1q") {
		t.Errorf("testnet P2WPKH should start with tb1q, got %q", addr)
	}
}

func TestP2TR(t *testing.T) {
	// Use the generator point x-coordinate as a 32-byte test key.
	xOnly, _ := hex.DecodeString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798")
	addr, err := P2TR(xOnly, true)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(addr, "bc1p") {
		t.Errorf("mainnet P2TR should start with bc1p, got %q", addr)
	}
}

func TestDeriveAddresses(t *testing.T) {
	key := testKey(t)
	set, err := DeriveAddresses(key.PublicKey(), true)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(set.P2PKH, "1") {
		t.Errorf("P2PKH = %q", set.P2PKH)
	}
	if !strings.HasPrefix(set.P2SHP2WPKH, "3") {
		t.Errorf("P2SHP2WPKH = %q", set.P2SHP2WPKH)
	}
	if !strings.HasPrefix(set.P2WPKH, "bc1q") {
		t.Errorf("P2WPKH = %q", set.P2WPKH)
	}
}

func TestDeriveAddresses_Testnet(t *testing.T) {
	key := testKey(t)
	set, err := DeriveAddresses(key.PublicKey(), false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(set.P2WPKH, "tb1q") {
		t.Errorf("testnet P2WPKH = %q", set.P2WPKH)
	}
}
