package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// PackageRelay draws parent-child transaction fee evaluation for Ch09 L4.
// Parent TX (low fee) + Child TX (high fee) evaluated together.
type PackageRelay struct{}

func (d *PackageRelay) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	bw := pct(usable, 32)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(26))

	// Parent TX (low fee).
	y0 := gtx.Dp(unit.Dp(8))
	parentX := pad + (usable-bw)/2
	shadowBox(gtx, th, "Parent TX", image.Pt(parentX, y0), bw, bh, th.Color.InfoBg)
	colorCaption(gtx, th, i18n.T("diagram.fee")+": 1 sat/vB", image.Pt(parentX+bw+gtx.Dp(unit.Dp(8)), y0+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	// Child TX (high fee, pays for both).
	y1 := y0 + bh + rowGap
	childX := parentX
	shadowBox(gtx, th, "Child TX", image.Pt(childX, y1), bw, bh, th.Color.TipBg)
	colorCaption(gtx, th, i18n.T("diagram.fee")+": 50 sat/vB", image.Pt(childX+bw+gtx.Dp(unit.Dp(8)), y1+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.Primary)

	dirArrow(gtx, image.Pt(parentX+bw/2, y0+bh), image.Pt(childX+bw/2, y1), lw, lc)

	// Package evaluation.
	y2 := y1 + bh + rowGap
	pkgW := pct(usable, 44)
	pkgX := pad + (usable-pkgW)/2
	groupBox(gtx, th, image.Pt(pkgX-4, y0-4), pkgW+8, y1+bh-y0+8)
	shadowBox(gtx, th, "CPFP: 25 sat/vB", image.Pt(pkgX, y2), pkgW, bh, th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
