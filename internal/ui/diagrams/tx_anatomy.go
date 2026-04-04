package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TxAnatomy shows the structure of a transaction: inputs on the left,
// outputs on the right, connected by a central TX box. Ch06 diagram.
type TxAnatomy struct{}

func (d *TxAnatomy) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(210))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	colW := pct(w, 30)
	centerW := pct(w, 18)
	bh := gtx.Dp(unit.Dp(28))
	usable := w - 2*pad
	gap := (usable - 2*colW - centerW) / 2
	if gap < 8 {
		gap = 8
	}

	leftX := pad
	centerX := leftX + colW + gap
	rightX := centerX + centerW + gap

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Column headers.
	headerY := gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, i18n.T("diagram.tx_input")+"s", image.Pt(leftX, headerY), th.Color.TextMuted)
	colorCaption(gtx, th, i18n.T("diagram.tx_output")+"s", image.Pt(rightX, headerY), th.Color.TextMuted)

	// Input boxes (left column).
	inputY := headerY + gtx.Dp(unit.Dp(26))
	inputSpacing := bh + 6
	shadowBox(gtx, th, i18n.T("diagram.tx_prev_tx")+": abcd...", image.Pt(leftX, inputY), colW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, i18n.T("diagram.tx_scriptsig"), image.Pt(leftX, inputY+inputSpacing), colW, bh, th.Color.InfoBg)

	// Center TX box (vertically centered between inputs).
	txH := bh + gtx.Dp(unit.Dp(12))
	txY := inputY + inputSpacing/2 - txH/4
	shadowBox(gtx, th, "TX", image.Pt(centerX, txY), centerW, txH, th.Color.Primary)

	// Output boxes (right column).
	outY := inputY
	shadowBox(gtx, th, "0.3 BTC", image.Pt(rightX, outY), colW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "0.7 BTC", image.Pt(rightX, outY+inputSpacing), colW, bh, th.Color.TipBg)

	// Arrows: inputs -> TX.
	txMidY := txY + txH/2
	dirArrow(gtx, image.Pt(leftX+colW, inputY+bh/2), image.Pt(centerX, txMidY-4), lw, lc)
	dirArrow(gtx, image.Pt(leftX+colW, inputY+inputSpacing+bh/2), image.Pt(centerX, txMidY+4), lw, lc)

	// Arrows: TX -> outputs.
	dirArrow(gtx, image.Pt(centerX+centerW, txMidY-4), image.Pt(rightX, outY+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(centerX+centerW, txMidY+4), image.Pt(rightX, outY+inputSpacing+bh/2), lw, lc)

	// Metadata at bottom.
	metaY := outY + 2*inputSpacing + gtx.Dp(unit.Dp(14))
	caption(gtx, th, i18n.T("diagram.tx_version")+": 1    "+i18n.T("diagram.tx_locktime")+": 0", image.Pt(pad, metaY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
