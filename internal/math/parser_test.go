package math

import (
	"strings"
	"testing"
)

func TestParseSimpleText(t *testing.T) {
	node, err := Parse("abc")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeRow {
		t.Fatalf("expected NodeRow, got %d", node.Type)
	}
	if len(node.Children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(node.Children))
	}
	for i, ch := range node.Children {
		if ch.Type != NodeText {
			t.Errorf("child %d: expected NodeText, got %d", i, ch.Type)
		}
	}
}

func TestParseSingleChar(t *testing.T) {
	node, err := Parse("x")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeText {
		t.Errorf("expected NodeText, got %d", node.Type)
	}
	if node.Value != "x" {
		t.Errorf("value = %q, want %q", node.Value, "x")
	}
}

func TestParseSuperscript(t *testing.T) {
	node, err := Parse("x^2")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeSup {
		t.Fatalf("expected NodeSup, got %d", node.Type)
	}
	if node.Children[0].Value != "x" {
		t.Errorf("base = %q, want %q", node.Children[0].Value, "x")
	}
	if node.Children[1].Value != "2" {
		t.Errorf("sup = %q, want %q", node.Children[1].Value, "2")
	}
}

func TestParseSuperscriptBraced(t *testing.T) {
	node, err := Parse("x^{256}")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeSup {
		t.Fatalf("expected NodeSup, got %d", node.Type)
	}
	// {256} should parse as a row of 3 text nodes.
	sup := node.Children[1]
	if sup.Type != NodeRow && sup.Type != NodeGroup {
		// Single chars get unwrapped, multi-char stay as row.
		if sup.Type == NodeText {
			t.Fatalf("braced {256} should not be a single NodeText")
		}
	}
}

func TestParseSubscript(t *testing.T) {
	node, err := Parse("k_{priv}")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeSub {
		t.Fatalf("expected NodeSub, got %d", node.Type)
	}
	if node.Children[0].Value != "k" {
		t.Errorf("base = %q, want %q", node.Children[0].Value, "k")
	}
}

func TestParseBothScripts(t *testing.T) {
	node, err := Parse("x^2_3")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeSupSub {
		t.Fatalf("expected NodeSupSub, got %d", node.Type)
	}
	if len(node.Children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(node.Children))
	}
}

