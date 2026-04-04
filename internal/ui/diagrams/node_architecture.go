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
// "RPC API" (right) below, connected by curved lines with arrows.
type NodeArchitecture struct{}

func (d *NodeArchitecture) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	topW := pct(w, 42)
	bottomW := pct(w, 30)
	bh := gtx.Dp(unit.Dp(36))
	vertGap := gtx.Dp(unit.Dp(24))

	// Top: Full Node, centered with shadow.
	topX := (w - topW) / 2
	topY := gtx.Dp(unit.Dp(10))
	shadowBox(gtx, th, i18n.T("diagram.full_node"), image.Pt(topX, topY), topW, bh, th.Color.Primary)

	// Bottom row: Wallet (left), RPC API (right).
	bottomY := topY + bh + vertGap
	usable := w - 2*pad
	bottomGap := usable - 2*bottomW
	if bottomGap < 8 {
		bottomGap = 8
	}
	leftX := pad + (usable-2*bottomW-bottomGap)/2
	rightX := leftX + bottomW + bottomGap

	shadowBox(gtx, th, i18n.T("diagram.wallet"), image.Pt(leftX, bottomY), bottomW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, i18n.T("diagram.rpc_api"), image.Pt(rightX, bottomY), bottomW, bh, th.Color.InfoBg)

	// Directional arrows from Full Node to each component.
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	from1 := image.Pt(topX+topW/3, topY+bh)
	to1 := image.Pt(leftX+bottomW/2, bottomY)
	curvedLine(gtx, from1, to1, lw, lc)

	from2 := image.Pt(topX+2*topW/3, topY+bh)
	to2 := image.Pt(rightX+bottomW/2, bottomY)
	curvedLine(gtx, from2, to2, lw, lc)

	// Subtle group box around the Full Node area.
	groupPad := gtx.Dp(unit.Dp(6))
	groupBox(gtx, th, image.Pt(pad-groupPad, topY-groupPad), usable+2*groupPad, h-topY+groupPad-gtx.Dp(unit.Dp(4)))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
