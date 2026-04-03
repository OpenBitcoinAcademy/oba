package components

import (
	"encoding/hex"
	"fmt"

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

// htlcStep represents one step in a 3-hop HTLC payment flow.
type htlcStep struct {
	label   string
	from    string
	to      string
	amount  int64
	timeout int
	detail  string
}

// HTLCTracer walks through a 3-hop Lightning payment step by step:
// Alice -> Bob -> Carol, then preimage propagation back.
type HTLCTracer struct {
	Theme    *theme.Theme
	nextBtn  widget.Clickable
	resetBtn widget.Clickable

	step       int
	steps      []htlcStep
	preimage   string
	payHash    string
	totalSteps int
}

// NewHTLCTracer creates an HTLC payment tracer widget.
func NewHTLCTracer(th *theme.Theme, cfg *content.ExerciseConfig) *HTLCTracer {
	ht := &HTLCTracer{Theme: th}
	ht.buildSteps()
	return ht
}

// Layout renders the HTLC tracer.
func (ht *HTLCTracer) Layout(gtx layout.Context) layout.Dimensions {
	th := ht.Theme

	if ht.nextBtn.Clicked(gtx) && ht.step < ht.totalSteps {
		ht.step++
	}
	if ht.resetBtn.Clicked(gtx) {
		ht.step = 0
		ht.buildSteps()
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.htlc_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Prompt.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.htlc_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),

			// Payment hash display.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Caption, "Payment Hash")
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(2)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return ht.codeBox(gtx, ht.payHash[:32]+"...", th.Color.InfoBg)
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Steps rendered so far.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ht.renderSteps(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Buttons.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if ht.step >= ht.totalSteps {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.Body, "Payment complete.")
							lbl.Color = th.Color.Success
							lbl.Font.Weight = font.Bold
							return lbl.Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: th.Space.Medium}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th.Material, &ht.resetBtn, i18n.T("exercise.reset"))
							btn.Background = th.Color.Surface
							btn.Color = th.Color.Primary
							return btn.Layout(gtx)
						}),
					)
				}
				return layout.Flex{}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &ht.nextBtn, i18n.T("exercise.next_step"))
						btn.Background = th.Color.Primary
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Caption,
							fmt.Sprintf("Step %d / %d", ht.step, ht.totalSteps))
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
				)
			}),
		)
	})
}

func (ht *HTLCTracer) buildSteps() {
	// Generate a random preimage and compute the payment hash.
	entropy, err := bitcoin.GenerateEntropy(32)
	if err != nil {
		// Fallback: use a static value. Should not happen.
		entropy = make([]byte, 32)
	}
	ht.preimage = hex.EncodeToString(entropy)
	hash := bitcoin.SHA256(entropy)
	ht.payHash = hex.EncodeToString(hash[:])

	ht.steps = []htlcStep{
		{
			label:   "1. Alice creates invoice",
			from:    "Carol",
			to:      "Alice",
			amount:  0,
			timeout: 0,
			detail:  "Carol generates preimage, shares payment hash with Alice",
		},
		{
			label:   "2. Alice -> Bob: HTLC offer",
			from:    "Alice",
			to:      "Bob",
			amount:  10200,
			timeout: 300,
			detail:  "Alice locks 10,200 sats with Bob. Timelock: 300 blocks.",
		},
		{
			label:   "3. Bob -> Carol: HTLC offer",
			from:    "Bob",
			to:      "Carol",
			amount:  10100,
			timeout: 250,
			detail:  "Bob locks 10,100 sats with Carol. Timelock: 250 blocks. Bob keeps 100 sats routing fee.",
		},
		{
			label:   "4. Carol reveals preimage to Bob",
			from:    "Carol",
			to:      "Bob",
			amount:  10100,
			timeout: 0,
			detail:  fmt.Sprintf("Carol reveals preimage: %s...  Bob verifies hash match, settles HTLC.", ht.preimage[:16]),
		},
		{
			label:   "5. Bob reveals preimage to Alice",
			from:    "Bob",
			to:      "Alice",
			amount:  10200,
			timeout: 0,
			detail:  "Bob uses preimage to claim 10,200 sats from Alice. Net gain: 100 sats fee.",
		},
		{
			label:   "6. Settlement complete",
			from:    "",
			to:      "",
			amount:  0,
			timeout: 0,
			detail:  fmt.Sprintf("Alice paid 10,200 sats. Carol received 10,100 sats. Bob earned 100 sats. Preimage: %s", ht.preimage[:32]+"..."),
		},
	}
	ht.totalSteps = len(ht.steps)
	ht.step = 0
}

func (ht *HTLCTracer) renderSteps(gtx layout.Context) layout.Dimensions {
	th := ht.Theme
	var children []layout.FlexChild

	visible := ht.step
	if visible > len(ht.steps) {
		visible = len(ht.steps)
	}

	for i := 0; i < visible; i++ {
		s := ht.steps[i]
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall, s.label)
					lbl.Color = th.Color.Primary
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if s.from != "" && s.to != "" && s.amount > 0 {
						text := fmt.Sprintf("%s -> %s: %d sats", s.from, s.to, s.amount)
						if s.timeout > 0 {
							text += fmt.Sprintf(" (timelock: %d blocks)", s.timeout)
						}
						return ht.codeBox(gtx, text, th.Color.InfoBg)
					}
					return layout.Dimensions{}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.Caption, s.detail)
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			)
		}))
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
}

func (ht *HTLCTracer) codeBox(gtx layout.Context, text string, bg interface{}) layout.Dimensions {
	th := ht.Theme
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			paint.FillShape(gtx.Ops, th.Color.InfoBg,
				clip.Rect{Max: gtx.Constraints.Min}.Op())
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: unit.Dp(4), Bottom: unit.Dp(4),
				Left: unit.Dp(8), Right: unit.Dp(8),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Code, text)
				lbl.Color = th.Color.Text
				lbl.MaxLines = 2
				return lbl.Layout(gtx)
			})
		}),
	)
}
