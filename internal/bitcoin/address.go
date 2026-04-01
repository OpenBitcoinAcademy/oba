package bitcoin

// P2PKH returns a legacy Pay-to-Public-Key-Hash address (starts with 1 or m/n).
func P2PKH(pub *PublicKey, mainnet bool) string {
	version := byte(0x6F) // testnet
	if mainnet {
		version = 0x00
	}
	hash := pub.Hash160()
	return Base58CheckEncode(version, hash[:])
}

// P2SHP2WPKH returns a P2SH-wrapped SegWit address (starts with 3 or 2).
// This is the "nested SegWit" format for backward compatibility.
func P2SHP2WPKH(pub *PublicKey, mainnet bool) string {
	version := byte(0xC4) // testnet
	if mainnet {
		version = 0x05
	}

	// Build the witness script: OP_0 <20-byte-hash>.
	pkHash := pub.Hash160()
	witnessScript := make([]byte, 22)
	witnessScript[0] = 0x00 // OP_0
	witnessScript[1] = 0x14 // push 20 bytes
	copy(witnessScript[2:], pkHash[:])

	// P2SH wraps the witness script hash.
	scriptHash := HASH160(witnessScript)
	return Base58CheckEncode(version, scriptHash[:])
}

// P2WPKH returns a native SegWit address (bech32, starts with bc1q or tb1q).
func P2WPKH(pub *PublicKey, mainnet bool) (string, error) {
	hrp := "tb" // testnet
	if mainnet {
		hrp = "bc"
	}
	hash := pub.Hash160()
	return SegwitEncode(hrp, 0, hash[:])
}

// P2TR returns a Taproot address (bech32m, starts with bc1p or tb1p).
// The tweakedPubkey must be the 32-byte x-only public key.
func P2TR(tweakedPubkey []byte, mainnet bool) (string, error) {
	hrp := "tb"
	if mainnet {
		hrp = "bc"
	}
	return SegwitEncode(hrp, 1, tweakedPubkey)
}

// DeriveAddresses returns all address types for a public key.
// Useful for the address builder exercise.
type AddressSet struct {
	P2PKH      string
	P2SHP2WPKH string
	P2WPKH     string
}

// DeriveAddresses computes all standard address types for a public key.
func DeriveAddresses(pub *PublicKey, mainnet bool) (*AddressSet, error) {
	p2wpkh, err := P2WPKH(pub, mainnet)
	if err != nil {
		return nil, err
	}
	return &AddressSet{
		P2PKH:      P2PKH(pub, mainnet),
		P2SHP2WPKH: P2SHP2WPKH(pub, mainnet),
		P2WPKH:     p2wpkh,
	}, nil
}
