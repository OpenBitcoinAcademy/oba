package bitcoin

import (
	"encoding/hex"
	"testing"
)

// Known mainnet transaction (Satoshi's first TX to Hal Finney):
// txid: f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16
// This is transaction #170 in block 170.
var knownTxHex = "0100000001c997a5e56e104102fa209c6a852dd90660a20b2d9c352423edce25857fcd3704000000004847304402204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd410220181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d0901ffffffff0200ca9a3b00000000434104ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84cac00286bee0000000043410411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3ac00000000"

func TestDeserializeTx_KnownTransaction(t *testing.T) {
	data, err := hex.DecodeString(knownTxHex)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := DeserializeTx(data)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Version != 1 {
		t.Errorf("version = %d", tx.Version)
	}
	if len(tx.Inputs) != 1 {
		t.Fatalf("inputs = %d", len(tx.Inputs))
	}
	if len(tx.Outputs) != 2 {
		t.Fatalf("outputs = %d", len(tx.Outputs))
	}
	// First output: 10 BTC = 1,000,000,000 satoshis.
	if tx.Outputs[0].Value != 1000000000 {
		t.Errorf("output[0].value = %d", tx.Outputs[0].Value)
	}
	if tx.Locktime != 0 {
		t.Errorf("locktime = %d", tx.Locktime)
	}
}

func TestTx_TxID_KnownTransaction(t *testing.T) {
	data, _ := hex.DecodeString(knownTxHex)
	tx, err := DeserializeTx(data)
	if err != nil {
		t.Fatal(err)
	}

	txid := tx.TxID()
	want := "f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16"
	if txid != want {
		t.Errorf("txid = %s, want %s", txid, want)
	}
}

func TestTx_SerializeRoundTrip(t *testing.T) {
	data, _ := hex.DecodeString(knownTxHex)
	tx, err := DeserializeTx(data)
	if err != nil {
		t.Fatal(err)
	}

	reserialized := tx.SerializedHex()
	if reserialized != knownTxHex {
		t.Errorf("roundtrip mismatch:\ngot  %s\nwant %s", reserialized, knownTxHex)
	}
}

func TestTx_Size(t *testing.T) {
	data, _ := hex.DecodeString(knownTxHex)
	tx, _ := DeserializeTx(data)
	if tx.Size() != len(data) {
		t.Errorf("size = %d, want %d", tx.Size(), len(data))
	}
}

func TestTx_SimpleConstruction(t *testing.T) {
	var prevTxID [32]byte
	prevTxID[0] = 0xaa

	tx := &Tx{
		Version: 1,
		Inputs: []TxIn{
			{PrevTxID: prevTxID, PrevIndex: 0, ScriptSig: []byte{0x00}, Sequence: 0xffffffff},
		},
		Outputs: []TxOut{
			{Value: 50000, ScriptPubKey: []byte{OP_DUP, OP_HASH160}},
		},
		Locktime: 0,
	}

	serialized := tx.Serialize()
	if len(serialized) == 0 {
		t.Error("serialized should not be empty")
	}

	// Roundtrip.
	tx2, err := DeserializeTx(serialized)
	if err != nil {
		t.Fatal(err)
	}
	if tx2.Version != 1 {
		t.Error("version mismatch")
	}
	if tx2.Outputs[0].Value != 50000 {
		t.Error("value mismatch")
	}
	if tx2.TxID() != tx.TxID() {
		t.Error("txid mismatch after roundtrip")
	}
}

func TestDeserializeTx_TooShort(t *testing.T) {
	_, err := DeserializeTx([]byte{0x01, 0x00})
	if err == nil {
		t.Error("expected error for short input")
	}
}

func TestVarInt_Encoding(t *testing.T) {
	tests := []struct {
		val  uint64
		size int
	}{
		{0, 1},
		{252, 1},
		{253, 3},
		{0xffff, 3},
		{0x10000, 5},
	}
	for _, tt := range tests {
		buf := appendVarInt(nil, tt.val)
		if len(buf) != tt.size {
			t.Errorf("varint(%d) size = %d, want %d", tt.val, len(buf), tt.size)
		}
	}
}
