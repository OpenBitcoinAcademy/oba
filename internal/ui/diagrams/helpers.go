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

// withAlpha returns a copy of c with the given alpha value.
func withAlpha(c color.NRGBA, a uint8) color.NRGBA {
	c.A = a
	return c
}

// ptf converts an image.Point to an f32.Point.
func ptf(p image.Point) f32.Point {
	return f32.Pt(float32(p.X), float32(p.Y))
}

// dirArrow draws a line with an arrowhead from any point to any point.
func dirArrow(gtx layout.Context, from, to image.Point, width float32, c color.NRGBA) {
	headSize := float32(7)
	dx := float32(to.X - from.X)
	dy := float32(to.Y - from.Y)
	length := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	if length < 1 {
		return
	}

	ux, uy := dx/length, dy/length
	px, py := -uy, ux

	lineEnd := f32.Pt(float32(to.X)-ux*headSize, float32(to.Y)-uy*headSize)

	var lp clip.Path
	lp.Begin(gtx.Ops)
	lp.MoveTo(ptf(from))
	lp.LineTo(lineEnd)
	paint.FillShape(gtx.Ops, c, clip.Stroke{Path: lp.End(), Width: width}.Op())

	tip := ptf(to)
	base1 := f32.Pt(lineEnd.X+px*headSize*0.5, lineEnd.Y+py*headSize*0.5)
	base2 := f32.Pt(lineEnd.X-px*headSize*0.5, lineEnd.Y-py*headSize*0.5)

	var hp clip.Path
	hp.Begin(gtx.Ops)
	hp.MoveTo(tip)
	hp.LineTo(base1)
	hp.LineTo(base2)
	hp.Close()
	paint.FillShape(gtx.Ops, c, clip.Outline{Path: hp.End()}.Op())
}

// curvedLine draws a smooth cubic bezier from one point to another.
// The curve drops vertically before sweeping horizontally, making it
// ideal for tree-branch connectors (parent bottom to child top).
func curvedLine(gtx layout.Context, from, to image.Point, width float32, c color.NRGBA) {
	midY := float32(from.Y+to.Y) / 2
	ctrl0 := f32.Pt(float32(from.X), midY)
	ctrl1 := f32.Pt(float32(to.X), midY)

	var p clip.Path
	p.Begin(gtx.Ops)
	p.MoveTo(ptf(from))
	p.CubeTo(ctrl0, ctrl1, ptf(to))
	paint.FillShape(gtx.Ops, c, clip.Stroke{Path: p.End(), Width: width}.Op())
}

// shadowBox draws a box with a subtle drop shadow behind it.
func shadowBox(gtx layout.Context, th *theme.Theme, label string, pos image.Point, w, h int, bg color.NRGBA) {
	off := 3
	shadowColor := color.NRGBA{A: 30}
	r := float32(h) / 6
	ri := int(r)

	func() {
		defer op.Offset(image.Pt(pos.X+off, pos.Y+off)).Push(gtx.Ops).Pop()
		rr := clip.RRect{Rect: image.Rect(0, 0, w, h), SE: ri, SW: ri, NE: ri, NW: ri}
		paint.FillShape(gtx.Ops, shadowColor, rr.Op(gtx.Ops))
	}()

	box(gtx, th, label, pos, w, h, bg)
}

// dashedLine draws a dashed line between two points.
func dashedLine(gtx layout.Context, from, to image.Point, width float32, c color.NRGBA) {
	dx := float64(to.X - from.X)
	dy := float64(to.Y - from.Y)
	length := math.Sqrt(dx*dx + dy*dy)
	if length < 1 {
		return
	}

	ux, uy := dx/length, dy/length
	dashLen := 6.0
	gapLen := 4.0

	for pos := 0.0; pos < length; pos += dashLen + gapLen {
		end := pos + dashLen
		if end > length {
			end = length
		}
		f := image.Pt(from.X+int(ux*pos), from.Y+int(uy*pos))
		t := image.Pt(from.X+int(ux*end), from.Y+int(uy*end))
		line(gtx, f, t, width, c)
	}
}

