package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// BIP32Derivation draws an HD wallet derivation path tree for Ch05 L3.
// m → 44' → 0' → 0' → 0 → addresses (BIP44 standard path).
// Hardened levels shown with solid lines, normal with dashed.
type BIP32Derivation struct{}

func (d *BIP32Derivation) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	bw := pct(usable, 18)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(20))

	// Row 0: Master (m).
	y0 := gtx.Dp(unit.Dp(8))
	mW := pct(usable, 22)
	mX := pad + (usable-mW)/2
	shadowBox(gtx, th, i18n.T("diagram.master_key")+" (m)", image.Pt(mX, y0), mW, bh, th.Color.Primary)

	// Row 1: Hardened purpose/coin/account (44'/0'/0').
	y1 := y0 + bh + rowGap
	gap3 := (usable - 3*bw) / 4
	x1 := pad + gap3
	x2 := x1 + bw + gap3
	x3 := x2 + bw + gap3

	shadowBox(gtx, th, "44'", image.Pt(x1, y1), bw, bh, th.Color.WarningBg)
	shadowBox(gtx, th, "0'", image.Pt(x2, y1), bw, bh, th.Color.WarningBg)
	shadowBox(gtx, th, "0'", image.Pt(x3, y1), bw, bh, th.Color.WarningBg)

	// Solid lines for hardened derivation.
	curvedLine(gtx, image.Pt(mX+mW/2, y0+bh), image.Pt(x1+bw/2, y1), lw, lc)
	dirArrow(gtx, image.Pt(x1+bw, y1+bh/2), image.Pt(x2, y1+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(x2+bw, y1+bh/2), image.Pt(x3, y1+bh/2), lw, lc)

	// Row 2: Normal chain (0 = external, 1 = change).
	y2 := y1 + bh + rowGap
	chainGap := (usable - 2*bw) / 3
	xExt := pad + chainGap
	xChg := xExt + bw + chainGap

	shadowBox(gtx, th, "0 ("+i18n.T("diagram.external")+")", image.Pt(xExt, y2), bw, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "1 ("+i18n.T("diagram.change")+")", image.Pt(xChg, y2), bw, bh, th.Color.InfoBg)

	// Dashed lines for normal derivation.
	dashedLine(gtx, image.Pt(x3+bw/2, y1+bh), image.Pt(xExt+bw/2, y2), lw, lc)
	dashedLine(gtx, image.Pt(x3+bw/2, y1+bh), image.Pt(xChg+bw/2, y2), lw, lc)

	// Row 3: Addresses.
	y3 := y2 + bh + rowGap
	addrW := pct(usable, 16)
	addrGap := pct(usable, 3)
	a0X := xExt
	a1X := a0X + addrW + addrGap

	shadowBox(gtx, th, i18n.T("diagram.hd_address")+" 0", image.Pt(a0X, y3), addrW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.hd_address")+" 1", image.Pt(a1X, y3), addrW, bh, th.Color.TipBg)

	dashedLine(gtx, image.Pt(xExt+bw/2, y2+bh), image.Pt(a0X+addrW/2, y3), lw, withAlpha(lc, 100))
	dashedLine(gtx, image.Pt(xExt+bw/2, y2+bh), image.Pt(a1X+addrW/2, y3), lw, withAlpha(lc, 100))

	// Legend.
	legendY := y3 + bh + gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, "' = "+i18n.T("diagram.hardened"), image.Pt(pad, legendY), th.Color.TextMuted)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
