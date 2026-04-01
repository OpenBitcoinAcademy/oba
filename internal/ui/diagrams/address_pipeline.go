package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// AddressPipeline draws the full key-to-address derivation pipeline.
// Private Key -> Public Key -> HASH160 -> Encode -> Address.
type AddressPipeline struct{}

func (d *AddressPipeline) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(80))
	w := gtx.Constraints.Max.X

	paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: image.Pt(w, h)}.Op())

	bw := gtx.Dp(unit.Dp(80))
	bh := gtx.Dp(unit.Dp(32))
	aw := gtx.Dp(unit.Dp(20))

	steps := 5
	totalW := steps*bw + (steps-1)*aw
	cx := (w - totalW) / 2
	if cx < 0 {
		cx = 4
	}
	y := (h - bh) / 2

	labels := []string{"Priv Key", "Pub Key", "HASH160", "Encode", "Address"}
	colors := []func() image.Uniform{} // unused, use direct colors
	bgColors := []struct{ c interface{} }{}
	_ = bgColors

	for i, label := range labels {
		x := cx + i*(bw+aw)
		bg := th.Color.InfoBg
		switch i {
		case 0:
			bg = th.Color.WarningBg
		case 2:
			bg = th.Color.Primary
		case 4:
			bg = th.Color.TipBg
		}
		box(gtx, th, label, image.Pt(x, y), bw, bh, bg)
		if i < len(labels)-1 {
			arrow(gtx, th, image.Pt(x+bw, y+bh/2), aw)
		}
	}
	_ = colors

	return layout.Dimensions{Size: image.Pt(w, h)}
}
