package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TaprootMAST draws a Merkelized Alternative Script Tree for Ch14 L3.
// Three levels: root (Merkle Root), two TapBranch nodes, and three
// script leaves (A, B, C). Script A is highlighted as the spending path.
type TaprootMAST struct{}

func (d *TaprootMAST) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(180))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bh := gtx.Dp(unit.Dp(26))
	lineColor := withAlpha(th.Color.TextMuted, 120)
	dimColor := withAlpha(th.Color.TextMuted, 60)
	lw := float32(1.5)

	usable := w - 2*pad

	// Four rows: root, branches, leaves, caption.
	rowGap := (h - 3*bh - gtx.Dp(unit.Dp(14))) / 4
	rootY := rowGap
	branchY := rootY + bh + rowGap
	leafY := branchY + bh + rowGap

	// --- Root (Merkle Root, Primary color) ---
	rootW := pct(w, 24)
	rootX := pad + (usable-rootW)/2
	rootPos := image.Pt(rootX, rootY)
	box(gtx, th, i18n.T("diagram.merkle_root"), rootPos, rootW, bh, th.Color.Primary)

	// --- Level 1: two TapBranch nodes ---
	// Left TapBranch sits over leaves A and B. Right TapBranch sits over leaf C.
	// Layout: left branch at ~30% center, right branch at ~70% center.
	leftBranchX := pad + pct(usable, 15)
	rightBranchX := pad + pct(usable, 65)
	branchW := pct(w, 20)

	leftBranchPos := image.Pt(leftBranchX, branchY)
	rightBranchPos := image.Pt(rightBranchX, branchY)

	box(gtx, th, i18n.T("diagram.tap_branch"), leftBranchPos, branchW, bh, th.Color.WarningBg)
	box(gtx, th, i18n.T("diagram.tap_branch"), rightBranchPos, branchW, bh, th.Color.WarningBg)

	// --- Leaves ---
	// Script A and Script B are children of left TapBranch.
	// Script C is sibling of left TapBranch (child of right TapBranch).
	leafW := pct(w, 18)

	// Position leaves: A under left side of left branch, B under right side.
	leafAX := pad + pct(usable, 5)
	leafBX := pad + pct(usable, 28)
	leafCX := pad + pct(usable, 65) + (branchW-leafW)/2

	// Script A: highlighted as spending path (TipBg).
	box(gtx, th, i18n.T("diagram.script_leaf")+" A", image.Pt(leafAX, leafY), leafW, bh, th.Color.TipBg)
	// Script B: dimmed sibling hash (InfoBg).
	box(gtx, th, i18n.T("diagram.script_leaf")+" B", image.Pt(leafBX, leafY), leafW, bh, th.Color.InfoBg)
	// Script C: dimmed sibling hash (InfoBg).
	box(gtx, th, i18n.T("diagram.script_leaf")+" C", image.Pt(leafCX, leafY), leafW, bh, th.Color.InfoBg)

	// --- Lines: root to branches ---
	rootBottom := image.Pt(rootX+rootW/2, rootY+bh)
	line(gtx, rootBottom, image.Pt(leftBranchX+branchW/2, branchY), lw, lineColor)
	line(gtx, rootBottom, image.Pt(rightBranchX+branchW/2, branchY), lw, lineColor)

	// --- Lines: left branch to leaves A and B ---
	leftBranchBottom := image.Pt(leftBranchX+branchW/2, branchY+bh)
	// Spending path (A): full opacity line.
	line(gtx, leftBranchBottom, image.Pt(leafAX+leafW/2, leafY), lw, lineColor)
	// Sibling hash (B): dimmed line.
	line(gtx, leftBranchBottom, image.Pt(leafBX+leafW/2, leafY), lw, dimColor)

	// --- Lines: right branch to leaf C ---
	rightBranchBottom := image.Pt(rightBranchX+branchW/2, branchY+bh)
	line(gtx, rightBranchBottom, image.Pt(leafCX+leafW/2, leafY), lw, dimColor)

	// --- Caption: spending path indicator ---
	captionY := leafY + bh + gtx.Dp(unit.Dp(4))
	colorCaption(gtx, th, i18n.T("diagram.spending_path"), image.Pt(leafAX, captionY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
