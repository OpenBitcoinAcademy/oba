package bitcoin

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	var s Stack
	s.Push([]byte{1, 2, 3})
	s.Push([]byte{4, 5})

	if s.Len() != 2 {
		t.Fatalf("len = %d, want 2", s.Len())
	}

	top, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(top, []byte{4, 5}) {
		t.Errorf("top = %x", top)
	}

	top, err = s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(top, []byte{1, 2, 3}) {
		t.Errorf("second = %x", top)
	}

	_, err = s.Pop()
	if err == nil {
		t.Error("expected underflow")
	}
}

func TestOpcodeName(t *testing.T) {
	if OpcodeName(OP_DUP) != "OP_DUP" {
		t.Error("OP_DUP")
	}
	if OpcodeName(OP_HASH160) != "OP_HASH160" {
		t.Error("OP_HASH160")
	}
	if OpcodeName(20) != "OP_PUSH(20)" {
		t.Errorf("push 20 = %q", OpcodeName(20))
	}
}

func TestEngine_P2PKH_Valid(t *testing.T) {
	// Generate a key pair.
	key, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	pub := key.PublicKey()
	pubBytes := pub.SerializeCompressed()
	pubHash := pub.Hash160()

	// Build scripts.
	fakeSig := bytes.Repeat([]byte{0x30}, 71) // DER-ish signature
	scriptSig := BuildP2PKHScriptSig(fakeSig, pubBytes)
	scriptPubKey := BuildP2PKHScript(pubHash)
	combined := CombineScripts(scriptSig, scriptPubKey)

	// Execute.
	engine := NewEngine(combined)
	steps, err := engine.Execute()
	if err != nil {
		t.Fatalf("execute: %v", err)
	}

	if !engine.Result() {
		t.Error("P2PKH should verify successfully")
	}

	// Should have 7 steps: push sig, push pubkey, dup, hash160, push hash,
	// equalverify, checksig.
	if len(steps) != 7 {
		t.Errorf("steps = %d, want 7", len(steps))
		for i, s := range steps {
			t.Logf("  step %d: %s - %s", i, s.Op, s.Desc)
		}
	}
}

func TestEngine_P2PKH_WrongKey(t *testing.T) {
	// Use one key for the script, a different key for the sig.
	key1, _ := GenerateKey()
	key2, _ := GenerateKey()
	pub1Hash := key1.PublicKey().Hash160()
	pub2Bytes := key2.PublicKey().SerializeCompressed()

	fakeSig := bytes.Repeat([]byte{0x30}, 71)
	scriptSig := BuildP2PKHScriptSig(fakeSig, pub2Bytes)
	scriptPubKey := BuildP2PKHScript(pub1Hash)
	combined := CombineScripts(scriptSig, scriptPubKey)

	engine := NewEngine(combined)
	_, err := engine.Execute()
	// Should fail at OP_EQUALVERIFY because pubkey hashes don't match.
	if err == nil {
		t.Fatal("expected verification failure")
	}
	scriptErr, ok := err.(*ScriptError)
	if !ok {
		t.Fatalf("expected ScriptError, got %T", err)
	}
	if scriptErr.Op != "OP_EQUALVERIFY" {
		t.Errorf("error at %s, expected OP_EQUALVERIFY", scriptErr.Op)
	}
}

func TestEngine_StepByStep(t *testing.T) {
	key, _ := GenerateKey()
	pub := key.PublicKey()
	pubBytes := pub.SerializeCompressed()
	pubHash := pub.Hash160()

	fakeSig := []byte{0x30, 0x44} // minimal
	scriptSig := BuildP2PKHScriptSig(fakeSig, pubBytes)
	scriptPubKey := BuildP2PKHScript(pubHash)
	combined := CombineScripts(scriptSig, scriptPubKey)

	engine := NewEngine(combined)
	count := 0
	for !engine.Done() {
		step, err := engine.Step()
		if err != nil {
			t.Fatalf("step %d: %v", count, err)
		}
		if step == nil {
			break
		}
		count++
		if count > 20 {
			t.Fatal("infinite loop")
		}
	}
	if count != 7 {
		t.Errorf("steps = %d", count)
	}
}

func TestBuildP2PKHScript(t *testing.T) {
	hash, _ := hex.DecodeString("89abcdefabbaabbaabbaabbaabbaabbaabbaabba")
	var h [20]byte
	copy(h[:], hash)
	script := BuildP2PKHScript(h)

	// Should be: OP_DUP OP_HASH160 OP_PUSH(20) <20 bytes> OP_EQUALVERIFY OP_CHECKSIG
	if len(script) != 25 {
		t.Fatalf("script length = %d, want 25", len(script))
	}
	if script[0] != OP_DUP {
		t.Error("byte 0 should be OP_DUP")
	}
	if script[1] != OP_HASH160 {
		t.Error("byte 1 should be OP_HASH160")
	}
	if script[2] != 20 {
		t.Error("byte 2 should be push 20")
	}
	if script[23] != OP_EQUALVERIFY {
		t.Error("byte 23 should be OP_EQUALVERIFY")
	}
	if script[24] != OP_CHECKSIG {
		t.Error("byte 24 should be OP_CHECKSIG")
	}
}

func TestEngine_UnsupportedOpcode(t *testing.T) {
	engine := NewEngine([]byte{0xff})
	_, err := engine.Execute()
	if err == nil {
		t.Error("expected error for unsupported opcode")
	}
}

func TestEngine_EmptyScript(t *testing.T) {
	engine := NewEngine([]byte{})
	steps, err := engine.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(steps) != 0 {
		t.Errorf("steps = %d", len(steps))
	}
	if engine.Result() {
		t.Error("empty script should be falsy")
	}
}

func TestEngine_OP_EQUAL(t *testing.T) {
	// Push two identical values, then OP_EQUAL.
	script := []byte{
		3, 0xaa, 0xbb, 0xcc, // push 3 bytes
		3, 0xaa, 0xbb, 0xcc, // push same 3 bytes
		OP_EQUAL,
	}
	engine := NewEngine(script)
	_, err := engine.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if !engine.Result() {
		t.Error("equal values should result in true")
	}
}

func TestEngine_OP_EQUAL_False(t *testing.T) {
	script := []byte{
		2, 0xaa, 0xbb,
		2, 0xcc, 0xdd,
		OP_EQUAL,
	}
	engine := NewEngine(script)
	_, err := engine.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if engine.Result() {
		t.Error("different values should result in false")
	}
}
