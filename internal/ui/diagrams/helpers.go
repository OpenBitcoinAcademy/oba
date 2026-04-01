package diagrams

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// box draws a rounded filled rectangle with a centered, width-constrained label.
func box(gtx layout.Context, th *theme.Theme, label string, pos image.Point, w, h int, bg color.NRGBA) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	r := float32(h) / 6 // corner radius

	// Rounded rect fill.
	rr := clip.RRect{Rect: image.Rect(0, 0, w, h), SE: int(r), SW: int(r), NE: int(r), NW: int(r)}
	paint.FillShape(gtx.Ops, bg, rr.Op(gtx.Ops))

	// Border.
	paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1.2}.Op())

	// Label constrained to box width.
	cgtx := gtx
	cgtx.Constraints.Max.X = w
	cgtx.Constraints.Min = image.Point{}
	m := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.BodySmall, label)
	lbl.Color = th.Color.Text
	lbl.Font.Weight = font.Bold
	lbl.MaxLines = 1
	dims := lbl.Layout(cgtx)
	call := m.Stop()

	lx := (w - dims.Size.X) / 2
	if lx < 2 {
		lx = 2
	}
	ly := (h - dims.Size.Y) / 2
	defer op.Offset(image.Pt(lx, ly)).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// arrow draws a horizontal line with a proper triangular arrowhead.
func arrow(gtx layout.Context, th *theme.Theme, from, to image.Point) {
	c := th.Color.TextMuted
	headSize := 6

	// Line body (stop short of arrowhead).
	lineEnd := to.X - headSize
	if lineEnd < from.X {
		lineEnd = from.X
	}

	// Draw line.
	lineH := 2
	defer op.Offset(image.Pt(from.X, from.Y-lineH/2)).Push(gtx.Ops).Pop()
	paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(lineEnd-from.X, lineH)}.Op())

	// Arrowhead triangle via clip.Path.
	tipX := float32(to.X - from.X)
	baseX := tipX - float32(headSize)
	midY := float32(lineH) / 2
	hs := float32(headSize)

	var p clip.Path
	p.Begin(gtx.Ops)
	p.MoveTo(f32.Pt(tipX, midY))
	p.LineTo(f32.Pt(baseX, midY-hs/2))
	p.LineTo(f32.Pt(baseX, midY+hs/2))
	p.Close()
	paint.FillShape(gtx.Ops, c, clip.Outline{Path: p.End()}.Op())
}

// circle draws a filled circle at a center point with given radius.
func circle(gtx layout.Context, center image.Point, radius int, c color.NRGBA) {
	defer op.Offset(image.Pt(center.X-radius, center.Y-radius)).Push(gtx.Ops).Pop()
	d := radius * 2
	el := clip.Ellipse{Max: image.Pt(d, d)}
	paint.FillShape(gtx.Ops, c, el.Op(gtx.Ops))
}

// smallCircle draws a small dot, useful for P2P node indicators.
func smallCircle(gtx layout.Context, center image.Point, c color.NRGBA) {
	circle(gtx, center, 4, c)
}

// line draws a line between two points.
func line(gtx layout.Context, from, to image.Point, width float32, c color.NRGBA) {
	dx := float64(to.X - from.X)
	dy := float64(to.Y - from.Y)
	length := math.Sqrt(dx*dx + dy*dy)
	if length < 1 {
		return
	}

	angle := math.Atan2(dy, dx)

	defer op.Offset(from).Push(gtx.Ops).Pop()
	defer op.Affine(f32.Affine2D{}.Rotate(f32.Pt(0, 0), float32(-angle))).Push(gtx.Ops).Pop()

	paint.FillShape(gtx.Ops, c, clip.Rect{
		Min: image.Pt(0, int(-width/2)),
		Max: image.Pt(int(length), int(width/2)+1),
	}.Op())
}

// caption draws small bold colored text at a position.
func caption(gtx layout.Context, th *theme.Theme, text string, pos image.Point) {
	m := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.Caption, text)
	lbl.Color = th.Color.TextMuted
	lbl.Font.Weight = font.Bold
	lbl.Layout(gtx)
	call := m.Stop()
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// colorCaption draws small bold text in a specific color.
func colorCaption(gtx layout.Context, th *theme.Theme, text string, pos image.Point, c color.NRGBA) {
	m := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.BodySmall, text)
	lbl.Color = c
	lbl.Layout(gtx)
	call := m.Stop()
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// processBox draws a box with a distinct "process" appearance (primary color, smaller).
func processBox(gtx layout.Context, th *theme.Theme, label string, pos image.Point, w, h int) {
	box(gtx, th, label, pos, w, h, th.Color.Primary)
}

// pct calculates a percentage of a total.
func pct(total, percent int) int {
	return total * percent / 100
}

// Ensure imports are used.
var _ = unit.Dp(0)
