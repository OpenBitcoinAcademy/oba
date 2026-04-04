package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TaprootSpending draws key path vs script path comparison for Ch07 L4.
type TaprootSpending struct{}

func (d *TaprootSpending) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	colW := pct(usable, 42)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// P2TR output at top center.
	outW := pct(usable, 30)
	outX := pad + (usable-outW)/2
	y0 := gtx.Dp(unit.Dp(8))
	shadowBox(gtx, th, "P2TR Output", image.Pt(outX, y0), outW, bh, th.Color.Primary)

	// Two paths below.
	y1 := y0 + bh + gtx.Dp(unit.Dp(28))
	leftX := pad
	rightX := pad + usable - colW

	// Key Path (common, fast).
	colorCaption(gtx, th, i18n.T("diagram.key_path"), image.Pt(leftX, y1), th.Color.Primary)
	y1cap := y1 + gtx.Dp(unit.Dp(22))
	shadowBox(gtx, th, i18n.T("diagram.signature")+" (64B)", image.Pt(leftX, y1cap), colW, bh, th.Color.TipBg)
	curvedLine(gtx, image.Pt(outX+outW/4, y0+bh), image.Pt(leftX+colW/2, y1cap), lw, lc)

	// Script Path (fallback).
	colorCaption(gtx, th, i18n.T("diagram.script_path"), image.Pt(rightX, y1), th.Color.TextMuted)
	shadowBox(gtx, th, i18n.T("diagram.script_leaf")+" + proof", image.Pt(rightX, y1cap), colW, bh, th.Color.InfoBg)
	curvedLine(gtx, image.Pt(outX+3*outW/4, y0+bh), image.Pt(rightX+colW/2, y1cap), lw, withAlpha(lc, 80))

	// Captions.
	capY := y1cap + bh + gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, i18n.T("diagram.common_case"), image.Pt(leftX, capY), th.Color.Primary)
	caption(gtx, th, i18n.T("diagram.fallback"), image.Pt(rightX, capY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
