package diagrams

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// box draws a filled rectangle with a centered label.
func box(gtx layout.Context, th *theme.Theme, label string, pos image.Point, w, h int, bg color.NRGBA) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()

	// Fill.
	rect := clip.Rect{Max: image.Pt(w, h)}.Op()
	paint.FillShape(gtx.Ops, bg, rect)

	// Border.
	paint.FillShape(gtx.Ops, th.Color.Text, clip.Stroke{
		Path:  clip.Rect{Max: image.Pt(w, h)}.Path(),
		Width: 1.5,
	}.Op())

	// Label.
	m := op.Record(gtx.Ops)
	lbl := material.Label(th.Material, th.Text.BodySmall, label)
	lbl.Color = th.Color.Text
	lbl.Font.Weight = font.Bold
	dims := lbl.Layout(gtx)
	call := m.Stop()

	lx := (w - dims.Size.X) / 2
	ly := (h - dims.Size.Y) / 2
	defer op.Offset(image.Pt(lx, ly)).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// arrow draws a horizontal line with a small arrowhead.
func arrow(gtx layout.Context, th *theme.Theme, pos image.Point, w int) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	c := th.Color.TextMuted

	// Line.
	paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(w, 2)}.Op())

	// Arrowhead (small filled triangle approximated as a rect).
	s := 6
	defer op.Offset(image.Pt(w-s, -s/2+1)).Push(gtx.Ops).Pop()
	paint.FillShape(gtx.Ops, c, clip.Rect{Max: image.Pt(s, s)}.Op())
}

// caption draws small bold text at a position.
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
