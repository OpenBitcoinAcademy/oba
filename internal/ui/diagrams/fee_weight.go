package diagrams

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// FeeWeight compares legacy vs segwit transaction weight for Ch09.
// Two side-by-side columns: Legacy (400 WU) and Segwit (250 WU),
// each with a proportional bar to visualize the difference.
type FeeWeight struct{}

func (d *FeeWeight) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	colW := pct(w, 40)
	colGap := w - 2*pad - 2*colW
	if colGap < 8 {
		colGap = 8
	}

	bh := gtx.Dp(unit.Dp(34))
	barH := gtx.Dp(unit.Dp(20))
	barMaxW := colW

	leftX := pad
	rightX := pad + colW + colGap

	// Title row.
	titleY := gtx.Dp(unit.Dp(4))
	box(gtx, th, i18n.T("diagram.legacy_tx"), image.Pt(leftX, titleY), colW, bh, th.Color.WarningBg)
	box(gtx, th, i18n.T("diagram.segwit_tx"), image.Pt(rightX, titleY), colW, bh, th.Color.InfoBg)

	// Weight labels.
	weightY := titleY + bh + gtx.Dp(unit.Dp(10))
	colorCaption(gtx, th, i18n.T("diagram.weight")+": 400 WU", image.Pt(leftX, weightY), th.Color.Warning)
	colorCaption(gtx, th, i18n.T("diagram.weight")+": 250 WU", image.Pt(rightX, weightY), th.Color.Accent)

	// Proportional bars. Legacy = 400/400 = 100%, Segwit = 250/400 = 62.5%.
	barY := weightY + gtx.Dp(unit.Dp(20))
	legacyBarW := barMaxW
	segwitBarW := barMaxW * 250 / 400

	drawBar(gtx, th, image.Pt(leftX, barY), legacyBarW, barH, th.Color.Warning)
	drawBar(gtx, th, image.Pt(rightX, barY), segwitBarW, barH, th.Color.Accent)

	// Savings caption.
	savingsY := barY + barH + gtx.Dp(unit.Dp(10))
	caption(gtx, th, "-37.5%", image.Pt(rightX, savingsY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}

// drawBar draws a rounded, semi-transparent horizontal bar.
func drawBar(gtx layout.Context, _ *theme.Theme, pos image.Point, w, h int, c color.NRGBA) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	r := float32(h) / 4
	rr := clip.RRect{
		Rect: image.Rect(0, 0, w, h),
		SE:   int(r), SW: int(r), NE: int(r), NW: int(r),
	}
	// Fill with reduced alpha for a bar effect.
	fill := withAlpha(c, 180)
	paint.FillShape(gtx.Ops, fill, rr.Op(gtx.Ops))

	// Border.
	paint.FillShape(gtx.Ops, c, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())
}
