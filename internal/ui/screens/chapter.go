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

// Chapter displays the lesson list for a selected chapter.
type Chapter struct {
	state       *app.State
	lessonBtns  []widget.Clickable
	backBtn     widget.Clickable
	list        widget.List
	lastChapter int
}

// NewChapter creates the chapter lesson-list screen.
func NewChapter(state *app.State) *Chapter {
	c := &Chapter{
		state:       state,
		lastChapter: -1,
	}
	c.list.Axis = layout.Vertical
	c.resetButtons()
	return c
}

func (c *Chapter) resetButtons() {
	if c.state.Chapter != nil {
		c.lessonBtns = make([]widget.Clickable, len(c.state.Chapter.Lessons))
	}
	c.lastChapter = c.state.CurrentChapter
}

// Layout renders the chapter screen.
func (c *Chapter) Layout(gtx layout.Context) layout.Dimensions {
	// Reset if chapter changed.
	if c.lastChapter != c.state.CurrentChapter {
		c.resetButtons()
	}

	th := c.state.Theme
	ch := c.state.Chapter
	if ch == nil {
		return layout.Dimensions{}
	}

	// Handle clicks.
	for i := range c.lessonBtns {
		if c.lessonBtns[i].Clicked(gtx) {
			c.state.NavigateToLesson(i)
		}
	}
	if c.backBtn.Clicked(gtx) {
		c.state.NavigateHome()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Header.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: th.Space.Large, Bottom: th.Space.Medium,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					// Back + settings row.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th.Material, &c.backBtn, i18n.T("nav.home"))
								btn.Background = th.Color.Surface
								btn.Color = th.Color.Primary
								btn.TextSize = th.Text.Caption
								return btn.Layout(gtx)
							}),
						)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					// Title.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H1, i18n.T(ch.TitleKey))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Body, i18n.T(ch.DescKey))
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					// Progress bar.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						bar := &components.ProgressBar{
							Completed: c.state.Progress.CompletedCountForChapter(ch),
							Total:     len(ch.Lessons),
							Theme:     th,
						}
						return bar.Layout(gtx)
					}),
				)
			})
		}),
		// Lesson list.
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.List(th.Material, &c.list).Layout(gtx, len(ch.Lessons),
				func(gtx layout.Context, i int) layout.Dimensions {
					return c.lessonCard(gtx, i)
				},
			)
		}),
	)
}

func (c *Chapter) lessonCard(gtx layout.Context, idx int) layout.Dimensions {
	th := c.state.Theme
	lesson := c.state.Chapter.Lessons[idx]
	complete := c.state.Progress.IsLessonComplete(lesson.ID)

	return layout.Inset{
		Left: th.Space.Medium, Right: th.Space.Medium,
		Bottom: th.Space.Small,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Clickable(gtx, &c.lessonBtns[idx], func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: gtx.Constraints.Max}.Op())

			return layout.Inset{
				Top: th.Space.Medium, Bottom: th.Space.Medium,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return components.CompletionDot(gtx, complete, th.Color)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.H3, i18n.T(lesson.TitleKey))
								lbl.Color = th.Color.Text
								return lbl.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								secs := len(lesson.Sections)
								text := i18n.TPlural("lessons.sections", secs, map[string]string{
									"count": itoa(secs),
								})
								lbl := material.Label(th.Material, th.Text.Caption, text)
								lbl.Color = th.Color.TextMuted
								return lbl.Layout(gtx)
							}),
						)
					}),
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

var _ = image.Point{} // keep import
