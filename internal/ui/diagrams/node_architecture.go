package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// NodeArchitecture draws Bitcoin Core's major components for Ch03.
// A primary "Full Node" box on top, with "Wallet" (left) and
// "RPC API" (right) below, connected by lines.
type NodeArchitecture struct{}

func (d *NodeArchitecture) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(120))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	topW := pct(w, 40)
	bottomW := pct(w, 30)
	bh := gtx.Dp(unit.Dp(36))
	vertGap := gtx.Dp(unit.Dp(20))

	// Top row: Full Node, centered.
	topX := (w - topW) / 2
	topY := gtx.Dp(unit.Dp(6))
	box(gtx, th, i18n.T("diagram.full_node"), image.Pt(topX, topY), topW, bh, th.Color.Primary)

	// Bottom row: Wallet (left) and RPC API (right).
	bottomY := topY + bh + vertGap
	usable := w - 2*pad
	bottomGap := usable - 2*bottomW
	if bottomGap < 8 {
		bottomGap = 8
	}
	leftX := pad + (usable-2*bottomW-bottomGap)/2
	rightX := leftX + bottomW + bottomGap

	box(gtx, th, i18n.T("diagram.wallet"), image.Pt(leftX, bottomY), bottomW, bh, th.Color.InfoBg)
	box(gtx, th, i18n.T("diagram.rpc_api"), image.Pt(rightX, bottomY), bottomW, bh, th.Color.InfoBg)

	// Lines from Full Node bottom edge to each component top edge.
	lc := th.Color.TextMuted
	lw := float32(1.5)

	// Full Node center-left to Wallet center-top.
	line(gtx, image.Pt(topX+topW/3, topY+bh), image.Pt(leftX+bottomW/2, bottomY), lw, lc)
	// Full Node center-right to RPC API center-top.
	line(gtx, image.Pt(topX+2*topW/3, topY+bh), image.Pt(rightX+bottomW/2, bottomY), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
