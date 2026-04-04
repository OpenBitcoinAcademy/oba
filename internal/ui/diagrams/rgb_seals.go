package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// RGBSeals draws single-use seal concept for Ch19 L2.
// UTXO as seal → state transition → new UTXO seals new state.
type RGBSeals struct{}

func (d *RGBSeals) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	bw := pct(usable, 22)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	gap := (usable - 4*bw) / 3
	y := (h - bh) / 2

	x1 := pad
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap
	x4 := x3 + bw + gap

	shadowBox(gtx, th, "UTXO (seal)", image.Pt(x1, y), bw, bh, th.Color.WarningBg)
	dirArrow(gtx, image.Pt(x1+bw, y+bh/2), image.Pt(x2, y+bh/2), lw, lc)
	shadowBox(gtx, th, "State \u2192", image.Pt(x2, y), bw, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x2+bw, y+bh/2), image.Pt(x3, y+bh/2), lw, lc)
	shadowBox(gtx, th, "New UTXO", image.Pt(x3, y), bw, bh, th.Color.TipBg)
	dirArrow(gtx, image.Pt(x3+bw, y+bh/2), image.Pt(x4, y+bh/2), lw, lc)
	shadowBox(gtx, th, "New State", image.Pt(x4, y), bw, bh, th.Color.TipBg)

	// Caption.
	captionY := y + bh + gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "Client-side validated", image.Pt(pad, captionY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
