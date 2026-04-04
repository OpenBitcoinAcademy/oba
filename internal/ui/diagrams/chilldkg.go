package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ChillDKG draws distributed key generation vs trusted dealer for Ch17 L3.
type ChillDKG struct{}

func (d *ChillDKG) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	colW := pct(usable, 42)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Left: Trusted Dealer (bad).
	leftX := pad
	y0 := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.trusted_dealer"), image.Pt(leftX, y0), th.Color.Warning)
	y1 := y0 + gtx.Dp(unit.Dp(24))
	box(gtx, th, i18n.T("diagram.dealer"), image.Pt(leftX, y1), colW, bh, th.Color.WarningBg)

	y2 := y1 + bh + gtx.Dp(unit.Dp(16))
	shareW := pct(colW, 45)
	for i := 0; i < 3; i++ {
		sy := y2 + i*(bh+4)
		box(gtx, th, i18n.T("diagram.share")+" "+string(rune('1'+i)), image.Pt(leftX, sy), shareW, bh, th.Color.InfoBg)
		dirArrow(gtx, image.Pt(leftX+colW/2, y1+bh), image.Pt(leftX+shareW, sy+bh/2), lw, withAlpha(lc, 80))
	}
	caption(gtx, th, "\u2717 Single point of failure", image.Pt(leftX, y2+3*(bh+4)))

	// Right: ChillDKG (good).
	rightX := pad + usable - colW
	colorCaption(gtx, th, "ChillDKG", image.Pt(rightX, y0), th.Color.Primary)
	for i := 0; i < 3; i++ {
		sy := y1 + i*(bh+4)
		shadowBox(gtx, th, i18n.T("diagram.signer")+" "+string(rune('1'+i)), image.Pt(rightX, sy), colW, bh, th.Color.TipBg)
	}
	// Mesh lines between signers.
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			sy1 := y1 + i*(bh+4) + bh/2
			sy2 := y1 + j*(bh+4) + bh/2
			dashedLine(gtx, image.Pt(rightX+colW, sy1), image.Pt(rightX+colW, sy2), lw, withAlpha(th.Color.Primary, 80))
		}
	}
	colorCaption(gtx, th, "\u2713 No single dealer", image.Pt(rightX, y1+3*(bh+4)), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
