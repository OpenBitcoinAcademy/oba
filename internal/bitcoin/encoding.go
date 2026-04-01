package bitcoin

import (
	"fmt"
	"math/big"
	"strings"
)

// Base58 alphabet used by Bitcoin (no 0, O, I, l to avoid ambiguity).
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58Encode encodes a byte slice into a Base58 string.
func Base58Encode(input []byte) string {
	x := new(big.Int).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result []byte
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	// Preserve leading zeros.
	for _, b := range input {
		if b != 0 {
			break
		}
		result = append(result, base58Alphabet[0])
	}

	// Reverse.
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Base58CheckEncode adds a 4-byte checksum and encodes in Base58.
// Used for legacy Bitcoin addresses and WIF private keys.
func Base58CheckEncode(version byte, payload []byte) string {
	data := make([]byte, 1+len(payload))
	data[0] = version
	copy(data[1:], payload)

	checksum := HASH256(data)
	data = append(data, checksum[:4]...)
	return Base58Encode(data)
}

// Base58Decode decodes a Base58 string into bytes.
func Base58Decode(input string) ([]byte, error) {
	result := big.NewInt(0)
	base := big.NewInt(58)

	for _, c := range input {
		idx := strings.IndexRune(base58Alphabet, c)
		if idx < 0 {
			return nil, fmt.Errorf("invalid Base58 character: %c", c)
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(idx)))
	}

	decoded := result.Bytes()

	// Restore leading zeros.
	for _, c := range input {
		if c != rune(base58Alphabet[0]) {
			break
		}
		decoded = append([]byte{0}, decoded...)
	}

	return decoded, nil
}

// Base58CheckDecode decodes a Base58Check string, verifies the checksum,
// and returns the version byte and payload.
func Base58CheckDecode(input string) (byte, []byte, error) {
	decoded, err := Base58Decode(input)
	if err != nil {
		return 0, nil, err
	}
	if len(decoded) < 5 {
		return 0, nil, fmt.Errorf("Base58Check data too short")
	}

	payload := decoded[:len(decoded)-4]
	checksum := decoded[len(decoded)-4:]

	computed := HASH256(payload)
	for i := 0; i < 4; i++ {
		if checksum[i] != computed[i] {
			return 0, nil, fmt.Errorf("Base58Check checksum mismatch")
		}
	}

	return payload[0], payload[1:], nil
}

// Bech32 encoding constants.
const bech32Charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// Bech32Encode encodes a human-readable part and data into a Bech32 string.
// spec is 1 for Bech32 (BIP-173) or 2 for Bech32m (BIP-350).
func Bech32Encode(hrp string, data []byte, spec int) (string, error) {
	values := make([]byte, len(data))
	copy(values, data)
	checksum := bech32Checksum(hrp, values, spec)
	combined := append(values, checksum...)

	var result strings.Builder
	result.WriteString(hrp)
	result.WriteByte('1')
	for _, v := range combined {
		if int(v) >= len(bech32Charset) {
			return "", fmt.Errorf("invalid data value: %d", v)
		}
		result.WriteByte(bech32Charset[v])
	}
	return result.String(), nil
}

// SegwitEncode encodes a witness version and program into a Bech32/Bech32m address.
func SegwitEncode(hrp string, version byte, program []byte) (string, error) {
	conv, err := convertBits(program, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("convert bits: %w", err)
	}
	data := append([]byte{version}, conv...)

	spec := 1 // Bech32 for witness v0
	if version > 0 {
		spec = 2 // Bech32m for witness v1+
	}

	return Bech32Encode(hrp, data, spec)
}

// bech32Checksum computes the Bech32/Bech32m checksum.
func bech32Checksum(hrp string, data []byte, spec int) []byte {
	values := append(bech32HRPExpand(hrp), data...)
	values = append(values, []byte{0, 0, 0, 0, 0, 0}...)

	constant := uint32(1) // Bech32
	if spec == 2 {
		constant = 0x2bc830a3 // Bech32m
	}

	polymod := bech32Polymod(values) ^ constant
	checksum := make([]byte, 6)
	for i := 0; i < 6; i++ {
		checksum[i] = byte((polymod >> uint(5*(5-i))) & 31)
	}
	return checksum
}

func bech32HRPExpand(hrp string) []byte {
	result := make([]byte, 0, len(hrp)*2+1)
	for _, c := range hrp {
		result = append(result, byte(c>>5))
	}
	result = append(result, 0)
	for _, c := range hrp {
		result = append(result, byte(c&31))
	}
	return result
}

func bech32Polymod(values []byte) uint32 {
	gen := [5]uint32{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
	chk := uint32(1)
	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(v)
		for i := 0; i < 5; i++ {
			if (top>>uint(i))&1 == 1 {
				chk ^= gen[i]
			}
		}
	}
	return chk
}

// convertBits converts a byte slice from fromBits to toBits grouping.
func convertBits(data []byte, fromBits, toBits uint, pad bool) ([]byte, error) {
	acc := uint32(0)
	bits := uint(0)
	maxV := uint32((1 << toBits) - 1)
	var result []byte

	for _, b := range data {
		if uint32(b)>>fromBits != 0 {
			return nil, fmt.Errorf("invalid data: %d exceeds %d bits", b, fromBits)
		}
		acc = (acc << fromBits) | uint32(b)
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			result = append(result, byte((acc>>bits)&maxV))
		}
	}

	if pad {
		if bits > 0 {
			result = append(result, byte((acc<<(toBits-bits))&maxV))
		}
	} else if bits >= fromBits {
		return nil, fmt.Errorf("excess padding")
	} else if (acc<<(toBits-bits))&maxV != 0 {
		return nil, fmt.Errorf("non-zero padding")
	}

	return result, nil
}
