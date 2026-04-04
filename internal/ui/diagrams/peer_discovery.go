package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// PeerDiscovery draws the network bootstrap process for Ch10 L2.
// DNS seeds → initial peers → addr gossip → connected mesh.
type PeerDiscovery struct{}

func (d *PeerDiscovery) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(200))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)
	rowGap := gtx.Dp(unit.Dp(28))

	// Row 0: New node.
	y0 := gtx.Dp(unit.Dp(8))
	nodeW := pct(usable, 30)
	nodeX := pad + (usable-nodeW)/2
	shadowBox(gtx, th, i18n.T("diagram.new_node"), image.Pt(nodeX, y0), nodeW, bh, th.Color.WarningBg)

	// Row 1: DNS Seeds.
	y1 := y0 + bh + rowGap
	seedW := pct(usable, 26)
	seedGap := (usable - 3*seedW) / 4
	xs1 := pad + seedGap
	xs2 := xs1 + seedW + seedGap
	xs3 := xs2 + seedW + seedGap

	shadowBox(gtx, th, "DNS Seed", image.Pt(xs1, y1), seedW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "DNS Seed", image.Pt(xs2, y1), seedW, bh, th.Color.InfoBg)
	shadowBox(gtx, th, "DNS Seed", image.Pt(xs3, y1), seedW, bh, th.Color.InfoBg)

	// Arrows from node to seeds.
	curvedLine(gtx, image.Pt(nodeX+nodeW/4, y0+bh), image.Pt(xs1+seedW/2, y1), lw, lc)
	dirArrow(gtx, image.Pt(nodeX+nodeW/2, y0+bh), image.Pt(xs2+seedW/2, y1), lw, lc)
	curvedLine(gtx, image.Pt(nodeX+3*nodeW/4, y0+bh), image.Pt(xs3+seedW/2, y1), lw, lc)

	// Row 2: Connected peers.
	y2 := y1 + bh + rowGap
	peerW := pct(usable, 20)
	peerGap := (usable - 4*peerW) / 5
	peers := make([]int, 4)
	for i := range peers {
		peers[i] = pad + peerGap + i*(peerW+peerGap)
	}

	for _, px := range peers {
		shadowBox(gtx, th, i18n.T("diagram.full_node"), image.Pt(px, y2), peerW, bh, th.Color.TipBg)
	}

	// Arrows from seeds to peers.
	dashedLine(gtx, image.Pt(xs1+seedW/2, y1+bh), image.Pt(peers[0]+peerW/2, y2), lw, lc)
	dashedLine(gtx, image.Pt(xs2+seedW/2, y1+bh), image.Pt(peers[1]+peerW/2, y2), lw, lc)
	dashedLine(gtx, image.Pt(xs2+seedW/2, y1+bh), image.Pt(peers[2]+peerW/2, y2), lw, lc)
	dashedLine(gtx, image.Pt(xs3+seedW/2, y1+bh), image.Pt(peers[3]+peerW/2, y2), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
