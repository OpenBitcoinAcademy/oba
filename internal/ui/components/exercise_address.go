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

// AddressBuilder walks users step-by-step through address derivation.
type AddressBuilder struct {
	Theme    *theme.Theme
	nextBtn  widget.Clickable
	resetBtn widget.Clickable
	step     int // 0=not started, 1..7=derivation steps
	steps    []addressStep
}

type addressStep struct {
	label  string
	input  string
	output string
}

// NewAddressBuilder creates an address builder from an exercise config.
func NewAddressBuilder(th *theme.Theme, cfg *content.ExerciseConfig) *AddressBuilder {
	return &AddressBuilder{Theme: th}
}

// Layout renders the address builder.
func (a *AddressBuilder) Layout(gtx layout.Context) layout.Dimensions {
	th := a.Theme

	if a.nextBtn.Clicked(gtx) {
		a.advance()
	}
	if a.resetBtn.Clicked(gtx) {
		a.step = 0
		a.steps = nil
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, "Address Builder")
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Steps display.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if a.step == 0 {
					lbl := material.Label(th.Material, th.Text.Body, "Press Next Step to generate a key pair and walk through address derivation.")
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}
				return a.renderSteps(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Buttons.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Spacing: layout.SpaceStart}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if a.step >= 7 {
							// All done, show reset.
							btn := material.Button(th.Material, &a.resetBtn, i18n.T("exercise.reset"))
							btn.Background = th.Color.Surface
							btn.Color = th.Color.Primary
							return btn.Layout(gtx)
						}
						btn := material.Button(th.Material, &a.nextBtn, i18n.T("exercise.next_step"))
						btn.Background = th.Color.Primary
						btn.Color = th.Color.OnPrimary
						return btn.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if a.step > 0 && a.step < 7 {
							return layout.Inset{Left: th.Space.Small}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Label(th.Material, th.Text.Caption,
									fmt.Sprintf("Step %d of 7", a.step))
								lbl.Color = th.Color.TextMuted
								return lbl.Layout(gtx)
							})
						}
						return layout.Dimensions{}
					}),
				)
			}),

			// Safety note.
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if a.step > 0 {
					lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise.safety_note"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}
				return layout.Dimensions{}
			}),
		)
	})
}

func (a *AddressBuilder) advance() {
	a.step++
	switch a.step {
	case 1: // Generate key pair.
		key, err := bitcoin.GenerateKey()
		if err != nil {
			return
		}
		pub := key.PublicKey()
		a.steps = append(a.steps, addressStep{
			label:  "1. Generate Key Pair",
			input:  "Random 256-bit number",
			output: "Public Key: " + hex.EncodeToString(pub.SerializeCompressed()),
		})
		// Pre-compute all remaining steps from this key.
		a.computeRemaining(pub)

	case 2, 3, 4, 5, 6, 7:
		// Steps were pre-computed in step 1. Just reveal them.
	}
}

func (a *AddressBuilder) computeRemaining(pub *bitcoin.PublicKey) {
	compressed := pub.SerializeCompressed()

	// Step 2: SHA-256 of public key.
	sha := bitcoin.SHA256(compressed)
	a.steps = append(a.steps, addressStep{
		label:  "2. SHA-256(Public Key)",
		input:  hex.EncodeToString(compressed),
		output: hex.EncodeToString(sha[:]),
	})

	// Step 3: RIPEMD-160 of SHA-256.
	ripe := bitcoin.RIPEMD160(sha[:])
	a.steps = append(a.steps, addressStep{
		label:  "3. RIPEMD-160(SHA-256)",
		input:  hex.EncodeToString(sha[:]),
		output: hex.EncodeToString(ripe[:]),
	})

	// Step 4: Add version byte (0x6F for testnet).
	versioned := append([]byte{0x6F}, ripe[:]...)
	a.steps = append(a.steps, addressStep{
		label:  "4. Add Version Byte (0x6F = testnet)",
		input:  hex.EncodeToString(ripe[:]),
		output: hex.EncodeToString(versioned),
	})

	// Step 5: Double SHA-256 for checksum.
	checkHash := bitcoin.HASH256(versioned)
	a.steps = append(a.steps, addressStep{
		label:  "5. HASH256(versioned) for checksum",
		input:  hex.EncodeToString(versioned),
		output: hex.EncodeToString(checkHash[:]),
	})

	// Step 6: Take first 4 bytes as checksum.
	checksum := checkHash[:4]
	a.steps = append(a.steps, addressStep{
		label:  "6. Take First 4 Bytes",
		input:  hex.EncodeToString(checkHash[:]),
		output: hex.EncodeToString(checksum),
	})

	// Step 7: Base58Check encode.
	address := bitcoin.Base58CheckEncode(0x6F, ripe[:])
	a.steps = append(a.steps, addressStep{
		label:  "7. Base58Check Encode",
		input:  hex.EncodeToString(versioned) + hex.EncodeToString(checksum),
		output: address,
	})
}

func (a *AddressBuilder) renderSteps(gtx layout.Context) layout.Dimensions {
	th := a.Theme
	var children []layout.FlexChild

	// Show steps up to current step.
	visible := a.step
	if visible > len(a.steps) {
		visible = len(a.steps)
	}

	for i := 0; i < visible; i++ {
		s := a.steps[i]
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// Step label.
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.BodySmall, s.label)
					lbl.Color = th.Color.Primary
					lbl.Font.Weight = font.Bold
					return lbl.Layout(gtx)
				}),
				// Output.
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
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
								lbl := material.Label(th.Material, th.Text.Code, s.output)
								lbl.Color = th.Color.Text
								lbl.MaxLines = 3
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
