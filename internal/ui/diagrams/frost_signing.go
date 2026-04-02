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
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	usable := w - 2*pad
	bh := gtx.Dp(unit.Dp(26))
	signerW := pct(usable, 16)
	coordW := pct(usable, 20)
	sigW := pct(usable, 18)
	lineColor := withAlpha(th.Color.TextMuted, 120)
	lw := float32(1.5)

	// Signer positions (left side, stacked vertically in each row).
	signerGap := gtx.Dp(unit.Dp(4))
	signerBlockH := 3*bh + 2*signerGap

	// Row 1 vertical space.
	row1Top := gtx.Dp(unit.Dp(16))
	row1CaptionY := row1Top - gtx.Dp(unit.Dp(2))
	row1SignerY := row1Top + gtx.Dp(unit.Dp(2))
	row1CoordY := row1SignerY + (signerBlockH-bh)/2

	// Row 2 vertical space.
	row2Top := row1SignerY + signerBlockH + gtx.Dp(unit.Dp(16))
	row2CaptionY := row2Top - gtx.Dp(unit.Dp(2))
	row2SignerY := row2Top + gtx.Dp(unit.Dp(2))
	row2CoordY := row2SignerY + (signerBlockH-bh)/2

	signerX := pad
	coordX := pad + pct(usable, 35)
	sigX := pad + pct(usable, 75)

	signers := []string{"A", "B", "C"}

	// --- Row 1: Commitments to Coordinator ---
	caption(gtx, th, i18n.T("diagram.round_1"), image.Pt(pad, row1CaptionY))

	for idx, s := range signers {
		sy := row1SignerY + idx*(bh+signerGap)
		box(gtx, th, s, image.Pt(signerX, sy), signerW, bh, th.Color.InfoBg)
		// Arrow from signer to coordinator.
		line(gtx,
			image.Pt(signerX+signerW, sy+bh/2),
			image.Pt(coordX, row1CoordY+bh/2),
			lw, lineColor)
	}

	// Coordinator box.
	box(gtx, th, i18n.T("diagram.coordinator"), image.Pt(coordX, row1CoordY), coordW, bh, th.Color.WarningBg)

	// Commitment label on the arrows.
	commitLabelY := row1CoordY - gtx.Dp(unit.Dp(12))
	caption(gtx, th, i18n.T("diagram.commitment"), image.Pt(signerX+signerW+pct(usable, 4), commitLabelY))

	// --- Row 2: Challenge -> Shares -> Signature ---
	caption(gtx, th, i18n.T("diagram.round_2"), image.Pt(pad, row2CaptionY))

	// Coordinator sends challenge.
	box(gtx, th, i18n.T("diagram.coordinator"), image.Pt(coordX, row2CoordY), coordW, bh, th.Color.WarningBg)

	// Challenge label.
	challengeLabelY := row2CoordY - gtx.Dp(unit.Dp(12))
	caption(gtx, th, i18n.T("diagram.challenge"), image.Pt(coordX+pct(usable, 2), challengeLabelY))

	for idx, s := range signers {
		sy := row2SignerY + idx*(bh+signerGap)
		box(gtx, th, s, image.Pt(signerX, sy), signerW, bh, th.Color.InfoBg)
		// Line from coordinator back to signer (challenge).
		line(gtx,
			image.Pt(coordX, row2CoordY+bh/2),
			image.Pt(signerX+signerW, sy+bh/2),
			lw, lineColor)
	}

	// Share label.
	shareLabelY := row2SignerY + signerBlockH + gtx.Dp(unit.Dp(2))
	caption(gtx, th, i18n.T("diagram.share"), image.Pt(signerX+signerW+pct(usable, 4), shareLabelY))

	// Arrow from coordinator to final signature.
	arrow(gtx, th,
		image.Pt(coordX+coordW, row2CoordY+bh/2),
		image.Pt(sigX, row2CoordY+bh/2))

	// Final Signature box.
	box(gtx, th, "Signature", image.Pt(sigX, row2CoordY), sigW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
