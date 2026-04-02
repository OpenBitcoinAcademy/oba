package screens

import (
	"image"
	"sort"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/app"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// Settings displays language selection, theme toggle, and progress reset.
type Settings struct {
	state      *app.State
	backBtn    widget.Clickable
	localeBtns map[string]*widget.Clickable
	lightBtn   widget.Clickable
	darkBtn    widget.Clickable
	resetBtn   widget.Clickable
	confirmBtn widget.Clickable
	cancelBtn  widget.Clickable
	confirming bool
}

// NewSettings creates the settings screen.
func NewSettings(state *app.State) *Settings {
	s := &Settings{
		state:      state,
		localeBtns: make(map[string]*widget.Clickable),
	}
	for _, loc := range i18n.Available() {
		s.localeBtns[loc] = new(widget.Clickable)
	}
	return s
}

// Layout renders the settings screen.
func (s *Settings) Layout(gtx layout.Context) layout.Dimensions {
	th := s.state.Theme

	if s.backBtn.Clicked(gtx) {
		s.state.NavigateHome()
		s.confirming = false
	}

	// Handle locale clicks.
	for loc, btn := range s.localeBtns {
		if btn.Clicked(gtx) {
			s.state.SetLocale(loc)
		}
	}

	// Handle theme toggle.
	if s.lightBtn.Clicked(gtx) {
		s.state.SetThemeMode(theme.ModeLight)
	}
	if s.darkBtn.Clicked(gtx) {
		s.state.SetThemeMode(theme.ModeDark)
	}

	// Handle reset flow.
	if s.resetBtn.Clicked(gtx) {
		s.confirming = true
	}
	if s.confirmBtn.Clicked(gtx) {
		s.state.Progress.Reset()
		s.state.SaveProgress()
		s.confirming = false
	}
	if s.cancelBtn.Clicked(gtx) {
		s.confirming = false
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Top bar.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: th.Space.Small, Bottom: th.Space.Small,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &s.backBtn, i18n.T("nav.home"))
						btn.Background = th.Color.Surface
						btn.Color = th.Color.Primary
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T("settings.language"))
						lbl.Color = th.Color.Text
						return lbl.Layout(gtx)
					}),
				)
			})
		}),
		// Divider.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, th.Color.Divider,
				clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, 1)}.Op())
			return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, 1)}
		}),
		// Content.
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:  th.Space.Large,
				Left: th.Space.Medium, Right: th.Space.Medium,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					// Language section.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T("settings.language"))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return s.renderLocaleButtons(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.XLarge}.Layout),

					// Theme section.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T("settings.theme"))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return s.renderThemeButtons(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.XLarge}.Layout),

					// Reset section.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.H3, i18n.T("settings.reset_progress"))
						lbl.Color = th.Color.Text
						lbl.Font.Weight = font.Bold
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if s.confirming {
							return s.renderConfirm(gtx)
						}
						btn := material.Button(th.Material, &s.resetBtn, i18n.T("settings.reset_progress"))
						btn.Background = th.Color.Error
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}),
				)
			})
		}),
	)
}

func (s *Settings) renderLocaleButtons(gtx layout.Context) layout.Dimensions {
	th := s.state.Theme
	current := i18n.Locale()

	locales := i18n.Available()
	sort.Strings(locales)

	var children []layout.FlexChild
	for _, loc := range locales {
		loc := loc
		btn := s.localeBtns[loc]
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			b := material.Button(th.Material, btn, localeDisplayName(loc))
			if loc == current {
				b.Background = th.Color.Primary
				b.Color = th.Color.OnPrimary
			} else {
				b.Background = th.Color.Surface
				b.Color = th.Color.Text
			}
			return layout.Inset{Right: th.Space.Small, Bottom: th.Space.Small}.Layout(gtx,
				b.Layout)
		}))
	}

	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
}

func (s *Settings) renderThemeButtons(gtx layout.Context) layout.Dimensions {
	th := s.state.Theme
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			b := material.Button(th.Material, &s.lightBtn, i18n.T("settings.theme_light"))
			if th.Mode == theme.ModeLight {
				b.Background = th.Color.Primary
				b.Color = th.Color.OnPrimary
			} else {
				b.Background = th.Color.Surface
				b.Color = th.Color.Text
			}
			return layout.Inset{Right: th.Space.Small}.Layout(gtx, b.Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			b := material.Button(th.Material, &s.darkBtn, i18n.T("settings.theme_dark"))
			if th.Mode == theme.ModeDark {
				b.Background = th.Color.Primary
				b.Color = th.Color.OnPrimary
			} else {
				b.Background = th.Color.Surface
				b.Color = th.Color.Text
			}
			return b.Layout(gtx)
		}),
	)
}

func (s *Settings) renderConfirm(gtx layout.Context) layout.Dimensions {
	th := s.state.Theme
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Body, i18n.T("settings.reset_confirm"))
			lbl.Color = th.Color.Error
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(th.Material, &s.confirmBtn, i18n.T("confirm.yes_reset"))
					btn.Background = th.Color.Error
					btn.Color = th.Color.OnPrimary
					return btn.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(th.Material, &s.cancelBtn, i18n.T("confirm.cancel"))
					btn.Background = th.Color.Surface
					btn.Color = th.Color.Text
					return btn.Layout(gtx)
				}),
			)
		}),
	)
}

func localeDisplayName(code string) string {
	switch code {
	case "en":
		return "English"
	case "es":
		return "Español"
	case "pt":
		return "Português"
	case "sw":
		return "Kiswahili"
	case "fr":
		return "Français"
	default:
		return code
	}
}
