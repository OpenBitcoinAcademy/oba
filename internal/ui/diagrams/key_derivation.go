package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// KeyDerivation draws Private Key --(one-way)--> Public Key.
// The middle shows the operation as a process step, with a "one-way"
// label on the arrow to emphasize irreversibility.
type KeyDerivation struct{}

func (d *KeyDerivation) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(110))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	// No background fill: inherits page background.

	// Responsive: two endpoint boxes (28%), one process box (26%).
	endW := pct(w, 28)
	procW := pct(w, 26)
	bh := gtx.Dp(unit.Dp(38))

	usable := w - 2*pad
	gapTotal := usable - 2*endW - procW
	gap := gapTotal / 2
	if gap < 8 {
		gap = 8
	}

	y := h/2 - bh/2 + gtx.Dp(unit.Dp(4))
	x1 := pad
	x2 := x1 + endW + gap
	x3 := x2 + procW + gap

	lc := withAlpha(th.Color.TextMuted, 160)
	shadowBox(gtx, th, i18n.T("diagram.private_key"), image.Pt(x1, y), endW, bh, th.Color.WarningBg)
	dirArrow(gtx, image.Pt(x1+endW, y+bh/2), image.Pt(x2, y+bh/2), 1.8, lc)
	shadowBox(gtx, th, i18n.T("diagram.elliptic_curve"), image.Pt(x2, y), procW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x2+procW, y+bh/2), image.Pt(x3, y+bh/2), 1.8, lc)
	shadowBox(gtx, th, i18n.T("diagram.public_key"), image.Pt(x3, y), endW, bh, th.Color.TipBg)

	captionX := x1 + endW + (x3-x1-endW)/2 - pct(w, 6)
	captionY := y - gtx.Dp(unit.Dp(16))
	colorCaption(gtx, th, i18n.T("diagram.one_way"), image.Pt(captionX, captionY), th.Color.TextMuted)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