// groupBox draws a subtle rounded border for visually grouping elements.
func groupBox(gtx layout.Context, th *theme.Theme, pos image.Point, w, h int) {
	borderColor := withAlpha(th.Color.TextMuted, 40)
	ri := 8

	defer op.Offset(pos).Push(gtx.Ops).Pop()
	rr := clip.RRect{Rect: image.Rect(0, 0, w, h), SE: ri, SW: ri, NE: ri, NW: ri}
	paint.FillShape(gtx.Ops, borderColor, clip.Stroke{Path: rr.Path(gtx.Ops), Width: 1}.Op())
}

// badge draws a labeled circle, useful for step numbers.
func badge(gtx layout.Context, th *theme.Theme, label string, center image.Point, radius int, bg color.NRGBA) {
	circle(gtx, center, radius, bg)

	m := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.Caption, label)
	lbl.Color = th.Color.OnPrimary
	lbl.Font.Weight = font.Bold
	dims := lbl.Layout(gtx)
	call := m.Stop()

	defer op.Offset(image.Pt(center.X-dims.Size.X/2, center.Y-dims.Size.Y/2)).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// diamond draws a decision diamond with a centered label.
func diamond(gtx layout.Context, th *theme.Theme, label string, center image.Point, w, h int, bg color.NRGBA) {
	halfW, halfH := float32(w/2), float32(h/2)
	cx, cy := float32(center.X), float32(center.Y)

	// Fill.
	var p clip.Path
	p.Begin(gtx.Ops)
	p.MoveTo(f32.Pt(cx, cy-halfH))
	p.LineTo(f32.Pt(cx+halfW, cy))
	p.LineTo(f32.Pt(cx, cy+halfH))
	p.LineTo(f32.Pt(cx-halfW, cy))
	p.Close()
	paint.FillShape(gtx.Ops, bg, clip.Outline{Path: p.End()}.Op())

	// Border.
	var p2 clip.Path
	p2.Begin(gtx.Ops)
	p2.MoveTo(f32.Pt(cx, cy-halfH))
	p2.LineTo(f32.Pt(cx+halfW, cy))
	p2.LineTo(f32.Pt(cx, cy+halfH))
	p2.LineTo(f32.Pt(cx-halfW, cy))
	p2.Close()
	paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{Path: p2.End(), Width: 1.2}.Op())

	// Centered label.
	mr := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.Caption, label)
	lbl.Color = th.Color.Text
	lbl.Font.Weight = font.Bold
	lbl.MaxLines = 1
	cgtx := gtx
	cgtx.Constraints.Max.X = w
	cgtx.Constraints.Min = image.Point{}
	dims := lbl.Layout(cgtx)
	call := mr.Stop()

	defer op.Offset(image.Pt(center.X-dims.Size.X/2, center.Y-dims.Size.Y/2)).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// roundedPath draws a polyline with rounded corners at each vertex.
// The radius controls how much the corners are smoothed.
func roundedPath(gtx layout.Context, points []image.Point, radius float32, width float32, c color.NRGBA) {
	n := len(points)
	if n < 2 {
		return
	}
	if n == 2 {
		line(gtx, points[0], points[1], width, c)
		return
	}

	var p clip.Path
	p.Begin(gtx.Ops)
	p.MoveTo(ptf(points[0]))

	for i := 1; i < n-1; i++ {
		curr := ptf(points[i])
		prev := ptf(points[i-1])
		next := ptf(points[i+1])

		toPrev := f32.Pt(prev.X-curr.X, prev.Y-curr.Y)
		toNext := f32.Pt(next.X-curr.X, next.Y-curr.Y)

		lenPrev := float32(math.Sqrt(float64(toPrev.X*toPrev.X + toPrev.Y*toPrev.Y)))
		lenNext := float32(math.Sqrt(float64(toNext.X*toNext.X + toNext.Y*toNext.Y)))

		if lenPrev < 1 || lenNext < 1 {
			p.LineTo(curr)
			continue
		}

		r := radius
		if maxR := min32(lenPrev, lenNext) / 2; r > maxR {
			r = maxR
		}

		arcStart := f32.Pt(curr.X+toPrev.X/lenPrev*r, curr.Y+toPrev.Y/lenPrev*r)
		arcEnd := f32.Pt(curr.X+toNext.X/lenNext*r, curr.Y+toNext.Y/lenNext*r)

		p.LineTo(arcStart)
		p.QuadTo(curr, arcEnd)
	}

	p.LineTo(ptf(points[n-1]))
	paint.FillShape(gtx.Ops, c, clip.Stroke{Path: p.End(), Width: width}.Op())
}

func min32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
