package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MiniscriptPipeline draws a three-stage horizontal flow for Ch15:
// Policy -> Compile -> Miniscript -> Map -> Bitcoin Script.
// Below the flow, a concrete example shows the transformation in
// caption text.
type MiniscriptPipeline struct{}

func (d *MiniscriptPipeline) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	usable := w - 2*pad
	bh := gtx.Dp(unit.Dp(30))
	procW := pct(usable, 14)
	dataW := pct(usable, 20)

	// Layout: data(20) gap proc(14) gap data(20) gap proc(14) gap data(20)
	// Total boxes: 3*20 + 2*14 = 88%, leaving 12% for 4 gaps.
	gapTotal := usable - 3*dataW - 2*procW
	gap := gapTotal / 4
	if gap < 4 {
		gap = 4
	}

	// Vertical center for the main flow row.
	y := h/3 - bh/2

	x1 := pad
	x2 := x1 + dataW + gap
	x3 := x2 + procW + gap
	x4 := x3 + dataW + gap
	x5 := x4 + procW + gap

	// Stage 1: Policy.
	box(gtx, th, i18n.T("diagram.policy"), image.Pt(x1, y), dataW, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(x1+dataW, y+bh/2), image.Pt(x2, y+bh/2))

	// Stage 2: Compile (process).
	processBox(gtx, th, i18n.T("diagram.compile"), image.Pt(x2, y), procW, bh)
	arrow(gtx, th, image.Pt(x2+procW, y+bh/2), image.Pt(x3, y+bh/2))

	// Stage 3: Miniscript.
	box(gtx, th, i18n.T("diagram.miniscript"), image.Pt(x3, y), dataW, bh, th.Color.WarningBg)
	arrow(gtx, th, image.Pt(x3+dataW, y+bh/2), image.Pt(x4, y+bh/2))

	// Stage 4: Map to Script (process).
	processBox(gtx, th, i18n.T("diagram.map_to_script"), image.Pt(x4, y), procW, bh)
	arrow(gtx, th, image.Pt(x4+procW, y+bh/2), image.Pt(x5, y+bh/2))

	// Stage 5: Bitcoin Script.
	box(gtx, th, i18n.T("diagram.bitcoin_script"), image.Pt(x5, y), dataW, bh, th.Color.TipBg)

	// Example line below, in caption text.
	exY := y + bh + gtx.Dp(unit.Dp(14))
	exText := "and(pk(A),older(1000)) \u2192 and_v(v:pk(A),older(1000)) \u2192 <A> CHECKSIGVERIFY <1000> CSV"
	caption(gtx, th, exText, image.Pt(pad, exY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
