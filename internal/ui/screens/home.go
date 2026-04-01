// Package screens implements the main application views.
package screens

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/components"
)

// Home displays the lesson list with progress indicators.
type Home struct {
	state       *app.State
	lessonBtns  []widget.Clickable
	settingsBtn widget.Clickable
	list        widget.List
}

// NewHome creates the home screen.
func NewHome(state *app.State) *Home {
	h := &Home{
		state:      state,
		lessonBtns: make([]widget.Clickable, len(state.Chapter.Lessons)),
	}
	h.list.Axis = layout.Vertical
	return h
}

// Layout renders the home screen.
func (h *Home) Layout(gtx layout.Context) layout.Dimensions {
	th := h.state.Theme

	// Handle clicks.
	for i := range h.lessonBtns {
		if h.lessonBtns[i].Clicked(gtx) {
			h.state.NavigateToLesson(i)
		}
	}
	if h.settingsBtn.Clicked(gtx) {
		h.state.NavigateSettings()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Header.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: th.Space.Large, Bottom: th.Space.Medium,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H1, i18n.T(h.state.Chapter.TitleKey))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Body, i18n.T(h.state.Chapter.DescKey))
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								bar := &components.ProgressBar{
									Completed: h.state.Progress.CompletedCount(),
									Total:     len(h.state.Chapter.Lessons),
									Theme:     th,
								}
								return bar.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th.Material, &h.settingsBtn, i18n.T("nav.settings"))
								btn.Background = th.Color.Surface
								btn.Color = th.Color.TextMuted
								btn.TextSize = th.Text.Caption
								return btn.Layout(gtx)
							}),
						)
					}),
				)
			})
		}),
		// Lesson list.
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.List(th.Material, &h.list).Layout(gtx, len(h.state.Chapter.Lessons),
				func(gtx layout.Context, i int) layout.Dimensions {
					return h.lessonCard(gtx, i)
				},
			)
		}),
	)
}

func (h *Home) lessonCard(gtx layout.Context, idx int) layout.Dimensions {
	th := h.state.Theme
	lesson := h.state.Chapter.Lessons[idx]
	complete := h.state.Progress.IsLessonComplete(lesson.ID)

	return layout.Inset{
		Left: th.Space.Medium, Right: th.Space.Medium,
		Bottom: th.Space.Small,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Clickable(gtx, &h.lessonBtns[idx], func(gtx layout.Context) layout.Dimensions {
			// Card background.
			paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: gtx.Constraints.Max}.Op())

			return layout.Inset{
				Top: th.Space.Medium, Bottom: th.Space.Medium,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					// Completion dot.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return components.CompletionDot(gtx, complete, th.Color)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
					// Lesson info.
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.H3, i18n.T(lesson.TitleKey))
								lbl.Color = th.Color.Text
								return lbl.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								secs := len(lesson.Sections)
								lbl := material.Label(th.Material, th.Text.Caption, itoa(secs)+" sections")
								lbl.Color = th.Color.TextMuted
								return lbl.Layout(gtx)
							}),
						)
					}),
					// Chevron.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, ">")
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
				)
			})
		})
	})
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var d []byte
	for n > 0 {
		d = append([]byte{byte('0' + n%10)}, d...)
		n /= 10
	}
	return string(d)
}

// Enforce minimum touch target size.
var _ = image.Point{}
