package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ForkResolution draws two competing chains diverging from a common
// ancestor, with the longer chain winning. Ch12 L3 diagram.
type ForkResolution struct{}

func (d *ForkResolution) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bw := pct(usable, 16)
	bh := gtx.Dp(unit.Dp(26))
	gap := pct(usable, 3)
	lc := withAlpha(th.Color.TextMuted, 160)
	dimLC := withAlpha(th.Color.TextMuted, 60)
	lw := float32(1.8)

	// Shared ancestor blocks (left side).
	y := h/2 - bh/2
	x0 := pad
	x1 := x0 + bw + gap

	shadowBox(gtx, th, i18n.T("diagram.block")+" N", image.Pt(x0, y), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x0+bw, y+bh/2), image.Pt(x1, y+bh/2), lw, lc)

	// Fork point.
	forkX := x1
	shadowBox(gtx, th, i18n.T("diagram.block")+" N+1", image.Pt(forkX, y), bw, bh, th.Color.InfoBg)

	// Chain A (top, winner): 3 blocks after fork.
	yA := y - bh - gtx.Dp(unit.Dp(12))
	xA1 := forkX + bw + gap
	xA2 := xA1 + bw + gap
	xA3 := xA2 + bw + gap

	shadowBox(gtx, th, "A1", image.Pt(xA1, yA), bw, bh, th.Color.TipBg)
	dirArrow(gtx, image.Pt(xA1+bw, yA+bh/2), image.Pt(xA2, yA+bh/2), lw, lc)
	shadowBox(gtx, th, "A2", image.Pt(xA2, yA), bw, bh, th.Color.TipBg)
	dirArrow(gtx, image.Pt(xA2+bw, yA+bh/2), image.Pt(xA3, yA+bh/2), lw, lc)
	shadowBox(gtx, th, "A3", image.Pt(xA3, yA), bw, bh, th.Color.TipBg)

	// Arrow from fork to chain A.
	curvedLine(gtx, image.Pt(forkX+bw, y+bh/4), image.Pt(xA1, yA+bh/2), lw, lc)

	// Chain B (bottom, orphaned): 2 blocks, dimmed.
	yB := y + bh + gtx.Dp(unit.Dp(12))
	xB1 := forkX + bw + gap
	xB2 := xB1 + bw + gap

	box(gtx, th, "B1", image.Pt(xB1, yB), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(xB1+bw, yB+bh/2), image.Pt(xB2, yB+bh/2), lw, dimLC)
	box(gtx, th, "B2", image.Pt(xB2, yB), bw, bh, th.Color.InfoBg)

	// Arrow from fork to chain B (dimmed).
	dashedLine(gtx, image.Pt(forkX+bw, y+3*bh/4), image.Pt(xB1, yB+bh/2), lw, dimLC)

	// Labels.
	colorCaption(gtx, th, i18n.T("diagram.valid_block"), image.Pt(xA3+bw+gap/2, yA+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.Primary)
	caption(gtx, th, i18n.T("diagram.orphaned"), image.Pt(xB2+bw+gap/2, yB+(bh-gtx.Dp(unit.Dp(12)))/2))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
