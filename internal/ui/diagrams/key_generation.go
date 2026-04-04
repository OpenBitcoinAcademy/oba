package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// KeyGeneration draws independent key generation vs HD for Ch05 L2.
// Top row: random → key (repeated, no backup link).
// Bottom row: single seed → all keys.
type KeyGeneration struct{}

func (d *KeyGeneration) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	bw := pct(usable, 20)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Row 0: Independent keys (bad).
	y0 := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.independent"), image.Pt(pad, y0), th.Color.Warning)
	y0 += gtx.Dp(unit.Dp(24))

	gap := pct(usable, 3)
	for i := 0; i < 3; i++ {
		x := pad + i*(bw+gap)
		box(gtx, th, "Random", image.Pt(x, y0), bw, bh, th.Color.InfoBg)
		dirArrow(gtx, image.Pt(x+bw/2, y0+bh), image.Pt(x+bw/2, y0+bh+gtx.Dp(unit.Dp(14))), lw, withAlpha(lc, 80))
		box(gtx, th, i18n.T("diagram.private_key"), image.Pt(x, y0+bh+gtx.Dp(unit.Dp(14))), bw, bh, th.Color.WarningBg)
	}
	caption(gtx, th, "\u2717 "+i18n.T("diagram.no_backup"), image.Pt(pad+3*(bw+gap)+gap, y0+bh/2))

	// Row 1: HD keys (good).
	y1 := y0 + 2*bh + gtx.Dp(unit.Dp(36))
	colorCaption(gtx, th, "HD", image.Pt(pad, y1), th.Color.Primary)
	y1 += gtx.Dp(unit.Dp(24))

	seedW := pct(usable, 22)
	shadowBox(gtx, th, i18n.T("diagram.seed"), image.Pt(pad, y1), seedW, bh, th.Color.Primary)

	for i := 0; i < 3; i++ {
		kx := pad + seedW + gap + i*(bw+gap)
		shadowBox(gtx, th, i18n.T("diagram.private_key"), image.Pt(kx, y1), bw, bh, th.Color.TipBg)
		curvedLine(gtx, image.Pt(pad+seedW, y1+bh/2), image.Pt(kx, y1+bh/2), lw, lc)
	}
	colorCaption(gtx, th, "\u2713 1 "+i18n.T("diagram.backup"), image.Pt(pad, y1+bh+gtx.Dp(unit.Dp(6))), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
