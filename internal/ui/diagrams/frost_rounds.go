package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// FrostRounds draws the detailed two-round FROST protocol for Ch17 L2.
// Round 1: nonce commitments. Round 2: signature shares.
type FrostRounds struct{}

func (d *FrostRounds) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(26))
	nodeW := pct(usable, 18)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Three signers across top.
	y0 := gtx.Dp(unit.Dp(8))
	nodeGap := (usable - 3*nodeW) / 4
	xs := make([]int, 3)
	for i := 0; i < 3; i++ {
		xs[i] = pad + nodeGap + i*(nodeW+nodeGap)
		shadowBox(gtx, th, i18n.T("diagram.signer")+" "+string(rune('A'+i)), image.Pt(xs[i], y0), nodeW, bh, th.Color.InfoBg)
	}

	// Round 1: Nonce commitments (each signer broadcasts).
	y1 := y0 + bh + gtx.Dp(unit.Dp(20))
	groupBox(gtx, th, image.Pt(pad-4, y1-4), usable+8, bh+8)
	colorCaption(gtx, th, i18n.T("diagram.round_1")+": Nonce (R\u1d62)", image.Pt(pad, y1-gtx.Dp(unit.Dp(2))), th.Color.Primary)
	r1Y := y1 + gtx.Dp(unit.Dp(2))
	for i := 0; i < 3; i++ {
		badge(gtx, th, "R", image.Pt(xs[i]+nodeW/2, r1Y+bh/2), 10, th.Color.Primary)
		dirArrow(gtx, image.Pt(xs[i]+nodeW/2, y0+bh), image.Pt(xs[i]+nodeW/2, r1Y), lw, lc)
	}

	// Round 2: Signature shares.
	y2 := r1Y + bh + gtx.Dp(unit.Dp(20))
	groupBox(gtx, th, image.Pt(pad-4, y2-4), usable+8, bh+8)
	colorCaption(gtx, th, i18n.T("diagram.round_2")+": Share (z\u1d62)", image.Pt(pad, y2-gtx.Dp(unit.Dp(2))), th.Color.Primary)
	r2Y := y2 + gtx.Dp(unit.Dp(2))
	for i := 0; i < 3; i++ {
		badge(gtx, th, "z", image.Pt(xs[i]+nodeW/2, r2Y+bh/2), 10, th.Color.WarningBg)
	}

	// Final signature.
	y3 := r2Y + bh + gtx.Dp(unit.Dp(20))
	sigW := pct(usable, 36)
	sigX := pad + (usable-sigW)/2
	shadowBox(gtx, th, i18n.T("diagram.signature")+" (R, z)", image.Pt(sigX, y3), sigW, bh, th.Color.TipBg)
	curvedLine(gtx, image.Pt(xs[1]+nodeW/2, r2Y+bh), image.Pt(sigX+sigW/2, y3), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
