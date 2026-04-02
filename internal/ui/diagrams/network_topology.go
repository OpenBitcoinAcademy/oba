package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// NetworkTopology draws a P2P network with full nodes in a mesh
// and SPV clients connected to edge nodes. Ch10 diagram.
type NetworkTopology struct{}

func (d *NetworkTopology) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	fullR := 10 // full node radius
	spvR := 5   // SPV client radius
	fullColor := th.Color.Primary
	spvColor := th.Color.Accent
	lineColor := withAlpha(th.Color.TextMuted, 80)
	lw := float32(1.2)

	// Place 7 full nodes in a rough mesh layout.
	cx := w / 2
	cy := h/2 + gtx.Dp(unit.Dp(4))

	spreadX := pct(w, 18)
	spreadY := pct(h, 22)

	// Full nodes: two rows with offset for a mesh feel.
	fullNodes := []image.Point{
		{cx - spreadX, cy - spreadY},   // 0: top-left
		{cx, cy - spreadY - spreadY/3}, // 1: top-center
		{cx + spreadX, cy - spreadY},   // 2: top-right
		{cx - spreadX - spreadX/3, cy}, // 3: mid-left
		{cx + spreadX + spreadX/3, cy}, // 4: mid-right
		{cx - spreadX/2, cy + spreadY}, // 5: bottom-left
		{cx + spreadX/2, cy + spreadY}, // 6: bottom-right
	}

	// Mesh connections between full nodes (not all-to-all, but dense).
	meshEdges := [][2]int{
		{0, 1}, {1, 2}, {0, 3}, {2, 4},
		{3, 5}, {4, 6}, {5, 6},
		{0, 5}, {1, 5}, {1, 6}, {2, 6},
		{3, 1}, {4, 1},
	}

	// Draw mesh lines first (behind nodes).
	for _, e := range meshEdges {
		line(gtx, fullNodes[e[0]], fullNodes[e[1]], lw, lineColor)
	}

	// SPV clients at the edges, each connected to 1-2 full nodes.
	type spvNode struct {
		pos   image.Point
		peers []int // indices into fullNodes
	}

	spvSpread := pct(w, 8)
	spvNodes := []spvNode{
		{image.Pt(fullNodes[3].X-spvSpread, fullNodes[3].Y-spvSpread/2), []int{3}},
		{image.Pt(fullNodes[0].X-spvSpread/2, fullNodes[0].Y-spvSpread), []int{0}},
		{image.Pt(fullNodes[4].X+spvSpread, fullNodes[4].Y-spvSpread/2), []int{4}},
		{image.Pt(fullNodes[2].X+spvSpread/2, fullNodes[2].Y-spvSpread), []int{2, 4}},
	}

	// Draw SPV connection lines.
	spvLineColor := withAlpha(spvColor, 80)
	for _, s := range spvNodes {
		for _, pi := range s.peers {
			line(gtx, s.pos, fullNodes[pi], lw, spvLineColor)
		}
	}

	// Draw full nodes on top.
	for _, fn := range fullNodes {
		circle(gtx, fn, fullR, fullColor)
	}

	// Draw SPV clients on top.
	for _, s := range spvNodes {
		circle(gtx, s.pos, spvR, spvColor)
	}

	// Captions at bottom.
	captionY := h - gtx.Dp(unit.Dp(18))
	colorCaption(gtx, th, i18n.T("diagram.full_node"), image.Pt(pad, captionY), fullColor)
	colorCaption(gtx, th, i18n.T("diagram.spv_client"), image.Pt(w/2, captionY), spvColor)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
