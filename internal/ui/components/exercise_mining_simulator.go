package components

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

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

// MiningSimulator lets users find a nonce that produces a hash below a target.
type MiningSimulator struct {
	Theme      *theme.Theme
	mineBtn    widget.Clickable
	resetBtn   widget.Clickable
	targetBits int
	target     [32]byte
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
		targetBits: 8,
	}
	if cfg != nil {
		ms.targetBits = int(cfg.ConfigInt("target_bits", 8))
	}
	ms.buildTarget()
	ms.refreshBlockData()
	return ms
}

// buildTarget sets the target from targetBits leading zero bits.
// targetBits=8: target[0]=0, target[1]=0xff, rest 0xff. ~256 avg attempts.
func (ms *MiningSimulator) buildTarget() {
	// Fill with 0xff, then zero out the leading bytes/bits.
	for i := range ms.target {
		ms.target[i] = 0xff
	}
	fullBytes := ms.targetBits / 8
	for i := 0; i < fullBytes && i < 32; i++ {
		ms.target[i] = 0x00
	}
	remainder := ms.targetBits % 8
	if remainder > 0 && fullBytes < 32 {
		ms.target[fullBytes] = 0xff >> remainder
	}
}

func (ms *MiningSimulator) refreshBlockData() {
	ms.blockData = fmt.Sprintf("ver:02000000 prev:000000000019d6...e26f merkle:4a5e1e...3bce ts:%d bits:1d00ffff", time.Now().UnixNano())
}

func (ms *MiningSimulator) targetHex() string {
	return hex.EncodeToString(ms.target[:])
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
		ms.refreshBlockData()
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
					lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise_ui.mining_target"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.Code, ms.targetHex())
					lbl.Color = th.Color.Primary
					lbl.MaxLines = 2
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

	if bytes.Compare(hash[:], ms.target[:]) < 0 {
		ms.found = true
	}
}
