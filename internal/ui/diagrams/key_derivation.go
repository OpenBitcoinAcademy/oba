package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// KeyDerivation draws Random Number -> [EC Multiply] -> Public Key.
type KeyDerivation struct{}

func (d *KeyDerivation) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(80))
	w := gtx.Constraints.Max.X

	paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: image.Pt(w, h)}.Op())

	bw := gtx.Dp(unit.Dp(100))
	bh := gtx.Dp(unit.Dp(36))
	aw := gtx.Dp(unit.Dp(30))

	cx := (w - (3*bw + 2*aw)) / 2
	if cx < 0 {
		cx = 4
	}
	y := (h - bh) / 2

	box(gtx, th, "Private Key", image.Pt(cx, y), bw, bh, th.Color.WarningBg)
	arrow(gtx, th, image.Pt(cx+bw, y+bh/2), aw)
	box(gtx, th, "EC Multiply", image.Pt(cx+bw+aw, y), bw, bh, th.Color.Primary)
	arrow(gtx, th, image.Pt(cx+2*bw+aw, y+bh/2), aw)
	box(gtx, th, "Public Key", image.Pt(cx+2*bw+2*aw, y), bw, bh, th.Color.TipBg)

	caption(gtx, th, "one-way", image.Pt(cx+bw+aw/2, y-16))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
