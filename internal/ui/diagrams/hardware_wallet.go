package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// HardwareWallet draws the watch-only + hardware signing workflow
// for Ch16 L4. PSBT flows from watch-only wallet to hardware and back.
type HardwareWallet struct{}

func (d *HardwareWallet) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	bw := pct(usable, 30)
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Watch-only wallet (left).
	y := (h - bh*2 - gtx.Dp(unit.Dp(20))) / 2
	leftX := pad
	shadowBox(gtx, th, i18n.T("diagram.watch_only"), image.Pt(leftX, y), bw, bh, th.Color.InfoBg)

	// Hardware device (right).
	rightX := pad + usable - bw
	shadowBox(gtx, th, i18n.T("diagram.hardware"), image.Pt(rightX, y), bw, bh, th.Color.WarningBg)

	// PSBT arrow: watch-only → hardware (unsigned).
	midY := y + bh/2
	psbW := pct(usable, 24)
	psbX := pad + (usable-psbW)/2
	dirArrow(gtx, image.Pt(leftX+bw, midY-4), image.Pt(rightX, midY-4), lw, lc)
	caption(gtx, th, "PSBT (unsigned)", image.Pt(psbX, midY-gtx.Dp(unit.Dp(16))))

	// Signed PSBT arrow: hardware → watch-only (return).
	y2 := y + bh + gtx.Dp(unit.Dp(20))
	shadowBox(gtx, th, i18n.T("diagram.broadcast"), image.Pt(leftX, y2), bw, bh, th.Color.TipBg)
	dirArrow(gtx, image.Pt(rightX, y2+bh/2), image.Pt(leftX+bw, y2+bh/2), lw, lc)
	caption(gtx, th, "PSBT (signed)", image.Pt(psbX, y2+bh+gtx.Dp(unit.Dp(4))))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
