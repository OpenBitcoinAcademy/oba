package components

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Callout renders a highlighted text box with a colored left border.
type Callout struct {
	Content string
	Style   string // "info", "warning", "tip"
	Theme   *theme.Theme
}

// Layout renders the callout.
func (c *Callout) Layout(gtx layout.Context) layout.Dimensions {
	bg, border := c.colors()
	borderWidth := gtx.Dp(unit.Dp(4))

	return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{}.Layout(gtx,
				// Background.
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					rect := clip.Rect{Max: gtx.Constraints.Min}.Op()
					paint.FillShape(gtx.Ops, bg, rect)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				// Left border.
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					rect := clip.Rect{Max: image.Pt(borderWidth, gtx.Constraints.Min.Y)}.Op()
					paint.FillShape(gtx.Ops, border, rect)
					return layout.Dimensions{}
				}),
				// Content.
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(12), Bottom: unit.Dp(12),
						Left: unit.Dp(16), Right: unit.Dp(12),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(c.Theme.Material, c.Theme.Text.Body, c.Content)
						lbl.Color = c.Theme.Color.Text
						return lbl.Layout(gtx)
					})
				}),
			)
		},
	)
}

func (c *Callout) colors() (bg, border color.NRGBA) {
	switch c.Style {
	case "warning":
		return c.Theme.Color.WarningBg, c.Theme.Color.Warning
	case "tip":
		return c.Theme.Color.TipBg, c.Theme.Color.Success
	default: // "info"
		return c.Theme.Color.InfoBg, c.Theme.Color.Accent
	}
}
