package bitcoin

import (
	"bytes"
	"errors"
	"fmt"
)

// Opcode constants for the Bitcoin Script subset we support.
// Enough for P2PKH verification, which is the foundation for
// understanding how Bitcoin transactions work.
const (
	OP_0           byte = 0x00
	OP_DUP         byte = 0x76
	OP_EQUAL       byte = 0x87
	OP_EQUALVERIFY byte = 0x88
	OP_HASH160     byte = 0xa9
	OP_CHECKSIG    byte = 0xac
)

// OpcodeName returns a human-readable name for an opcode.
func OpcodeName(op byte) string {
	switch op {
	case OP_0:
		return "OP_0"
	case OP_DUP:
		return "OP_DUP"
	case OP_HASH160:
		return "OP_HASH160"
	case OP_EQUAL:
		return "OP_EQUAL"
	case OP_EQUALVERIFY:
		return "OP_EQUALVERIFY"
	case OP_CHECKSIG:
		return "OP_CHECKSIG"
	default:
		if op >= 1 && op <= 75 {
			return fmt.Sprintf("OP_PUSH(%d)", op)
		}
		return fmt.Sprintf("0x%02x", op)
	}
}

// ScriptError represents a script execution failure.
type ScriptError struct {
	Op      string
	Message string
}

func (e *ScriptError) Error() string {
	return fmt.Sprintf("script error at %s: %s", e.Op, e.Message)
}

// Stack is the script execution stack.
type Stack struct {
	data [][]byte
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item []byte) {
	s.data = append(s.data, item)
}

// Pop removes and returns the top element.
func (s *Stack) Pop() ([]byte, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack underflow")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, nil
}

// Peek returns the top element without removing it.
func (s *Stack) Peek() ([]byte, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack empty")
	}
	return s.data[len(s.data)-1], nil
}

// Len returns the stack depth.
func (s *Stack) Len() int {
	return len(s.data)
}

// Items returns a copy of the stack contents (bottom to top).
func (s *Stack) Items() [][]byte {
	out := make([][]byte, len(s.data))
	copy(out, s.data)
	return out
}

// ScriptStep records one step of script execution for the step-by-step
// visualization in the UI.
type ScriptStep struct {
	Op    string   // opcode name
	Stack [][]byte // stack state after this step
	Desc  string   // human-readable description
}

// Engine executes a Bitcoin Script and records each step.
type Engine struct {
	script []byte
	stack  Stack
	steps  []ScriptStep
	pos    int
}

// NewEngine creates a script engine for the given script bytes.
func NewEngine(script []byte) *Engine {
	return &Engine{script: script}
}

// Execute runs the entire script and returns the recorded steps.
// Returns an error if the script fails verification.
func (e *Engine) Execute() ([]ScriptStep, error) {
	for e.pos < len(e.script) {
		if err := e.step(); err != nil {
			return e.steps, err
		}
	}
	return e.steps, nil
}

// Step executes one instruction and returns it.
func (e *Engine) Step() (*ScriptStep, error) {
	if e.pos >= len(e.script) {
		return nil, nil
	}
	if err := e.step(); err != nil {
		return nil, err
	}
	return &e.steps[len(e.steps)-1], nil
}

// Done returns true when the script has been fully executed.
func (e *Engine) Done() bool {
	return e.pos >= len(e.script)
}

// Result returns true if the top of the stack is truthy after execution.
func (e *Engine) Result() bool {
	if e.stack.Len() == 0 {
		return false
	}
	top, _ := e.stack.Peek()
	return isTruthy(top)
}

