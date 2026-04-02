package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// LayerStack draws three stacked horizontal bars for Ch19, widest at
// the bottom: Layer 1 (Bitcoin), Layer 2 (Lightning), Layer 3+ (RGB,
// Fedimint, Ark). Each bar is centered, narrowing toward the top.
type LayerStack struct{}

func (d *LayerStack) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	usable := w - 2*pad
	bh := gtx.Dp(unit.Dp(30))

	// Three bars with gaps between.
	gap := (h - 3*bh) / 4
	if gap < 4 {
		gap = 4
	}

	// Bar widths: bottom 100%, middle 80%, top 60% of usable.
	layer1W := usable
	layer2W := pct(usable, 80)
	layer3W := pct(usable, 60)

	// Y positions: top bar first, bottom bar last.
	y3 := gap
	y2 := y3 + bh + gap
	y1 := y2 + bh + gap

	// Center each bar within usable width.
	x1 := pad + (usable-layer1W)/2
	x2 := pad + (usable-layer2W)/2
	x3 := pad + (usable-layer3W)/2

	// Layer 1: Bitcoin Blockchain (bottom, full width, Primary).
	box(gtx, th, i18n.T("diagram.layer_1"), image.Pt(x1, y1), layer1W, bh, th.Color.Primary)

	// Layer 2: Lightning Network (middle, 80%, WarningBg).
	box(gtx, th, i18n.T("diagram.layer_2"), image.Pt(x2, y2), layer2W, bh, th.Color.WarningBg)

	// Layer 3+: RGB, Fedimint, Ark (top, 60%, InfoBg).
	box(gtx, th, i18n.T("diagram.layer_3"), image.Pt(x3, y3), layer3W, bh, th.Color.InfoBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
