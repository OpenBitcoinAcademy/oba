package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SPVProof draws the SPV Merkle proof verification for Ch10 L4.
// Header chain → merkle root → proof path → target TX.
type SPVProof struct{}

func (d *SPVProof) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(24))

	// Row 0: Block headers (light client stores these).
	y0 := gtx.Dp(unit.Dp(8))
	hdrW := pct(usable, 18)
	hdrGap := pct(usable, 3)
	for i := 0; i < 4; i++ {
		x := pad + i*(hdrW+hdrGap)
		shadowBox(gtx, th, "Hdr", image.Pt(x, y0), hdrW, bh, th.Color.InfoBg)
		if i < 3 {
			dirArrow(gtx, image.Pt(x+hdrW, y0+bh/2), image.Pt(x+hdrW+hdrGap, y0+bh/2), lw, lc)
		}
	}
	colorCaption(gtx, th, i18n.T("diagram.header_chain"), image.Pt(pad+4*(hdrW+hdrGap), y0+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	// Row 1: Merkle proof path.
	y1 := y0 + bh + rowGap
	proofW := pct(usable, 22)
	proofGap := pct(usable, 4)
	p1X := pad + pct(usable, 10)
	p2X := p1X + proofW + proofGap

	shadowBox(gtx, th, i18n.T("diagram.merkle_root"), image.Pt(p1X, y1), proofW, bh, th.Color.Primary)
	shadowBox(gtx, th, i18n.T("diagram.hash_pair"), image.Pt(p2X, y1), proofW, bh, th.Color.WarningBg)

	// Arrow from header to merkle root.
	dirArrow(gtx, image.Pt(pad+hdrW/2, y0+bh), image.Pt(p1X+proofW/2, y1), lw, lc)
	dirArrow(gtx, image.Pt(p1X+proofW, y1+bh/2), image.Pt(p2X, y1+bh/2), lw, lc)

	// Row 2: Target TX.
	y2 := y1 + bh + rowGap
	txW := pct(usable, 30)
	txX := pad + (usable-txW)/2
	shadowBox(gtx, th, i18n.T("diagram.target_tx"), image.Pt(txX, y2), txW, bh, th.Color.TipBg)

	curvedLine(gtx, image.Pt(p2X+proofW/2, y1+bh), image.Pt(txX+txW/2, y2), lw, lc)
	colorCaption(gtx, th, i18n.T("diagram.verified"), image.Pt(txX+txW+gtx.Dp(unit.Dp(8)), y2+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
