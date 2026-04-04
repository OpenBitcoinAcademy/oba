package diagrams

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// AddressPipeline draws the public-key-to-address derivation pipeline.
// Row 1: Public Key -> SHA-256 -> RIPEMD-160
// Connector curves from RIPEMD-160 down to Add Version.
// Row 2: Add Version -> Base58Check -> Address
type AddressPipeline struct{}

func (d *AddressPipeline) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(220))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)

	bw := pct(w, 28)
	procW := pct(w, 24)
	bh := gtx.Dp(unit.Dp(32))
	usable := w - 2*pad

	lc := withAlpha(th.Color.TextMuted, 160)
	lw := float32(1.8)

	// Row 1: Public Key -> SHA-256 -> RIPEMD-160
	gap1 := (usable - bw - procW - procW) / 2
	if gap1 < 6 {
		gap1 = 6
	}
	y1 := gtx.Dp(unit.Dp(32))
	x1 := pad
	x2 := x1 + bw + gap1
	x3 := x2 + procW + gap1

	colorCaption(gtx, th, i18n.T("diagram.hash_pubkey"), image.Pt(pad, y1-gtx.Dp(unit.Dp(24))), th.Color.TextMuted)
	shadowBox(gtx, th, i18n.T("diagram.public_key"), image.Pt(x1, y1), bw, bh, th.Color.InfoBg)
	dirArrow(gtx, image.Pt(x1+bw, y1+bh/2), image.Pt(x2, y1+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.sha256"), image.Pt(x2, y1), procW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x2+procW, y1+bh/2), image.Pt(x3, y1+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.ripemd160"), image.Pt(x3, y1), procW, bh, th.Color.Primary)

	// Row 2: Add Version -> Base58Check -> Address
	gap2 := (usable - procW - procW - bw) / 2
	if gap2 < 6 {
		gap2 = 6
	}
	y2 := y1 + bh + gtx.Dp(unit.Dp(50))
	x4 := pad
	x5 := x4 + procW + gap2
	x6 := x5 + procW + gap2

	colorCaption(gtx, th, i18n.T("diagram.encode_address"), image.Pt(pad, y2-gtx.Dp(unit.Dp(24))), th.Color.TextMuted)
	shadowBox(gtx, th, i18n.T("diagram.add_version"), image.Pt(x4, y2), procW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x4+procW, y2+bh/2), image.Pt(x5, y2+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.base58check"), image.Pt(x5, y2), procW, bh, th.Color.Primary)
	dirArrow(gtx, image.Pt(x5+procW, y2+bh/2), image.Pt(x6, y2+bh/2), lw, lc)
	shadowBox(gtx, th, i18n.T("diagram.address"), image.Pt(x6, y2), bw, bh, th.Color.TipBg)

	// Curved connector: RIPEMD-160 bottom -> Add Version top.
	fromPt := image.Pt(x3+procW/2, y1+bh)
	toPt := image.Pt(x4+procW/2, y2)
	curvedLine(gtx, fromPt, toPt, lw, lc)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
