package bitcoin

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Tx represents a Bitcoin transaction.
type Tx struct {
	Version  uint32
	Inputs   []TxIn
	Outputs  []TxOut
	Locktime uint32
}

// TxIn represents a transaction input.
type TxIn struct {
	PrevTxID  [32]byte // txid of the transaction being spent
	PrevIndex uint32   // output index in the previous transaction
	ScriptSig []byte   // unlocking script
	Sequence  uint32
}

// TxOut represents a transaction output.
type TxOut struct {
	Value        uint64 // satoshis
	ScriptPubKey []byte // locking script
}

// Serialize encodes the transaction in Bitcoin's wire format.
func (tx *Tx) Serialize() []byte {
	var buf []byte

	// Version (4 bytes, little-endian).
	buf = appendUint32(buf, tx.Version)

	// Input count (varint).
	buf = appendVarInt(buf, uint64(len(tx.Inputs)))

	// Inputs.
	for _, in := range tx.Inputs {
		buf = append(buf, reverseBytes(in.PrevTxID[:])...) // txid is reversed in wire format
		buf = appendUint32(buf, in.PrevIndex)
		buf = appendVarInt(buf, uint64(len(in.ScriptSig)))
		buf = append(buf, in.ScriptSig...)
		buf = appendUint32(buf, in.Sequence)
	}

	// Output count (varint).
	buf = appendVarInt(buf, uint64(len(tx.Outputs)))

	// Outputs.
	for _, out := range tx.Outputs {
		buf = appendUint64(buf, out.Value)
		buf = appendVarInt(buf, uint64(len(out.ScriptPubKey)))
		buf = append(buf, out.ScriptPubKey...)
	}

	// Locktime (4 bytes, little-endian).
	buf = appendUint32(buf, tx.Locktime)

	return buf
}

// TxID returns the double-SHA256 hash of the serialized transaction,
// displayed in reversed byte order (Bitcoin convention).
func (tx *Tx) TxID() string {
	serialized := tx.Serialize()
	hash := HASH256(serialized)
	reversed := reverseBytes(hash[:])
	return hex.EncodeToString(reversed)
}

// SerializedHex returns the hex-encoded serialized transaction.
func (tx *Tx) SerializedHex() string {
	return hex.EncodeToString(tx.Serialize())
}

// Size returns the serialized transaction size in bytes.
func (tx *Tx) Size() int {
	return len(tx.Serialize())
}

// DeserializeTx parses a transaction from Bitcoin wire format bytes.
func DeserializeTx(data []byte) (*Tx, error) {
	if len(data) < 10 {
		return nil, fmt.Errorf("transaction too short")
	}

	pos := 0
	read := func(n int) ([]byte, error) {
		if pos+n > len(data) {
			return nil, fmt.Errorf("unexpected end of transaction at byte %d", pos)
		}
		b := data[pos : pos+n]
		pos += n
		return b, nil
	}

	readUint32 := func() (uint32, error) {
		b, err := read(4)
		if err != nil {
			return 0, err
		}
		return binary.LittleEndian.Uint32(b), nil
	}

	readUint64 := func() (uint64, error) {
		b, err := read(8)
		if err != nil {
			return 0, err
		}
		return binary.LittleEndian.Uint64(b), nil
	}

	readVarInt := func() (uint64, error) {
		if pos >= len(data) {
			return 0, fmt.Errorf("unexpected end reading varint")
		}
		first := data[pos]
		pos++
		switch {
		case first < 0xfd:
			return uint64(first), nil
		case first == 0xfd:
			b, err := read(2)
			if err != nil {
				return 0, err
			}
			return uint64(binary.LittleEndian.Uint16(b)), nil
		case first == 0xfe:
			b, err := read(4)
			if err != nil {
				return 0, err
			}
			return uint64(binary.LittleEndian.Uint32(b)), nil
		default: // 0xff
			b, err := read(8)
			if err != nil {
				return 0, err
			}
			return binary.LittleEndian.Uint64(b), nil
		}
	}

	tx := &Tx{}

	// Version.
	var err error
	tx.Version, err = readUint32()
	if err != nil {
		return nil, fmt.Errorf("version: %w", err)
	}

	// Inputs.
	inCount, err := readVarInt()
	if err != nil {
		return nil, fmt.Errorf("input count: %w", err)
	}
	for i := uint64(0); i < inCount; i++ {
		txidBytes, err := read(32)
		if err != nil {
			return nil, fmt.Errorf("input %d txid: %w", i, err)
		}
		var in TxIn
		copy(in.PrevTxID[:], reverseBytes(txidBytes)) // reverse back to internal order
		in.PrevIndex, err = readUint32()
		if err != nil {
			return nil, fmt.Errorf("input %d index: %w", i, err)
		}
		scriptLen, err := readVarInt()
		if err != nil {
			return nil, fmt.Errorf("input %d script len: %w", i, err)
		}
		in.ScriptSig, err = read(int(scriptLen))
		if err != nil {
			return nil, fmt.Errorf("input %d script: %w", i, err)
		}
		in.Sequence, err = readUint32()
		if err != nil {
			return nil, fmt.Errorf("input %d sequence: %w", i, err)
		}
		tx.Inputs = append(tx.Inputs, in)
	}

	// Outputs.
	outCount, err := readVarInt()
	if err != nil {
		return nil, fmt.Errorf("output count: %w", err)
	}
	for i := uint64(0); i < outCount; i++ {
		var out TxOut
		out.Value, err = readUint64()
		if err != nil {
			return nil, fmt.Errorf("output %d value: %w", i, err)
		}
		scriptLen, err := readVarInt()
		if err != nil {
			return nil, fmt.Errorf("output %d script len: %w", i, err)
		}
		out.ScriptPubKey, err = read(int(scriptLen))
		if err != nil {
			return nil, fmt.Errorf("output %d script: %w", i, err)
		}
		tx.Outputs = append(tx.Outputs, out)
	}

	// Locktime.
	tx.Locktime, err = readUint32()
	if err != nil {
		return nil, fmt.Errorf("locktime: %w", err)
	}

	return tx, nil
}

func appendUint32(buf []byte, v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return append(buf, b...)
}

func appendUint64(buf []byte, v uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return append(buf, b...)
}

func appendVarInt(buf []byte, v uint64) []byte {
	switch {
	case v < 0xfd:
		return append(buf, byte(v))
	case v <= 0xffff:
		buf = append(buf, 0xfd)
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v))
		return append(buf, b...)
	case v <= 0xffffffff:
		buf = append(buf, 0xfe)
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v))
		return append(buf, b...)
	default:
		buf = append(buf, 0xff)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, v)
		return append(buf, b...)
	}
}

func reverseBytes(b []byte) []byte {
	r := make([]byte, len(b))
	for i, v := range b {
		r[len(b)-1-i] = v
	}
	return r
}
