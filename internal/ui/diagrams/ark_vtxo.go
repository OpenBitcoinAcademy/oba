package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ArkVTXO draws the Ark VTXO tree structure for Ch19 L4.
// ASP pools UTXOs into a shared UTXO, users get VTXOs.
type ArkVTXO struct{}

func (d *ArkVTXO) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(24))

	// Row 0: ASP (Ark Service Provider).
	y0 := gtx.Dp(unit.Dp(8))
	aspW := pct(usable, 36)
	aspX := pad + (usable-aspW)/2
	shadowBox(gtx, th, "ASP", image.Pt(aspX, y0), aspW, bh, th.Color.Primary)

	// Row 1: Shared UTXO (on-chain).
	y1 := y0 + bh + rowGap
	sharedW := pct(usable, 44)
	sharedX := pad + (usable-sharedW)/2
	shadowBox(gtx, th, i18n.T("diagram.shared_utxo"), image.Pt(sharedX, y1), sharedW, bh, th.Color.WarningBg)
	dirArrow(gtx, image.Pt(aspX+aspW/2, y0+bh), image.Pt(sharedX+sharedW/2, y1), lw, lc)

	// Row 2: VTXOs (virtual, off-chain).
	y2 := y1 + bh + rowGap
	vtxoW := pct(usable, 20)
	vtxoGap := (usable - 4*vtxoW) / 5

	for i := 0; i < 4; i++ {
		x := pad + vtxoGap + i*(vtxoW+vtxoGap)
		shadowBox(gtx, th, "VTXO", image.Pt(x, y2), vtxoW, bh, th.Color.TipBg)
		curvedLine(gtx, image.Pt(sharedX+sharedW/2, y1+bh), image.Pt(x+vtxoW/2, y2), lw, withAlpha(lc, 120))
	}

	colorCaption(gtx, th, i18n.T("diagram.off_chain")+" \u2192 "+i18n.T("diagram.on_chain")+" exit", image.Pt(pad, y2+bh+gtx.Dp(unit.Dp(6))), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
