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

// Home displays the chapter selector with progress for each chapter.
type Home struct {
	state       *app.State
	chapterBtns []widget.Clickable
	settingsBtn widget.Clickable
	list        widget.List
}

// NewHome creates the home screen.
func NewHome(state *app.State) *Home {
	h := &Home{
		state:       state,
		chapterBtns: make([]widget.Clickable, len(state.Chapters)),
	}
	h.list.Axis = layout.Vertical
	return h
}

// Layout renders the home screen.
func (h *Home) Layout(gtx layout.Context) layout.Dimensions {
	th := h.state.Theme

	// Handle clicks.
	for i := range h.chapterBtns {
		if h.chapterBtns[i].Clicked(gtx) {
			h.state.NavigateToChapter(i)
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
						lbl := material.Label(th.Material, th.Text.H1, i18n.T("nav.app_title"))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					// Settings button.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &h.settingsBtn, i18n.T("nav.settings"))
						btn.Background = th.Color.Surface
						btn.Color = th.Color.TextMuted
						btn.TextSize = th.Text.Caption
						return btn.Layout(gtx)
					}),
				)
			})
		}),
		// Chapter list.
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.List(th.Material, &h.list).Layout(gtx, len(h.state.Chapters),
				func(gtx layout.Context, i int) layout.Dimensions {
					return h.chapterCard(gtx, i)
				},
			)
		}),
	)
}

func (h *Home) chapterCard(gtx layout.Context, idx int) layout.Dimensions {
	th := h.state.Theme
	ch := h.state.Chapters[idx]
	completed := h.state.Progress.CompletedCountForChapter(ch)
	total := len(ch.Lessons)

	return layout.Inset{
		Left: th.Space.Medium, Right: th.Space.Medium,
		Bottom: th.Space.Small,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Clickable(gtx, &h.chapterBtns[idx], func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, th.Color.Surface, clip.Rect{Max: gtx.Constraints.Max}.Op())

			return layout.Inset{
				Top: th.Space.Medium, Bottom: th.Space.Medium,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					// Title.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T(ch.TitleKey))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
					// Description.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Body, i18n.T(ch.DescKey))
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
					// Progress.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						bar := &components.ProgressBar{
							Completed: completed,
							Total:     total,
							Theme:     th,
						}
						return bar.Layout(gtx)
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

var _ = image.Point{} // keep import
