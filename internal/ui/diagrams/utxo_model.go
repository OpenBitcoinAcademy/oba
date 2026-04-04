package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// UTXOModel shows how transactions consume and create UTXOs.
// TX A has two outputs (one spent, one unspent). TX B spends TX A's output 0.
type UTXOModel struct{}

func (d *UTXOModel) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(215))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bw := pct(w, 28)
	bh := gtx.Dp(unit.Dp(28))
	outW := pct(w, 22)
	gap := pct(w, 5)

	// Row 1: TX A with outputs.
	y1 := gtx.Dp(unit.Dp(10))
	colorCaption(gtx, th, i18n.T("diagram.utxo_tx_a"), image.Pt(pad, y1), th.Color.TextMuted)
	y1 += gtx.Dp(unit.Dp(24))

	txAx := pad
	shadowBox(gtx, th, "TX A", image.Pt(txAx, y1), bw, bh, th.Color.Primary)

	outAx := txAx + bw + gap/2
	// Output 0: spent (dimmed).
	shadowBox(gtx, th, "0.5 BTC", image.Pt(outAx, y1), outW, bh, th.Color.Divider)
	// Output 1: unspent (green).
	shadowBox(gtx, th, "0.3 BTC", image.Pt(outAx+outW+4, y1), outW, bh, th.Color.TipBg)

	dirArrow(gtx, image.Pt(txAx+bw, y1+bh/2), image.Pt(outAx, y1+bh/2), 1.8, withAlpha(th.Color.TextMuted, 160))

	// Labels.
	labelY := y1 + bh + 2
	caption(gtx, th, i18n.T("diagram.utxo_spent"), image.Pt(outAx+outW/4, labelY))
	caption(gtx, th, i18n.T("diagram.utxo_unspent"), image.Pt(outAx+outW+4+outW/6, labelY))

	// Connecting line from TX A output 0 to TX B input.
	midY := y1 + bh + gtx.Dp(unit.Dp(25))
	dashedLine(gtx, image.Pt(outAx+outW/2, y1+bh), image.Pt(outAx+outW/2, midY), 1.8, withAlpha(th.Color.TextMuted, 120))

	// Row 2: TX B with outputs.
	y2 := midY + gtx.Dp(unit.Dp(10))
	colorCaption(gtx, th, i18n.T("diagram.utxo_tx_b"), image.Pt(pad, y2), th.Color.TextMuted)
	y2 += gtx.Dp(unit.Dp(24))

	txBx := pad
	shadowBox(gtx, th, "TX B", image.Pt(txBx, y2), bw, bh, th.Color.Primary)

	outBx := txBx + bw + gap/2
	shadowBox(gtx, th, "0.3 BTC", image.Pt(outBx, y2), outW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "0.19 BTC", image.Pt(outBx+outW+4, y2), outW, bh, th.Color.TipBg)

	dirArrow(gtx, image.Pt(txBx+bw, y2+bh/2), image.Pt(outBx, y2+bh/2), 1.8, withAlpha(th.Color.TextMuted, 160))

	// Connecting line into TX B.
	curvedLine(gtx, image.Pt(outAx+outW/2, midY), image.Pt(txBx+bw/2, y2), 1.8, withAlpha(th.Color.TextMuted, 160))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
