package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TapscriptVerify draws OP_CHECKSIGADD accumulation for Ch14 L4.
// Three signature checks building up a counter, then threshold check.
type TapscriptVerify struct{}

func (d *TapscriptVerify) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(110))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	bw := pct(usable, 20)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	y := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "OP_CHECKSIGADD", image.Pt(pad, y), th.Color.Primary)
	y += gtx.Dp(unit.Dp(24))

	gap := (usable - 4*bw) / 5

	x1 := pad + gap
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap
	x4 := x3 + bw + gap

	shadowBox(gtx, th, "sig\u2081 pk\u2081", image.Pt(x1, y), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x1+bw, y+bh/2), image.Pt(x2, y+bh/2), lw, lc)
	badge(gtx, th, "1", image.Pt(x2-4, y+bh/2), 8, th.Color.Primary)

	shadowBox(gtx, th, "sig\u2082 pk\u2082", image.Pt(x2, y), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x2+bw, y+bh/2), image.Pt(x3, y+bh/2), lw, lc)
	badge(gtx, th, "2", image.Pt(x3-4, y+bh/2), 8, th.Color.Primary)

	shadowBox(gtx, th, "sig\u2083 pk\u2083", image.Pt(x3, y), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x3+bw, y+bh/2), image.Pt(x4, y+bh/2), lw, lc)
	badge(gtx, th, "3", image.Pt(x4-4, y+bh/2), 8, th.Color.TipBg)

	shadowBox(gtx, th, "\u2265 2 ?", image.Pt(x4, y), bw, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
