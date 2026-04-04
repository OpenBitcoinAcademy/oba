package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// LightningRouting draws a multi-hop Lightning payment path for
// Ch18 L3. Alice → Node1 → Node2 → Bob with onion layers.
type LightningRouting struct{}

func (d *LightningRouting) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(30))
	nodeW := pct(usable, 18)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Four nodes in a row.
	gap := (usable - 4*nodeW) / 3
	if gap < 6 {
		gap = 6
	}
	y := (h - bh) / 2

	xA := pad
	x1 := xA + nodeW + gap
	x2 := x1 + nodeW + gap
	xB := x2 + nodeW + gap

	shadowBox(gtx, th, "Alice", image.Pt(xA, y), nodeW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "Node 1", image.Pt(x1, y), nodeW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "Node 2", image.Pt(x2, y), nodeW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "Bob", image.Pt(xB, y), nodeW, bh, th.Color.TipBg)

	// Payment arrows.
	dirArrow(gtx, image.Pt(xA+nodeW, y+bh/2), image.Pt(x1, y+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(x1+nodeW, y+bh/2), image.Pt(x2, y+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(x2+nodeW, y+bh/2), image.Pt(xB, y+bh/2), lw, lc)

	// Onion layers (concentric arcs above the path, visualized as badges).
	arcY := y - gtx.Dp(unit.Dp(14))
	badge(gtx, th, "3", image.Pt(xA+nodeW+gap/2, arcY), 9, withAlpha(th.Color.Primary, 200))
	badge(gtx, th, "2", image.Pt(x1+nodeW+gap/2, arcY), 9, withAlpha(th.Color.Primary, 150))
	badge(gtx, th, "1", image.Pt(x2+nodeW+gap/2, arcY), 9, withAlpha(th.Color.Primary, 100))

	// Caption.
	captionY := y + bh + gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "Onion layers \u2192", image.Pt(pad, captionY), th.Color.TextMuted)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
