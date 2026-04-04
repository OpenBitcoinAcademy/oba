package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// OutputDescriptors draws an annotated descriptor string for Ch15 L3.
type OutputDescriptors struct{}

func (d *OutputDescriptors) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))

	// Full descriptor string.
	y := gtx.Dp(unit.Dp(8))
	shadowBox(gtx, th, "wpkh([d34db33f/84'/0'/0']xpub.../0/*)", image.Pt(pad, y), usable, bh, th.Color.InfoBg)

	// Annotated parts below.
	y += bh + gtx.Dp(unit.Dp(16))
	partW := pct(usable, 22)
	gap := (usable - 4*partW) / 5

	x1 := pad + gap
	x2 := x1 + partW + gap
	x3 := x2 + partW + gap
	x4 := x3 + partW + gap

	shadowBox(gtx, th, "wpkh()", image.Pt(x1, y), partW, bh, th.Color.Primary)
	shadowBox(gtx, th, "[fingerprint]", image.Pt(x2, y), partW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, "xpub...", image.Pt(x3, y), partW, bh, th.Color.TipBg)
	shadowBox(gtx, th, "/0/*", image.Pt(x4, y), partW, bh, th.Color.InfoBg)

	// Labels.
	labelY := y + bh + gtx.Dp(unit.Dp(4))
	caption(gtx, th, "Script type", image.Pt(x1, labelY))
	caption(gtx, th, "Origin", image.Pt(x2, labelY))
	caption(gtx, th, "Key", image.Pt(x3, labelY))
	caption(gtx, th, "Derivation", image.Pt(x4, labelY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
