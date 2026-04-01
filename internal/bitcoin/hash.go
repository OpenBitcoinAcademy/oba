// Package bitcoin provides pure Go cryptographic primitives used in
// Bitcoin: hashing, key generation, address derivation, and encoding.
// All functions are tested against known BIP test vectors.
package bitcoin

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// SHA256 returns the SHA-256 hash of data.
func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

// HASH256 returns the double-SHA-256 hash of data (SHA-256 of SHA-256).
// Used for transaction IDs and block hashes.
func HASH256(data []byte) [32]byte {
	first := sha256.Sum256(data)
	return sha256.Sum256(first[:])
}

// RIPEMD160 returns the RIPEMD-160 hash of data.
func RIPEMD160(data []byte) [20]byte {
	h := ripemd160.New()
	h.Write(data)
	var out [20]byte
	copy(out[:], h.Sum(nil))
	return out
}

// HASH160 returns RIPEMD-160(SHA-256(data)).
// Used in Bitcoin address derivation from public keys.
func HASH160(data []byte) [20]byte {
	sha := sha256.Sum256(data)
	return RIPEMD160(sha[:])
}
