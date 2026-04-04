package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TxWalkthrough draws the Alice-pays-Bob transaction flow for Ch02 L2.
// Alice's UTXO splits into payment to Bob and change back to Alice.
type TxWalkthrough struct{}

func (d *TxWalkthrough) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(26))

	// Row 0: Alice's UTXO (input).
	y0 := gtx.Dp(unit.Dp(8))
	inputW := pct(usable, 36)
	inputX := pad + (usable-inputW)/2
	shadowBox(gtx, th, "Alice: 0.015 BTC", image.Pt(inputX, y0), inputW, bh, th.Color.WarningBg)

	// Row 1: TX processing.
	y1 := y0 + bh + rowGap
	txW := pct(usable, 24)
	txX := pad + (usable-txW)/2
	shadowBox(gtx, th, "TX", image.Pt(txX, y1), txW, bh, th.Color.Primary)

	dirArrow(gtx, image.Pt(inputX+inputW/2, y0+bh), image.Pt(txX+txW/2, y1), lw, lc)

	// Row 2: Two outputs (Bob + change).
	y2 := y1 + bh + rowGap
	outW := pct(usable, 32)
	outGap := (usable - 2*outW) / 3
	bobX := pad + outGap
	changeX := bobX + outW + outGap

	shadowBox(gtx, th, "Bob: 0.01 BTC", image.Pt(bobX, y2), outW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.change")+": 0.0049", image.Pt(changeX, y2), outW, bh, th.Color.InfoBg)

	// Arrows from TX to outputs.
	curvedLine(gtx, image.Pt(txX+txW/3, y1+bh), image.Pt(bobX+outW/2, y2), lw, lc)
	curvedLine(gtx, image.Pt(txX+2*txW/3, y1+bh), image.Pt(changeX+outW/2, y2), lw, lc)

	// Fee caption.
	feeY := y2 + bh + gtx.Dp(unit.Dp(8))
	caption(gtx, th, i18n.T("diagram.fee")+": 0.0001 BTC", image.Pt(pad+(usable-pct(usable, 30))/2, feeY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
