package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MiniscriptCompiler draws the three-layer compilation for Ch15 L2.
// Policy (human) → Miniscript (structured) → Bitcoin Script (opcodes).
type MiniscriptCompiler struct{}

func (d *MiniscriptCompiler) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(30))
	bw := pct(usable, 60)
	bx := pad + (usable-bw)/2
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(22))

	y := gtx.Dp(unit.Dp(8))
	shadowBox(gtx, th, i18n.T("diagram.policy")+": and(pk(A),older(1000))", image.Pt(bx, y), bw, bh, th.Color.InfoBg)

	y += bh + rowGap
	dirArrow(gtx, image.Pt(bx+bw/2, y-rowGap), image.Pt(bx+bw/2, y), lw, lc)
	shadowBox(gtx, th, "and_v(v:pk(A),older(1000))", image.Pt(bx, y), bw, bh, th.Color.WarningBg)
	colorCaption(gtx, th, i18n.T("diagram.miniscript"), image.Pt(pad, y+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	y += bh + rowGap
	dirArrow(gtx, image.Pt(bx+bw/2, y-rowGap), image.Pt(bx+bw/2, y), lw, lc)
	shadowBox(gtx, th, "<A> CHECKSIGVERIFY <1000> CSV", image.Pt(bx, y), bw, bh, th.Color.TipBg)
	colorCaption(gtx, th, i18n.T("diagram.bitcoin_script"), image.Pt(pad, y+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
