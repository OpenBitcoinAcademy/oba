package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// NonceReuse draws the nonce reuse attack for Ch17 L4.
// Two signatures with the same nonce → private key extracted.
type NonceReuse struct{}

func (d *NonceReuse) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Two signatures with same R.
	y0 := gtx.Dp(unit.Dp(8))
	sigW := pct(usable, 36)
	sigGap := (usable - 2*sigW) / 3
	s1X := pad + sigGap
	s2X := s1X + sigW + sigGap

	shadowBox(gtx, th, "sig\u2081 = (R, z\u2081)", image.Pt(s1X, y0), sigW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "sig\u2082 = (R, z\u2082)", image.Pt(s2X, y0), sigW, bh, th.Color.InfoBg)

	// Same R highlighted.
	badge(gtx, th, "R", image.Pt(s1X-6, y0+bh/2), 10, th.Color.Warning)
	badge(gtx, th, "R", image.Pt(s2X-6, y0+bh/2), 10, th.Color.Warning)

	// Arrow down: attacker computes.
	y1 := y0 + bh + gtx.Dp(unit.Dp(24))
	compW := pct(usable, 50)
	compX := pad + (usable-compW)/2
	shadowBox(gtx, th, "d = (z\u2081\u2212z\u2082) / (m\u2081\u2212m\u2082)", image.Pt(compX, y1), compW, bh, th.Color.WarningBg)

	curvedLine(gtx, image.Pt(s1X+sigW/2, y0+bh), image.Pt(compX+compW/3, y1), lw, lc)
	curvedLine(gtx, image.Pt(s2X+sigW/2, y0+bh), image.Pt(compX+2*compW/3, y1), lw, lc)

	// Result: private key leaked.
	y2 := y1 + bh + gtx.Dp(unit.Dp(24))
	leakW := pct(usable, 36)
	leakX := pad + (usable-leakW)/2
	shadowBox(gtx, th, i18n.T("diagram.private_key")+" d", image.Pt(leakX, y2), leakW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(compX+compW/2, y1+bh), image.Pt(leakX+leakW/2, y2), lw, withAlpha(th.Color.Warning, 200))

	// Warning.
	colorCaption(gtx, th, "\u26a0 "+i18n.T("diagram.key_compromised"), image.Pt(leakX, y2+bh+gtx.Dp(unit.Dp(4))), th.Color.Warning)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
