package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// FedimintArch draws the Fedimint federation architecture for Ch19 L3.
// Threshold guardians hold keys, users get ecash tokens.
type FedimintArch struct{}

func (d *FedimintArch) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(24))

	// Row 0: Guardians (threshold multisig).
	y0 := gtx.Dp(unit.Dp(8))
	gW := pct(usable, 20)
	gGap := (usable - 4*gW) / 5

	for i := 0; i < 4; i++ {
		x := pad + gGap + i*(gW+gGap)
		shadowBox(gtx, th, i18n.T("diagram.guardian")+" "+string(rune('1'+i)), image.Pt(x, y0), gW, bh, th.Color.WarningBg)
	}

	// Row 1: Federation (multisig wallet).
	y1 := y0 + bh + rowGap
	fedW := pct(usable, 40)
	fedX := pad + (usable-fedW)/2
	shadowBox(gtx, th, "3-of-4 Federation", image.Pt(fedX, y1), fedW, bh, th.Color.Primary)

	for i := 0; i < 4; i++ {
		x := pad + gGap + i*(gW+gGap) + gW/2
		curvedLine(gtx, image.Pt(x, y0+bh), image.Pt(fedX+fedW/2, y1), lw, lc)
	}

	// Row 2: Users with ecash.
	y2 := y1 + bh + rowGap
	uW := pct(usable, 22)
	uGap := (usable - 3*uW) / 4

	for i := 0; i < 3; i++ {
		x := pad + uGap + i*(uW+uGap)
		shadowBox(gtx, th, "e-cash", image.Pt(x, y2), uW, bh, th.Color.TipBg)
		dashedLine(gtx, image.Pt(fedX+fedW/2, y1+bh), image.Pt(x+uW/2, y2), lw, withAlpha(lc, 100))
	}

	colorCaption(gtx, th, "Blind signatures \u2192 privacy", image.Pt(pad, y2+bh+gtx.Dp(unit.Dp(6))), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
