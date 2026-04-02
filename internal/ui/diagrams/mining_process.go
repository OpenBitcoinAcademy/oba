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
	h := gtx.Dp(unit.Dp(100))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bh := gtx.Dp(unit.Dp(30))
	usable := w - 2*pad

	// Four boxes in a row: Candidate, SHA-256, Hash<Target?, Valid Block.
	// Widths as percentages of usable space.
	candW := pct(usable, 22)
	shaW := pct(usable, 16)
	checkW := pct(usable, 24)
	validW := pct(usable, 22)

	gapTotal := usable - candW - shaW - checkW - validW
	gap := gapTotal / 3
	if gap < 6 {
		gap = 6
	}

	y := (h - bh) / 2

	x1 := pad
	x2 := x1 + candW + gap
	x3 := x2 + shaW + gap
	x4 := x3 + checkW + gap

	// Candidate Block.
	box(gtx, th, i18n.T("diagram.candidate_block"), image.Pt(x1, y), candW, bh, th.Color.InfoBg)

	// SHA-256 process.
	processBox(gtx, th, "SHA-256", image.Pt(x2, y), shaW, bh)

	// Hash < Target? (decision, using Warning color to distinguish).
	box(gtx, th, i18n.T("diagram.hash_below_target"), image.Pt(x3, y), checkW, bh, th.Color.WarningBg)

	// Valid Block.
	box(gtx, th, i18n.T("diagram.valid_block"), image.Pt(x4, y), validW, bh, th.Color.TipBg)

	// Arrows between stages.
	arrow(gtx, th, image.Pt(x1+candW, y+bh/2), image.Pt(x2, y+bh/2))
	arrow(gtx, th, image.Pt(x2+shaW, y+bh/2), image.Pt(x3, y+bh/2))
	arrow(gtx, th, image.Pt(x3+checkW, y+bh/2), image.Pt(x4, y+bh/2))

	// "No" loop: line from bottom of decision box, curving back to candidate block.
	// Draw as two line segments: down from check, then left back to start.
	loopY := y + bh + gtx.Dp(unit.Dp(10))
	checkMidX := x3 + checkW/2
	lineColor := withAlpha(th.Color.Warning, 160)
	loopLw := float32(1.5)

	// Down from decision.
	line(gtx, image.Pt(checkMidX, y+bh), image.Pt(checkMidX, loopY), loopLw, lineColor)
	// Left along the bottom.
	line(gtx, image.Pt(checkMidX, loopY), image.Pt(x1+candW/2, loopY), loopLw, lineColor)
	// Up into candidate box.
	line(gtx, image.Pt(x1+candW/2, loopY), image.Pt(x1+candW/2, y+bh), loopLw, lineColor)

	// "No" / retry label on loop.
	caption(gtx, th, i18n.T("diagram.retry"), image.Pt(x2, loopY-gtx.Dp(unit.Dp(10))))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
