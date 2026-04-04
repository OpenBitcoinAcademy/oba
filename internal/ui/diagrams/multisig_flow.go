package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MultisigFlow draws a 2-of-3 multisig signing flow for Ch07 L2.
// Three signers at top, two provide signatures, which combine
// into a P2SH redeem script that unlocks the output.
type MultisigFlow struct{}

func (d *MultisigFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(220))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	badgeR := 12
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(28))

	// Row 0: Three signers.
	signerW := pct(usable, 24)
	signerGap := (usable - 3*signerW) / 4
	y0 := gtx.Dp(unit.Dp(8))

	x1 := pad + signerGap
	x2 := x1 + signerW + signerGap
	x3 := x2 + signerW + signerGap

	shadowBox(gtx, th, i18n.T("diagram.signer")+" 1", image.Pt(x1, y0), signerW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.signer")+" 2", image.Pt(x2, y0), signerW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.signer")+" 3", image.Pt(x3, y0), signerW, bh, th.Color.InfoBg)

	// Badges showing which sign.
	badge(gtx, th, "\u2713", image.Pt(x1+signerW+6, y0+bh/2), badgeR, th.Color.Primary)
	badge(gtx, th, "\u2713", image.Pt(x2+signerW+6, y0+bh/2), badgeR, th.Color.Primary)

	// Row 1: 2-of-3.
	y1 := y0 + bh + rowGap
	reqW := pct(usable, 30)
	reqX := pad + (usable-reqW)/2
	shadowBox(gtx, th, "2-of-3", image.Pt(reqX, y1), reqW, bh, th.Color.Primary)

	// Arrows from signers 1 and 2 to 2-of-3.
	curvedLine(gtx, image.Pt(x1+signerW/2, y0+bh), image.Pt(reqX+reqW/3, y1), lw, lc)
	curvedLine(gtx, image.Pt(x2+signerW/2, y0+bh), image.Pt(reqX+2*reqW/3, y1), lw, lc)
	// Dashed from signer 3 (not signing this time).
	dashedLine(gtx, image.Pt(x3+signerW/2, y0+bh), image.Pt(reqX+reqW, y1), lw, withAlpha(lc, 60))

	// Row 2: P2SH.
	y2 := y1 + bh + rowGap
	p2shW := pct(usable, 36)
	p2shX := pad + (usable-p2shW)/2
	shadowBox(gtx, th, "P2SH "+i18n.T("diagram.address"), image.Pt(p2shX, y2), p2shW, bh, th.Color.WarningBg)

	dirArrow(gtx, image.Pt(reqX+reqW/2, y1+bh), image.Pt(p2shX+p2shW/2, y2), lw, lc)

	// Row 3: Unlocked.
	y3 := y2 + bh + rowGap
	outW := pct(usable, 28)
	outX := pad + (usable-outW)/2
	shadowBox(gtx, th, i18n.T("diagram.utxo_unspent"), image.Pt(outX, y3), outW, bh, th.Color.TipBg)

	dirArrow(gtx, image.Pt(p2shX+p2shW/2, y2+bh), image.Pt(outX+outW/2, y3), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
