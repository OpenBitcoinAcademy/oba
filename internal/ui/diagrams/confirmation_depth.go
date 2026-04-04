package diagrams

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ConfirmationDepth draws a chain of blocks with a highlighted
// transaction showing increasing confirmation depth. Ch02 L3.
type ConfirmationDepth struct{}

func (d *ConfirmationDepth) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bw := pct(usable, 14)
	bh := gtx.Dp(unit.Dp(32))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	blocks := 6
	gapTotal := usable - blocks*bw
	gap := gapTotal / (blocks - 1)
	if gap < 4 {
		gap = 4
	}

	y := gtx.Dp(unit.Dp(8))

	// Draw 6 blocks, leftmost contains "your TX".
	for i := 0; i < blocks; i++ {
		x := pad + i*(bw+gap)
		label := fmt.Sprintf("%s %d", i18n.T("diagram.block"), i+1)
		alpha := uint8(255 - i*30)
		if i == 0 {
			// The block containing the transaction.
			shadowBox(gtx, th, "TX", image.Pt(x, y), bw, bh, th.Color.Primary)
		} else {
			shadowBox(gtx, th, label, image.Pt(x, y), bw, bh, withAlpha(th.Color.InfoBg, alpha))
		}

		// Arrow to next block.
		if i < blocks-1 {
			nextX := pad + (i+1)*(bw+gap)
			dirArrow(gtx, image.Pt(x+bw, y+bh/2), image.Pt(nextX, y+bh/2), lw, lc)
		}
	}

	// Confirmation labels below.
	labelY := y + bh + gtx.Dp(unit.Dp(10))
	colorCaption(gtx, th, "1 conf", image.Pt(pad+bw+gap+bw/4, labelY), th.Color.TextMuted)
	colorCaption(gtx, th, "3 conf", image.Pt(pad+3*(bw+gap)+bw/4, labelY), th.Color.TextMuted)
	colorCaption(gtx, th, "6 conf", image.Pt(pad+5*(bw+gap)+bw/4, labelY), th.Color.Primary)

	// Arrow showing depth direction.
	arrowY := labelY + gtx.Dp(unit.Dp(16))
	caption(gtx, th, i18n.T("diagram.confirmation_depth")+" \u2192", image.Pt(pad, arrowY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
