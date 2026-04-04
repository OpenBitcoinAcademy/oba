package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SeedRecovery draws the BIP39 recovery flow for Ch05 L4.
// 24 words → seed → master key → wallet.
type SeedRecovery struct{}

func (d *SeedRecovery) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(100))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(30))
	bw := pct(usable, 22)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	gap := (usable - 4*bw) / 3
	if gap < 6 {
		gap = 6
	}
	y := (h - bh) / 2

	x1 := pad
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap
	x4 := x3 + bw + gap

	shadowBox(gtx, th, "24 "+i18n.T("diagram.words"), image.Pt(x1, y), bw, bh, th.Color.WarningBg)
	dirArrow(gtx, image.Pt(x1+bw, y+bh/2), image.Pt(x2, y+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.seed"), image.Pt(x2, y), bw, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x2+bw, y+bh/2), image.Pt(x3, y+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.master_key"), image.Pt(x3, y), bw, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x3+bw, y+bh/2), image.Pt(x4, y+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.wallet"), image.Pt(x4, y), bw, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
