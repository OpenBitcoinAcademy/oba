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
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bh := gtx.Dp(unit.Dp(28))
	activeColor := withAlpha(th.Color.Primary, 180)
	dimColor := withAlpha(th.Color.TextMuted, 50)
	activeLine := float32(2.0)
	dimLine := float32(1.2)

	usable := w - 2*pad

	// Four rows: root, branches, leaves, caption.
	rowGap := (h - 3*bh - gtx.Dp(unit.Dp(18))) / 4
	rootY := rowGap
	branchY := rootY + bh + rowGap
	leafY := branchY + bh + rowGap

	// Root (Merkle Root).
	rootW := pct(w, 26)
	rootX := pad + (usable-rootW)/2
	shadowBox(gtx, th, i18n.T("diagram.merkle_root"), image.Pt(rootX, rootY), rootW, bh, th.Color.Primary)

	// Two TapBranch nodes.
	branchW := pct(w, 22)
	leftBranchX := pad + pct(usable, 13)
	rightBranchX := pad + pct(usable, 63)

	shadowBox(gtx, th, i18n.T("diagram.tap_branch"), image.Pt(leftBranchX, branchY), branchW, bh, th.Color.WarningBg)
	shadowBox(gtx, th, i18n.T("diagram.tap_branch"), image.Pt(rightBranchX, branchY), branchW, bh, th.Color.WarningBg)

	// Three script leaves.
	leafW := pct(w, 19)
	leafAX := pad + pct(usable, 3)
	leafBX := pad + pct(usable, 26)
	leafCX := pad + pct(usable, 63) + (branchW-leafW)/2

	// Script A: highlighted as spending path.
	shadowBox(gtx, th, i18n.T("diagram.script_leaf")+" A", image.Pt(leafAX, leafY), leafW, bh, th.Color.TipBg)
	// Scripts B and C: dimmed.
	box(gtx, th, i18n.T("diagram.script_leaf")+" B", image.Pt(leafBX, leafY), leafW, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.script_leaf")+" C", image.Pt(leafCX, leafY), leafW, bh, th.Color.InfoBg)

	// Curved lines: root to branches.
	rootBottom := image.Pt(rootX+rootW/2, rootY+bh)
	curvedLine(gtx, rootBottom, image.Pt(leftBranchX+branchW/2, branchY), activeLine, activeColor)
	curvedLine(gtx, rootBottom, image.Pt(rightBranchX+branchW/2, branchY), dimLine, dimColor)

	// Curved lines: left branch to leaves A and B.
	leftBottom := image.Pt(leftBranchX+branchW/2, branchY+bh)
	curvedLine(gtx, leftBottom, image.Pt(leafAX+leafW/2, leafY), activeLine, activeColor)
	dashedLine(gtx, leftBottom, image.Pt(leafBX+leafW/2, leafY), dimLine, dimColor)

	// Dashed line: right branch to leaf C.
	rightBottom := image.Pt(rightBranchX+branchW/2, branchY+bh)
	dashedLine(gtx, rightBottom, image.Pt(leafCX+leafW/2, leafY), dimLine, dimColor)

	// Spending path indicator.
	captionY := leafY + bh + gtx.Dp(unit.Dp(6))
	colorCaption(gtx, th, i18n.T("diagram.spending_path"), image.Pt(leafAX, captionY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
