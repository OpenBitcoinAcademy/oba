package components

import (
	"encoding/hex"
	"fmt"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/openbitcoinacademy/oba/internal/bitcoin"
	"github.com/openbitcoinacademy/oba/internal/content"
	"github.com/openbitcoinacademy/oba/internal/i18n"
	"github.com/openbitcoinacademy/oba/internal/ui/theme"
)

// MiningSimulator lets users find a nonce that produces a hash with leading zeros.
type MiningSimulator struct {
	Theme      *theme.Theme
	mineBtn    widget.Clickable
	resetBtn   widget.Clickable
	difficulty int
	nonce      uint32
	blockData  string
	lastHash   string
	found      bool
	attempts   int
}

// NewMiningSimulator creates a mining simulator widget.
func NewMiningSimulator(th *theme.Theme, cfg *content.ExerciseConfig) *MiningSimulator {
	ms := &MiningSimulator{
		Theme:      th,
		difficulty: 1,
		blockData:  "OBA Block Data 2024",
	}
	if cfg != nil {
		d := cfg.ConfigInt("initial_difficulty", 1)
		ms.difficulty = int(d)
	}
	return ms
}

// Layout renders the mining simulator.
func (ms *MiningSimulator) Layout(gtx layout.Context) layout.Dimensions {
	th := ms.Theme

	if ms.mineBtn.Clicked(gtx) && !ms.found {
		ms.tryNonce()
	}
	if ms.resetBtn.Clicked(gtx) {
		ms.nonce = 0
		ms.lastHash = ""
		ms.found = false
		ms.attempts = 0
	}

	return layout.Inset{Top: th.Space.Medium, Bottom: th.Space.Medium}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.mining_title"))
					lbl.Color = th.Color.Text
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.mining_prompt"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
				// Target display.
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					target := fmt.Sprintf("%s %s", i18n.T("exercise_ui.mining_target"), strings.Repeat("0", ms.difficulty)+"...")
					lbl := material.Label(th.Material, th.Text.BodySmall, target)
					lbl.Color = th.Color.Primary
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
				// Current attempt.
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if ms.lastHash == "" {
						return layout.Dimensions{}
					}
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.Caption,
								fmt.Sprintf("Nonce: %d  Attempts: %d", ms.nonce, ms.attempts))
							lbl.Color = th.Color.TextMuted
							return lbl.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							bg := th.Color.InfoBg
							c := th.Color.Text
							if ms.found {
								bg = th.Color.TipBg
								c = th.Color.Success
							}
							return layout.Stack{}.Layout(gtx,
								layout.Expanded(func(gtx layout.Context) layout.Dimensions {
									paint.FillShape(gtx.Ops, bg,
										clip.Rect{Max: gtx.Constraints.Min}.Op())
									return layout.Dimensions{Size: gtx.Constraints.Min}
								}),
								layout.Stacked(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{
										Top: unit.Dp(4), Bottom: unit.Dp(4),
										Left: unit.Dp(8), Right: unit.Dp(8),
									}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										lbl := material.Label(th.Material, th.Text.Code, ms.lastHash)
										lbl.Color = c
										lbl.MaxLines = 2
										return lbl.Layout(gtx)
									})
								}),
							)
						}),
					)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
				// Buttons.
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if ms.found {
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.mining_found"))
								lbl.Color = th.Color.Success
								return lbl.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th.Material, &ms.resetBtn, i18n.T("exercise.reset"))
								btn.Background = th.Color.Surface
								btn.Color = th.Color.Primary
								return btn.Layout(gtx)
							}),
						)
					}
					btn := material.Button(th.Material, &ms.mineBtn, i18n.T("exercise_ui.mining_try"))
					btn.Background = th.Color.Primary
					btn.Color = th.Color.OnPrimary
					return btn.Layout(gtx)
				}),
			)
		})
}

func (ms *MiningSimulator) tryNonce() {
	ms.nonce++
	ms.attempts++
	data := fmt.Sprintf("%s|%d", ms.blockData, ms.nonce)
	hash := bitcoin.SHA256([]byte(data))
	ms.lastHash = hex.EncodeToString(hash[:])

	// Check if hash starts with enough zeros.
	prefix := strings.Repeat("0", ms.difficulty)
	if strings.HasPrefix(ms.lastHash, prefix) {
		ms.found = true
	}
}
