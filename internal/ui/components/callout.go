package components

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Callout renders a highlighted text box with a type-distinguishing icon.
type Callout struct {
	Content string
	Style   string // "info", "warning", "tip"
	Theme   *theme.Theme
}

// Layout renders the callout.
func (c *Callout) Layout(gtx layout.Context) layout.Dimensions {
	bg := c.Theme.Color.InfoBg // All types use the same overlay

	return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{}.Layout(gtx,
				// Background with rounded corners.
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					r := gtx.Dp(unit.Dp(6))
					rect := clip.RRect{
						Rect: image.Rectangle{Max: gtx.Constraints.Min},
						NE:   r, NW: r, SE: r, SW: r,
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, bg, rect)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				// Icon + content.
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(12), Bottom: unit.Dp(12),
						Left: unit.Dp(12), Right: unit.Dp(12),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{Right: unit.Dp(10), Top: unit.Dp(2)}.Layout(gtx,
									func(gtx layout.Context) layout.Dimensions {
										return c.drawIcon(gtx)
									},
								)
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								md := &Markdown{Content: c.Content, Theme: c.Theme}
								return md.Layout(gtx)
							}),
						)
					})
				}),
			)
		},
	)
}

func (c *Callout) drawIcon(gtx layout.Context) layout.Dimensions {
	size := gtx.Dp(unit.Dp(18))
	col := c.Theme.Color.TextMuted

	switch c.Style {
	case "warning":
		drawWarningTriangle(gtx.Ops, size, col)
	case "tip":
		drawLightbulb(gtx.Ops, size, col)
	default:
		drawInfoCircle(gtx.Ops, size, col)
	}
	return layout.Dimensions{Size: image.Pt(size, size)}
}

// drawInfoCircle draws a circle with "i" (dot + line).
func drawInfoCircle(ops *op.Ops, size int, col color.NRGBA) {
	cx := float32(size) / 2
	cy := float32(size) / 2
	r := float32(size) / 2

	drawCircleOutline(ops, cx, cy, r, 1.5, col)

	// Dot.
	dotR := float32(size) / 8
	drawFilledCircle(ops, cx, cy-r*0.3, dotR, col)

	// Line.
	lw := float32(size) / 6
	lh := r * 0.45
	ly := cy + r*0.05
	rect := clip.Rect{
		Min: image.Pt(int(cx-lw/2), int(ly)),
		Max: image.Pt(int(cx+lw/2), int(ly+lh)),
	}.Op()
	paint.FillShape(ops, col, rect)
}

// drawWarningTriangle draws a triangle with "!".
func drawWarningTriangle(ops *op.Ops, size int, col color.NRGBA) {
	s := float32(size)
	pad := s * 0.05

	var p clip.Path
	p.Begin(ops)
	p.MoveTo(f32.Pt(s/2, pad))
	p.LineTo(f32.Pt(s-pad, s-pad))
	p.LineTo(f32.Pt(pad, s-pad))
	p.Close()
	outline := clip.Stroke{Path: p.End(), Width: 1.5}.Op()
	paint.FillShape(ops, col, outline)

	// Exclamation line.
	cx := s / 2
	lw := s / 7
	rect := clip.Rect{
		Min: image.Pt(int(cx-lw/2), int(s*0.32)),
		Max: image.Pt(int(cx+lw/2), int(s*0.6)),
	}.Op()
	paint.FillShape(ops, col, rect)

	// Exclamation dot.
	drawFilledCircle(ops, cx, s*0.72, s/8, col)
}

// drawLightbulb draws a simplified lightbulb.
func drawLightbulb(ops *op.Ops, size int, col color.NRGBA) {
	s := float32(size)
	cx := s / 2
	r := s * 0.32

	// Bulb circle.
	drawCircleOutline(ops, cx, s*0.35, r, 1.5, col)

	// Base lines.
	lw := r * 0.8
	for i := 0; i < 2; i++ {
		y := s*0.7 + float32(i)*s*0.08
		rect := clip.Rect{
			Min: image.Pt(int(cx-lw/2), int(y)),
			Max: image.Pt(int(cx+lw/2), int(y+1.5)),
		}.Op()
		paint.FillShape(ops, col, rect)
	}

	// Neck lines.
	for _, dx := range []float32{-0.18, 0.18} {
		x := cx + s*dx
		rect := clip.Rect{
			Min: image.Pt(int(x-0.75), int(s*0.58)),
			Max: image.Pt(int(x+0.75), int(s*0.7)),
		}.Op()
		paint.FillShape(ops, col, rect)
	}
}

func drawCircleOutline(ops *op.Ops, cx, cy, r, width float32, col color.NRGBA) {
	const segments = 32
	var p clip.Path
	p.Begin(ops)
	for i := 0; i <= segments; i++ {
		angle := float64(i) * 2 * math.Pi / segments
		x := cx + r*float32(math.Cos(angle))
		y := cy + r*float32(math.Sin(angle))
		if i == 0 {
			p.MoveTo(f32.Pt(x, y))
		} else {
			p.LineTo(f32.Pt(x, y))
		}
	}
	p.Close()
	outline := clip.Stroke{Path: p.End(), Width: width}.Op()
	paint.FillShape(ops, col, outline)
}

func drawFilledCircle(ops *op.Ops, cx, cy, r float32, col color.NRGBA) {
	const segments = 16
	var p clip.Path
	p.Begin(ops)
	for i := 0; i <= segments; i++ {
		angle := float64(i) * 2 * math.Pi / segments
		x := cx + r*float32(math.Cos(angle))
		y := cy + r*float32(math.Sin(angle))
		if i == 0 {
			p.MoveTo(f32.Pt(x, y))
		} else {
			p.LineTo(f32.Pt(x, y))
		}
	}
	p.Close()
	fill := clip.Outline{Path: p.End()}.Op()
	paint.FillShape(ops, col, fill)
}
