package diagrams

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// WireFormat draws the transaction serialization layout for Ch06 L3.
// Legacy vs SegWit field ordering comparison.
type WireFormat struct{}

func (d *WireFormat) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(130))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))

	// Row 0: Legacy format.
	y0 := gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, "Legacy", image.Pt(pad, y0), th.Color.TextMuted)
	y0 += gtx.Dp(unit.Dp(22))

	fw := pct(usable, 22)
	gap := (usable - 4*fw) / 5
	x1 := pad + gap

	legacyFields := []struct {
		label string
		bg    color.NRGBA
	}{
		{i18n.T("diagram.version"), th.Color.InfoBg},
		{i18n.T("diagram.tx_input") + "s", th.Color.WarningBg},
		{i18n.T("diagram.tx_output") + "s", th.Color.TipBg},
		{i18n.T("diagram.tx_locktime"), th.Color.InfoBg},
	}
	for i, f := range legacyFields {
		shadowBox(gtx, th, f.label, image.Pt(x1+i*(fw+gap), y0), fw, bh, f.bg)
	}

	// Row 1: SegWit format.
	y1 := y0 + bh + gtx.Dp(unit.Dp(16))
	colorCaption(gtx, th, "SegWit", image.Pt(pad, y1), th.Color.Primary)
	y1 += gtx.Dp(unit.Dp(22))

	sw := pct(usable, 14)
	sgap := (usable - 6*sw) / 7

	segwitFields := []struct {
		label string
		bg    color.NRGBA
	}{
		{i18n.T("diagram.version"), th.Color.InfoBg},
		{"0x00", th.Color.Primary},
		{"0x01", th.Color.Primary},
		{i18n.T("diagram.tx_input") + "s", th.Color.WarningBg},
		{i18n.T("diagram.tx_output") + "s", th.Color.TipBg},
		{i18n.T("diagram.witness_data"), th.Color.Primary},
	}
	for i, f := range segwitFields {
		shadowBox(gtx, th, f.label, image.Pt(pad+sgap+i*(sw+sgap), y1), sw, bh, f.bg)
	}

	return layout.Dimensions{Size: image.Pt(w, h)}
}
