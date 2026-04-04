package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// RBFBumping draws the Replace-By-Fee mechanism for Ch09 L3.
// Shows original TX being replaced by a higher-fee version in the mempool.
type RBFBumping struct{}

func (d *RBFBumping) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Left column: original TX (dimmed, being replaced).
	colW := pct(usable, 38)
	colGap := pct(usable, 8)
	leftX := pad
	rightX := pad + colW + colGap

	// Original TX (crossed out / dimmed).
	y0 := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.original"), image.Pt(leftX, y0), th.Color.TextMuted)
	y1 := y0 + gtx.Dp(unit.Dp(24))
	box(gtx, th, "TX v1", image.Pt(leftX, y1), colW, bh, th.Color.InfoBg)
	y2 := y1 + bh + 4
	box(gtx, th, i18n.T("diagram.fee")+": 500 sat", image.Pt(leftX, y2), colW, bh, th.Color.InfoBg)

	// Replacement TX (highlighted).
	colorCaption(gtx, th, i18n.T("diagram.replacement"), image.Pt(rightX, y0), th.Color.Primary)
	shadowBox(gtx, th, "TX v2", image.Pt(rightX, y1), colW, bh, th.Color.TipBg)
	yr2 := y1 + bh + 4
	shadowBox(gtx, th, i18n.T("diagram.fee")+": 2000 sat", image.Pt(rightX, yr2), colW, bh, th.Color.Primary)

	// Arrow from v1 to v2 (replaces).
	arrowY := y1 + bh/2
	dirArrow(gtx, image.Pt(leftX+colW, arrowY), image.Pt(rightX, arrowY), lw, lc)

	// Mempool box below.
	mempoolY := yr2 + bh + gtx.Dp(unit.Dp(20))
	mempoolW := pct(usable, 50)
	mempoolX := pad + (usable-mempoolW)/2
	shadowBox(gtx, th, i18n.T("diagram.mempool"), image.Pt(mempoolX, mempoolY), mempoolW, bh, th.Color.WarningBg)

	// Dashed from old TX (rejected).
	dashedLine(gtx, image.Pt(leftX+colW/2, y2+bh), image.Pt(mempoolX+mempoolW/4, mempoolY), lw, withAlpha(lc, 60))
	// Solid from new TX (accepted).
	dirArrow(gtx, image.Pt(rightX+colW/2, yr2+bh), image.Pt(mempoolX+3*mempoolW/4, mempoolY), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
