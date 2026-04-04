package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SegwitLayout draws the SegWit transaction structure for Ch07 L3.
// Shows the witness data separated from the transaction, with the
// weight discount visualized.
type SegwitLayout struct{}

func (d *SegwitLayout) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(150))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	rowGap := gtx.Dp(unit.Dp(20))

	// Row 0: Label.
	y0 := gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, "SegWit TX", image.Pt(pad, y0), th.Color.Primary)

	// Row 1: Transaction fields (full weight).
	y1 := y0 + gtx.Dp(unit.Dp(26))
	fieldW := pct(usable, 22)
	fieldGap := (usable - 4*fieldW) / 5

	x1 := pad + fieldGap
	x2 := x1 + fieldW + fieldGap
	x3 := x2 + fieldW + fieldGap
	x4 := x3 + fieldW + fieldGap

	shadowBox(gtx, th, i18n.T("diagram.version"), image.Pt(x1, y1), fieldW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, i18n.T("diagram.tx_input")+"s", image.Pt(x2, y1), fieldW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, i18n.T("diagram.tx_output")+"s", image.Pt(x3, y1), fieldW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.tx_locktime"), image.Pt(x4, y1), fieldW, bh, th.Color.InfoBg)

	// Divider.
	divY := y1 + bh + rowGap/2
	paint.FillShape(gtx.Ops, th.Color.Divider,
		clip.Rect{Min: image.Pt(pad, divY), Max: image.Pt(w-pad, divY+1)}.Op())

	// Row 2: Witness data (discounted weight).
	y2 := divY + rowGap/2
	witW := pct(usable, 50)
	witX := pad + (usable-witW)/2
	shadowBox(gtx, th, i18n.T("diagram.witness_data"), image.Pt(witX, y2), witW, bh, th.Color.Primary)

	// Discount label.
	discY := y2 + bh + gtx.Dp(unit.Dp(8))
	func() {
		defer op.Offset(image.Pt(witX, discY)).Push(gtx.Ops).Pop()
		colorCaption(gtx, th, "\u00bc "+i18n.T("diagram.weight"), image.Pt(0, 0), th.Color.Primary)
	}()

	return layout.Dimensions{Size: image.Pt(w, h)}
}
