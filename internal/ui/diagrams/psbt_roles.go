package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// PSBTRoles draws the PSBT role pipeline for Ch16 L2.
// Two rows: Creator→Updater→Signer on top,
// Combiner→Finalizer→Extractor on bottom.
type PSBTRoles struct{}

func (d *PSBTRoles) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	bw := pct(usable, 28)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	gap := (usable - 3*bw) / 4

	// Row 1: Creator → Updater → Signer.
	y1 := gtx.Dp(unit.Dp(8))
	x1 := pad + gap
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap

	shadowBox(gtx, th, i18n.T("diagram.creator"), image.Pt(x1, y1), bw, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x1+bw, y1+bh/2), image.Pt(x2, y1+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.updater"), image.Pt(x2, y1), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x2+bw, y1+bh/2), image.Pt(x3, y1+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.signer"), image.Pt(x3, y1), bw, bh, th.Color.WarningBg)

	// Connector from row 1 to row 2.
	y2 := y1 + bh + gtx.Dp(unit.Dp(36))
	roundedPath(gtx, []image.Point{
		{x3 + bw/2, y1 + bh},
		{x3 + bw/2, y1 + bh + gtx.Dp(unit.Dp(18))},
		{x1 + bw/2, y1 + bh + gtx.Dp(unit.Dp(18))},
		{x1 + bw/2, y2},
	}, 12, lw, lc)

	// Row 2: Combiner → Finalizer → Extractor.
	shadowBox(gtx, th, i18n.T("diagram.combiner"), image.Pt(x1, y2), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x1+bw, y2+bh/2), image.Pt(x2, y2+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.finalizer"), image.Pt(x2, y2), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x2+bw, y2+bh/2), image.Pt(x3, y2+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.extractor"), image.Pt(x3, y2), bw, bh, th.Color.TipBg)

	// Step badges.
	badge(gtx, th, "1", image.Pt(x1-8, y1+bh/2), 8, th.Color.Primary)
	badge(gtx, th, "2", image.Pt(x2-8, y1+bh/2), 8, th.Color.Primary)
	badge(gtx, th, "3", image.Pt(x3-8, y1+bh/2), 8, th.Color.Primary)
	badge(gtx, th, "4", image.Pt(x1-8, y2+bh/2), 8, th.Color.Primary)
	badge(gtx, th, "5", image.Pt(x2-8, y2+bh/2), 8, th.Color.Primary)
	badge(gtx, th, "6", image.Pt(x3-8, y2+bh/2), 8, th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
