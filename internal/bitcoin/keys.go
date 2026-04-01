package bitcoin

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

// PrivateKey wraps a 32-byte secret scalar on the secp256k1 curve.
type PrivateKey struct {
	key *btcec.PrivateKey
}

// GenerateKey creates a new random private key using the system CSPRNG.
func GenerateKey() (*PrivateKey, error) {
	key, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}
	return &PrivateKey{key: key}, nil
}

// PrivateKeyFromBytes creates a private key from a 32-byte slice.
func PrivateKeyFromBytes(b []byte) (*PrivateKey, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("private key must be 32 bytes, got %d", len(b))
	}

	// Verify the scalar is valid (1 <= key < curve order).
	scalar := new(big.Int).SetBytes(b)
	if scalar.Sign() == 0 {
		return nil, fmt.Errorf("private key must not be zero")
	}
	if scalar.Cmp(btcec.S256().N) >= 0 {
		return nil, fmt.Errorf("private key exceeds curve order")
	}

	key, _ := btcec.PrivKeyFromBytes(b)
	return &PrivateKey{key: key}, nil
}

// Bytes returns the 32-byte big-endian representation of the private key.
func (pk *PrivateKey) Bytes() []byte {
	b := pk.key.Serialize()
	return b
}

// PublicKey returns the corresponding public key.
func (pk *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{key: pk.key.PubKey()}
}

// WIF returns the Wallet Import Format encoding of the private key.
// mainnet=true uses version 0x80, testnet uses 0xEF. compressed adds
// the 0x01 suffix before checksum.
func (pk *PrivateKey) WIF(mainnet, compressed bool) string {
	version := byte(0xEF)
	if mainnet {
		version = 0x80
	}

	payload := pk.Bytes()
	if compressed {
		payload = append(payload, 0x01)
	}

	return Base58CheckEncode(version, payload)
}

// PublicKey wraps a secp256k1 public key point.
type PublicKey struct {
	key *btcec.PublicKey
}

// SerializeCompressed returns the 33-byte SEC compressed encoding.
func (pub *PublicKey) SerializeCompressed() []byte {
	return pub.key.SerializeCompressed()
}

// SerializeUncompressed returns the 65-byte SEC uncompressed encoding.
func (pub *PublicKey) SerializeUncompressed() []byte {
	return pub.key.SerializeUncompressed()
}

// HASH160 returns RIPEMD160(SHA256(compressed pubkey)).
// This is the public key hash used in address derivation.
func (pub *PublicKey) Hash160() [20]byte {
	return HASH160(pub.SerializeCompressed())
}

// GenerateEntropy returns n bytes of cryptographically secure random data.
// Useful for displaying randomness source in the key generator exercise.
func GenerateEntropy(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("generate entropy: %w", err)
	}
	return b, nil
}
