package components

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// CodeBlock renders a code section with a dark background and monospace font.
type CodeBlock struct {
	Content string
	Theme   *theme.Theme
}

// Layout renders the code block.
func (cb *CodeBlock) Layout(gtx layout.Context) layout.Dimensions {
	bg := cb.codeBg()

	return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					r := gtx.Dp(unit.Dp(6))
					rect := clip.RRect{
						Rect: image.Rectangle{Max: gtx.Constraints.Min},
						NE:   r, NW: r, SE: r, SW: r,
					}.Op(gtx.Ops)
					paint.FillShape(gtx.Ops, bg, rect)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(12), Bottom: unit.Dp(12),
						Left: unit.Dp(14), Right: unit.Dp(14),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(cb.Theme.Material, cb.Theme.Text.Code, cb.Content)
						lbl.Font.Typeface = "monospace"
						lbl.Font.Weight = font.Normal
						lbl.Color = cb.codeText()
						return lbl.Layout(gtx)
					})
				}),
			)
		},
	)
}

func (cb *CodeBlock) codeBg() color.NRGBA {
	if cb.Theme.Mode == theme.ModeDark {
		return color.NRGBA{R: 0x10, G: 0x12, B: 0x30, A: 0xFF}
	}
	return color.NRGBA{R: 0x1E, G: 0x22, B: 0x4C, A: 0xFF}
}

func (cb *CodeBlock) codeText() color.NRGBA {
	if cb.Theme.Mode == theme.ModeDark {
		return color.NRGBA{R: 0xC8, G: 0xCC, B: 0xE0, A: 0xFF}
	}
	return color.NRGBA{R: 0xE0, G: 0xE2, B: 0xF0, A: 0xFF}
}
