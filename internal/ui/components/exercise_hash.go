package components

import (
	"encoding/base64"
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

// HashExplorer lets users type text and see its SHA-256 hash live.
type HashExplorer struct {
	Theme  *theme.Theme
	editor widget.Editor

	// Output format toggle: 0=hex, 1=binary, 2=base64.
	format    int
	formatBtn widget.Clickable

	// Preloaded example buttons (variable length from config).
	exampleBtns []widget.Clickable
	examples    []string
}

// NewHashExplorer creates a hash explorer from an exercise config.
// Falls back to defaults if config is nil.
func NewHashExplorer(th *theme.Theme, cfg *content.ExerciseConfig) *HashExplorer {
	examples := []string{"bitcoin", "Bitcoin", "bitcoin!"}
	if cfg != nil {
		if ex := cfg.ConfigStrings("preloaded_examples"); len(ex) > 0 {
			examples = ex
		}
	}
	h := &HashExplorer{
		Theme:       th,
		examples:    examples,
		exampleBtns: make([]widget.Clickable, len(examples)),
	}
	h.editor.SingleLine = true
	h.editor.Submit = false
	return h
}

// Layout renders the hash explorer.
func (h *HashExplorer) Layout(gtx layout.Context) layout.Dimensions {
	th := h.Theme

	// Handle format toggle.
	if h.formatBtn.Clicked(gtx) {
		h.format = (h.format + 1) % 3
	}

	// Handle example buttons.
	for i := range h.exampleBtns {
		if h.exampleBtns[i].Clicked(gtx) {
			h.editor.SetText(h.examples[i])
		}
	}

	input := h.editor.Text()
	hash := bitcoin.SHA256([]byte(input))
	output := h.formatHash(hash[:])
	formatLabel := h.formatName()

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.hash_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Input field.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise_ui.hash_input_label"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// Input background.
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
							Top: unit.Dp(8), Bottom: unit.Dp(8),
							Left: unit.Dp(12), Right: unit.Dp(12),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							ed := material.Editor(th.Material, &h.editor, "Hello")
							ed.Color = th.Color.Text
							ed.TextSize = th.Text.Body
							return ed.Layout(gtx)
						})
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Example buttons (dynamic count from config).
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				var children []layout.FlexChild
				children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise_ui.hash_try"))
					lbl.Color = th.Color.TextMuted
					return lbl.Layout(gtx)
				}))
				for i := range h.examples {
					children = append(children, layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout))
					children = append(children, layout.Rigid(h.exampleButton(i)))
				}
				return layout.Flex{Spacing: layout.SpaceStart}.Layout(gtx, children...)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Output label.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Caption, i18n.T("exercise_ui.hash_sha256"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),

			// Hash output.
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
							lbl := material.Label(th.Material, th.Text.Code, output)
							lbl.Color = th.Color.Text
							lbl.MaxLines = 4
							return lbl.Layout(gtx)
						})
					}),
				)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),

			// Stats row.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th.Material, &h.formatBtn, formatLabel)
						btn.Background = th.Color.Surface
						btn.Color = th.Color.Primary
						btn.TextSize = th.Text.Caption
						return btn.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						text := i18n.TFmt("exercise_ui.hash_stats", map[string]string{"bytes": fmt.Sprintf("%d", len(input))})
						lbl := material.Label(th.Material, th.Text.Caption, text)
						lbl.Color = th.Color.TextMuted
						return lbl.Layout(gtx)
					}),
				)
			}),
		)
	})
}

func (h *HashExplorer) exampleButton(idx int) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(h.Theme.Material, &h.exampleBtns[idx], h.examples[idx])
		btn.Background = h.Theme.Color.Surface
		btn.Color = h.Theme.Color.Accent
		btn.TextSize = h.Theme.Text.Caption
		return btn.Layout(gtx)
	}
}

func (h *HashExplorer) formatHash(hash []byte) string {
	switch h.format {
	case 1:
		return toBinary(hash)
	case 2:
		return base64.StdEncoding.EncodeToString(hash)
	default:
		return hex.EncodeToString(hash)
	}
}

func (h *HashExplorer) formatName() string {
	switch h.format {
	case 1:
		return i18n.T("exercise_ui.hash_binary")
	case 2:
		return i18n.T("exercise_ui.hash_base64")
	default:
		return i18n.T("exercise_ui.hash_hex")
	}
}

func toBinary(data []byte) string {
	var result []byte
	for i, b := range data {
		if i > 0 && i%4 == 0 {
			result = append(result, ' ')
		}
		result = append(result, []byte(fmt.Sprintf("%08b", b))...)
	}
	return string(result)
}

// Ensure image is used.
