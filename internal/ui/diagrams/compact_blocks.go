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

// CompactBlocks draws a size comparison between a full block and
// a compact block for Ch10 L3.
type CompactBlocks struct{}

func (d *CompactBlocks) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 3)
	usable := w - 2*pad

	barH := gtx.Dp(unit.Dp(32))
	rowGap := gtx.Dp(unit.Dp(22))

	// Full block bar (100% width).
	y1 := gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, i18n.T("diagram.full_block"), image.Pt(pad, y1), th.Color.TextMuted)
	y1 += gtx.Dp(unit.Dp(24))
	func() {
		defer op.Offset(image.Pt(pad, y1)).Push(gtx.Ops).Pop()
		r := float32(barH) / 4
		ri := int(r)
		rr := clip.RRect{Rect: image.Rect(0, 0, usable, barH), SE: ri, SW: ri, NE: ri, NW: ri}
		paint.FillShape(gtx.Ops, th.Color.WarningBg, rr.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())
	}()
	colorCaption(gtx, th, "~1-2 MB", image.Pt(pad+usable/2-pct(usable, 5), y1+(barH-gtx.Dp(unit.Dp(12)))/2), th.Color.Text)

	// Compact block bar (~6% width = header + short IDs).
	y2 := y1 + barH + rowGap
	colorCaption(gtx, th, i18n.T("diagram.compact_block"), image.Pt(pad, y2-gtx.Dp(unit.Dp(16))), th.Color.TextMuted)
	compactW := pct(usable, 15)
	func() {
		defer op.Offset(image.Pt(pad, y2)).Push(gtx.Ops).Pop()
		r := float32(barH) / 4
		ri := int(r)
		rr := clip.RRect{Rect: image.Rect(0, 0, compactW, barH), SE: ri, SW: ri, NE: ri, NW: ri}
		paint.FillShape(gtx.Ops, th.Color.TipBg, rr.Op(gtx.Ops))
		paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())
	}()
	colorCaption(gtx, th, "~20 KB", image.Pt(pad+compactW+gtx.Dp(unit.Dp(8)), y2+(barH-gtx.Dp(unit.Dp(12)))/2), th.Color.Text)

	// Savings label.
	savingsY := y2 + barH + gtx.Dp(unit.Dp(8))
	colorCaption(gtx, th, "~99% "+i18n.T("diagram.savings"), image.Pt(pad, savingsY), th.Color.Primary)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
