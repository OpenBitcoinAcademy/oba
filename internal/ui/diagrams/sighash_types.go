package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SighashTypes draws which parts of a transaction each SIGHASH type
// commits to for Ch08 L4. Three rows showing ALL, NONE, SINGLE.
type SighashTypes struct{}

func (d *SighashTypes) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(24))
	labelW := pct(usable, 28)
	inW := pct(usable, 28)
	outW := pct(usable, 28)
	gap := (usable - labelW - inW - outW) / 4
	rowGap := 6

	xLabel := pad + gap
	xIn := xLabel + labelW + gap
	xOut := xIn + inW + gap

	// Header row.
	y := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "SIGHASH", image.Pt(xLabel, y), th.Color.TextMuted)
	colorCaption(gtx, th, "Inputs", image.Pt(xIn, y), th.Color.TextMuted)
	colorCaption(gtx, th, "Outputs", image.Pt(xOut, y), th.Color.TextMuted)
	y += gtx.Dp(unit.Dp(22))

	// ALL: signs all inputs and outputs.
	shadowBox(gtx, th, "ALL", image.Pt(xLabel, y), labelW, bh, th.Color.Primary)
	shadowBox(gtx, th, "\u2713 All", image.Pt(xIn, y), inW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "\u2713 All", image.Pt(xOut, y), outW, bh, th.Color.TipBg)
	y += bh + rowGap

	// NONE: signs inputs, no outputs.
	shadowBox(gtx, th, "NONE", image.Pt(xLabel, y), labelW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, "\u2713 All", image.Pt(xIn, y), inW, bh, th.Color.TipBg)
	box(gtx, th, "\u2717 None", image.Pt(xOut, y), outW, bh, th.Color.InfoBg)
	y += bh + rowGap

	// SINGLE: signs inputs, matching output only.
	shadowBox(gtx, th, "SINGLE", image.Pt(xLabel, y), labelW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "\u2713 All", image.Pt(xIn, y), inW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "\u2713 One", image.Pt(xOut, y), outW, bh, th.Color.WarningBg)
	y += bh + rowGap

	// ANYONECANPAY modifier.
	shadowBox(gtx, th, "| ANYONECANPAY", image.Pt(xLabel, y), labelW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "\u2713 One", image.Pt(xIn, y), inW, bh, th.Color.WarningBg)
	caption(gtx, th, "(varies)", image.Pt(xOut+outW/4, y+(bh-gtx.Dp(unit.Dp(12)))/2))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
