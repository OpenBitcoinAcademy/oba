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

// SchnorrVsECDSA draws a side-by-side comparison for Ch08 L3.
// Left: ECDSA (r,s + DER, variable length).
// Right: Schnorr (64 bytes, fixed, aggregatable).
type SchnorrVsECDSA struct{}

func (d *SchnorrVsECDSA) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	colW := pct(usable, 44)
	divX := pad + colW + pct(usable, 4)
	rightX := divX + pct(usable, 4)
	spacing := 5

	// Divider line.
	paint.FillShape(gtx.Ops, th.Color.Divider,
		clip.Rect{Min: image.Pt(divX, gtx.Dp(unit.Dp(4))), Max: image.Pt(divX+1, h-gtx.Dp(unit.Dp(4)))}.Op())

	// Left column: ECDSA.
	y := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "ECDSA", image.Pt(pad, y), th.Color.Warning)
	y += gtx.Dp(unit.Dp(24))

	shadowBox(gtx, th, "(r, s)", image.Pt(pad, y), colW, bh, th.Color.WarningBg)
	y += bh + spacing
	shadowBox(gtx, th, "DER "+i18n.T("diagram.encode_address"), image.Pt(pad, y), colW, bh, th.Color.InfoBg)
	y += bh + spacing
	shadowBox(gtx, th, "71\u201373 "+i18n.T("diagram.bytes"), image.Pt(pad, y), colW, bh, th.Color.InfoBg)

	// Right column: Schnorr.
	y = gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "Schnorr", image.Pt(rightX, y), th.Color.Primary)
	y += gtx.Dp(unit.Dp(24))

	shadowBox(gtx, th, "(R, s)", image.Pt(rightX, y), colW, bh, th.Color.TipBg)
	y += bh + spacing
	shadowBox(gtx, th, i18n.T("diagram.fixed_length"), image.Pt(rightX, y), colW, bh, th.Color.TipBg)
	y += bh + spacing
	shadowBox(gtx, th, "64 "+i18n.T("diagram.bytes"), image.Pt(rightX, y), colW, bh, th.Color.Primary)

	// "Aggregatable" badge on Schnorr side.
	y += bh + gtx.Dp(unit.Dp(8))
	func() {
		defer op.Offset(image.Pt(rightX, y)).Push(gtx.Ops).Pop()
		colorCaption(gtx, th, "\u2295 "+i18n.T("diagram.aggregatable"), image.Pt(0, 0), th.Color.Primary)
	}()

	return layout.Dimensions{Size: image.Pt(w, h)}
}
