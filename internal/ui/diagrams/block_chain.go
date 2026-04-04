package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// BlockChain draws what a block is and how blocks link for Ch11 L1.
// Three blocks linked by prev_hash, each containing header + TXs.
type BlockChain struct{}

func (d *BlockChain) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bw := pct(usable, 26)
	bh := gtx.Dp(unit.Dp(32))
	txH := gtx.Dp(unit.Dp(20))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	gap := (usable - 3*bw) / 2

	y := gtx.Dp(unit.Dp(8))

	for i := 0; i < 3; i++ {
		x := pad + i*(bw+gap)

		// Block header.
		shadowBox(gtx, th, i18n.T("diagram.block_header"), image.Pt(x, y), bw, bh, th.Color.Primary)

		// TX area below header.
		txY := y + bh + 2
		box(gtx, th, "TX, TX, ...", image.Pt(x, txY), bw, txH, th.Color.InfoBg)

		// Group border around block.
		groupBox(gtx, th, image.Pt(x-3, y-3), bw+6, bh+txH+8)

		// Arrow to next block.
		if i < 2 {
			nextX := pad + (i+1)*(bw+gap)
			dirArrow(gtx, image.Pt(x+bw, y+bh/2), image.Pt(nextX, y+bh/2), lw, lc)
		}
	}

	// prev_hash label on arrows.
	labelY := y + bh + txH + gtx.Dp(unit.Dp(12))
	caption(gtx, th, i18n.T("diagram.prev_hash")+" \u2192", image.Pt(pad, labelY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
