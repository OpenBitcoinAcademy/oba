package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SignatureFlow draws digital signature creation and verification for Ch08.
// Row 1 (Sign):   Private Key + Message -> [Sign] -> Signature
// Row 2 (Verify): Public Key + Message + Signature -> [Verify] -> True/False
type SignatureFlow struct{}

func (d *SignatureFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	inputW := pct(w, 22)
	procW := pct(w, 16)
	outputW := pct(w, 22)
	bh := gtx.Dp(unit.Dp(28))
	inputSpacing := 4
	rowH := h / 2

	lc := th.Color.TextMuted
	lw := float32(1.5)

	// === Row 1: Signing ===
	r1LabelY := gtx.Dp(unit.Dp(2))
	colorCaption(gtx, th, i18n.T("diagram.sign"), image.Pt(pad, r1LabelY), th.Color.Primary)

	// Two input boxes stacked: Private Key, Message.
	inputX := pad
	inputY1 := r1LabelY + gtx.Dp(unit.Dp(16))
	inputY2 := inputY1 + bh + inputSpacing

	box(gtx, th, i18n.T("diagram.private_key"), image.Pt(inputX, inputY1), inputW, bh, th.Color.WarningBg)
	box(gtx, th, i18n.T("diagram.message"), image.Pt(inputX, inputY2), inputW, bh, th.Color.InfoBg)

	// Process box: Sign.
	procX := pad + inputW + pct(w, 4)
	procY := inputY1 + (inputY2+bh-inputY1-bh)/2
	processBox(gtx, th, i18n.T("diagram.sign"), image.Pt(procX, procY), procW, bh)

	// Lines from inputs to process.
	line(gtx, image.Pt(inputX+inputW, inputY1+bh/2), image.Pt(procX, procY+bh/2), lw, lc)
	line(gtx, image.Pt(inputX+inputW, inputY2+bh/2), image.Pt(procX, procY+bh/2), lw, lc)

	// Output: Signature.
	outX := procX + procW + pct(w, 4)
	arrow(gtx, th, image.Pt(procX+procW, procY+bh/2), image.Pt(outX, procY+bh/2))
	box(gtx, th, i18n.T("diagram.signature"), image.Pt(outX, procY), outputW, bh, th.Color.TipBg)

	// === Row 2: Verification ===
	r2LabelY := rowH + gtx.Dp(unit.Dp(2))
	colorCaption(gtx, th, i18n.T("diagram.verify"), image.Pt(pad, r2LabelY), th.Color.Primary)

	// Three input boxes stacked: Public Key, Message, Signature.
	v1Y := r2LabelY + gtx.Dp(unit.Dp(16))
	v2Y := v1Y + bh + inputSpacing
	v3Y := v2Y + bh + inputSpacing

	box(gtx, th, i18n.T("diagram.public_key"), image.Pt(inputX, v1Y), inputW, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.message"), image.Pt(inputX, v2Y), inputW, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.signature"), image.Pt(inputX, v3Y), inputW, bh, th.Color.TipBg)

	// Process box: Verify.
	vprocY := v2Y
	processBox(gtx, th, i18n.T("diagram.verify"), image.Pt(procX, vprocY), procW, bh)

	// Lines from inputs to process.
	line(gtx, image.Pt(inputX+inputW, v1Y+bh/2), image.Pt(procX, vprocY+bh/2), lw, lc)
	line(gtx, image.Pt(inputX+inputW, v2Y+bh/2), image.Pt(procX, vprocY+bh/2), lw, lc)
	line(gtx, image.Pt(inputX+inputW, v3Y+bh/2), image.Pt(procX, vprocY+bh/2), lw, lc)

	// Output: Result.
	arrow(gtx, th, image.Pt(procX+procW, vprocY+bh/2), image.Pt(outX, vprocY+bh/2))
	box(gtx, th, i18n.T("diagram.result"), image.Pt(outX, vprocY), outputW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
