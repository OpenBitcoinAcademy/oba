package components

import (
	"fmt"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// hdLevel defines one component of a BIP44/84 derivation path.
type hdLevel struct {
	name     string
	options  []string
	hardened bool
	selected int
}

// HDPathExplorer lets users build a BIP44/84 derivation path by
// selecting each level: purpose, coin type, account, change, index.
type HDPathExplorer struct {
	Theme    *theme.Theme
	levels   []hdLevel
	levelBtn []widget.Clickable
}

// NewHDPathExplorer creates an HD path explorer widget.
func NewHDPathExplorer(th *theme.Theme, cfg *content.ExerciseConfig) *HDPathExplorer {
	levels := []hdLevel{
		{name: "purpose", options: []string{"44'", "49'", "84'", "86'"}, hardened: true, selected: 2},
		{name: "coin", options: []string{"0'", "1'"}, hardened: true, selected: 0},
		{name: "account", options: []string{"0'", "1'", "2'"}, hardened: true, selected: 0},
		{name: "change", options: []string{"0", "1"}, hardened: false, selected: 0},
		{name: "index", options: []string{"0", "1", "2", "3", "4"}, hardened: false, selected: 0},
	}

	return &HDPathExplorer{
		Theme:    th,
		levels:   levels,
		levelBtn: make([]widget.Clickable, len(levels)),
	}
}

// Layout renders the HD path explorer.
func (hp *HDPathExplorer) Layout(gtx layout.Context) layout.Dimensions {
	th := hp.Theme

	// Handle level button clicks: cycle to next option.
	for i := range hp.levelBtn {
		if hp.levelBtn[i].Clicked(gtx) {
			hp.levels[i].selected = (hp.levels[i].selected + 1) % len(hp.levels[i].options)
		}
	}

	path := hp.buildPath()

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.hd_path_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Prompt.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.hd_path_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Level buttons: each click cycles the value.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				var children []layout.FlexChild
				for i := range hp.levels {
					idx := i
					children = append(children,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							label := fmt.Sprintf("%s: %s", hp.levels[idx].name, hp.levels[idx].options[hp.levels[idx].selected])
							btn := material.Button(th.Material, &hp.levelBtn[idx], label)
							btn.Background = th.Color.Surface
							btn.Color = th.Color.Primary
							btn.TextSize = th.Text.BodySmall
							return layout.Inset{Right: th.Space.Small}.Layout(gtx, btn.Layout)
						}),
					)
				}
				return layout.Flex{Spacing: layout.SpaceStart}.Layout(gtx, children...)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Path display.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Caption, "Derivation Path")
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						paint.FillShape(gtx.Ops, th.Color.InfoBg,
							clip.Rect{Max: gtx.Constraints.Min}.Op())
						return layout.Dimensions{Size: gtx.Constraints.Min}
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: unit.Dp(8), Bottom: unit.Dp(8),
							Left: unit.Dp(12), Right: unit.Dp(12),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.H3, path)
							lbl.Color = th.Color.Text
							lbl.Font.Typeface = "monospace"
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Description of each level.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				desc := hp.describeLevel()
				lbl := material.Label(th.Material, th.Text.BodySmall, desc)
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
		)
	})
}

func (hp *HDPathExplorer) buildPath() string {
	parts := make([]string, len(hp.levels))
	for i, lvl := range hp.levels {
		parts[i] = lvl.options[lvl.selected]
	}
	return "m/" + strings.Join(parts, "/")
}

func (hp *HDPathExplorer) describeLevel() string {
	names := []string{
		"Purpose: 44=Legacy, 49=SegWit-compat, 84=Native SegWit, 86=Taproot",
		"Coin: 0=Bitcoin mainnet, 1=Bitcoin testnet",
		"Account: separate accounts under one seed",
		"Change: 0=receiving addresses, 1=change addresses",
		"Index: address number within the chain",
	}
	var lines []string
	for _, n := range names {
		lines = append(lines, n)
	}
	return strings.Join(lines, "\n")
}
