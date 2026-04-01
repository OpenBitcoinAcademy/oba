package math

import "image"

// Box holds the computed layout for a single AST node.
// Origin is the top-left corner. Bounds is relative to Origin.
// Baseline is the distance from the top of the box to the text baseline,
// used for vertical alignment of adjacent elements.
type Box struct {
	Node     *Node
	Origin   image.Point
	Bounds   image.Rectangle // (0,0) to (width, height)
	Baseline int             // distance from top to baseline
	Children []*Box
}

// Width returns the box width.
func (b *Box) Width() int { return b.Bounds.Dx() }

// Height returns the box height.
func (b *Box) Height() int { return b.Bounds.Dy() }

// LayoutConfig holds font metrics needed for layout calculations.
type LayoutConfig struct {
	FontSize    int     // base font size in pixels
	ScriptScale float64 // scale factor for super/subscripts (0.7)
	FracBarH    int     // fraction bar thickness in pixels
	FracGap     int     // gap between fraction bar and numerator/denominator
	SqrtPad     int     // padding inside sqrt radical
	DelimPad    int     // padding inside delimiters
	CharWidth   int     // average character width at base font size
	Ascent      int     // font ascent from baseline
	Descent     int     // font descent below baseline
}

// DefaultConfig returns a layout config for 14sp equivalent at ~2x density.
func DefaultConfig(fontSize int) LayoutConfig {
	charW := fontSize * 6 / 10 // approximate: 60% of font size
	return LayoutConfig{
		FontSize:    fontSize,
		ScriptScale: 0.7,
		FracBarH:    max(1, fontSize/14),
		FracGap:     max(2, fontSize/7),
		SqrtPad:     max(2, fontSize/7),
		DelimPad:    max(2, fontSize/10),
		CharWidth:   charW,
		Ascent:      fontSize * 8 / 10,
		Descent:     fontSize * 2 / 10,
	}
}

// Layout computes bounding boxes for the entire AST tree.
func Layout(node *Node, cfg LayoutConfig) *Box {
	return layoutNode(node, cfg, cfg.FontSize)
}

func layoutNode(node *Node, cfg LayoutConfig, fontSize int) *Box {
	switch node.Type {
	case NodeRow:
		return layoutRow(node, cfg, fontSize)
	case NodeText:
		return layoutText(node, cfg, fontSize)
	case NodeSymbol:
		return layoutSymbol(node, cfg, fontSize)
	case NodeTextCmd:
		return layoutTextCmd(node, cfg, fontSize)
	case NodeMod:
		return layoutMod(node, cfg, fontSize)
	case NodeGroup:
		return layoutRow(node, cfg, fontSize)
	case NodeSup:
		return layoutSup(node, cfg, fontSize)
	case NodeSub:
		return layoutSub(node, cfg, fontSize)
	case NodeSupSub:
		return layoutSupSub(node, cfg, fontSize)
	case NodeFrac:
		return layoutFrac(node, cfg, fontSize)
	case NodeSqrt:
		return layoutSqrt(node, cfg, fontSize)
	case NodeDelimited:
		return layoutDelimited(node, cfg, fontSize)
	default:
		return &Box{Node: node, Bounds: image.Rect(0, 0, 0, 0)}
	}
}

func layoutRow(node *Node, cfg LayoutConfig, fontSize int) *Box {
	var children []*Box
	for _, child := range node.Children {
		children = append(children, layoutNode(child, cfg, fontSize))
	}
	return arrangeHorizontal(node, children, fontSize, cfg)
}

// arrangeHorizontal places children left-to-right, aligned on baseline.
func arrangeHorizontal(node *Node, children []*Box, fontSize int, cfg LayoutConfig) *Box {
	if len(children) == 0 {
		ascent := fontSize * 8 / 10
		descent := fontSize * 2 / 10
		return &Box{
			Node:     node,
			Bounds:   image.Rect(0, 0, 0, ascent+descent),
			Baseline: ascent,
		}
	}

	// Find the maximum ascent (above baseline) and descent (below baseline).
	maxAscent := 0
	maxDescent := 0
	for _, ch := range children {
		asc := ch.Baseline
		desc := ch.Height() - ch.Baseline
		if asc > maxAscent {
			maxAscent = asc
		}
		if desc > maxDescent {
			maxDescent = desc
		}
	}

	// Position each child.
	x := 0
	for _, ch := range children {
		ch.Origin.X = x
		ch.Origin.Y = maxAscent - ch.Baseline
		x += ch.Width()
	}

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, x, maxAscent+maxDescent),
		Baseline: maxAscent,
		Children: children,
	}
}

func layoutText(node *Node, cfg LayoutConfig, fontSize int) *Box {
	w := textWidth(node.Value, fontSize, cfg)
	ascent := fontSize * 8 / 10
	descent := fontSize * 2 / 10
	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, w, ascent+descent),
		Baseline: ascent,
	}
}

func layoutSymbol(node *Node, cfg LayoutConfig, fontSize int) *Box {
	w := textWidth(node.Value, fontSize, cfg)
	// Operators get extra spacing.
	if _, ok := Operators[reverseSymbol(node.Value)]; ok {
		pad := fontSize / 5
		w += 2 * pad
	}
	ascent := fontSize * 8 / 10
	descent := fontSize * 2 / 10
	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, w, ascent+descent),
		Baseline: ascent,
	}
}

func layoutTextCmd(node *Node, cfg LayoutConfig, fontSize int) *Box {
	w := textWidth(node.Value, fontSize, cfg)
	ascent := fontSize * 8 / 10
	descent := fontSize * 2 / 10
	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, w, ascent+descent),
		Baseline: ascent,
	}
}

