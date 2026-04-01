package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// AddressPipeline draws the public-key-to-address derivation pipeline.
// Starts from Public Key (not Private Key, since that's covered in
// key_derivation). Shows both hash steps separately:
//
//	Public Key -> SHA-256 -> RIPEMD-160 -> Base58Check -> Address
//
// This matches the lesson text which explains each step.
type AddressPipeline struct{}

func (d *AddressPipeline) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(190))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	// No background fill: inherits page background.

	// Two rows to avoid overflow on narrow screens.
	// Row 1: Public Key -> SHA-256 -> RIPEMD-160
	// Row 2: Version + Checksum -> Base58Check -> Address

	bw := pct(w, 28)
	procW := pct(w, 24)
	bh := gtx.Dp(unit.Dp(32))
	rowH := h / 2

	usable := w - 2*pad

	// Row 1.
	gap1 := (usable - bw - procW - procW) / 2
	if gap1 < 6 {
		gap1 = 6
	}
	y1 := rowH/2 - bh/2
	x1 := pad
	x2 := x1 + bw + gap1
	x3 := x2 + procW + gap1

	colorCaption(gtx, th, "Hash the public key", image.Pt(pad, y1-gtx.Dp(unit.Dp(14))), th.Color.TextMuted)
	box(gtx, th, "Public Key", image.Pt(x1, y1), bw, bh, th.Color.InfoBg)
	arrow(gtx, th, image.Pt(x1+bw, y1+bh/2), image.Pt(x2, y1+bh/2))
	processBox(gtx, th, "SHA-256", image.Pt(x2, y1), procW, bh)
	arrow(gtx, th, image.Pt(x2+procW, y1+bh/2), image.Pt(x3, y1+bh/2))
	processBox(gtx, th, "RIPEMD-160", image.Pt(x3, y1), procW, bh)

	// Connecting arrow between rows (vertical).
	midX := x3 + procW/2
	line(gtx, image.Pt(midX, y1+bh), image.Pt(midX, rowH+y1), 1.5, th.Color.TextMuted)

	// Row 2.
	gap2 := (usable - procW - procW - bw) / 2
	if gap2 < 6 {
		gap2 = 6
	}
	y2 := rowH + rowH/2 - bh/2
	x4 := pad
	x5 := x4 + procW + gap2
	x6 := x5 + procW + gap2

	colorCaption(gtx, th, "Encode the address", image.Pt(pad, y2-gtx.Dp(unit.Dp(14))), th.Color.TextMuted)
	processBox(gtx, th, "Add Version", image.Pt(x4, y2), procW, bh)
	arrow(gtx, th, image.Pt(x4+procW, y2+bh/2), image.Pt(x5, y2+bh/2))
	processBox(gtx, th, "Base58Check", image.Pt(x5, y2), procW, bh)
	arrow(gtx, th, image.Pt(x5+procW, y2+bh/2), image.Pt(x6, y2+bh/2))
	box(gtx, th, "Address", image.Pt(x6, y2), bw, bh, th.Color.TipBg)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
