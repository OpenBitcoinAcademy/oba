package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MiningProcess draws the mining flow: Candidate Block -> SHA-256 ->
// Hash < Target? -> (No: loop back, Yes: Valid Block). Ch12 diagram.
type MiningProcess struct{}

func (d *MiningProcess) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bh := gtx.Dp(unit.Dp(30))
	usable := w - 2*pad

	// Four elements in a row.
	candW := pct(usable, 22)
	shaW := pct(usable, 16)
	diamondW := pct(usable, 22)
	diamondH := bh + 6
	validW := pct(usable, 22)

	gapTotal := usable - candW - shaW - diamondW - validW
	gap := gapTotal / 3
	if gap < 6 {
		gap = 6
	}

	y := gtx.Dp(unit.Dp(10))

	x1 := pad
	x2 := x1 + candW + gap
	x3 := x2 + shaW + gap
	x4 := x3 + diamondW + gap

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Candidate Block.
	shadowBox(gtx, th, i18n.T("diagram.candidate_block"), image.Pt(x1, y), candW, bh, th.Color.InfoBg)

	// SHA-256 process.
	shadowBox(gtx, th, "SHA-256", image.Pt(x2, y), shaW, bh, th.Color.Primary)

	// Decision diamond: Hash < Target?
	diamondCenter := image.Pt(x3+diamondW/2, y+diamondH/2)
	diamond(gtx, th, i18n.T("diagram.hash_below_target"), diamondCenter, diamondW, diamondH, th.Color.WarningBg)

	// Valid Block.
	shadowBox(gtx, th, i18n.T("diagram.valid_block"), image.Pt(x4, y), validW, bh, th.Color.TipBg)

	// Arrows between stages.
	dirArrow(gtx, image.Pt(x1+candW, y+bh/2), image.Pt(x2, y+bh/2), lw, lc)
	dirArrow(gtx, image.Pt(x2+shaW, y+bh/2), image.Pt(x3, y+diamondH/2), lw, lc)
	dirArrow(gtx, image.Pt(x3+diamondW, y+diamondH/2), image.Pt(x4, y+bh/2), lw, lc)

	// "No" loop: rounded path from bottom of diamond, back to candidate.
	loopY := y + bh + gtx.Dp(unit.Dp(16))
	loopColor := withAlpha(th.Color.Warning, 160)

	roundedPath(gtx, []image.Point{
		{x3 + diamondW/2, y + diamondH},
		{x3 + diamondW/2, loopY},
		{x1 + candW/2, loopY},
		{x1 + candW/2, y + bh},
	}, 10, 1.8, loopColor)

	// Retry label on loop.
	caption(gtx, th, i18n.T("diagram.retry"), image.Pt(x2, loopY-gtx.Dp(unit.Dp(12))))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
