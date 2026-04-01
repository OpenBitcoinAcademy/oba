package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// CentralizedVsP2P draws Alice -> Bank -> Bob vs Alice -> Network -> Bob.
type CentralizedVsP2P struct{}

func (d *CentralizedVsP2P) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X

	paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: image.Pt(w, h)}.Op())

	bw := gtx.Dp(unit.Dp(80))
	bh := gtx.Dp(unit.Dp(32))
	aw := gtx.Dp(unit.Dp(40))
	rowH := h / 2

	cx := (w - (3*bw + 2*aw)) / 2
	if cx < 0 {
		cx = 4
	}

	// Row 1: Centralized.
	y1 := (rowH - bh) / 2
	caption(gtx, th, "Centralized", image.Pt(cx, y1-16))
	box(gtx, th, "Alice", image.Pt(cx, y1), bw, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(cx+bw, y1+bh/2), aw)
	box(gtx, th, "Bank", image.Pt(cx+bw+aw, y1), bw, bh, th.Color.WarningBg)
	arrow(gtx, th, image.Pt(cx+2*bw+aw, y1+bh/2), aw)
	box(gtx, th, "Bob", image.Pt(cx+2*bw+2*aw, y1), bw, bh, th.Color.InfoBg)

	// Row 2: Peer-to-peer.
	y2 := rowH + (rowH-bh)/2
	caption(gtx, th, "Peer-to-Peer", image.Pt(cx, y2-16))
	box(gtx, th, "Alice", image.Pt(cx, y2), bw, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(cx+bw, y2+bh/2), aw)
	box(gtx, th, "Network", image.Pt(cx+bw+aw, y2), bw, bh, th.Color.TipBg)
	arrow(gtx, th, image.Pt(cx+2*bw+aw, y2+bh/2), aw)
	box(gtx, th, "Bob", image.Pt(cx+2*bw+2*aw, y2), bw, bh, th.Color.InfoBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
