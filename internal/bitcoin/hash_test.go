package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestSHA256(t *testing.T) {
	// Known vector: SHA256("abc") from NIST.
	input := []byte("abc")
	got := SHA256(input)
	want := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if hex.EncodeToString(got[:]) != want {
		t.Errorf("SHA256(abc) = %x, want %s", got, want)
	}
}

func TestSHA256_Empty(t *testing.T) {
	got := SHA256([]byte{})
	want := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	if hex.EncodeToString(got[:]) != want {
		t.Errorf("SHA256('') = %x, want %s", got, want)
	}
}

func TestHASH256(t *testing.T) {
	// Double-SHA256 of "hello".
	input := []byte("hello")
	got := HASH256(input)
	// Verified independently.
	first := SHA256(input)
	second := SHA256(first[:])
	if got != second {
		t.Errorf("HASH256 != SHA256(SHA256(input))")
	}
}

func TestRIPEMD160(t *testing.T) {
	// Known vector: RIPEMD160("abc").
	input := []byte("abc")
	got := RIPEMD160(input)
	want := "8eb208f7e05d987a9b044a8e98c6b087f15a0bfc"
	if hex.EncodeToString(got[:]) != want {
		t.Errorf("RIPEMD160(abc) = %x, want %s", got, want)
	}
}

func TestHASH160(t *testing.T) {
	// HASH160 = RIPEMD160(SHA256(data)).
	input := []byte("hello")
	got := HASH160(input)

	sha := SHA256(input)
	ripe := RIPEMD160(sha[:])
	if got != ripe {
		t.Errorf("HASH160 != RIPEMD160(SHA256(input))")
	}
}

func TestSHA256_Bitcoin(t *testing.T) {
	// SHA256("bitcoin") - useful for the hash explorer exercise.
	input := []byte("bitcoin")
	got := SHA256(input)
	want := "6b88c087247aa2f07ee1c5956b8e1a9f4c7f892a70e324f1bb3d161e05ca107b"
	if hex.EncodeToString(got[:]) != want {
		t.Errorf("SHA256(bitcoin) = %x, want %s", got, want)
	}
}
