package components

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"

	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// ProgressBar renders a lesson progress indicator.
type ProgressBar struct {
	Completed int
	Total     int
	Theme     *theme.Theme
}

// Layout renders the progress bar with label.
func (p *ProgressBar) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Bar.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			height := gtx.Dp(unit.Dp(6))
			width := gtx.Constraints.Max.X

			// Track.
			track := clip.Rect{Max: image.Pt(width, height)}.Op()
			paint.FillShape(gtx.Ops, p.Theme.Color.Divider, track)

			// Fill.
			if p.Total > 0 {
				fillW := width * p.Completed / p.Total
				if fillW > width {
					fillW = width
				}
				fill := clip.Rect{Max: image.Pt(fillW, height)}.Op()
				paint.FillShape(gtx.Ops, p.Theme.Color.Primary, fill)
			}

			return layout.Dimensions{Size: image.Pt(width, height)}
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
		// Label.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			text := i18n.TFmt("nav.progress", map[string]string{
				"completed": itoa(p.Completed),
				"total":     itoa(p.Total),
			})
			lbl := material.Label(p.Theme.Material, p.Theme.Text.Caption, text)
			lbl.Color = p.Theme.Color.TextMuted
			return lbl.Layout(gtx)
		}),
	)
}

// CompletionDot renders a small colored dot indicating completion status.
func CompletionDot(gtx layout.Context, complete bool, c Colors) layout.Dimensions {
	size := gtx.Dp(unit.Dp(10))
	col := c.Divider
	if complete {
		col = c.Success
	}
	ellipse := clip.Ellipse{Max: image.Pt(size, size)}
	defer ellipse.Push(gtx.Ops).Pop()
	paint.FillShape(gtx.Ops, col, ellipse.Op(gtx.Ops))
	return layout.Dimensions{Size: image.Pt(size, size)}
}

// Colors is a re-export for use by callers without importing theme.
type Colors = theme.Colors

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
