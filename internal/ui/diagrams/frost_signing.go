package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// FrostSigning draws the two-round FROST threshold signing protocol
// for Ch17. Row 1: signers send commitments to the coordinator.
// Row 2: coordinator sends a challenge, signers return shares,
// producing a final signature.
type FrostSigning struct{}

func (d *FrostSigning) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	usable := w - 2*pad
	bh := gtx.Dp(unit.Dp(26))
	signerW := pct(usable, 16)
	coordW := pct(usable, 20)
	sigW := pct(usable, 18)

	lineColor := withAlpha(th.Color.TextMuted, 140)
	dashColor := withAlpha(th.Color.Warning, 120)
	lw := float32(1.8)

	// Signer block dimensions.
	signerGap := gtx.Dp(unit.Dp(4))
	signerBlockH := 3*bh + 2*signerGap

	// Row 1 layout.
	row1Top := gtx.Dp(unit.Dp(18))
	row1SignerY := row1Top + gtx.Dp(unit.Dp(4))
	row1CoordY := row1SignerY + (signerBlockH-bh)/2

	// Row 2 layout.
	row2Top := row1SignerY + signerBlockH + gtx.Dp(unit.Dp(18))
	row2SignerY := row2Top + gtx.Dp(unit.Dp(4))
	row2CoordY := row2SignerY + (signerBlockH-bh)/2

	signerX := pad
	coordX := pad + pct(usable, 35)
	sigX := pad + pct(usable, 75)

	signers := []string{"A", "B", "C"}

	// --- Row 1: Commitments to Coordinator ---
	// Group box around round 1.
	groupBox(gtx, th, image.Pt(pad-4, row1Top-gtx.Dp(unit.Dp(4))), usable+8, signerBlockH+gtx.Dp(unit.Dp(12)))
	colorCaption(gtx, th, i18n.T("diagram.round_1"), image.Pt(pad, row1Top-gtx.Dp(unit.Dp(2))), th.Color.Primary)

	for idx, s := range signers {
		sy := row1SignerY + idx*(bh+signerGap)
		shadowBox(gtx, th, s, image.Pt(signerX, sy), signerW, bh, th.Color.InfoBg)
		// Directional arrow from signer to coordinator.
		dirArrow(gtx,
			image.Pt(signerX+signerW, sy+bh/2),
			image.Pt(coordX, row1CoordY+bh/2),
			lw, lineColor)
	}

	shadowBox(gtx, th, i18n.T("diagram.coordinator"), image.Pt(coordX, row1CoordY), coordW, bh, th.Color.WarningBg)

	// Commitment label.
	commitLabelY := row1CoordY - gtx.Dp(unit.Dp(12))
	caption(gtx, th, i18n.T("diagram.commitment"), image.Pt(signerX+signerW+pct(usable, 4), commitLabelY))

	// --- Row 2: Challenge -> Shares -> Signature ---
	groupBox(gtx, th, image.Pt(pad-4, row2Top-gtx.Dp(unit.Dp(4))), usable+8, signerBlockH+gtx.Dp(unit.Dp(12)))
	colorCaption(gtx, th, i18n.T("diagram.round_2"), image.Pt(pad, row2Top-gtx.Dp(unit.Dp(2))), th.Color.Primary)

	shadowBox(gtx, th, i18n.T("diagram.coordinator"), image.Pt(coordX, row2CoordY), coordW, bh, th.Color.WarningBg)

	for idx, s := range signers {
		sy := row2SignerY + idx*(bh+signerGap)
		shadowBox(gtx, th, s, image.Pt(signerX, sy), signerW, bh, th.Color.InfoBg)
		// Dashed line from coordinator back to signer (challenge).
		dashedLine(gtx,
			image.Pt(coordX, row2CoordY+bh/2),
			image.Pt(signerX+signerW, sy+bh/2),
			lw, dashColor)
	}

	// Challenge label.
	challengeLabelY := row2CoordY - gtx.Dp(unit.Dp(12))
	caption(gtx, th, i18n.T("diagram.challenge"), image.Pt(coordX+pct(usable, 2), challengeLabelY))

	// Arrow from coordinator to final signature.
	dirArrow(gtx,
		image.Pt(coordX+coordW, row2CoordY+bh/2),
		image.Pt(sigX, row2CoordY+bh/2),
		lw, lineColor)

	shadowBox(gtx, th, "Signature", image.Pt(sigX, row2CoordY), sigW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
