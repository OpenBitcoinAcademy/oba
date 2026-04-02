package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TxAnatomy shows the structure of a transaction: inputs on the left,
// outputs on the right, connected by a central TX box.
type TxAnatomy struct{}

func (d *TxAnatomy) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	colW := pct(w, 35)
	centerW := pct(w, 20)
	bh := gtx.Dp(unit.Dp(30))
	gap := (w - 2*pad - 2*colW - centerW) / 2

	leftX := pad
	centerX := leftX + colW + gap
	rightX := centerX + centerW + gap

	// Input column.
	y1 := gtx.Dp(unit.Dp(20))
	colorCaption(gtx, th, i18n.T("diagram.tx_input")+"s", image.Pt(leftX, y1), th.Color.TextMuted)

	inputY := y1 + gtx.Dp(unit.Dp(20))
	box(gtx, th, i18n.T("diagram.tx_prev_tx")+": abcd...", image.Pt(leftX, inputY), colW, bh, th.Color.WarningBg)
	box(gtx, th, i18n.T("diagram.tx_scriptsig"), image.Pt(leftX, inputY+bh+4), colW, bh, th.Color.InfoBg)

	// Center TX box.
	txY := inputY + bh/2
	processBox(gtx, th, "TX", image.Pt(centerX, txY), centerW, bh+gtx.Dp(unit.Dp(10)))

	// Arrows: input -> TX -> output.
	arrow(gtx, th, image.Pt(leftX+colW, inputY+bh/2), image.Pt(centerX, inputY+bh/2))
	arrow(gtx, th, image.Pt(centerX+centerW, txY+bh/2), image.Pt(rightX, txY-gtx.Dp(unit.Dp(5))+bh/2))

	// Output column.
	colorCaption(gtx, th, i18n.T("diagram.tx_output")+"s", image.Pt(rightX, y1), th.Color.TextMuted)

	outY1 := y1 + gtx.Dp(unit.Dp(20))
	box(gtx, th, "0.3 BTC", image.Pt(rightX, outY1), colW, bh, th.Color.TipBg)
	box(gtx, th, "0.7 BTC", image.Pt(rightX, outY1+bh+4), colW, bh, th.Color.TipBg)

	// Metadata at bottom.
	metaY := outY1 + 2*bh + gtx.Dp(unit.Dp(20))
	caption(gtx, th, i18n.T("diagram.tx_version")+": 1    "+i18n.T("diagram.tx_locktime")+": 0", image.Pt(pad, metaY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