func layoutMod(node *Node, cfg LayoutConfig, fontSize int) *Box {
	// "mod" with spacing on both sides.
	w := textWidth("mod", fontSize, cfg) + fontSize/3
	ascent := fontSize * 8 / 10
	descent := fontSize * 2 / 10
	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, w, ascent+descent),
		Baseline: ascent,
	}
}

func layoutSup(node *Node, cfg LayoutConfig, fontSize int) *Box {
	base := layoutNode(node.Children[0], cfg, fontSize)
	scriptSize := int(float64(fontSize) * cfg.ScriptScale)
	sup := layoutNode(node.Children[1], cfg, scriptSize)

	// Superscript: bottom of sup aligns with ~60% up from base bottom.
	// A negative supShift moves the sup above the base origin.
	supShift := base.Baseline/2 - sup.Height()

	totalW := base.Width() + sup.Width()
	topExtent := 0
	if supShift < 0 {
		topExtent = -supShift
	}
	totalH := topExtent + base.Height()

	baseY := topExtent
	base.Origin = image.Pt(0, baseY)
	sup.Origin = image.Pt(base.Width(), baseY+supShift)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, totalW, totalH),
		Baseline: baseY + base.Baseline,
		Children: []*Box{base, sup},
	}
}

func layoutSub(node *Node, cfg LayoutConfig, fontSize int) *Box {
	base := layoutNode(node.Children[0], cfg, fontSize)
	scriptSize := int(float64(fontSize) * cfg.ScriptScale)
	sub := layoutNode(node.Children[1], cfg, scriptSize)

	// Subscript sits below base's baseline by a fraction of its height.
	subShift := base.Baseline - sub.Baseline + sub.Height()/3

	totalW := base.Width() + sub.Width()
	totalH := subShift + sub.Height()
	if base.Height() > totalH {
		totalH = base.Height()
	}

	base.Origin = image.Pt(0, 0)
	sub.Origin = image.Pt(base.Width(), subShift)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, totalW, totalH),
		Baseline: base.Baseline,
		Children: []*Box{base, sub},
	}
}

func layoutSupSub(node *Node, cfg LayoutConfig, fontSize int) *Box {
	base := layoutNode(node.Children[0], cfg, fontSize)
	scriptSize := int(float64(fontSize) * cfg.ScriptScale)
	sup := layoutNode(node.Children[1], cfg, scriptSize)
	sub := layoutNode(node.Children[2], cfg, scriptSize)

	supShift := base.Baseline/2 - sup.Height()
	subShift := base.Baseline - sub.Baseline + sub.Height()/3

	topExtent := 0
	if supShift < 0 {
		topExtent = -supShift
	}

	scriptW := sup.Width()
	if sub.Width() > scriptW {
		scriptW = sub.Width()
	}
	totalW := base.Width() + scriptW

	baseY := topExtent
	totalH := baseY + subShift + sub.Height()
	if baseY+base.Height() > totalH {
		totalH = baseY + base.Height()
	}

	base.Origin = image.Pt(0, baseY)
	sup.Origin = image.Pt(base.Width(), baseY+supShift)
	sub.Origin = image.Pt(base.Width(), baseY+subShift)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, totalW, totalH),
		Baseline: baseY + base.Baseline,
		Children: []*Box{base, sup, sub},
	}
}

func layoutFrac(node *Node, cfg LayoutConfig, fontSize int) *Box {
	num := layoutNode(node.Children[0], cfg, fontSize)
	den := layoutNode(node.Children[1], cfg, fontSize)

	w := num.Width()
	if den.Width() > w {
		w = den.Width()
	}
	w += cfg.FracGap * 2 // padding on sides

	// Center numerator and denominator.
	numX := (w - num.Width()) / 2
	denX := (w - den.Width()) / 2

	// Fraction bar at baseline level.
	barY := num.Height() + cfg.FracGap
	denY := barY + cfg.FracBarH + cfg.FracGap

	totalH := denY + den.Height()

	num.Origin = image.Pt(numX, 0)
	den.Origin = image.Pt(denX, denY)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, w, totalH),
		Baseline: barY + cfg.FracBarH/2,
		Children: []*Box{num, den},
	}
}

func layoutSqrt(node *Node, cfg LayoutConfig, fontSize int) *Box {
	content := layoutNode(node.Children[0], cfg, fontSize)

	// Radical sign width + padding.
	radW := fontSize / 2
	padTop := cfg.SqrtPad
	padRight := cfg.SqrtPad

	totalW := radW + content.Width() + padRight
	totalH := padTop + content.Height()

	content.Origin = image.Pt(radW, padTop)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, totalW, totalH),
		Baseline: padTop + content.Baseline,
		Children: []*Box{content},
	}
}

func layoutDelimited(node *Node, cfg LayoutConfig, fontSize int) *Box {
	content := layoutNode(node.Children[0], cfg, fontSize)

	delimW := fontSize / 2
	pad := cfg.DelimPad

	totalW := delimW + pad + content.Width() + pad + delimW
	totalH := content.Height()

	content.Origin = image.Pt(delimW+pad, 0)

	return &Box{
		Node:     node,
		Bounds:   image.Rect(0, 0, totalW, totalH),
		Baseline: content.Baseline,
		Children: []*Box{content},
	}
}

// textWidth estimates the pixel width of a string at a given font size.
func textWidth(s string, fontSize int, cfg LayoutConfig) int {
	charW := cfg.CharWidth * fontSize / cfg.FontSize
	n := 0
	for range s {
		n++
	}
	if n == 0 {
		return 0
	}
	return n * charW
}

// reverseSymbol finds the LaTeX command for a Unicode symbol string.
// Used to detect operators for spacing.
func reverseSymbol(s string) string {
	for cmd, r := range Operators {
		if string(r) == s {
			return cmd
		}
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
