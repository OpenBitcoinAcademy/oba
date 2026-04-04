package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ECDSAFlow draws the ECDSA signing algorithm for Ch08 L2.
// Three inputs (private key, message hash, random k) feed into
// elliptic curve multiplication and modular arithmetic to produce
// the (r, s) signature pair.
type ECDSAFlow struct{}

func (d *ECDSAFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(280))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	smallW := pct(usable, 28)
	procW := pct(usable, 28)
	wideW := pct(usable, 38)
	outW := pct(usable, 34)

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(26))

	// Row 0: Three inputs across the top.
	y0 := gtx.Dp(unit.Dp(8))
	inputGap := (usable - 3*smallW) / 4
	xPriv := pad + inputGap
	xMsg := xPriv + smallW + inputGap
	xK := xMsg + smallW + inputGap

	shadowBox(gtx, th, i18n.T("diagram.private_key")+" (d)", image.Pt(xPriv, y0), smallW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, i18n.T("diagram.message")+" (z)", image.Pt(xMsg, y0), smallW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "Random k", image.Pt(xK, y0), smallW, bh, th.Color.InfoBg)

	// Row 1: k × G (under Random k).
	y1 := y0 + bh + rowGap
	kgX := xK + (smallW-procW)/2
	shadowBox(gtx, th, "k \u00d7 G", image.Pt(kgX, y1), procW, bh, th.Color.Primary)

	// Arrow: Random k → k×G (straight down).
	dirArrow(gtx, image.Pt(xK+smallW/2, y0+bh), image.Pt(xK+smallW/2, y1), lw, lc)

	// Row 2: r = x-coordinate (centered under k×G).
	y2 := y1 + bh + rowGap
	rX := kgX
	shadowBox(gtx, th, "r = x-coord", image.Pt(rX, y2), procW, bh, th.Color.InfoBg)

	// Arrow: k×G → r (straight down).
	dirArrow(gtx, image.Pt(kgX+procW/2, y1+bh), image.Pt(rX+procW/2, y2), lw, lc)

	// Row 3: s = k⁻¹(z + r·d) (centered in diagram).
	y3 := y2 + bh + rowGap
	sX := pad + (usable-wideW)/2
	shadowBox(gtx, th, "s = k\u207b\u00b9(z + r\u00b7d)", image.Pt(sX, y3), wideW, bh, th.Color.Primary)

	// Arrow: Private key → s (curved, from bottom of privkey to left side of s).
	curvedLine(gtx, image.Pt(xPriv+smallW/2, y0+bh), image.Pt(sX+wideW/5, y3), lw, lc)

	// Arrow: Message hash → s (curved, to center of s).
	curvedLine(gtx, image.Pt(xMsg+smallW/2, y0+bh), image.Pt(sX+wideW/2, y3), lw, lc)

	// Arrow: r → s (from bottom of r to right side of s).
	dirArrow(gtx, image.Pt(rX+procW/2, y2+bh), image.Pt(sX+4*wideW/5, y3), lw, lc)

	// Row 4: Signature output.
	y4 := y3 + bh + rowGap
	sigX := pad + (usable-outW)/2
	shadowBox(gtx, th, i18n.T("diagram.signature")+" (r, s)", image.Pt(sigX, y4), outW, bh, th.Color.TipBg)

	// Arrow: s → signature (straight down).
	dirArrow(gtx, image.Pt(sX+wideW/2, y3+bh), image.Pt(sigX+outW/2, y4), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
