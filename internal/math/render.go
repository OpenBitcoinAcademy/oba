package math

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Renderer draws a laid-out math Box tree using Gio drawing operations.
type Renderer struct {
	Theme    *material.Theme
	FontSize unit.Sp
	Color    color.NRGBA
}

// NewRenderer creates a renderer with the given theme and font size.
func NewRenderer(th *material.Theme, fontSize unit.Sp) *Renderer {
	return &Renderer{
		Theme:    th,
		FontSize: fontSize,
		Color:    color.NRGBA{A: 255},
	}
}

// Layout renders a formula string. Parses, lays out, and draws in one call.
func (r *Renderer) Layout(gtx layout.Context, formula string) layout.Dimensions {
	node, err := Parse(formula)
	if err != nil {
		return r.renderError(gtx, err.Error())
	}

	pxSize := gtx.Sp(r.FontSize)
	cfg := DefaultConfig(pxSize)
	box := Layout(node, cfg)

	r.drawBox(gtx, box, image.Point{}, pxSize)

	return layout.Dimensions{
		Size: image.Pt(box.Width(), box.Height()),
	}
}

// drawBox recursively renders a Box and its children.
func (r *Renderer) drawBox(gtx layout.Context, box *Box, offset image.Point, fontSize int) {
	pos := offset.Add(box.Origin)

	switch box.Node.Type {
	case NodeText:
		r.drawText(gtx, box.Node.Value, pos, fontSize)

	case NodeSymbol:
		r.drawText(gtx, box.Node.Value, pos, fontSize)

	case NodeTextCmd:
		r.drawTextUpright(gtx, box.Node.Value, pos, fontSize)

	case NodeMod:
		pad := fontSize / 6
		r.drawTextUpright(gtx, "mod", image.Pt(pos.X+pad, pos.Y), fontSize)

	case NodeFrac:
		// Draw numerator and denominator.
		for _, child := range box.Children {
			r.drawBox(gtx, child, pos, fontSize)
		}
		// Draw fraction bar.
		barY := box.Children[0].Origin.Y + box.Children[0].Height() + r.fracGap(fontSize)
		barW := box.Width()
		barH := max(2, fontSize/10)
		r.drawRect(gtx, image.Pt(pos.X, pos.Y+barY), barW, barH)

	case NodeSqrt:
		content := box.Children[0]
		r.drawBox(gtx, content, pos, fontSize)
		// Draw radical: vertical line on left, horizontal bar on top.
		radW := fontSize / 2
		// Vertical stroke of radical.
		r.drawRect(gtx, image.Pt(pos.X+radW-2, pos.Y), max(2, fontSize/14), box.Height())
		// Horizontal overline.
		r.drawRect(gtx, image.Pt(pos.X+radW-2, pos.Y), box.Width()-radW+2, max(1, fontSize/14))
		// Small diagonal tick at bottom-left.
		tickH := box.Height() / 3
		r.drawRect(gtx, image.Pt(pos.X+radW/3, pos.Y+box.Height()-tickH), max(2, fontSize/14), tickH)

	case NodeDelimited:
		content := box.Children[0]
		r.drawBox(gtx, content, pos, fontSize)
		delimW := fontSize / 2
		// Draw left delimiter.
		r.drawDelimChar(gtx, box.Node.Left, image.Pt(pos.X, pos.Y), delimW, box.Height(), fontSize)
		// Draw right delimiter.
		rightX := pos.X + box.Width() - delimW
		r.drawDelimChar(gtx, box.Node.Right, image.Pt(rightX, pos.Y), delimW, box.Height(), fontSize)

	case NodeSup, NodeSub, NodeSupSub:
		base := box.Children[0]
		r.drawBox(gtx, base, pos, fontSize)
		scriptSize := int(float64(fontSize) * 0.7)
		for i := 1; i < len(box.Children); i++ {
			r.drawBox(gtx, box.Children[i], pos, scriptSize)
		}

	default:
		// Row, Group: draw children sequentially.
		for _, child := range box.Children {
			r.drawBox(gtx, child, pos, fontSize)
		}
	}
}

// drawText renders math-italic text at the given position and font size.
func (r *Renderer) drawText(gtx layout.Context, s string, pos image.Point, fontSize int) {
	if s == "" {
		return
	}
	macro := op.Record(gtx.Ops)
	lbl := material.Label(r.Theme, unit.Sp(fontSize), s)
	lbl.Color = r.Color
	lbl.Font.Style = font.Italic
	lbl.Layout(gtx)
	call := macro.Stop()

	defer op.Offset(pos).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// drawTextUpright renders upright (roman) text for \text{} and \mod.
func (r *Renderer) drawTextUpright(gtx layout.Context, s string, pos image.Point, fontSize int) {
	if s == "" {
		return
	}
	macro := op.Record(gtx.Ops)
	lbl := material.Label(r.Theme, unit.Sp(fontSize), s)
	lbl.Color = r.Color
	lbl.Font.Style = font.Regular
	lbl.Alignment = text.Start
	lbl.Layout(gtx)
	call := macro.Stop()

	defer op.Offset(pos).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

// drawRect draws a filled rectangle (for fraction bars, radical strokes).
func (r *Renderer) drawRect(gtx layout.Context, pos image.Point, w, h int) {
	defer op.Offset(pos).Push(gtx.Ops).Pop()
	rect := clip.Rect{Max: image.Pt(w, h)}.Op()
	paint.FillShape(gtx.Ops, r.Color, rect)
}

// drawDelimChar draws a delimiter character scaled to the content height.
func (r *Renderer) drawDelimChar(gtx layout.Context, delim string, pos image.Point, w, h, fontSize int) {
	if delim == "." {
		return // invisible delimiter
	}
	// For now, draw the delimiter character at a larger size.
	// A proper implementation would stretch the glyph, but this is
	// sufficient for the POC.
	macro := op.Record(gtx.Ops)
	lbl := material.Label(r.Theme, unit.Sp(h*8/10), delim)
	lbl.Color = r.Color
	lbl.Alignment = text.Middle
	lbl.Layout(gtx)
	call := macro.Stop()

	defer op.Offset(pos).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)
}

func (r *Renderer) renderError(gtx layout.Context, msg string) layout.Dimensions {
	lbl := material.Label(r.Theme, r.FontSize, "[math error: "+msg+"]")
	lbl.Color = color.NRGBA{R: 200, A: 255}
	return lbl.Layout(gtx)
}

func (r *Renderer) fracGap(fontSize int) int {
	return max(2, fontSize/7)
}
