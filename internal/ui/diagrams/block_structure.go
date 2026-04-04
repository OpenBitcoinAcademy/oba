package diagrams

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// BlockStructure draws the 80-byte block header as stacked fields
// for Ch11 L2. Six fields from version to nonce, with byte sizes.
type BlockStructure struct{}

func (d *BlockStructure) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(230))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	spacing := 4

	// Header title.
	colorCaption(gtx, th, "80 "+i18n.T("diagram.bytes"), image.Pt(pad, gtx.Dp(unit.Dp(4))), th.Color.Primary)

	fieldsTop := gtx.Dp(unit.Dp(26))
	fieldsH := 6*(bh+spacing) + spacing
	groupBox(gtx, th, image.Pt(pad-4, fieldsTop-4), usable+8, fieldsH+8)

	labelW := pct(usable, 62)
	sizeX := pad + labelW + pct(usable, 4)

	type field struct {
		label string
		bytes string
		bg    color.NRGBA
	}

	fields := []field{
		{i18n.T("diagram.version"), "4", th.Color.InfoBg},
		{i18n.T("diagram.prev_hash"), "32", th.Color.WarningBg},
		{i18n.T("diagram.merkle_root"), "32", th.Color.Primary},
		{i18n.T("diagram.timestamp"), "4", th.Color.InfoBg},
		{i18n.T("diagram.target_bits"), "4", th.Color.InfoBg},
		{i18n.T("diagram.nonce"), "4", th.Color.TipBg},
	}

	y := fieldsTop + spacing
	for _, f := range fields {
		shadowBox(gtx, th, f.label, image.Pt(pad, y), labelW, bh, f.bg)
		caption(gtx, th, f.bytes+" B", image.Pt(sizeX, y+(bh-gtx.Dp(unit.Dp(12)))/2))
		y += bh + spacing
	}

	return layout.Dimensions{Size: image.Pt(w, h)}
}
