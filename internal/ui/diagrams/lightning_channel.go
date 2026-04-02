package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// LightningChannel draws a horizontal capacity bar for Ch18, split
// into Alice's and Bob's balances. Above: channel capacity caption.
// Below: commitment number caption.
type LightningChannel struct{}

func (d *LightningChannel) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(100))
	w := gtx.Constraints.Max.X
	pad := pct(w, 6)

	usable := w - 2*pad
	barH := gtx.Dp(unit.Dp(36))
	r := float32(barH) / 6

	// Vertical positioning.
	captionAboveY := gtx.Dp(unit.Dp(8))
	barY := gtx.Dp(unit.Dp(30))
	captionBelowY := barY + barH + gtx.Dp(unit.Dp(8))

	// Balance split: Alice 70%, Bob 30%.
	aliceW := pct(usable, 70)
	bobW := usable - aliceW

	// --- Caption above: Channel Capacity ---
	caption(gtx, th, i18n.T("diagram.channel_capacity")+": 1.0 BTC", image.Pt(pad, captionAboveY))

	// --- Alice's portion (left, InfoBg) ---
	func() {
		defer op.Offset(image.Pt(pad, barY)).Push(gtx.Ops).Pop()
		rr := clip.RRect{
			Rect: image.Rect(0, 0, aliceW, barH),
			NW:   int(r), SW: int(r),
		}
		paint.FillShape(gtx.Ops, th.Color.InfoBg, rr.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())
	}()

	// Alice label, centered in her portion.
	aliceLabelX := pad + aliceW/4
	colorCaption(gtx, th, "Alice: 0.7 BTC", image.Pt(aliceLabelX, barY+(barH-gtx.Dp(unit.Dp(14)))/2), th.Color.Text)

	// --- Bob's portion (right, TipBg) ---
	func() {
		defer op.Offset(image.Pt(pad+aliceW, barY)).Push(gtx.Ops).Pop()
		rr := clip.RRect{
			Rect: image.Rect(0, 0, bobW, barH),
			NE:   int(r), SE: int(r),
		}
		paint.FillShape(gtx.Ops, th.Color.TipBg, rr.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())
	}()

	// Bob label, centered in his portion.
	bobLabelX := pad + aliceW + bobW/6
	colorCaption(gtx, th, "Bob: 0.3 BTC", image.Pt(bobLabelX, barY+(barH-gtx.Dp(unit.Dp(14)))/2), th.Color.Text)

	// --- Caption below: Commitment number ---
	caption(gtx, th, i18n.T("diagram.commitment_num")+" #47", image.Pt(pad, captionBelowY))

	return layout.Dimensions{Size: image.Pt(w, h)}
}
