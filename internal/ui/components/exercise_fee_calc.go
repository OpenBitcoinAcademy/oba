package components

import (
	"fmt"
	"strconv"

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

// Fee estimation constants in virtual bytes.
const (
	feeInputWeight    = 68 // Typical P2WPKH input weight in vB.
	feeOutputWeight   = 31 // Typical P2WPKH output weight in vB.
	feeOverheadWeight = 10 // Transaction overhead (version, locktime, etc).
)

// feeRate defines a fee tier for display.
type feeRate struct {
	label     string
	satsPerVB int
}

// FeeCalculator computes estimated transaction fees based on
// input/output count and common fee rates.
type FeeCalculator struct {
	Theme        *theme.Theme
	inputEditor  widget.Editor
	outputEditor widget.Editor
	calcBtn      widget.Clickable
	resetBtn     widget.Clickable

	numInputs  int
	numOutputs int
	calculated bool
	weightVB   int
	rates      []feeRate
}

// NewFeeCalculator creates a fee calculator widget.
func NewFeeCalculator(th *theme.Theme, cfg *content.ExerciseConfig) *FeeCalculator {
	fc := &FeeCalculator{
		Theme:      th,
		numInputs:  1,
		numOutputs: 2,
		rates: []feeRate{
			{label: "Low (1 sat/vB)", satsPerVB: 1},
			{label: "Medium (5 sat/vB)", satsPerVB: 5},
			{label: "High (20 sat/vB)", satsPerVB: 20},
		},
	}
	fc.inputEditor.SingleLine = true
	fc.inputEditor.SetText("1")
	fc.outputEditor.SingleLine = true
	fc.outputEditor.SetText("2")
	return fc
}

// Layout renders the fee calculator.
func (fc *FeeCalculator) Layout(gtx layout.Context) layout.Dimensions {
	th := fc.Theme

	if fc.calcBtn.Clicked(gtx) {
		fc.calculate()
	}
	if fc.resetBtn.Clicked(gtx) {
		fc.calculated = false
		fc.inputEditor.SetText("1")
		fc.outputEditor.SetText("2")
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.fee_calc_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Prompt.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.fee_calc_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Input count editor.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return fc.editorRow(gtx, "Inputs", &fc.inputEditor)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Output count editor.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return fc.editorRow(gtx, "Outputs", &fc.outputEditor)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Calculate button.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th.Material, &fc.calcBtn, "Calculate")
				btn.Background = th.Color.Primary
				btn.Color = th.Color.OnPrimary
				return btn.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Results table.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !fc.calculated {
					return layout.Dimensions{}
				}
				return fc.renderTable(gtx)
			}),

			// Reset button.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !fc.calculated {
					return layout.Dimensions{}
				}
				return layout.Inset{Top: th.Space.Small}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(th.Material, &fc.resetBtn, i18n.T("exercise.reset"))
					btn.Background = th.Color.Surface
					btn.Color = th.Color.Primary
					return btn.Layout(gtx)
				})
			}),
		)
	})
}

func (fc *FeeCalculator) editorRow(gtx layout.Context, label string, ed *widget.Editor) layout.Dimensions {
	th := fc.Theme
	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Body, label+": ")
			lbl.Color = th.Color.Text
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					paint.FillShape(gtx.Ops, th.Color.Surface,
						clip.Rect{Max: gtx.Constraints.Min}.Op())
					paint.FillShape(gtx.Ops, th.Color.Divider,
						clip.Stroke{Path: clip.Rect{Max: gtx.Constraints.Min}.Path(), Width: 1}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(6), Bottom: unit.Dp(6),
						Left: unit.Dp(8), Right: unit.Dp(8),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						e := material.Editor(th.Material, ed, "1")
						e.Color = th.Color.Text
						e.TextSize = th.Text.Body
						return e.Layout(gtx)
					})
				}),
			)
		}),
	)
}

func (fc *FeeCalculator) calculate() {
	inputs, err := strconv.Atoi(fc.inputEditor.Text())
	if err != nil || inputs < 1 {
		inputs = 1
	}
	outputs, err := strconv.Atoi(fc.outputEditor.Text())
	if err != nil || outputs < 1 {
		outputs = 1
	}
	// Clamp to reasonable values.
	if inputs > 100 {
		inputs = 100
	}
	if outputs > 100 {
		outputs = 100
	}

	fc.numInputs = inputs
	fc.numOutputs = outputs
	fc.weightVB = inputs*feeInputWeight + outputs*feeOutputWeight + feeOverheadWeight
	fc.calculated = true
}

func (fc *FeeCalculator) renderTable(gtx layout.Context) layout.Dimensions {
	th := fc.Theme

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Weight line.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			text := fmt.Sprintf("Estimated size: %d vB (%d inputs x %d + %d outputs x %d + %d overhead)",
				fc.weightVB, fc.numInputs, feeInputWeight, fc.numOutputs, feeOutputWeight, feeOverheadWeight)
			lbl := material.Label(th.Material, th.Text.BodySmall, text)
			lbl.Color = th.Color.Primary
			lbl.Font.Weight = font.Bold
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
		// Fee rows.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			var children []layout.FlexChild
			for _, rate := range fc.rates {
				r := rate
				children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					fee := fc.weightVB * r.satsPerVB
					text := fmt.Sprintf("%s: %d sats", r.label, fee)
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							paint.FillShape(gtx.Ops, th.Color.InfoBg,
								clip.Rect{Max: gtx.Constraints.Min}.Op())
							return layout.Dimensions{Size: gtx.Constraints.Min}
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Top: unit.Dp(6), Bottom: unit.Dp(6),
								Left: unit.Dp(8), Right: unit.Dp(8),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.Code, text)
								lbl.Color = th.Color.Text
								return lbl.Layout(gtx)
							})
						}),
					)
				}))
				children = append(children, layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout))
			}
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
		}),
	)
}
