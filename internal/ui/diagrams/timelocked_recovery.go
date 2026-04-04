package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// TimelockedRecovery draws a time-decayed spending policy for Ch15 L4.
// Primary key always works. After N blocks, recovery key also works.
type TimelockedRecovery struct{}

func (d *TimelockedRecovery) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	bh := gtx.Dp(unit.Dp(28))
	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Timeline bar.
	barH := gtx.Dp(unit.Dp(8))
	barY := gtx.Dp(unit.Dp(10))
	barW := usable

	// Left portion (primary only).
	primaryW := pct(barW, 55)
	func() {
		shadowBox(gtx, th, "", image.Pt(pad, barY), primaryW, barH, th.Color.Primary)
	}()
	// Right portion (recovery available).
	func() {
		shadowBox(gtx, th, "", image.Pt(pad+primaryW, barY), barW-primaryW, barH, th.Color.TipBg)
	}()

	// Timelock marker.
	markerX := pad + primaryW
	caption(gtx, th, "\u25bc 1000 blocks", image.Pt(markerX-pct(usable, 10), barY+barH+2))

	// Primary key (always).
	y1 := barY + barH + gtx.Dp(unit.Dp(28))
	keyW := pct(usable, 36)
	shadowBox(gtx, th, i18n.T("diagram.primary_key"), image.Pt(pad, y1), keyW, bh, th.Color.Primary)
	colorCaption(gtx, th, i18n.T("diagram.always"), image.Pt(pad+keyW+gtx.Dp(unit.Dp(8)), y1+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.Primary)

	// Recovery key (after timelock).
	y2 := y1 + bh + gtx.Dp(unit.Dp(10))
	shadowBox(gtx, th, i18n.T("diagram.recovery_key"), image.Pt(pad, y2), keyW, bh, th.Color.TipBg)
	colorCaption(gtx, th, "after 1000 blocks", image.Pt(pad+keyW+gtx.Dp(unit.Dp(8)), y2+(bh-gtx.Dp(unit.Dp(12)))/2), th.Color.TextMuted)

	// Arrow from UTXO to both keys.
	utxoX := pad + usable - pct(usable, 24)
	utxoY := y1 + (y2+bh-y1)/2 - bh/2
	shadowBox(gtx, th, "UTXO", image.Pt(utxoX, utxoY), pct(usable, 24), bh, th.Color.WarningBg)
	dashedLine(gtx, image.Pt(pad+keyW, y1+bh/2), image.Pt(utxoX, utxoY+bh/3), lw, lc)
	dashedLine(gtx, image.Pt(pad+keyW, y2+bh/2), image.Pt(utxoX, utxoY+2*bh/3), lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
