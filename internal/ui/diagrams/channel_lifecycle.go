package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ChannelLifecycle draws the Lightning channel open/close flow
// for Ch18 L4. Funding TX → commitment TXs → close.
type ChannelLifecycle struct{}

func (d *ChannelLifecycle) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(28))

	// Row 0: Funding TX (on-chain).
	y0 := gtx.Dp(unit.Dp(8))
	fundW := pct(usable, 40)
	fundX := pad + (usable-fundW)/2
	shadowBox(gtx, th, i18n.T("diagram.funding_tx"), image.Pt(fundX, y0), fundW, bh, th.Color.Primary)
	colorCaption(gtx, th, i18n.T("diagram.on_chain"), image.Pt(fundX+fundW+gtx.Dp(unit.Dp(8)), y0+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	// Row 1: Commitment TXs (off-chain, multiple updates).
	y1 := y0 + bh + rowGap
	cmtW := pct(usable, 26)
	cmtGap := pct(usable, 4)
	c1X := pad + (usable-3*cmtW-2*cmtGap)/2
	c2X := c1X + cmtW + cmtGap
	c3X := c2X + cmtW + cmtGap

	shadowBox(gtx, th, "#1", image.Pt(c1X, y1), cmtW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "#2", image.Pt(c2X, y1), cmtW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "#N", image.Pt(c3X, y1), cmtW, bh, th.Color.InfoBg)

	dirArrow(gtx, image.Pt(c1X+cmtW, y1+bh/2), image.Pt(c2X, y1+bh/2), lw, lc)
	dashedLine(gtx, image.Pt(c2X+cmtW, y1+bh/2), image.Pt(c3X, y1+bh/2), lw, lc)

	// Arrow from funding to commitments.
	dirArrow(gtx, image.Pt(fundX+fundW/2, y0+bh), image.Pt(c2X+cmtW/2, y1), lw, lc)

	colorCaption(gtx, th, i18n.T("diagram.off_chain"), image.Pt(c3X+cmtW+gtx.Dp(unit.Dp(8)), y1+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	// Row 2: Close options.
	y2 := y1 + bh + rowGap
	closeW := pct(usable, 34)
	closeGap := (usable - 2*closeW) / 3
	coopX := pad + closeGap
	forceX := coopX + closeW + closeGap

	shadowBox(gtx, th, i18n.T("diagram.cooperative_close"), image.Pt(coopX, y2), closeW, bh, th.Color.TipBg)
	shadowBox(gtx, th, i18n.T("diagram.force_close"), image.Pt(forceX, y2), closeW, bh, th.Color.WarningBg)

	curvedLine(gtx, image.Pt(c2X+cmtW/4, y1+bh), image.Pt(coopX+closeW/2, y2), lw, lc)
	curvedLine(gtx, image.Pt(c2X+3*cmtW/4, y1+bh), image.Pt(forceX+closeW/2, y2), lw, lc)

	colorCaption(gtx, th, i18n.T("diagram.on_chain"), image.Pt(forceX+closeW+gtx.Dp(unit.Dp(8)), y2+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
