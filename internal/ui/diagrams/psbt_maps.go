package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// PSBTMaps draws the three-map PSBT structure for Ch16 L3.
// Global map, per-input maps, per-output maps.
type PSBTMaps struct{}

func (d *PSBTMaps) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Global map.
	y0 := gtx.Dp(unit.Dp(8))
	globalW := pct(usable, 40)
	globalX := pad + (usable-globalW)/2
	shadowBox(gtx, th, i18n.T("diagram.global_map"), image.Pt(globalX, y0), globalW, bh, th.Color.Primary)

	// Input maps.
	y1 := y0 + bh + gtx.Dp(unit.Dp(26))
	mapW := pct(usable, 26)
	mapGap := (usable - 3*mapW) / 4
	for i := 0; i < 3; i++ {
		x := pad + mapGap + i*(mapW+mapGap)
		label := "Input " + string(rune('0'+i))
		shadowBox(gtx, th, label, image.Pt(x, y1), mapW, bh, th.Color.WarningBg)
		curvedLine(gtx, image.Pt(globalX+globalW/2, y0+bh), image.Pt(x+mapW/2, y1), lw, lc)
	}

	// Output maps.
	y2 := y1 + bh + gtx.Dp(unit.Dp(26))
	for i := 0; i < 2; i++ {
		outGap := (usable - 2*mapW) / 3
		x := pad + outGap + i*(mapW+outGap)
		label := "Output " + string(rune('0'+i))
		shadowBox(gtx, th, label, image.Pt(x, y2), mapW, bh, th.Color.TipBg)
		curvedLine(gtx, image.Pt(globalX+globalW/2, y0+bh), image.Pt(x+mapW/2, y2), lw, withAlpha(lc, 100))
	}

	return layout.Dimensions{Size: image.Pt(w, h)}
}
