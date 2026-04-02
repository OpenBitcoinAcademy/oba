package components

import (
	"bytes"
	"encoding/hex"

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

// ScriptDebugger steps through P2PKH script execution.
type ScriptDebugger struct {
	Theme    *theme.Theme
	stepBtn  widget.Clickable
	resetBtn widget.Clickable
	engine   *bitcoin.Engine
	steps    []bitcoin.ScriptStep
	pos      int
	done     bool
	result   bool
}

// NewScriptDebugger creates a script debugger widget.
func NewScriptDebugger(th *theme.Theme, cfg *content.ExerciseConfig) *ScriptDebugger {
	sd := &ScriptDebugger{Theme: th}
	sd.generateScript()
	return sd
}

func (sd *ScriptDebugger) generateScript() {
	key, err := bitcoin.GenerateKey()
	if err != nil {
		return
	}
	pub := key.PublicKey()
	pubBytes := pub.SerializeCompressed()
	pubHash := pub.Hash160()

	fakeSig := bytes.Repeat([]byte{0x30}, 71)
	scriptSig := bitcoin.BuildP2PKHScriptSig(fakeSig, pubBytes)
	scriptPubKey := bitcoin.BuildP2PKHScript(pubHash)
	combined := bitcoin.CombineScripts(scriptSig, scriptPubKey)

	sd.engine = bitcoin.NewEngine(combined)
	sd.steps = nil
	sd.pos = 0
	sd.done = false
	sd.result = false
}

// Layout renders the script debugger.
func (sd *ScriptDebugger) Layout(gtx layout.Context) layout.Dimensions {
	th := sd.Theme

	if sd.stepBtn.Clicked(gtx) && !sd.done {
		step, err := sd.engine.Step()
		if err != nil {
			sd.done = true
			sd.result = false
		} else if step != nil {
			sd.steps = append(sd.steps, *step)
			sd.pos++
		}
		if sd.engine.Done() {
			sd.done = true
			sd.result = sd.engine.Result()
		}
	}
	if sd.resetBtn.Clicked(gtx) {
		sd.generateScript()
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.script_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.script_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			// Steps.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return sd.renderSteps(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			// Result or buttons.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if sd.done {
					return sd.renderResult(gtx)
				}
				return layout.Flex{}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &sd.stepBtn, i18n.T("exercise.next_step"))
						btn.Background = th.Color.Primary
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &sd.resetBtn, i18n.T("exercise.reset"))
						btn.Background = th.Color.Surface
						btn.Color = th.Color.Primary
						return btn.Layout(gtx)
					}),
				)
			}),
		)
	})
}

func (sd *ScriptDebugger) renderSteps(gtx layout.Context) layout.Dimensions {
	th := sd.Theme
	var children []layout.FlexChild

	for _, step := range sd.steps {
		step := step
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall, step.Op)
					lbl.Color = th.Color.Primary
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.Caption, step.Desc)
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					stackStr := "Stack: ["
					for i, item := range step.Stack {
						if i > 0 {
							stackStr += ", "
						}
						h := hex.EncodeToString(item)
						if len(h) > 16 {
							h = h[:16] + "..."
						}
						stackStr += h
					}
					stackStr += "]"
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							paint.FillShape(gtx.Ops, th.Color.InfoBg,
								clip.Rect{Max: gtx.Constraints.Min}.Op())
							return layout.Dimensions{Size: gtx.Constraints.Min}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Top: unit.Dp(2), Bottom: unit.Dp(2),
								Left: unit.Dp(4), Right: unit.Dp(4),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.Code, stackStr)
								lbl.Color = th.Color.Text
								lbl.MaxLines = 2
								return lbl.Layout(gtx)
							})
						}),
					)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			)
		}))
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

func (sd *ScriptDebugger) renderResult(gtx layout.Context) layout.Dimensions {
	th := sd.Theme
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			msg := i18n.T("exercise_ui.script_result_fail")
			c := th.Color.Error
			if sd.result {
				msg = i18n.T("exercise_ui.script_result_pass")
				c = th.Color.Success
			}
			lbl := material.Label(th.Material, th.Text.H3, msg)
			lbl.Color = c
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th.Material, &sd.resetBtn, i18n.T("exercise.reset"))
			btn.Background = th.Color.Surface
			btn.Color = th.Color.Primary
			return btn.Layout(gtx)
		}),
	)
}
