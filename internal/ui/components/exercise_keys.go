package components

import (
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

// KeyGenerator lets users generate a random private key and see the
// derived public key.
type KeyGenerator struct {
	Theme       *theme.Theme
	generateBtn widget.Clickable
	generated   bool
	privKey     string
	pubKey      string
}

// NewKeyGenerator creates a key generator from an exercise config.
func NewKeyGenerator(th *theme.Theme, cfg *content.ExerciseConfig) *KeyGenerator {
	return &KeyGenerator{Theme: th}
}

// Layout renders the key generator.
func (k *KeyGenerator) Layout(gtx layout.Context) layout.Dimensions {
	th := k.Theme

	if k.generateBtn.Clicked(gtx) {
		key, err := bitcoin.GenerateKey()
		if err == nil {
			k.generated = true
			k.privKey = hex.EncodeToString(key.Bytes())
			k.pubKey = hex.EncodeToString(key.PublicKey().SerializeCompressed())
		}
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.keys_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Generate button.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th.Material, &k.generateBtn, i18n.T("exercise.generate"))
				btn.Background = th.Color.Primary
				btn.Color = th.Color.OnPrimary
				return btn.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Results.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !k.generated {
					lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.keys_prompt"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(k.keyDisplay(i18n.T("exercise_ui.keys_private"), k.privKey)),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					layout.Rigid(k.keyDisplay(i18n.T("exercise_ui.keys_public"), k.pubKey)),
					layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),
					// Safety note.
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise.safety_note"))
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
				)
			}),
		)
	})
}

func (k *KeyGenerator) keyDisplay(label, value string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(k.Theme.Material, k.Theme.Text.Caption, label)
				lbl.Color = k.Theme.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						paint.FillShape(gtx.Ops, k.Theme.Color.InfoBg,
							clip.Rect{Max: gtx.Constraints.Min}.Op())
						return layout.Dimensions{Size: gtx.Constraints.Min}
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: unit.Dp(8), Bottom: unit.Dp(8),
							Left: unit.Dp(12), Right: unit.Dp(12),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(k.Theme.Material, k.Theme.Text.Code, value)
							lbl.Color = k.Theme.Color.Text
							lbl.MaxLines = 3
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
		)
	}
}