func TestParseFrac(t *testing.T) {
	node, err := Parse(`\frac{a}{b}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeFrac {
		t.Fatalf("expected NodeFrac, got %d", node.Type)
	}
	if len(node.Children) != 2 {
		t.Fatalf("expected 2 children (num/den), got %d", len(node.Children))
	}
	if node.Children[0].Value != "a" {
		t.Errorf("numerator = %q, want %q", node.Children[0].Value, "a")
	}
	if node.Children[1].Value != "b" {
		t.Errorf("denominator = %q, want %q", node.Children[1].Value, "b")
	}
}

func TestParseSqrt(t *testing.T) {
	node, err := Parse(`\sqrt{x}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeSqrt {
		t.Fatalf("expected NodeSqrt, got %d", node.Type)
	}
	if node.Children[0].Value != "x" {
		t.Errorf("content = %q, want %q", node.Children[0].Value, "x")
	}
}

func TestParseTextCmd(t *testing.T) {
	node, err := Parse(`\text{HASH256}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeTextCmd {
		t.Fatalf("expected NodeTextCmd, got %d", node.Type)
	}
	if node.Value != "HASH256" {
		t.Errorf("value = %q, want %q", node.Value, "HASH256")
	}
}

func TestParseGreekLetters(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`\alpha`, "α"},
		{`\sigma`, "σ"},
		{`\pi`, "π"},
		{`\omega`, "ω"},
		{`\Sigma`, "Σ"},
	}
	for _, tt := range tests {
		node, err := Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%q): %v", tt.input, err)
			continue
		}
		if node.Type != NodeSymbol {
			t.Errorf("Parse(%q): type = %d, want NodeSymbol", tt.input, node.Type)
			continue
		}
		if node.Value != tt.want {
			t.Errorf("Parse(%q): value = %q, want %q", tt.input, node.Value, tt.want)
		}
	}
}

func TestParseOperators(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`\cdot`, "·"},
		{`\times`, "×"},
		{`\neq`, "≠"},
		{`\equiv`, "≡"},
	}
	for _, tt := range tests {
		node, err := Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%q): %v", tt.input, err)
			continue
		}
		if node.Value != tt.want {
			t.Errorf("Parse(%q): value = %q, want %q", tt.input, node.Value, tt.want)
		}
	}
}

func TestParseMod(t *testing.T) {
	node, err := Parse(`\mod`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeMod {
		t.Fatalf("expected NodeMod, got %d", node.Type)
	}
	if node.Value != "mod" {
		t.Errorf("value = %q, want %q", node.Value, "mod")
	}
}

func TestParseMisc(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`\ldots`, "…"},
		{`\infty`, "∞"},
		{`\to`, "→"},
	}
	for _, tt := range tests {
		node, err := Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%q): %v", tt.input, err)
			continue
		}
		if node.Value != tt.want {
			t.Errorf("Parse(%q): value = %q, want %q", tt.input, node.Value, tt.want)
		}
	}
}

func TestParseDelimited(t *testing.T) {
	node, err := Parse(`\left( x \right)`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeDelimited {
		t.Fatalf("expected NodeDelimited, got %d", node.Type)
	}
	if node.Left != "(" {
		t.Errorf("left = %q, want %q", node.Left, "(")
	}
	if node.Right != ")" {
		t.Errorf("right = %q, want %q", node.Right, ")")
	}
}

func TestParseDelimitedInvisible(t *testing.T) {
	// \left. ... \right) is valid (invisible left delimiter).
	node, err := Parse(`\left. x \right)`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Left != "." {
		t.Errorf("left = %q, want %q", node.Left, ".")
	}
}

// --- POC formulas from ARCHITECTURE.md Section 7.3 ---

func TestParsePOC_EllipticCurve(t *testing.T) {
	// y^2 = x^3 + 7
	node, err := Parse("y^2 = x^3 + 7")
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeRow {
		t.Fatalf("expected NodeRow, got %d", node.Type)
	}
}

func TestParsePOC_KeyDerivation(t *testing.T) {
	// k_{priv} \cdot G = K_{pub}
	node, err := Parse(`k_{priv} \cdot G = K_{pub}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeRow {
		t.Fatalf("expected NodeRow, got %d", node.Type)
	}
}

func TestParsePOC_HalvingFormula(t *testing.T) {
	// \frac{reward}{2^n}
	node, err := Parse(`\frac{reward}{2^n}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeFrac {
		t.Fatalf("expected NodeFrac, got %d", node.Type)
	}
}

// --- Error cases ---

func TestParseUnsupportedCommand(t *testing.T) {
	_, err := Parse(`\matrix{a}`)
	if err == nil {
		t.Fatal("expected error for unsupported command")
	}
	if !strings.Contains(err.Error(), "unsupported") {
		t.Errorf("error = %q, want to contain 'unsupported'", err.Error())
	}
}

func TestParseUnclosedGroup(t *testing.T) {
	_, err := Parse(`{abc`)
	if err == nil {
		t.Fatal("expected error for unclosed group")
	}
}

func TestParseUnmatchedRight(t *testing.T) {
	_, err := Parse(`\right)`)
	if err == nil {
		t.Fatal("expected error for unmatched \\right")
	}
}

func TestParseEmptyCommand(t *testing.T) {
	_, err := Parse(`\ `)
	if err == nil {
		t.Fatal("expected error for empty command")
	}
}

func TestParseFracMissingArg(t *testing.T) {
	_, err := Parse(`\frac{a}`)
	if err == nil {
		t.Fatal("expected error for \\frac with missing denominator")
	}
}

func TestParseUnexpectedCloseBrace(t *testing.T) {
	_, err := Parse(`a}b`)
	if err == nil {
		t.Fatal("expected error for unexpected '}'")
	}
}

func TestParseNestedFrac(t *testing.T) {
	node, err := Parse(`\frac{\frac{a}{b}}{c}`)
	if err != nil {
		t.Fatal(err)
	}
	if node.Type != NodeFrac {
		t.Fatalf("expected NodeFrac, got %d", node.Type)
	}
	inner := node.Children[0]
	if inner.Type != NodeFrac {
		t.Fatalf("inner numerator: expected NodeFrac, got %d", inner.Type)
	}
}

func TestParseFracWithSuperscript(t *testing.T) {
	// \frac{reward}{2^n} -- the denominator contains a superscript.
	node, err := Parse(`\frac{reward}{2^n}`)
	if err != nil {
		t.Fatal(err)
	}
	den := node.Children[1]
	if den.Type != NodeSup {
		t.Fatalf("denominator: expected NodeSup, got %d", den.Type)
	}
}
