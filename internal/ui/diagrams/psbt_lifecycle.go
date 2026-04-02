package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// PSBTLifecycle draws a horizontal five-stage PSBT workflow for Ch16:
// Creator -> Updater -> Signer -> Finalizer -> Broadcast.
// Each stage is a processBox with arrows between them.
type PSBTLifecycle struct{}

func (d *PSBTLifecycle) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(80))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bw := pct(w, 16)
	bh := gtx.Dp(unit.Dp(34))

	usable := w - 2*pad
	gapTotal := usable - 5*bw
	gap := gapTotal / 4
	if gap < 4 {
		gap = 4
	}

	y := (h - bh) / 2

	x1 := pad
	x2 := x1 + bw + gap
	x3 := x2 + bw + gap
	x4 := x3 + bw + gap
	x5 := x4 + bw + gap

	// Creator.
	processBox(gtx, th, i18n.T("diagram.creator"), image.Pt(x1, y), bw, bh)
	arrow(gtx, th, image.Pt(x1+bw, y+bh/2), image.Pt(x2, y+bh/2))

	// Updater.
	processBox(gtx, th, i18n.T("diagram.updater"), image.Pt(x2, y), bw, bh)
	arrow(gtx, th, image.Pt(x2+bw, y+bh/2), image.Pt(x3, y+bh/2))

	// Signer.
	processBox(gtx, th, i18n.T("diagram.signer"), image.Pt(x3, y), bw, bh)
	arrow(gtx, th, image.Pt(x3+bw, y+bh/2), image.Pt(x4, y+bh/2))

	// Finalizer.
	processBox(gtx, th, i18n.T("diagram.finalizer"), image.Pt(x4, y), bw, bh)
	arrow(gtx, th, image.Pt(x4+bw, y+bh/2), image.Pt(x5, y+bh/2))

	// Broadcast.
	processBox(gtx, th, i18n.T("diagram.broadcast"), image.Pt(x5, y), bw, bh)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
