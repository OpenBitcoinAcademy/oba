package diagrams

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/openbitcoinacademy/oba/internal/i18n"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// CentralizedVsP2P compares centralized payment (Alice -> Bank -> Bob)
// with peer-to-peer (Alice and Bob both connect to a shared network).
// Row 2 shows a mesh topology, not a linear pipeline.
type CentralizedVsP2P struct{}

func (d *CentralizedVsP2P) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	totalH := gtx.Dp(unit.Dp(300))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	// No background fill: diagram sits on the page background naturally.

	rowH := totalH / 2

	// === Row 1: Centralized ===
	d.layoutCentralized(gtx, th, w, pad, rowH)

	// Divider.
	paint.FillShape(gtx.Ops, th.Color.Divider,
		clip.Rect{Min: image.Pt(pad, rowH), Max: image.Pt(w-pad, rowH+1)}.Op())

	// === Row 2: P2P ===
	d.layoutP2P(gtx, th, w, pad, rowH, totalH)

	return layout.Dimensions{Size: image.Pt(w, totalH)}
}

func (d *CentralizedVsP2P) layoutCentralized(gtx layout.Context, th *theme.Theme, w, pad, rowH int) {
	bw := pct(w, 22)
	bh := gtx.Dp(unit.Dp(34))
	usable := w - 2*pad
	gap := (usable - 3*bw) / 4

	y := rowH/2 - bh/2 + gtx.Dp(unit.Dp(10))
	x1 := pad + gap
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap

	colorCaption(gtx, th, i18n.T("diagram.centralized"), image.Pt(pad, y-gtx.Dp(unit.Dp(20))), th.Color.Warning)

	box(gtx, th, i18n.T("diagram.alice"), image.Pt(x1, y), bw, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(x1+bw, y+bh/2), image.Pt(x2, y+bh/2))
	box(gtx, th, i18n.T("diagram.bank"), image.Pt(x2, y), bw, bh, th.Color.WarningBg)
	arrow(gtx, th, image.Pt(x2+bw, y+bh/2), image.Pt(x3, y+bh/2))
	box(gtx, th, i18n.T("diagram.bob"), image.Pt(x3, y), bw, bh, th.Color.InfoBg)
}

func (d *CentralizedVsP2P) layoutP2P(gtx layout.Context, th *theme.Theme, w, pad, rowH, totalH int) {
	bw := pct(w, 22)
	bh := gtx.Dp(unit.Dp(34))

	// Center of the P2P row area.
	centerY := rowH + rowH/2
	centerX := w / 2
	captionH := gtx.Dp(unit.Dp(20))

	colorCaption(gtx, th, i18n.T("diagram.peer_to_peer"), image.Pt(pad, rowH+gtx.Dp(unit.Dp(8))), th.Color.Success)

	// Alice on left, Bob on right, vertically centered in row.
	aliceX := pad + pct(w, 3)
	bobX := w - pad - pct(w, 3) - bw
	boxY := centerY - bh/2 + captionH/2

	box(gtx, th, i18n.T("diagram.alice"), image.Pt(aliceX, boxY), bw, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.bob"), image.Pt(bobX, boxY), bw, bh, th.Color.InfoBg)

	// Box midpoints for connection lines.
	boxMidY := boxY + bh/2
	aliceEdge := image.Pt(aliceX+bw, boxMidY)
	bobEdge := image.Pt(bobX, boxMidY)

	// Network nodes: 4 nodes in a diamond pattern centered between Alice and Bob.
	nodeColor := th.Color.Success
	nodeR := 6
	hSpread := pct(w, 8)
	vSpread := bh / 2

	n1 := image.Pt(centerX, boxMidY-vSpread-4) // top
	n2 := image.Pt(centerX-hSpread, boxMidY)   // left
	n3 := image.Pt(centerX+hSpread, boxMidY)   // right
	n4 := image.Pt(centerX, boxMidY+vSpread+4) // bottom

	circle(gtx, n1, nodeR, nodeColor)
	circle(gtx, n2, nodeR, nodeColor)
	circle(gtx, n3, nodeR, nodeColor)
	circle(gtx, n4, nodeR, nodeColor)

	// Mesh lines between all nodes.
	lc := withAlpha(nodeColor, 100)
	lw := float32(1.5)
	line(gtx, n1, n2, lw, lc)
	line(gtx, n1, n3, lw, lc)
	line(gtx, n2, n4, lw, lc)
	line(gtx, n3, n4, lw, lc)
	line(gtx, n2, n3, lw, lc)
	line(gtx, n1, n4, lw, lc)

	// Lines from Alice to left nodes, Bob to right nodes.
	line(gtx, aliceEdge, n1, lw, lc)
	line(gtx, aliceEdge, n2, lw, lc)
	line(gtx, aliceEdge, n4, lw, lc)
	line(gtx, bobEdge, n1, lw, lc)
	line(gtx, bobEdge, n3, lw, lc)
	line(gtx, bobEdge, n4, lw, lc)
}

func withAlpha(c color.NRGBA, a uint8) color.NRGBA {
	c.A = a
	return c
}