func (e *Engine) step() error {
	op := e.script[e.pos]
	e.pos++

	// Data push: opcodes 1-75 push that many bytes.
	if op >= 1 && op <= 75 {
		end := e.pos + int(op)
		if end > len(e.script) {
			return &ScriptError{Op: OpcodeName(op), Message: "push exceeds script length"}
		}
		data := e.script[e.pos:end]
		e.pos = end
		e.stack.Push(data)
		e.record(OpcodeName(op), fmt.Sprintf("Push %d bytes onto stack", len(data)))
		return nil
	}

	switch op {
	case OP_0:
		e.stack.Push([]byte{})
		e.record("OP_0", "Push empty byte array (false)")

	case OP_DUP:
		top, err := e.stack.Peek()
		if err != nil {
			return &ScriptError{Op: "OP_DUP", Message: err.Error()}
		}
		dup := make([]byte, len(top))
		copy(dup, top)
		e.stack.Push(dup)
		e.record("OP_DUP", "Duplicate the top stack element")

	case OP_HASH160:
		top, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_HASH160", Message: err.Error()}
		}
		hash := HASH160(top)
		e.stack.Push(hash[:])
		e.record("OP_HASH160", "Replace top with RIPEMD160(SHA256(top))")

	case OP_EQUAL:
		a, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_EQUAL", Message: err.Error()}
		}
		b, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_EQUAL", Message: err.Error()}
		}
		if bytes.Equal(a, b) {
			e.stack.Push([]byte{1})
		} else {
			e.stack.Push([]byte{})
		}
		e.record("OP_EQUAL", "Compare top two elements, push result")

	case OP_EQUALVERIFY:
		a, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_EQUALVERIFY", Message: err.Error()}
		}
		b, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_EQUALVERIFY", Message: err.Error()}
		}
		if !bytes.Equal(a, b) {
			e.record("OP_EQUALVERIFY", "FAILED: top two elements are not equal")
			return &ScriptError{Op: "OP_EQUALVERIFY", Message: "values not equal"}
		}
		e.record("OP_EQUALVERIFY", "Verify top two elements are equal (pass)")

	case OP_CHECKSIG:
		pubkeyBytes, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_CHECKSIG", Message: err.Error()}
		}
		sigBytes, err := e.stack.Pop()
		if err != nil {
			return &ScriptError{Op: "OP_CHECKSIG", Message: err.Error()}
		}
		// Simplified: verify the signature is non-empty and the pubkey
		// is a valid compressed or uncompressed format.
		valid := len(sigBytes) > 0 && (len(pubkeyBytes) == 33 || len(pubkeyBytes) == 65)
		if valid {
			e.stack.Push([]byte{1})
			e.record("OP_CHECKSIG", "Verify signature against public key (pass)")
		} else {
			e.stack.Push([]byte{})
			e.record("OP_CHECKSIG", "Verify signature against public key (fail)")
		}

	default:
		return &ScriptError{Op: fmt.Sprintf("0x%02x", op), Message: "unsupported opcode"}
	}

	return nil
}

func (e *Engine) record(op, desc string) {
	e.steps = append(e.steps, ScriptStep{
		Op:    op,
		Stack: e.stack.Items(),
		Desc:  desc,
	})
}

func isTruthy(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return true
		}
	}
	return false
}

// BuildP2PKHScript builds a standard P2PKH scriptPubKey:
// OP_DUP OP_HASH160 <20-byte-hash> OP_EQUALVERIFY OP_CHECKSIG
func BuildP2PKHScript(pubKeyHash [20]byte) []byte {
	script := make([]byte, 0, 25)
	script = append(script, OP_DUP)
	script = append(script, OP_HASH160)
	script = append(script, 20) // push 20 bytes
	script = append(script, pubKeyHash[:]...)
	script = append(script, OP_EQUALVERIFY)
	script = append(script, OP_CHECKSIG)
	return script
}

// BuildP2PKHScriptSig builds a P2PKH scriptSig (unlocking script):
// <signature> <pubkey>
func BuildP2PKHScriptSig(sig, pubkey []byte) []byte {
	script := make([]byte, 0, 2+len(sig)+len(pubkey))
	script = append(script, byte(len(sig)))
	script = append(script, sig...)
	script = append(script, byte(len(pubkey)))
	script = append(script, pubkey...)
	return script
}

// CombineScripts concatenates scriptSig and scriptPubKey for execution.
// In real Bitcoin, scriptSig runs first, then scriptPubKey runs on the
// resulting stack. For educational simplicity we concatenate them.
func CombineScripts(scriptSig, scriptPubKey []byte) []byte {
	combined := make([]byte, len(scriptSig)+len(scriptPubKey))
	copy(combined, scriptSig)
	copy(combined[len(scriptSig):], scriptPubKey)
	return combined
}

// Ensure sha256 import is used.
