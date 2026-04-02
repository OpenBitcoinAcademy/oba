package diagrams

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// SecurityLayers draws three nested boxes representing Bitcoin's security
// model: Full Node (verification) > Hardware Wallet (key protection) >
// Seed Backup (recovery). Ch13 diagram.
type SecurityLayers struct{}

func (d *SecurityLayers) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	h := gtx.Dp(unit.Dp(140))
	w := gtx.Constraints.Max.X
	pad := pct(w, 4)

	usable := w - 2*pad
	cx := w / 2

	outerW := pct(usable, 90)
	outerH := pct(h, 80)
	midW := pct(usable, 58)
	midH := pct(h, 50)
	innerW := pct(usable, 30)
	innerH := pct(h, 24)

	// Vertical center of the diagram.
	cy := h / 2

	// Outer layer: Full Node (verification).
	outerX := cx - outerW/2
	outerY := cy - outerH/2
	drawSecurityLayer(gtx, image.Pt(outerX, outerY), outerW, outerH,
		withAlpha(th.Color.Accent, 30), th.Color.Accent)
	colorCaption(gtx, th,
		i18n.T("diagram.full_node")+": "+i18n.T("diagram.verification"),
		image.Pt(outerX+pad/2, outerY+gtx.Dp(unit.Dp(4))),
		th.Color.Accent)

	// Middle layer: Hardware Wallet (key protection).
	midX := cx - midW/2
	midY := cy - midH/2
	drawSecurityLayer(gtx, image.Pt(midX, midY), midW, midH,
		withAlpha(th.Color.Warning, 30), th.Color.Warning)
	colorCaption(gtx, th,
		i18n.T("diagram.hardware_wallet")+": "+i18n.T("diagram.key_protection"),
		image.Pt(midX+pad/2, midY+gtx.Dp(unit.Dp(4))),
		th.Color.Warning)

	// Inner layer: Seed Backup (recovery).
	innerX := cx - innerW/2
	innerY := cy - innerH/2
	drawSecurityLayer(gtx, image.Pt(innerX, innerY), innerW, innerH,
		withAlpha(th.Color.Success, 40), th.Color.Success)
	colorCaption(gtx, th,
		i18n.T("diagram.seed_backup")+": "+i18n.T("diagram.recovery"),
		image.Pt(innerX+pad/2, innerY+gtx.Dp(unit.Dp(4))),
		th.Color.Success)

	return layout.Dimensions{Size: image.Pt(w, h)}
}

// drawSecurityLayer draws a rounded rectangle with a light fill and a colored border.
func drawSecurityLayer(gtx layout.Context, pos image.Point, w, h int, fill, borderColor color.NRGBA) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	r := float32(h) / 6
	rr := clip.RRect{
		Rect: image.Rect(0, 0, w, h),
		SE:   int(r), SW: int(r), NE: int(r), NW: int(r),
	}
	paint.FillShape(gtx.Ops, fill, rr.Op(gtx.Ops))
	paint.FillShape(gtx.Ops, borderColor, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.5}.Op())
}
