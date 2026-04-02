package bitcoin

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

// BlockHeader holds the 6 fields of a Bitcoin block header.
type BlockHeader struct {
	Version    uint32
	PrevBlock  [32]byte
	MerkleRoot [32]byte
	Timestamp  uint32
	Bits       uint32
	Nonce      uint32
}

// ParseBlockHeader reads an 80-byte block header.
func ParseBlockHeader(data []byte) (*BlockHeader, error) {
	if len(data) != 80 {
		return nil, fmt.Errorf("block header must be 80 bytes, got %d", len(data))
	}
	h := &BlockHeader{
		Version:   binary.LittleEndian.Uint32(data[0:4]),
		Timestamp: binary.LittleEndian.Uint32(data[68:72]),
		Bits:      binary.LittleEndian.Uint32(data[72:76]),
		Nonce:     binary.LittleEndian.Uint32(data[76:80]),
	}
	copy(h.PrevBlock[:], data[4:36])
	copy(h.MerkleRoot[:], data[36:68])
	return h, nil
}

// PrevBlockHex returns the previous block hash in display order (reversed).
func (h *BlockHeader) PrevBlockHex() string {
	return hex.EncodeToString(reverseBytes(h.PrevBlock[:]))
}

// MerkleRootHex returns the merkle root in display order (reversed).
func (h *BlockHeader) MerkleRootHex() string {
	return hex.EncodeToString(reverseBytes(h.MerkleRoot[:]))
}

// Time returns the block timestamp as a time.Time.
func (h *BlockHeader) Time() time.Time {
	return time.Unix(int64(h.Timestamp), 0)
}

// BlockHash computes the double-SHA-256 hash of the header (the block ID).
func (h *BlockHeader) BlockHash() string {
	data := h.Serialize()
	hash := HASH256(data)
	return hex.EncodeToString(reverseBytes(hash[:]))
}

// Serialize encodes the header as 80 bytes.
func (h *BlockHeader) Serialize() []byte {
	data := make([]byte, 80)
	binary.LittleEndian.PutUint32(data[0:4], h.Version)
	copy(data[4:36], h.PrevBlock[:])
	copy(data[36:68], h.MerkleRoot[:])
	binary.LittleEndian.PutUint32(data[68:72], h.Timestamp)
	binary.LittleEndian.PutUint32(data[72:76], h.Bits)
	binary.LittleEndian.PutUint32(data[76:80], h.Nonce)
	return data
}

// DifficultyBits returns the target in compact "bits" notation decoded
// to a human-readable description.
func (h *BlockHeader) DifficultyDescription() string {
	exp := h.Bits >> 24
	coeff := h.Bits & 0x007fffff
	return fmt.Sprintf("0x%06x * 2^(8*(0x%02x-3)) [%d leading zero bits required]",
		coeff, exp, 8*(32-int(exp))+countLeadingZeroBits(coeff))
}

func countLeadingZeroBits(v uint32) int {
	if v == 0 {
		return 24
	}
	n := 0
	for i := 23; i >= 0; i-- {
		if v&(1<<uint(i)) == 0 {
			n++
		} else {
			break
		}
	}
	return n
}
