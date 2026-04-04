package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TweakedKey draws the Taproot key tweaking formula for Ch14 L1.
// P + H(P||m)·G = Q. Internal key + tweak = output key.
type TweakedKey struct{}

func (d *TweakedKey) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(28))

	// Row 0: Internal Key P and Script Tree (m).
	inputW := pct(usable, 30)
	inputGap := (usable - 2*inputW) / 3
	y0 := gtx.Dp(unit.Dp(8))
	xP := pad + inputGap
	xM := xP + inputW + inputGap

	shadowBox(gtx, th, i18n.T("diagram.internal_key")+" P", image.Pt(xP, y0), inputW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, i18n.T("diagram.script_tree")+" m", image.Pt(xM, y0), inputW, bh, th.Color.InfoBg)

	// Row 1: Hash and tweak computation.
	y1 := y0 + bh + rowGap
	tweakW := pct(usable, 40)
	tweakX := pad + (usable-tweakW)/2
	shadowBox(gtx, th, "H(P \u2016 m) \u00b7 G", image.Pt(tweakX, y1), tweakW, bh, th.Color.Primary)

	// Arrows from inputs to tweak.
	curvedLine(gtx, image.Pt(xP+inputW/2, y0+bh), image.Pt(tweakX+tweakW/3, y1), lw, lc)
	curvedLine(gtx, image.Pt(xM+inputW/2, y0+bh), image.Pt(tweakX+2*tweakW/3, y1), lw, lc)

	// Row 2: Output Key Q = P + tweak.
	y2 := y1 + bh + rowGap
	outW := pct(usable, 36)
	outX := pad + (usable-outW)/2
	shadowBox(gtx, th, "Q = P + t\u00b7G", image.Pt(outX, y2), outW, bh, th.Color.TipBg)

	// Arrows: P and tweak both feed into Q.
	curvedLine(gtx, image.Pt(xP+inputW/2, y0+bh), image.Pt(outX+outW/4, y2), lw, withAlpha(lc, 100))
	dirArrow(gtx, image.Pt(tweakX+tweakW/2, y1+bh), image.Pt(outX+outW/2, y2), lw, lc)

	// Caption.
	captionY := y2 + bh + gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, i18n.T("diagram.tweaked_key"), image.Pt(outX, captionY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
