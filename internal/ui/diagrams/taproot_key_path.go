package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TaprootKeyPath draws the Taproot key path spend flow for Ch14 L2.
// Row 1 (Tweak): Internal Key P + Merkle Root -> [TapTweak] -> Tweaked Key Q
// Row 2 (Spend): Tweaked Private Key + Transaction -> [Sign] -> 64-byte Signature
type TaprootKeyPath struct{}

func (d *TaprootKeyPath) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(175))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	inputW := pct(w, 22)
	procW := pct(w, 16)
	outputW := pct(w, 22)
	bh := gtx.Dp(unit.Dp(28))
	inputSpacing := 4

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	rowH := h / 2

	// === Row 1: Tweak ===
	r1LabelY := gtx.Dp(unit.Dp(2))
	colorCaption(gtx, th, i18n.T("diagram.tap_tweak"), image.Pt(pad, r1LabelY), th.Color.Primary)

	// Two input boxes stacked: Internal Key P, Merkle Root.
	inputX := pad
	inputY1 := r1LabelY + gtx.Dp(unit.Dp(24))
	inputY2 := inputY1 + bh + inputSpacing

	shadowBox(gtx, th, i18n.T("diagram.internal_key"), image.Pt(inputX, inputY1), inputW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, i18n.T("diagram.merkle_root"), image.Pt(inputX, inputY2), inputW, bh, th.Color.InfoBg)

	// Process box: TapTweak.
	procX := pad + inputW + pct(w, 4)
	procY := inputY1 + (inputY2+bh-inputY1-bh)/2
	shadowBox(gtx, th, i18n.T("diagram.tap_tweak"), image.Pt(procX, procY), procW, bh, th.Color.Primary)

	// Lines from inputs to process.
	dirArrow(gtx, image.Pt(inputX+inputW, inputY1+bh/2), image.Pt(procX, procY+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(inputX+inputW, inputY2+bh/2), image.Pt(procX, procY+bh/2), lw, lc)

	// Output: Tweaked Key Q.
	outX := procX + procW + pct(w, 4)
	dirArrow(gtx, image.Pt(procX+procW, procY+bh/2), image.Pt(outX, procY+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.tweaked_key"), image.Pt(outX, procY), outputW, bh, th.Color.TipBg)

	// Vertical connector between rows.
	connX := outX + outputW/2
	connTopY := procY + bh
	connBotY := rowH + gtx.Dp(unit.Dp(18)) + gtx.Dp(unit.Dp(16))
	line(gtx, image.Pt(connX, connTopY), image.Pt(connX, connBotY), lw, lc)

	// === Row 2: Spend ===
	r2LabelY := rowH + gtx.Dp(unit.Dp(2))
	colorCaption(gtx, th, i18n.T("diagram.sign"), image.Pt(pad, r2LabelY), th.Color.Primary)

	// Two input boxes stacked: Tweaked Private Key, Transaction.
	v1Y := r2LabelY + gtx.Dp(unit.Dp(24))
	v2Y := v1Y + bh + inputSpacing

	shadowBox(gtx, th, i18n.T("diagram.private_key"), image.Pt(inputX, v1Y), inputW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, i18n.T("diagram.transaction"), image.Pt(inputX, v2Y), inputW, bh, th.Color.InfoBg)

	// Process box: Sign.
	vprocY := v1Y + (v2Y+bh-v1Y-bh)/2
	shadowBox(gtx, th, i18n.T("diagram.sign"), image.Pt(procX, vprocY), procW, bh, th.Color.Primary)

	// Lines from inputs to process.
	dirArrow(gtx, image.Pt(inputX+inputW, v1Y+bh/2), image.Pt(procX, vprocY+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(inputX+inputW, v2Y+bh/2), image.Pt(procX, vprocY+bh/2), lw, lc)

	// Output: 64-byte Signature.
	dirArrow(gtx, image.Pt(procX+procW, vprocY+bh/2), image.Pt(outX, vprocY+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.schnorr_sig"), image.Pt(outX, vprocY), outputW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
