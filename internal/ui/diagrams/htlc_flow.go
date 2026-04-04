package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// HTLCFlow draws a Hash Time-Locked Contract payment chain for
// Ch18 L2. Alice pays Carol through Bob using hash preimage reveals.
type HTLCFlow struct{}

func (d *HTLCFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	nodeW := pct(usable, 22)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(30))

	// Row 0: Three participants.
	nodeGap := (usable - 3*nodeW) / 4
	y0 := gtx.Dp(unit.Dp(8))
	xA := pad + nodeGap
	xB := xA + nodeW + nodeGap
	xC := xB + nodeW + nodeGap

	shadowBox(gtx, th, "Alice", image.Pt(xA, y0), nodeW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "Bob", image.Pt(xB, y0), nodeW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "Carol", image.Pt(xC, y0), nodeW, bh, th.Color.InfoBg)

	// Row 1: HTLC offers (left to right).
	y1 := y0 + bh + rowGap
	htlcW := pct(usable, 30)
	htlc1X := pad + pct(usable, 5)
	htlc2X := pad + pct(usable, 50)

	shadowBox(gtx, th, "HTLC: H(r)", image.Pt(htlc1X, y1), htlcW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, "HTLC: H(r)", image.Pt(htlc2X, y1), htlcW, bh, th.Color.WarningBg)

	// Arrows: Alice→HTLC1, Bob→HTLC2.
	dirArrow(gtx, image.Pt(xA+nodeW/2, y0+bh), image.Pt(htlc1X+htlcW/4, y1), lw, lc)
	dirArrow(gtx, image.Pt(xB+nodeW/2, y0+bh), image.Pt(htlc1X+3*htlcW/4, y1), lw, lc)
	dirArrow(gtx, image.Pt(xB+nodeW/2, y0+bh), image.Pt(htlc2X+htlcW/4, y1), lw, lc)
	dirArrow(gtx, image.Pt(xC+nodeW/2, y0+bh), image.Pt(htlc2X+3*htlcW/4, y1), lw, lc)

	// Row 2: Preimage reveal (right to left).
	y2 := y1 + bh + rowGap
	revealW := pct(usable, 50)
	revealX := pad + (usable-revealW)/2

	shadowBox(gtx, th, i18n.T("diagram.preimage")+" r \u2190", image.Pt(revealX, y2), revealW, bh, th.Color.TipBg)

	// Dashed arrows showing preimage flowing back.
	dashedLine(gtx, image.Pt(xC+nodeW/2, y0+bh), image.Pt(revealX+revealW, y2+bh/2), lw, withAlpha(th.Color.Primary, 160))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
