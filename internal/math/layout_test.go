package math

import "testing"

func cfg() LayoutConfig {
	return DefaultConfig(28) // 14sp at 2x
}

func TestLayoutSimpleText(t *testing.T) {
	node, _ := Parse("x")
	box := Layout(node, cfg())

	if box.Width() <= 0 {
		t.Errorf("width = %d, want > 0", box.Width())
	}
	if box.Height() <= 0 {
		t.Errorf("height = %d, want > 0", box.Height())
	}
	if box.Baseline <= 0 {
		t.Errorf("baseline = %d, want > 0", box.Baseline)
	}
}

func TestLayoutSuperscriptTallerThanBase(t *testing.T) {
	node, _ := Parse("x^2")
	box := Layout(node, cfg())

	// Superscript makes the overall height >= base height.
	baseNode, _ := Parse("x")
	baseBox := Layout(baseNode, cfg())
	if box.Height() < baseBox.Height() {
		t.Errorf("sup height %d < base height %d", box.Height(), baseBox.Height())
	}
}

func TestLayoutSubscriptExtendsBelowBaseline(t *testing.T) {
	node, _ := Parse("k_{priv}")
	box := Layout(node, cfg())

	// Subscript should extend below the base's baseline.
	baseNode, _ := Parse("k")
	baseBox := Layout(baseNode, cfg())
	if box.Height() <= baseBox.Height() {
		t.Errorf("sub height %d should be > base height %d", box.Height(), baseBox.Height())
	}
}

func TestLayoutFracWiderThanSingleChar(t *testing.T) {
	node, _ := Parse(`\frac{a}{b}`)
	box := Layout(node, cfg())

	charNode, _ := Parse("a")
	charBox := Layout(charNode, cfg())

	// Fraction should be wider than a single character (due to padding).
	if box.Width() <= charBox.Width() {
		t.Errorf("frac width %d should be > char width %d", box.Width(), charBox.Width())
	}
}

func TestLayoutFracBaselineAtBar(t *testing.T) {
	node, _ := Parse(`\frac{a}{b}`)
	c := cfg()
	box := Layout(node, c)

	// Baseline should be between top and bottom, near the bar.
	if box.Baseline <= 0 || box.Baseline >= box.Height() {
		t.Errorf("baseline = %d, height = %d: baseline should be between 0 and height", box.Baseline, box.Height())
	}
}

func TestLayoutSqrtWiderThanContent(t *testing.T) {
	node, _ := Parse(`\sqrt{x}`)
	box := Layout(node, cfg())

	contentNode, _ := Parse("x")
	contentBox := Layout(contentNode, cfg())

	if box.Width() <= contentBox.Width() {
		t.Errorf("sqrt width %d should be > content width %d", box.Width(), contentBox.Width())
	}
}

func TestLayoutRowWidthIsSum(t *testing.T) {
	// "ab" should be roughly twice as wide as "a".
	aNode, _ := Parse("a")
	aBox := Layout(aNode, cfg())

	abNode, _ := Parse("ab")
	abBox := Layout(abNode, cfg())

	// Allow some tolerance for spacing.
	expected := aBox.Width() * 2
	if abBox.Width() < expected-2 || abBox.Width() > expected+2 {
		t.Errorf("row width = %d, want ~%d (2x single char)", abBox.Width(), expected)
	}
}

func TestLayoutEllipticCurve(t *testing.T) {
	node, _ := Parse("y^2 = x^3 + 7")
	box := Layout(node, cfg())
	if box.Width() <= 0 || box.Height() <= 0 {
		t.Errorf("formula has zero dimensions: %dx%d", box.Width(), box.Height())
	}
}

func TestLayoutKeyDerivation(t *testing.T) {
	node, _ := Parse(`k_{priv} \cdot G = K_{pub}`)
	box := Layout(node, cfg())
	if box.Width() <= 0 || box.Height() <= 0 {
		t.Errorf("formula has zero dimensions: %dx%d", box.Width(), box.Height())
	}
}

func TestLayoutHalving(t *testing.T) {
	node, _ := Parse(`\frac{reward}{2^n}`)
	box := Layout(node, cfg())
	if box.Width() <= 0 || box.Height() <= 0 {
		t.Errorf("formula has zero dimensions: %dx%d", box.Width(), box.Height())
	}
}

func TestLayoutEmptyRow(t *testing.T) {
	node := &Node{Type: NodeRow}
	box := Layout(node, cfg())
	// Empty row should have zero width but nonzero height (line height).
	if box.Width() != 0 {
		t.Errorf("empty row width = %d, want 0", box.Width())
	}
	if box.Height() <= 0 {
		t.Errorf("empty row height = %d, want > 0", box.Height())
	}
}

func TestLayoutDelimitedWiderThanContent(t *testing.T) {
	node, _ := Parse(`\left( x \right)`)
	box := Layout(node, cfg())

	contentNode, _ := Parse("x")
	contentBox := Layout(contentNode, cfg())

	if box.Width() <= contentBox.Width() {
		t.Errorf("delimited width %d should be > content width %d", box.Width(), contentBox.Width())
	}
}

func TestLayoutSupSubBothScripts(t *testing.T) {
	node, _ := Parse("x^2_3")
	box := Layout(node, cfg())

	supNode, _ := Parse("x^2")
	supBox := Layout(supNode, cfg())

	// Both scripts should be at least as tall as superscript alone.
	if box.Height() < supBox.Height() {
		t.Errorf("supsub height %d < sup height %d", box.Height(), supBox.Height())
	}
}

func TestDefaultConfigValues(t *testing.T) {
	c := DefaultConfig(28)
	if c.FontSize != 28 {
		t.Errorf("FontSize = %d, want 28", c.FontSize)
	}
	if c.ScriptScale != 0.7 {
		t.Errorf("ScriptScale = %f, want 0.7", c.ScriptScale)
	}
	if c.CharWidth <= 0 {
		t.Errorf("CharWidth = %d, want > 0", c.CharWidth)
	}
	if c.Ascent <= 0 || c.Descent <= 0 {
		t.Errorf("Ascent=%d Descent=%d: both should be > 0", c.Ascent, c.Descent)
	}
}
