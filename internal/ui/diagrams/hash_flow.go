package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// HashFlow draws "Hello" -> [SHA-256] -> hash output.
type HashFlow struct{}

func (d *HashFlow) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(80))
	w := gtx.Constraints.Max.X

	paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: image.Pt(w, h)}.Op())

	bw := gtx.Dp(unit.Dp(90))
	bh := gtx.Dp(unit.Dp(36))
	aw := gtx.Dp(unit.Dp(30))
	hashW := gtx.Dp(unit.Dp(140))

	cx := (w - (bw + aw + bw + aw + hashW)) / 2
	if cx < 0 {
		cx = 4
	}
	y := (h - bh) / 2

	box(gtx, th, "\"Hello\"", image.Pt(cx, y), bw, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(cx+bw, y+bh/2), aw)
	box(gtx, th, "SHA-256", image.Pt(cx+bw+aw, y), bw, bh, th.Color.Primary)
	arrow(gtx, th, image.Pt(cx+2*bw+aw, y+bh/2), aw)
	box(gtx, th, "185f8db3...", image.Pt(cx+2*bw+2*aw, y), hashW, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
