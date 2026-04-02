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

// HalvingSchedule draws 5 horizontal bars representing block reward epochs:
// 50, 25, 12.5, 6.25, and 3.125 BTC. Each bar is proportional in width.
// Ch12 diagram.
type HalvingSchedule struct{}

func (d *HalvingSchedule) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	usable := w - 2*pad
	barH := gtx.Dp(unit.Dp(18))

	// 5 epochs: subsidy halves each time.
	type epoch struct {
		label   string  // subsidy amount
		widthPc float64 // percentage of full width (50 BTC = 100%)
		alpha   uint8   // decreasing opacity
	}

	epochs := []epoch{
		{"50 BTC", 1.0, 255},
		{"25 BTC", 0.5, 210},
		{"12.5 BTC", 0.25, 170},
		{"6.25 BTC", 0.125, 130},
		{"3.125 BTC", 0.0625, 100},
	}

	totalBars := len(epochs)
	vertSpace := h - gtx.Dp(unit.Dp(8))
	barGap := (vertSpace - totalBars*barH) / (totalBars + 1)
	if barGap < 2 {
		barGap = 2
	}

	baseColor := th.Color.Primary

	for idx, ep := range epochs {
		y := barGap + idx*(barH+barGap)
		barW := int(float64(usable) * ep.widthPc)
		if barW < gtx.Dp(unit.Dp(20)) {
			barW = gtx.Dp(unit.Dp(20))
		}

		// Draw filled bar with decreasing alpha.
		barColor := withAlpha(baseColor, ep.alpha)
		func() {
			defer op.Offset(image.Pt(pad, y)).Push(gtx.Ops).Pop()
			r := float32(barH) / 4
			rr := clip.RRect{
				Rect: image.Rect(0, 0, barW, barH),
				SE:   int(r), SW: int(r), NE: int(r), NW: int(r),
			}
			paint.FillShape(gtx.Ops, barColor, rr.Op(gtx.Ops))
		}()

		// Epoch label on the left inside the bar (or next to it for narrow bars).
		epochLabel := fmt.Sprintf("%s %d", i18n.T("diagram.epoch"), idx+1)
		labelX := pad + barW + gtx.Dp(unit.Dp(6))
		labelY := y + (barH-gtx.Dp(unit.Dp(12)))/2
		caption(gtx, th, epochLabel, image.Pt(labelX, labelY))

		// Subsidy amount inside the bar.
		subsidyLabel := fmt.Sprintf("%s: %s", i18n.T("diagram.subsidy"), ep.label)
		subsidyX := pad + gtx.Dp(unit.Dp(4))
		colorCaption(gtx, th, subsidyLabel, image.Pt(subsidyX, labelY), th.Color.Text)
	}

	return layout.Dimensions{Size: image.Pt(w, h)}
}
