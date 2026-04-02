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

// MASTBuilder shows a 2-leaf MAST tree. Users see two script leaves,
// their tagged hashes, and the computed Merkle root (TapBranch).
type MASTBuilder struct {
	Theme     *theme.Theme
	buildBtn  widget.Clickable
	resetBtn  widget.Clickable
	scriptAEd widget.Editor
	scriptBEd widget.Editor

	built      bool
	leafAHash  string
	leafBHash  string
	branchHash string
}

// NewMASTBuilder creates a MAST tree builder widget.
func NewMASTBuilder(th *theme.Theme, cfg *content.ExerciseConfig) *MASTBuilder {
	mb := &MASTBuilder{Theme: th}
	mb.scriptAEd.SingleLine = true
	mb.scriptAEd.SetText("OP_CHECKSIG <pubkeyA>")
	mb.scriptBEd.SingleLine = true
	mb.scriptBEd.SetText("OP_CHECKMULTISIG 2-of-3")
	return mb
}

// Layout renders the MAST tree builder.
func (mb *MASTBuilder) Layout(gtx layout.Context) layout.Dimensions {
	th := mb.Theme

	if mb.buildBtn.Clicked(gtx) {
		mb.computeTree()
	}
	if mb.resetBtn.Clicked(gtx) {
		mb.built = false
		mb.leafAHash = ""
		mb.leafBHash = ""
		mb.branchHash = ""
	}

	return layout.Inset{
		Top: th.Space.Medium, Bottom: th.Space.Medium,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// Title.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.H3, i18n.T("exercise_ui.mast_title"))
				lbl.Color = th.Color.Text
				lbl.Font.Weight = font.Bold
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Prompt.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lbl := material.Label(th.Material, th.Text.Body, i18n.T("exercise_ui.mast_prompt"))
				lbl.Color = th.Color.TextMuted
				return lbl.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Script A editor.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return mb.scriptRow(gtx, "Script A (Leaf)", &mb.scriptAEd)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),
			// Script B editor.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return mb.scriptRow(gtx, "Script B (Leaf)", &mb.scriptBEd)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Build button.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th.Material, &mb.buildBtn, "Build Tree")
				btn.Background = th.Color.Primary
				btn.Color = th.Color.OnPrimary
				return btn.Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: th.Space.Medium}.Layout),

			// Tree visualization.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !mb.built {
					return layout.Dimensions{}
				}
				return mb.renderTree(gtx)
			}),

			// Reset.
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if !mb.built {
					return layout.Dimensions{}
				}
				return layout.Inset{Top: th.Space.Small}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(th.Material, &mb.resetBtn, i18n.T("exercise.reset"))
					btn.Background = th.Color.Surface
					btn.Color = th.Color.Primary
					return btn.Layout(gtx)
				})
			}),
		)
	})
}

func (mb *MASTBuilder) scriptRow(gtx layout.Context, label string, ed *widget.Editor) layout.Dimensions {
	th := mb.Theme
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Caption, label)
			lbl.Color = th.Color.TextMuted
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
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
						e := material.Editor(th.Material, ed, "script")
						e.Color = th.Color.Text
						e.TextSize = th.Text.Code
						return e.Layout(gtx)
					})
				}),
			)
		}),
	)
}

// computeTree hashes each leaf with a TapLeaf tagged hash, then
// computes the TapBranch by sorting and concatenating the leaf hashes.
// This is a simplified demonstration: real Taproot uses BIP-340 tagged
// hashes, but this exercise uses SHA-256 to show the Merkle structure.
func (mb *MASTBuilder) computeTree() {
	scriptA := mb.scriptAEd.Text()
	scriptB := mb.scriptBEd.Text()

	// Tagged leaf hash: SHA256("TapLeaf" || script).
	tagLeaf := bitcoin.SHA256([]byte("TapLeaf"))
	leafAData := append(tagLeaf[:], []byte(scriptA)...)
	leafBData := append(tagLeaf[:], []byte(scriptB)...)

	hashA := bitcoin.SHA256(leafAData)
	hashB := bitcoin.SHA256(leafBData)

	mb.leafAHash = hex.EncodeToString(hashA[:])
	mb.leafBHash = hex.EncodeToString(hashB[:])

	// TapBranch: sort lexicographically, concatenate, hash.
	var combined []byte
	if mb.leafAHash < mb.leafBHash {
		combined = append(hashA[:], hashB[:]...)
	} else {
		combined = append(hashB[:], hashA[:]...)
	}
	tagBranch := bitcoin.SHA256([]byte("TapBranch"))
	branchData := append(tagBranch[:], combined...)
	branch := bitcoin.SHA256(branchData)
	mb.branchHash = hex.EncodeToString(branch[:])

	mb.built = true
}

func (mb *MASTBuilder) renderTree(gtx layout.Context) layout.Dimensions {
	th := mb.Theme
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Merkle root (TapBranch).
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Caption, "Merkle Root (TapBranch)")
			lbl.Color = th.Color.Primary
			lbl.Font.Weight = font.Bold
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout),
		layout.Rigid(mb.hashBox(mb.branchHash, th.Color.TipBg)),
		layout.Rigid(layout.Spacer{Height: th.Space.Small}.Layout),

		// Connector label.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Label(th.Material, th.Text.Caption, fmt.Sprintf("  %s         %s", "\u250c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524", "\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2510"))
			lbl.Color = th.Color.TextMuted
			lbl.Font.Typeface = "monospace"
			return lbl.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: th.Space.XSmall}.Layout),

		// Leaves side by side.
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.Caption, "Leaf A Hash")
							lbl.Color = th.Color.TextMuted
							return lbl.Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(2)}.Layout),
						layout.Rigid(mb.hashBox(mb.leafAHash[:16]+"...", th.Color.InfoBg)),
					)
				}),
				layout.Rigid(layout.Spacer{Width: th.Space.Small}.Layout),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lbl := material.Label(th.Material, th.Text.Caption, "Leaf B Hash")
							lbl.Color = th.Color.TextMuted
							return lbl.Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(2)}.Layout),
						layout.Rigid(mb.hashBox(mb.leafBHash[:16]+"...", th.Color.InfoBg)),
					)
				}),
			)
		}),
	)
}

func (mb *MASTBuilder) hashBox(value string, bg interface{}) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		th := mb.Theme
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
					lbl := material.Label(th.Material, th.Text.Code, value)
					lbl.Color = th.Color.Text
					lbl.MaxLines = 2
					return lbl.Layout(gtx)
				})
			}),
		)
	}
}
