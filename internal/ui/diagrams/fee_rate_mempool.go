package diagrams

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// FeeRateMempool draws a mempool fee priority visualization for
// Ch09 L2. Stacked bars sorted by fee rate, miner picks the top.
type FeeRateMempool struct{}

func (d *FeeRateMempool) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	barH := gtx.Dp(unit.Dp(22))
	spacing := 3

	type txBar struct {
		rate  int
		label string
		pct   float32 // width as fraction of max
	}

	bars := []txBar{
		{50, "TX A", 1.0},
		{20, "TX B", 0.7},
		{10, "TX C", 0.5},
		{5, "TX D", 0.35},
		{1, "TX E", 0.2},
	}

	// Title.
	y := gtx.Dp(unit.Dp(4))
	colorCaption(gtx, th, i18n.T("diagram.mempool"), image.Pt(pad, y), th.Color.TextMuted)
	y += gtx.Dp(unit.Dp(24))

	maxBarW := pct(usable, 65)
	labelX := pad + maxBarW + pct(usable, 3)

	for i, b := range bars {
		barW := int(float32(maxBarW) * b.pct)
		if barW < gtx.Dp(unit.Dp(20)) {
			barW = gtx.Dp(unit.Dp(20))
		}

		// Color: top bars are hotter.
		alpha := uint8(255 - i*35)
		bg := withAlpha(th.Color.Primary, alpha)

		// Bar with rounded corners.
		func() {
			defer op.Offset(image.Pt(pad, y)).Push(gtx.Ops).Pop()
			r := float32(barH) / 4
			ri := int(r)
			// Shadow.
			rr := clip.RRect{Rect: image.Rect(3, 3, barW+3, barH+3), SE: ri, SW: ri, NE: ri, NW: ri}
			paint.FillShape(gtx.Ops, withAlpha(th.Color.Text, 30), rr.Op(gtx.Ops))
			// Bar.
			rr2 := clip.RRect{Rect: image.Rect(0, 0, barW, barH), SE: ri, SW: ri, NE: ri, NW: ri}
			paint.FillShape(gtx.Ops, bg, rr2.Op(gtx.Ops))
		}()

		// Rate label.
		rateLabel := fmt.Sprintf("%s: %d sat/vB", b.label, b.rate)
		caption(gtx, th, rateLabel, image.Pt(labelX, y+(barH-gtx.Dp(unit.Dp(12)))/2))

		y += barH + spacing
	}

	// "Miner picks top" indicator.
	arrowY := gtx.Dp(unit.Dp(28)) + barH/2
	colorCaption(gtx, th, i18n.T("diagram.miner")+" \u2191", image.Pt(w-pad-pct(usable, 15), arrowY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
