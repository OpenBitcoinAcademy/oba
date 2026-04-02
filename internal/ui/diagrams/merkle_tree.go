package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MerkleTree draws a binary Merkle tree: 4 leaf TX hashes at the bottom,
// 2 intermediate hash pairs in the middle, and 1 root at the top. Ch11 diagram.
type MerkleTree struct{}

func (d *MerkleTree) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(160))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bw := pct(w, 18)
	bh := gtx.Dp(unit.Dp(26))
	lineColor := withAlpha(th.Color.TextMuted, 120)
	lw := float32(1.5)

	usable := w - 2*pad

	// Three rows: root (top), pair (middle), leaves (bottom).
	rowGap := (h - 3*bh) / 4
	rootY := rowGap
	pairY := rootY + bh + rowGap
	leafY := pairY + bh + rowGap

	// --- Leaves (4 TX hashes) ---
	leafGap := (usable - 4*bw) / 5
	leaves := make([]image.Point, 4)
	for i := 0; i < 4; i++ {
		x := pad + leafGap + i*(bw+leafGap)
		leaves[i] = image.Pt(x, leafY)
		box(gtx, th, i18n.T("diagram.tx_hash"), leaves[i], bw, bh, th.Color.InfoBg)
	}

	// --- Pair hashes (2 intermediate nodes) ---
	pairW := pct(w, 20)
	pairGap := (usable - 2*pairW) / 3
	pairs := make([]image.Point, 2)
	for i := 0; i < 2; i++ {
		x := pad + pairGap + i*(pairW+pairGap)
		pairs[i] = image.Pt(x, pairY)
		box(gtx, th, i18n.T("diagram.hash_pair"), pairs[i], pairW, bh, th.Color.WarningBg)
	}

	// --- Root (highlighted in Primary) ---
	rootW := pct(w, 24)
	rootX := pad + (usable-rootW)/2
	rootPos := image.Pt(rootX, rootY)
	box(gtx, th, i18n.T("diagram.merkle_root"), rootPos, rootW, bh, th.Color.Primary)

	// --- Lines: leaves to pairs ---
	// Leaves 0,1 -> Pair 0; Leaves 2,3 -> Pair 1.
	for i := 0; i < 2; i++ {
		pairCenter := image.Pt(pairs[i].X+pairW/2, pairs[i].Y+bh)
		leftLeaf := image.Pt(leaves[2*i].X+bw/2, leaves[2*i].Y)
		rightLeaf := image.Pt(leaves[2*i+1].X+bw/2, leaves[2*i+1].Y)
		line(gtx, pairCenter, leftLeaf, lw, lineColor)
		line(gtx, pairCenter, rightLeaf, lw, lineColor)
	}

	// --- Lines: pairs to root ---
	rootBottom := image.Pt(rootX+rootW/2, rootY+bh)
	line(gtx, rootBottom, image.Pt(pairs[0].X+pairW/2, pairs[0].Y), lw, lineColor)
	line(gtx, rootBottom, image.Pt(pairs[1].X+pairW/2, pairs[1].Y), lw, lineColor)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
